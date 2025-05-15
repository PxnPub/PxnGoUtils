package net;

import(
	OS      "os"
	Log     "log"
	Fmt     "fmt"
	Net     "net"
	Strings "strings"
	Errors  "errors"
	Binary  "encoding/binary"
	San     "github.com/PxnPub/pxnGoUtils/san"
);



func NewSock(bind string) (*Net.Listener, error) {
	if bind == "" { return nil, Errors.New("bind address required"); }
	if Strings.Contains(bind, "://") {
		parts := Strings.SplitN(bind, "://", 2);
		return NewSocket(parts[0], parts[1]);
	}
	return nil, Fmt.Errorf("Invalid protocol/address: %s", bind);
}

func NewSocket(protocol string, address string) (*Net.Listener, error) {
	if protocol == "" || address == "" { return nil, Errors.New("protocol and address are required"); }
	if !San.IsSafeAlphaLower(protocol) { return nil, Fmt.Errorf("Invalid protocol: %s",    protocol); }
	if len(address) < 5 {                return nil, Fmt.Errorf("Invalid unix socket: %s", address ); }
	switch protocol {
		case "unix":
			if !San.IsSafeFilePath(address) {
				return nil, Fmt.Errorf("Invalid address: %s", address); }
			if _, err := OS.Stat(address); err == nil {
				if err := OS.Remove(address); err != nil {
					Log.Panicf("Failed to remove existing socket file: %s", address); }}
			break;
		case "tcp":
			if !San.IsSafeDomainPort(address) {
				return nil, Fmt.Errorf("Invalid address: %s", address); }
			break;
		default: return nil, Fmt.Errorf("Unknown protocol: %s", protocol);
	}
	listen, err := Net.Listen(protocol, address);
	if err != nil { return nil, err; }
	return &listen, nil;
}



func IPToIntPair(address string) (uint64, uint64, error) {
	var ip Net.IP = Net.ParseIP(address);
	if ip == nil { return 0, 0, Fmt.Errorf("Invalid address: %s", address); }
	// ipv4
	if ip.To4() != nil {
		ip4 := ip.To4();
		return 0, uint64(Binary.BigEndian.Uint32(ip4)), nil;
	// ipv6
	} else {
		ip6 := ip.To16();
		return Binary.BigEndian.Uint64(ip6[0:8]),
			Binary.BigEndian.Uint64(ip6[8:16]),
			nil;
	}
}
