package ipsearch

// IPv4Range represents an IPv4 CIDR range.
type IPv4Range struct {
	firstSeg   uint8
	start, end uint32
	cidr       string
}

// NewIPv4Range creates a new IPv4 CIDR range
func NewIPv4Range(cidr string) *IPv4Range {
	start, end := IPv4CIDRRange(cidr)
	return &IPv4Range{
		firstSeg: GetIPv4Segment(cidr, 1),
		start:    start,
		end:      end,
		cidr:     cidr,
	}
}

// NewIPv4RangeSlice creates a new slice of IPv4Range.
func NewIPv4RangeSlice(lines []string) []*IPv4Range {
	ranges := make([]*IPv4Range, len(lines))
	for i, line := range lines {
		ranges[i] = NewIPv4Range(line)
	}
	return ranges
}

func (ip *IPv4Range) String() string {
	return ip.cidr
}

// Range returns the range of the IPv4 address in string format
func (ip *IPv4Range) Range() string {
	return IPv4IntToStr(ip.start) + " - " + IPv4IntToStr(ip.end)
}

// CIDR returns the CIDR of the IP range
func (ip *IPv4Range) CIDR() string {
	return ip.cidr
}
