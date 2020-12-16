package main

import (
	"errors"

	"github.com/jsimonetti/rtnetlink/internal/unix"
	"github.com/mdlayher/netlink"
	"github.com/mdlayher/netlink/nlenc"
)

var nativeEndian = nlenc.NativeEndian()

type Message struct {
	// Unique interface index, using a nonzero value with
	// NewLink will instruct the kernel to create a
	// device with the given index (kernel 3.7+ required)
	Index uint32
}

// MarshalBinary marshals a LinkMessage into a byte slice.
func (m *Message) MarshalBinary() ([]byte, error) {
	b := make([]byte, unix.SizeofIfInfomsg)

	b[0] = 0 // Family
	b[1] = 0 // reserved
	nativeEndian.PutUint16(b[2:4], 0)
	nativeEndian.PutUint32(b[4:8], m.Index)
	nativeEndian.PutUint32(b[8:12], 0)
	nativeEndian.PutUint32(b[12:16], 0)

	aeGenMode := netlink.NewAttributeEncoder()
	aeGenMode.Uint8(IFLA_INET6_ADDR_GEN_MODE, IN6_ADDR_GEN_MODE_NONE)

	aeAfInet := netlink.NewAttributeEncoder()
	aeAfInet.Bytes(unix.AF_INET6, aeEncode(aeGenMode))

	aeAfSpec := netlink.NewAttributeEncoder()
	aeAfSpec.Bytes(IFLA_AF_SPEC, aeEncode(aeAfInet))

	b = append(b, aeEncode(aeAfSpec)...)

	return b, nil
}

var errNotImplemented = errors.New("not implemented")

// UnmarshalBinary unmarshals the contents of a byte slice into a LinkMessage.
func (m *Message) UnmarshalBinary(b []byte) error {
	return errNotImplemented
}

// rtMessage is an empty method to sattisfy the Message interface.
func (m *Message) rtMessage() {}
