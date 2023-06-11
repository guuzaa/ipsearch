package ipsearch_test

import (
	"github.com/guuzaa/ipsearch"
	"github.com/stretchr/testify/assert"
	"strconv"
	"strings"
	"testing"
)

var (
	ipv6Lst = []string{
		"7d9d:9bcd:8f82:d66:f32c:25a7:5d53:fc8d",
		"9b98:f9e1:84a3:c99f:de51:4d75:2e2a:40c6",
		"4854:70f6:e7ef:8f26:7cb9:30b:6a5e:6a56",
		"5747:43f9:4cb7:5021:f49f:4267:5f24:a53d",
		"462:8f1c:8a4c:2e40:e75:281d:ef6f:5149",
		"bc98:4c23:44f2:a5cc:56d9:b9d1:4a16:9c75",
		"1425:94d:929c:b136:ea47:305c:91d:e1d9",
		"3efc:4296:c283:8934:3ea0:7566:9d4f:4b6a",
		"e8d5:d423:3d4e:294d:3b5a:8c16:ffe6:1d6c",
		"810:641d:b029:60c8:3329:3f4d:89b5:a3ba",
		"27:8e60:afdf:37d2:1dba:fd0e:ca59:d516",
		"e844:dd8:9cf5:fe5e:6e60:7dba:2c09:c5d3",
	}

	testcasesV6 = []struct {
		cidr     string
		ipStr    string
		rangeStr string
		flag     bool
	}{
		{"2001:b28::/32", "2001:b28::1", "2001:b28:: - 2001:b28:ffff:ffff:ffff:ffff:ffff:ffff", true},
		{"2605:9e40::/32", "2604:9e40::22", "2605:9e40:: - 2605:9e40:ffff:ffff:ffff:ffff:ffff:ffff", false},
		{"2620:6f:2000::/48", "2620:6f:2000::3:4", "2620:6f:2000:: - 2620:6f:2000:ffff:ffff:ffff:ffff:ffff", true},
		{"2620:95:a000::/48", "2619:95:a000::1", "2620:95:a000:: - 2620:95:a000:ffff:ffff:ffff:ffff:ffff", false},
		{"2a00:11d8::/32", "2a00:11d8:3:4:5:6:7:8", "2a00:11d8:: - 2a00:11d8:ffff:ffff:ffff:ffff:ffff:ffff", true},
		{"2a02:1e0::/29", "2a03:1e0::1", "2a02:1e0:: - 2a02:1e7:ffff:ffff:ffff:ffff:ffff:ffff", false},
		{"2a03:5f80::/32", "2a03:5f80::3e3", "2a03:5f80:: - 2a03:5f80:ffff:ffff:ffff:ffff:ffff:ffff", true},
	}
)

func TestIPv6(t *testing.T) {
	for _, ipv6 := range ipv6Lst {
		ret := ipsearch.IPv6StrToInt(ipv6)
		assert.Equal(t, ipv6, ipsearch.IPv6IntToStr(ret))
	}
}

func TestGetIPv6Segment(t *testing.T) {
	for _, ipv6 := range ipv6Lst {
		parts := strings.Split(ipv6, ":")
		segment := 1
		ret, _ := strconv.ParseUint(parts[segment-1], 16, 16)
		assert.Equal(t, uint16(ret), ipsearch.GetIPv6Segment(ipv6, segment))

		segment = 3
		ret, _ = strconv.ParseUint(parts[segment-1], 16, 16)
		assert.Equal(t, uint16(ret), ipsearch.GetIPv6Segment(ipv6, segment))
	}
}

func TestIPv6CIDR(t *testing.T) {
	for _, tt := range testcasesV6 {
		assert.Equal(t, ipsearch.IPv6InCIDR(tt.ipStr, tt.cidr), tt.flag)
	}
}
