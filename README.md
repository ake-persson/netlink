[![GoDoc](https://godoc.org/github.com/mickep76/netlink?status.svg)](https://godoc.org/github.com/mickep76/netlink)


# netlink
`import "github.com/mickep76/netlink"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Subdirectories](#pkg-subdirectories)

## <a name="pkg-overview">Overview</a>



## <a name="pkg-index">Index</a>
* [Constants](#pkg-constants)
* [func Interfaces() ([]Interface, error)](#Interfaces)
* [type Conn](#Conn)
  * [func Dial(family int, groups uint32) (*Conn, error)](#Dial)
  * [func (c *Conn) Bind() error](#Conn.Bind)
  * [func (c *Conn) Close() error](#Conn.Close)
  * [func (c *Conn) Receive() ([]syscall.NetlinkMessage, error)](#Conn.Receive)
* [type Flags](#Flags)
  * [func (f Flags) MarshalJSON() ([]byte, error)](#Flags.MarshalJSON)
  * [func (f Flags) Slice() []string](#Flags.Slice)
  * [func (f Flags) String() string](#Flags.String)
* [type HwAddr](#HwAddr)
  * [func (a HwAddr) MarshalJSON() ([]byte, error)](#HwAddr.MarshalJSON)
  * [func (a HwAddr) String() string](#HwAddr.String)
* [type Interface](#Interface)
  * [func ParseNewLink(ifim *syscall.IfInfomsg, attrs []syscall.NetlinkRouteAttr) *Interface](#ParseNewLink)


#### <a name="pkg-files">Package files</a>
[interfaces.go](/src/github.com/mickep76/netlink/interfaces.go) [netlink.go](/src/github.com/mickep76/netlink/netlink.go) [parse_newlink.go](/src/github.com/mickep76/netlink/parse_newlink.go) 


## <a name="pkg-constants">Constants</a>
``` go
const (
    // NetlinkRoute return socket descriptor.
    NetlinkRoute = 0

    // RtmGrpLink Rtnetlink multicast group.
    RtmGrpLink = 0x1
)
```



## <a name="Interfaces">func</a> [Interfaces](/src/target/interfaces.go?s=132:170#L10)
``` go
func Interfaces() ([]Interface, error)
```
Interfaces connect using rtnetlink and retrieve all network interfaces.




## <a name="Conn">type</a> [Conn](/src/target/netlink.go?s=276:400#L20)
``` go
type Conn struct {
    Family     int
    Groups     uint32
    FileDescr  int
    SocketAddr *unix.SockaddrNetlink
    Pid        uint32
}
```
Conn provides an interface for connecting to netlink socket.







### <a name="Dial">func</a> [Dial](/src/target/netlink.go?s=426:477#L29)
``` go
func Dial(family int, groups uint32) (*Conn, error)
```
Dial netlink socket.





### <a name="Conn.Bind">func</a> (\*Conn) [Bind](/src/target/netlink.go?s=729:756#L47)
``` go
func (c *Conn) Bind() error
```
Bind to netlink socket.




### <a name="Conn.Close">func</a> (\*Conn) [Close](/src/target/netlink.go?s=1172:1200#L70)
``` go
func (c *Conn) Close() error
```
Close netlink socket.




### <a name="Conn.Receive">func</a> (\*Conn) [Receive](/src/target/netlink.go?s=1279:1337#L75)
``` go
func (c *Conn) Receive() ([]syscall.NetlinkMessage, error)
```
Receive messages from netlink socket.




## <a name="Flags">type</a> [Flags](/src/target/parse_newlink.go?s=612:627#L34)
``` go
type Flags uint
```
Flags type for network interface state.


``` go
const (
    // FlagUp interface is up (administratively).
    FlagUp Flags = 1 << iota

    // FlagBroadcast broadcast address valid.
    FlagBroadcast

    // FlagDebug turn on debugging.
    FlagDebug

    // FlagLoopback is a loopback net.
    FlagLoopback

    // FlagPointToPoint interface is has p-p link.
    FlagPointToPoint

    // FlagNoTrailers avoid use of trailers.
    FlagNoTrailers

    // FlagRunning interface RFC2863 OPER_UP.
    FlagRunning

    // FlagNoArp no ARP protocol.
    FlagNoArp

    // FlagPromisc receive all packets.
    FlagPromisc

    // FlagAllMulti receive all multicast packets.
    FlagAllMulti

    // FlagMaster master of a load balancer.
    FlagMaster

    // FlagSlave slave of a load balancer.
    FlagSlave

    // FlagMulticast supports multicast.
    FlagMulticast

    // FlagPortSel can set media type.
    FlagPortSel

    // FlagAutoMedia auto media select active.
    FlagAutoMedia

    // FlagDynamic dialup device with changing addresses.
    FlagDynamic

    // FlagLowerUp driver signals L1 up.
    FlagLowerUp

    // FlagDormant driver signals dormant.
    FlagDormant

    // FlagEcho echo sent packets.
    FlagEcho
)
```









### <a name="Flags.MarshalJSON">func</a> (Flags) [MarshalJSON](/src/target/parse_newlink.go?s=3153:3197#L159)
``` go
func (f Flags) MarshalJSON() ([]byte, error)
```
MarshalJSON marshal flags into JSON.




### <a name="Flags.Slice">func</a> (Flags) [Slice](/src/target/parse_newlink.go?s=2959:2990#L148)
``` go
func (f Flags) Slice() []string
```
Slice return a list of all flags.




### <a name="Flags.String">func</a> (Flags) [String](/src/target/parse_newlink.go?s=2849:2879#L143)
``` go
func (f Flags) String() string
```
String return a string of all flags.




## <a name="HwAddr">type</a> [HwAddr](/src/target/parse_newlink.go?s=1988:2006#L118)
``` go
type HwAddr []byte
```
HwAddr hardware address type.










### <a name="HwAddr.MarshalJSON">func</a> (HwAddr) [MarshalJSON](/src/target/parse_newlink.go?s=4943:4988#L261)
``` go
func (a HwAddr) MarshalJSON() ([]byte, error)
```
MarshalJSON marshal hardware address into JSON.




### <a name="HwAddr.String">func</a> (HwAddr) [String](/src/target/parse_newlink.go?s=4624:4655#L245)
``` go
func (a HwAddr) String() string
```



## <a name="Interface">type</a> [Interface](/src/target/parse_newlink.go?s=2514:2807#L133)
``` go
type Interface struct {
    Index        int            `json:"index"`
    MTU          int            `json:"mtu"`
    Name         string         `json:"name"`
    HwAddr       HwAddr         `json:"hwaddr,omitempty"`
    Flags        Flags          `json:"flags"`
    NetInterface *net.Interface `json:"-"`
}
```
Interface provides information about a network interface.







### <a name="ParseNewLink">func</a> [ParseNewLink](/src/target/parse_newlink.go?s=5073:5160#L266)
``` go
func ParseNewLink(ifim *syscall.IfInfomsg, attrs []syscall.NetlinkRouteAttr) *Interface
```
ParseNewLink parse interface info message.









- - -
Generated by [godoc2md](http://godoc.org/github.com/davecheney/godoc2md)
# Credits

- Parse messages based on Go package [net](https://golang.org/src/net/interface_linux.go).
- Socket connection based on Go package [mdlayher/netlink](https://github.com/mdlayher/netlink).
