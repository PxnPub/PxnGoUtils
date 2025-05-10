package net;

import(
	Fmt     "fmt"
	Net     "net"
	Strings "strings"
	Errors  "errors"
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
				return nil, Fmt.Errorf("Invalid address: %s", address);
			}
//			if err := OS.RemoveAll(address); err != nil {
//				Log.Panicf("Failed to remove existing socket file: %s", address);
//			}
			break;
		case "tcp":
			if !San.IsSafeDomainPort(address) {
				return nil, Fmt.Errorf("Invalid address: %s", address);
			}
			break;
		default: return nil, Fmt.Errorf("Unknown protocol: %s", protocol);
	}
	listen, err := Net.Listen(protocol, address);
	if err != nil { return nil, err; }
	return &listen, nil;
}
