package netlink

import (
	"net"
	"syscall"
	"unsafe"
)

const (
	IFF_UP          = 0x1     // interface is up (administratively)
	IFF_BROADCAST   = 0x2     // broadcast address valid
	IFF_DEBUG       = 0x4     // turn on debugging
	IFF_LOOPBACK    = 0x8     // is a loopback net
	IFF_POINTOPOINT = 0x10    // interface is has p-p link
	IFF_NOTRAILERS  = 0x20    // avoid use of trailers
	IFF_RUNNING     = 0x40    // interface RFC2863 OPER_UP
	IFF_NOARP       = 0x80    // no ARP protocol
	IFF_PROMISC     = 0x100   // receive all packets
	IFF_ALLMULTI    = 0x200   // receive all multicast packets
	IFF_MASTER      = 0x400   // master of a load balancer
	IFF_SLAVE       = 0x800   // slave of a load balancer
	IFF_MULTICAST   = 0x1000  // Supports multicast
	IFF_PORTSEL     = 0x2000  // can set media type
	IFF_AUTOMEDIA   = 0x4000  // auto media select active
	IFF_DYNAMIC     = 0x8000  // dialup device with changing addresses
	IFF_LOWER_UP    = 0x10000 // driver signals L1 up
	IFF_DORMANT     = 0x20000 // driver signals dormant
	IFF_ECHO        = 0x40000 // echo sent packets
)

type ExtFlags uint

const (
	FlagUp ExtFlags = 1 << iota
	FlagBroadcast
	FlagDebug
	FlagLoopback
	FlagPointToPoint
	FlagNoTrailers
	FlagRunning
	FlagNoArp
	FlagPromisc
	FlagAllMulti
	FlagMaster
	FlagSlave
	FlagMulticast
	FlagPortSel
	FlagAutoMedia
	FlagDynamic
	FlagLowerUp
	FlagDormant
	FlagEcho
)

var extFlagNames = []string{
	"up",
	"broadcast",
	"debug",
	"loopback",
	"point_to_point",
	"no_trailers",
	"running",
	"no_arp",
	"promisc",
	"all_multi",
	"master",
	"slave",
	"multicast",
	"port_sel",
	"auto_media",
	"dynamic",
	"lower_up",
	"dormant",
	"echo",
}

const (
	// See linux/if_arp.h.
	// Note that Linux doesn't support IPv4 over IPv6 tunneling.
	sysARPHardwareIPv4IPv4 = 768 // IPv4 over IPv4 tunneling
	sysARPHardwareIPv6IPv6 = 769 // IPv6 over IPv6 tunneling
	sysARPHardwareIPv6IPv4 = 776 // IPv6 over IPv4 tunneling
	sysARPHardwareGREIPv4  = 778 // any over GRE over IPv4 tunneling
	sysARPHardwareGREIPv6  = 823 // any over GRE over IPv6 tunneling
)

type Interface struct {
	ExtFlags ExtFlags
	*net.Interface
}

func (f ExtFlags) String() string {
	s := ""
	for i, name := range extFlagNames {
		if f&(1<<uint(i)) != 0 {
			if s != "" {
				s += "|"
			}
			s += name
		}
	}
	if s == "" {
		s = "0"
	}
	return s
}

func linkFlags(rawFlags uint32) net.Flags {
	var f net.Flags
	if rawFlags&IFF_UP != 0 {
		f |= net.FlagUp
	}
	if rawFlags&IFF_BROADCAST != 0 {
		f |= net.FlagBroadcast
	}
	if rawFlags&IFF_LOOPBACK != 0 {
		f |= net.FlagLoopback
	}
	if rawFlags&IFF_POINTOPOINT != 0 {
		f |= net.FlagPointToPoint
	}
	if rawFlags&IFF_MULTICAST != 0 {
		f |= net.FlagMulticast
	}
	return f
}

func linkExtFlags(flags uint32) ExtFlags {
	var f ExtFlags
	if flags&IFF_UP != 0 {
		f |= FlagUp
	}
	if flags&IFF_BROADCAST != 0 {
		f |= FlagBroadcast
	}
	if flags&IFF_DEBUG != 0 {
		f |= FlagDebug
	}
	if flags&IFF_LOOPBACK != 0 {
		f |= FlagLoopback
	}
	if flags&IFF_POINTOPOINT != 0 {
		f |= FlagPointToPoint
	}
	if flags&IFF_NOTRAILERS != 0 {
		f |= FlagNoTrailers
	}
	if flags&IFF_RUNNING != 0 {
		f |= FlagRunning
	}
	if flags&IFF_NOARP != 0 {
		f |= FlagNoArp
	}
	if flags&IFF_PROMISC != 0 {
		f |= FlagPromisc
	}
	if flags&IFF_ALLMULTI != 0 {
		f |= FlagAllMulti
	}
	if flags&IFF_MASTER != 0 {
		f |= FlagMaster
	}
	if flags&IFF_SLAVE != 0 {
		f |= FlagSlave
	}
	if flags&IFF_MULTICAST != 0 {
		f |= FlagMulticast
	}
	if flags&IFF_PORTSEL != 0 {
		f |= FlagPortSel
	}
	if flags&IFF_AUTOMEDIA != 0 {
		f |= FlagAutoMedia
	}
	if flags&IFF_DYNAMIC != 0 {
		f |= FlagDynamic
	}
	if flags&IFF_LOWER_UP != 0 {
		f |= FlagLowerUp
	}
	if flags&IFF_DORMANT != 0 {
		f |= FlagDormant
	}
	if flags&IFF_ECHO != 0 {
		f |= FlagEcho
	}
	return f
}

func ParseNewLink(ifim *syscall.IfInfomsg, attrs []syscall.NetlinkRouteAttr) *Interface {
	ifi := Interface{
		ExtFlags:  linkExtFlags(ifim.Flags),
		Interface: &net.Interface{Index: int(ifim.Index), Flags: linkFlags(ifim.Flags)},
	}

	for _, a := range attrs {
		switch a.Attr.Type {
		case syscall.IFLA_ADDRESS:
			// We never return any /32 or /128 IP address
			// prefix on any IP tunnel interface as the
			// hardware address.
			switch len(a.Value) {
			case net.IPv4len:
				switch ifim.Type {
				case sysARPHardwareIPv4IPv4, sysARPHardwareGREIPv4, sysARPHardwareIPv6IPv4:
					continue
				}
			case net.IPv6len:
				switch ifim.Type {
				case sysARPHardwareIPv6IPv6, sysARPHardwareGREIPv6:
					continue
				}
			}
			var nonzero bool
			for _, b := range a.Value {
				if b != 0 {
					nonzero = true
					break
				}
			}
			if nonzero {
				ifi.HardwareAddr = a.Value[:]
			}
		case syscall.IFLA_IFNAME:
			ifi.Name = string(a.Value[:len(a.Value)-1])
		case syscall.IFLA_MTU:
			ifi.MTU = int(*(*uint32)(unsafe.Pointer(&a.Value[:4][0])))
		}
	}
	return &ifi
}