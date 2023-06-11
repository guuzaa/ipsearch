package ipsearch

import (
	log "github.com/sirupsen/logrus"
	"sort"
	"strings"
)

// IPv4RangeList is a list of IPv4 CIDR ranges.
type IPv4RangeList []*IPv4Range

func NewIPv4RangeList(lines []string) IPv4RangeList {
	list := make(IPv4RangeList, 0, len(lines))
	for _, line := range lines {
		list.Append(NewIPv4Range(line))
	}

	if log.GetLevel() == log.DebugLevel {
		for _, ipCIDR := range list {
			log.Debugf("Ranges: %s", ipCIDR.Range())
		}
	}

	return list
}

// Append adds a IPv4 CIDR to the list
func (list *IPv4RangeList) Append(rge *IPv4Range) {
	*list = append(*list, rge)
}

// InsertSorted inserts a IPv4 range to the list, keeping the list sorted
func (list *IPv4RangeList) InsertSorted(rge *IPv4Range) {
	start := 0
	end := len(*list) - 1
	for start <= end {
		mid := (start + end) / 2
		if (*list)[mid].start < rge.start {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}

	idx := start
	*list = append(*list, nil)
	copy((*list)[idx+1:], (*list)[idx:])
	(*list)[idx] = rge
}

// Sort sorts the list of IPv4 CIDR ranges
func (list *IPv4RangeList) Sort() {
	sort.Slice(*list, func(i, j int) bool {
		return (*list)[i].start < (*list)[j].start
	})
}

// Len returns the length of the list of IPv4 CIDR ranges
func (list *IPv4RangeList) Len() int {
	return len(*list)
}

// String returns a string representation of the list of IPv4 CIDR ranges
func (list *IPv4RangeList) String() string {
	var str strings.Builder
	for _, cidr := range *list {
		str.WriteString(cidr.String())
		str.WriteString("\n")
	}
	return str.String()
}

// Range returns a string representation of the list of IPv4 address ranges
func (list *IPv4RangeList) Range() string {
	var str strings.Builder
	for _, cidr := range *list {
		str.WriteString(cidr.Range())
		str.WriteString("\n")
	}
	return str.String()
}

// Search searches if an IP address in the list
func (list *IPv4RangeList) Search(ipStr string) *IPv4Range {
	ip := IPv4StrToInt(ipStr)

	start := 0
	end := len(*list) - 1
	for start <= end {
		mid := (start + end) / 2
		if ipInRange(ip, (*list)[mid].start, (*list)[mid].end) {
			log.Debugf("IP %s in Range %s", ipStr, (*list)[mid].Range())
			return (*list)[mid]
		}

		if ip < (*list)[mid].start {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}

	log.Debugf("IP %s is not in any following Ranges.\n%s", ipStr, list)
	return nil
}
