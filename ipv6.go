package ipsearch

import (
	num "github.com/shabbyrobe/go-num"
	"net"
	"strconv"
)

// IPv6StrToInt converts a string IPv6 address to an integer
func IPv6StrToInt(ipStr string) num.U128 {
	ip := net.ParseIP(ipStr).To16()
	if ip == nil || len(ip) != 16 {
		return num.U128{}
	}
	return ipv6ToInt(ip)
}

func ipv6ToInt(ip net.IP) num.U128 {
	step := uint(120)
	ret := num.U128{}
	for _, hex := range ip {
		ret = ret.Or(num.U128From8(hex).Lsh(step))
		step -= 8
	}
	return ret
}

// IPv6IntToStr converts an integer to a string IPv6 address
func IPv6IntToStr(ip num.U128) string {
	bytes := [16]byte{}
	step := uint(120)
	for i := range bytes {
		ele, _ := strconv.ParseUint(ip.Rsh(step).And64(0xff).String(), 10, 8)
		bytes[i] = byte(ele)
		step -= 8
	}

	return net.IP(bytes[:]).String()
}

// IPv6CIDRRange returns the start and end IPv6 address for a CIDR range
func IPv6CIDRRange(cidr string) (num.U128, num.U128) {
	var mask uint
	ip, ipNet, _ := net.ParseCIDR(cidr)
	if ipNet == nil {
		ip = net.ParseIP(cidr)
	} else {
		ones, _ := ipNet.Mask.Size()
		mask = uint(ones)
	}
	if mask == 0 || mask > 128 {
		mask = 128
	}

	start := ipv6ToInt(ip.To16())
	end := start.Or(num.U128From8(1).Lsh(128 - mask).Sub64(1))
	return start, end
}

// ipv6InRange checks if an IPv6 address in a range
func ipv6InRange(ip num.U128, start num.U128, end num.U128) bool {
	return ip.GreaterThan(start) && ip.LessThan(end)
}

// IPv6InCIDR checks if an IPv6 address is in a CIDR range
func IPv6InCIDR(ipStr string, cidr string) bool {
	start, end := IPv6CIDRRange(cidr)
	ip := IPv6StrToInt(ipStr)
	return ipv6InRange(ip, start, end)
}

// GetIPv6Segment returns the segment of an IPv6 address
func GetIPv6Segment(ipStr string, segment int) uint16 {
	ip := net.ParseIP(ipStr)
	if ip == nil {
		ip, _, _ = net.ParseCIDR(ipStr)
	}
	if ip == nil || len(ip) != 16 {
		return 0
	}

	ip = ip.To16()
	segment = 2 * (segment - 1)
	return uint16(ip[segment])<<8 | uint16(ip[segment+1])
}
