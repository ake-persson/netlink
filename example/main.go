package main

import (
	"fmt"
	"log"
	"syscall"
	"unsafe"

	"github.com/mickep76/netlink"
)

func listen() error {
	clnt, err := netlink.Dial(netlink.NETLINK_ROUTE, netlink.RTMGRP_LINK)
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
				log.Printf("%+v %+v", ift, ift.Interface)
			}
		}
	}
}

func main() {
	if err := listen(); err != nil {
		log.Fatal(err)
	}
}
