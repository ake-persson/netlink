package main

import (
	"encoding/json"
	"fmt"
	"log"
	"syscall"
	"unsafe"

	"github.com/mickep76/netlink"
)

func listen() error {
	clnt, err := netlink.Dial(netlink.NetlinkRoute, netlink.RtmGrpLink)
	if err != nil {
		log.Fatal(err)
	}

	if err := clnt.Bind(); err != nil {
		log.Fatal(err)
	}

	defer clnt.Close()

	for {
		msgs, err := clnt.Receive()
		if err != nil {
			log.Fatal(err)
		}

		for _, m := range msgs {
			switch m.Header.Type {
			case syscall.NLMSG_DONE:
				break
			case syscall.RTM_NEWLINK:
				ifim := (*syscall.IfInfomsg)(unsafe.Pointer(&m.Data[0]))
				attrs, err := syscall.ParseNetlinkRouteAttr(&m)
				if err != nil {
					return fmt.Errorf("parse netlink route attr: %v", err)
				}
				ift := *netlink.ParseNewLink(ifim, attrs)
				j, _ := json.MarshalIndent(ift, "", "  ")
				fmt.Printf("%s\n", j)
			}
		}
	}
}

func main() {
	ifs, err := netlink.Interfaces()
	if err != nil {
		log.Fatal(err)
	}

	for _, i := range ifs {
		j, _ := json.MarshalIndent(i, "", "  ")
		fmt.Printf("%s\n", j)
	}

	if err := listen(); err != nil {
		log.Fatal(err)
	}
}
