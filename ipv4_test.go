package ipsearch_test

import (
	"encoding/binary"
	"fmt"
	"github.com/guuzaa/ipsearch"
	"github.com/stretchr/testify/assert"
	"net"
	"testing"
)

const (
	ipFmt = "%d.%d.%d.%d"
)

var (
	ipLst = []string{
		"1.0.1.0",
		"1.0.2.0",
		"1.0.8.0",
		"1.0.32.0",
		"1.1.0.0",
		"1.1.2.0",
		"1.1.4.0",
		"1.1.8.0",
		"1.1.16.0",
		"1.1.32.0",
		"1.2.0.0",
		"1.2.4.0",
		"1.2.8.0",
		"1.2.16.0",
		"1.2.32.0",
		"1.2.64.0",
		"1.3.0.0",
		"1.4.1.0",
		"1.4.2.0",
		"1.4.4.0",
	}
)

func intToIPv4(ipInt uint32) string {
	ip := make(net.IP, 4)
	binary.BigEndian.PutUint32(ip, ipInt)
	return ip.String()
}

func TestIPv4(t *testing.T) {
	ipStr := "192.168.1.100"
	ip := ipsearch.IPv4StrToInt(ipStr)
	assert.Equal(t, ip, ipv4StrToInt1(ipStr))
	assert.Equal(t, ipsearch.IPv4IntToStr(ip), intToIPv4(ip))

	ipStr = "172.20.1.1"
	ip = ipsearch.IPv4StrToInt(ipStr)
	assert.Equal(t, ip, ipv4StrToInt1(ipStr))
	assert.Equal(t, ipsearch.IPv4IntToStr(ip), intToIPv4(ip))

	ipStr = "1.1.1.2"
	ip = ipsearch.IPv4StrToInt(ipStr)
	assert.Equal(t, ip, ipv4StrToInt1(ipStr))
	assert.Equal(t, ipsearch.IPv4IntToStr(ip), intToIPv4(ip))
}

func TestIPv4CIDR(t *testing.T) {
	ipCIDR := "192.168.1.0/24"
	start, end := ipsearch.IPv4CIDRRange(ipCIDR)

	assert.Equal(t, start, ipv4StrToInt1("192.168.1.0"))
	assert.Equal(t, end, ipv4StrToInt1("192.168.1.255"))

	ip := "192.168.1.1"
	assert.True(t, ipsearch.IPv4InCIDR(ip, ipCIDR))
	assert.False(t, ipsearch.IPv4InCIDR("192.168.10.10", ipCIDR))

	ipCIDR = "1.1.1.1"
	start, end = ipsearch.IPv4CIDRRange(ipCIDR)
	assert.Equal(t, start, ipv4StrToInt1(ipCIDR))
	assert.Equal(t, end, ipv4StrToInt1(ipCIDR))
}

func TestIPv4Segment(t *testing.T) {
	ip := "192.168.1.1"
	assert.Equal(t, ipsearch.GetIPv4Segment(ip, 1), uint8(192))

	ip = "172.20.1.1"
	assert.Equal(t, ipsearch.GetIPv4Segment(ip, 1), uint8(172))

	ip = "0.0.0.0"
	assert.Equal(t, ipsearch.GetIPv4Segment(ip, 3), uint8(0))

	ip = "255.255.255.0"
	assert.Equal(t, ipsearch.GetIPv4Segment(ip, 1), uint8(255))

	ip = "254.255.255.0"
	assert.Equal(t, ipsearch.GetIPv4Segment(ip, 1), uint8(254))

	// todo
	// ip = "256.255.255.0"
	// assert.Equal(t, ipis.GetIPv4Segment(ip, 1), uint8(0))
}

func ipv4StrToInt1(ipStr string) uint32 {
	ip := net.ParseIP(ipStr)
	return binary.BigEndian.Uint32(ip.To4())
}

func ipv4StrToInt2(ipStr string) uint32 {
	var (
		ip1, ip2, ip3, ip4 uint32
	)

	fmt.Sscanf(ipStr, ipFmt, &ip1, &ip2, &ip3, &ip4)
	return ip1<<24 | ip2<<16 | ip3<<8 | ip4
}

func BenchmarkIPv4StrToInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, ipStr := range ipLst {
			ipsearch.IPv4StrToInt(ipStr)
		}
	}
}

func BenchmarkIPv4StrToInt1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, ipStr := range ipLst {
			ipv4StrToInt1(ipStr)
		}
	}
}

func BenchmarkIPv4StrToInt2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, ipStr := range ipLst {
			ipv4StrToInt2(ipStr)
		}
	}
}
