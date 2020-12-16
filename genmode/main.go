package main

import (
	"log"
	"net"

	"github.com/jsimonetti/rtnetlink"
	"github.com/jsimonetti/rtnetlink/internal/unix"
	"github.com/mdlayher/netlink"
)

func main() {
	// Gather the interface Index
	iface, err := net.InterfaceByName("dummy0")
	if err != nil {
		log.Fatal(err)
	}

	// Dial a connection to the rtnetlink socket
	conn, err := rtnetlink.Dial(nil)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// Request the details of the interface
	msg, err := conn.Link.Get(uint32(iface.Index))
	if err != nil {
		log.Fatal(err)
	}

	state := msg.Attributes.OperationalState
	log.Println("link state:", state)

	req := &Message{
		Index: msg.Index,
	}

	flags := netlink.Request | netlink.Acknowledge
	_, err = conn.Execute(req, unix.RTM_NEWLINK, flags)

	log.Fatal(err)
}

func aeEncode(ae *netlink.AttributeEncoder) []byte {
	a, err := ae.Encode()
	if err != nil {
		log.Fatal(err)
	}

	return a
}
