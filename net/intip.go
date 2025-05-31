package net;

import(
	Fmt    "fmt"
	Net    "net"
	Binary "encoding/binary"
);



type IntIP struct {
	High uint64
	Low  uint64
}



func StringToIntIP(address string) (*IntIP, error) {
	var ip Net.IP = Net.ParseIP(address);
	if ip == nil { return nil, Fmt.Errorf("Invalid address: %s", address); }
	// ipv4
	if ip.To4() != nil {
		ip4 := ip.To4();
		return &IntIP{
			High: 0,
			Low: uint64(Binary.BigEndian.Uint32(ip4)),
		}, nil;
	// ipv6
	} else {
		ip6 := ip.To16();
		return &IntIP{
			High: Binary.BigEndian.Uint64(ip6[0:8]),
			Low:  Binary.BigEndian.Uint64(ip6[8:16]),
		}, nil;
	}
}
