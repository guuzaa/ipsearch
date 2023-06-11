package ipsearch

// IPv4RangeMapList is a map of lists of IPv4 CIDR ranges
type IPv4RangeMapList map[uint8]*IPv4RangeList

// NewIPv4RangeMapList creates a new map of lists of IPv4 CIDR ranges
func NewIPv4RangeMapList() IPv4RangeMapList {
	m := make(IPv4RangeMapList)
	return m
}

// AppendBatch adds a list of IPv4 CIDR ranges to the map of lists of IPv4 CIDR ranges
func (m IPv4RangeMapList) AppendBatch(ipRanges []*IPv4Range) {
	for _, ip := range ipRanges {
		m.Append(ip)
	}
}

// Append adds an IPv4 CIDR range to the map of lists of IPv4 CIDR ranges
func (m IPv4RangeMapList) Append(ipRange *IPv4Range) {
	ip1 := ipRange.firstSeg
	if _, ok := m[ip1]; !ok {
		m[ip1] = &IPv4RangeList{}
	}
	(*m[ip1]).Append(ipRange)
}

// Sort sorts the map of lists of IPv4 CIDR ranges
func (m IPv4RangeMapList) Sort() {
	for _, lst := range m {
		lst.Sort()
	}
}

// InsertSorted inserts a IPv4 CIDR range to the map of lists of IPv4 CIDR ranges, keeping the lists sorted
func (m IPv4RangeMapList) InsertSorted(ipRange *IPv4Range) {
	ip1 := ipRange.firstSeg
	if _, ok := m[ip1]; !ok {
		m[ip1] = &IPv4RangeList{}
	}
	(*m[ip1]).InsertSorted(ipRange)
}

// InsertSortedCIDRs inserts a list of IPv4 CIDR ranges to the map of IPv4 CIDR ranges lists, keeping the lists sorted
func (m IPv4RangeMapList) InsertSortedCIDRs(ipRanges []*IPv4Range) {
	for _, ip := range ipRanges {
		m.InsertSorted(ip)
	}
}

// Search searches if an IPv4 address is in the map of lists
func (m IPv4RangeMapList) Search(ipStr string) *IPv4Range {
	ip1 := GetIPv4Segment(ipStr, 1)
	if _, ok := m[ip1]; !ok {
		return nil
	}
	return (*m[ip1]).Search(ipStr)
}
