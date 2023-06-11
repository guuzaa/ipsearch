package ipsearch

import "github.com/shabbyrobe/go-num"

// IPv6Range represents an IPv6 CIDR range.
type IPv6Range struct {
	firstSeg   uint16
	start, end num.U128
	cidr       string
}

// NewIPv6Range creates a new IPv6 CIDR range
func NewIPv6Range(cidr string) *IPv6Range {
	start, end := IPv6CIDRRange(cidr)
	return &IPv6Range{
		firstSeg: GetIPv6Segment(cidr, 1),
		start:    start,
		end:      end,
		cidr:     cidr,
	}
}

func NewIPv6RangeSlice(cidrs []string) []*IPv6Range {
	ranges := make([]*IPv6Range, len(cidrs))
	for i, cidr := range cidrs {
		ranges[i] = NewIPv6Range(cidr)
	}
	return ranges
}

func (ip *IPv6Range) String() string {
	return ip.cidr
}

// Range returns the range of the IPv6 address in string format
func (ip *IPv6Range) Range() string {
	return IPv6IntToStr(ip.start) + " - " + IPv6IntToStr(ip.end)
}

// CIDR returns the CIDR of the IP range
func (ip *IPv6Range) CIDR() string {
	return ip.cidr
}
