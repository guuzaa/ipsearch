package ipsearch

// IPv6RangeMapList is a map of lists of IPv6 CIDR ranges
type IPv6RangeMapList map[uint16]*IPv6RangeList

// NewIPv6RangeMapList creates a new map of lists of IPv6 CIDR ranges
func NewIPv6RangeMapList() IPv6RangeMapList {
	m := make(IPv6RangeMapList)
	return m
}

// AppendBatch adds a list of IPv6 CIDR ranges to the map of lists of IPv6 CIDR
func (m IPv6RangeMapList) AppendBatch(ipRanges []*IPv6Range) {
	for _, ip := range ipRanges {
		m.Append(ip)
	}
}

// Append adds an IPv6 CIDR range to the map of lists of IPv6 CIDR ranges
func (m IPv6RangeMapList) Append(ipRange *IPv6Range) {
	ip1 := ipRange.firstSeg
	if _, ok := m[ip1]; !ok {
		m[ip1] = &IPv6RangeList{}
	}
	(*m[ip1]).Append(ipRange)
}

// Sort sorts the map of lists of IPv6 CIDR ranges
func (m IPv6RangeMapList) Sort() {
	for _, lst := range m {
		lst.Sort()
	}
}

func (m IPv6RangeMapList) InsertSorted(ipRange *IPv6Range) {
	ip1 := ipRange.firstSeg
	if _, ok := m[ip1]; !ok {
		m[ip1] = &IPv6RangeList{}
	}
	(*m[ip1]).InsertSorted(ipRange)
}

// InsertSortedCIDRs inserts a list of IPv6 CIDR ranges to the map of IPv6 CIDR lists
// keeping the lists sorted
func (m IPv6RangeMapList) InsertSortedCIDRs(ipRanges []*IPv6Range) {
	for _, ip := range ipRanges {
		m.InsertSorted(ip)
	}
}

// Search searches if an IPv6 address is in the map of lists
func (m IPv6RangeMapList) Search(ipStr string) *IPv6Range {
	ip1 := GetIPv6Segment(ipStr, 1)
	if _, ok := m[ip1]; !ok {
		return nil
	}
	return m[ip1].Search(ipStr)
}
