package ipsearch

import (
	"net"
)

// IPv4StrToInt converts a string IPv4 address to an integer
func IPv4StrToInt(ipStr string) uint32 {
	ip := net.ParseIP(ipStr).To4()
	if ip == nil {
		return 0
	}
	return ipv4ToInt(ip)
}

// IPv4IntToStr converts an integer to a string IPv4 address
func IPv4IntToStr(ip uint32) string {
	return net.IPv4(
		byte(ip>>24),
		byte(ip>>16&0xff),
		byte(ip>>8&0xff),
		byte(ip&0xff),
	).String()
}

// IPv4CIDRRange returns the start and end IP v4 address for a CIDR range
func IPv4CIDRRange(cidr string) (uint32, uint32) {
	var mask uint32
	ip, ipNet, _ := net.ParseCIDR(cidr)
	if ipNet == nil {
		ip = net.ParseIP(cidr)
	} else {
		ones, _ := ipNet.Mask.Size()
		mask = uint32(ones)
	}
	if mask == 0 || mask > 32 {
		mask = 32
	}

	start := ipv4ToInt(ip.To4())
	end := start | (1<<(32-mask) - 1)
	return start, end
}

func ipv4ToInt(ip net.IP) uint32 {
	return uint32(ip[0])<<24 | uint32(ip[1])<<16 | uint32(ip[2])<<8 | uint32(ip[3])
}

// ipInRange checks if an IPv4 address in a range
func ipInRange(ip uint32, start uint32, end uint32) bool {
	return ip >= start && ip <= end
}

// IPv4InCIDR checks if an IPv4 address is in a CIDR range
func IPv4InCIDR(ipStr string, cidr string) bool {
	start, end := IPv4CIDRRange(cidr)
	i := IPv4StrToInt(ipStr)
	return ipInRange(i, start, end)
}

// GetIPv4Segment returns the segment of an IPv4 address
func GetIPv4Segment(ipStr string, segment int) uint8 {
	ip := net.ParseIP(ipStr)
	if ip == nil {
		ip, _, _ = net.ParseCIDR(ipStr)
	}
	if ip == nil {
		return 0
	}

	ip = ip.To4()
	return ip[segment-1]
}
