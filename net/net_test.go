package net;

import(
	Testing "testing"
	Assert  "github.com/stretchr/testify/assert"
);



func Test_IPToIntPair(t *Testing.T) {
	var tests = []struct {
		ip       string
		expect_h uint64
		expect_l uint64
	}{
		// IP                H           L
		{ "127.0.0.1",       0, 2130706433 },
		{ "192.168.0.0",     0, 3232235520 },
		{ "192.168.1.1",     0, 3232235777 },
		{ "1.2.3.4",         0,   16909060 },
		{ "123.123.123.123", 0, 2071690107 },
		{ "0:0:0:0:0:0:0:1", 0,          1 },
		{ "::1",             0,          1 },
		{ "::",              0,          0 },
		{ "1:2:3:0:0:6::",                               281483566841856,         25769803776 },
		{ "1:2:3:0:0:6:0:0",                             281483566841856,         25769803776 },
		{ "1234:5678::8765:4321",                    1311768464867721216,          2271560481 },
		{ "1234:1234:1234:1234:1234:1234:1234:1234", 1311693406324658740, 1311693406324658740 },
	};
	for _, test := range tests {
		ip_h, ip_l, err := IPToIntPair(test.ip);
		Assert.Equal(t, test.expect_h, ip_h);
		Assert.Equal(t, test.expect_l, ip_l);
		Assert.Equal(t, nil, err);
	}
}
