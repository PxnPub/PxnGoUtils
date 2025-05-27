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
	protocol, address := SplitProtocolAddress(bind);
	if protocol == "" || address == "" {
		return nil, Fmt.Errorf("Invalid protocol/address: %s", bind);
	}
	return NewSocket(protocol, address);
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
		case "tcp":  fallthrough;
		case "tcp4": fallthrough;
		case "tcp6": fallthrough;
		case "udp":  fallthrough;
		case "udp4": fallthrough;
		case "udp6":
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



func SplitProtocolAddress(addr string) (string, string) {
	if addr == "" { return "", ""; }
	if Strings.Contains(addr, "://") {
		parts := Strings.SplitN(addr, "://", 2);
		return parts[0], parts[1];
	}
	return "", addr;
}



func RemoveOldSocket(file string) {
	// file exists
	if _, err := OS.Stat(file); err == nil {
		// file type
		info, err := OS.Lstat(file);
		if err != nil { panic(Fmt.Sprintf("Failed to stat file type: %v", err)); }
		// is a socket
		if info.Mode()&OS.ModeSocket != 0 {
			// remove old socket file
			if err := OS.Remove(file); err != nil {
				panic(Fmt.Sprintf("Failed to remove old socket file: %v\n", err));
			} else {
				panic(Fmt.Sprintf("Removed old socket file: %s\n", file));
			}
		} else {
			panic(Fmt.Sprintf("File exists but is not a socket: %s\n", file));
		}
	}
}
