package ipsearch

import (
	log "github.com/sirupsen/logrus"
	"sort"
	"strings"
)

// IPv6RangeList is a list of IPv6 CIDR ranges.
type IPv6RangeList []*IPv6Range

func NewIPv6RangeList(lines []string) IPv6RangeList {
	lst := make(IPv6RangeList, 0, len(lines))
	for _, line := range lines {
		lst.Append(NewIPv6Range(line))
	}

	if log.GetLevel() == log.DebugLevel {
		for _, ipCIDR := range lst {
			log.Debugf("Ranges: %s", ipCIDR.Range())
		}
	}

	return lst
}

// Append adds a IPv6 CIDR to the list
func (lst *IPv6RangeList) Append(rge *IPv6Range) {
	*lst = append(*lst, rge)
}

// InsertSorted inserts a IPv6 range to the list, keeping the list sorted
func (lst *IPv6RangeList) InsertSorted(rge *IPv6Range) {
	start := 0
	end := len(*lst) - 1
	for start <= end {
		mid := (start + end) / 2
		if (*lst)[mid].start.LessThan(rge.start) {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}

	idx := start
	*lst = append(*lst, nil)
	copy((*lst)[idx+1:], (*lst)[idx:])
	(*lst)[idx] = rge
}

// Sort sorts the list of IPv6 CIDR ranges
func (lst *IPv6RangeList) Sort() {
	sort.Slice(*lst, func(i, j int) bool {
		return (*lst)[i].start.LessThan((*lst)[j].start)
	})
}

// Len returns the length of the list of IPv6 CIDR ranges
func (lst *IPv6RangeList) Len() int {
	return len(*lst)
}

// String returns a string representation of the list of IPv6 CIDR ranges
func (lst *IPv6RangeList) String() string {
	var str strings.Builder
	for _, cidr := range *lst {
		str.WriteString(cidr.String())
		str.WriteString("\n")
	}
	return str.String()
}

// Range returns a string representation of the list of IPv6 address ranges
func (lst *IPv6RangeList) Range() string {
	var str strings.Builder
	for _, cidr := range *lst {
		str.WriteString(cidr.Range())
		str.WriteString("\n")
	}
	return str.String()
}

// Search searches if an IPv6 address in the list
func (lst *IPv6RangeList) Search(ipStr string) *IPv6Range {
	ip := IPv6StrToInt(ipStr)

	start := 0
	end := len(*lst) - 1
	for start <= end {
		mid := (start + end) / 2
		if ipv6InRange(ip, (*lst)[mid].start, (*lst)[mid].end) {
			log.Debugf("IP %s in Range %s", ipStr, (*lst)[mid].Range())
			return (*lst)[mid]
		}

		if ip.LessThan((*lst)[mid].start) {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	log.Debugf("IP %s is not in any following Ranges:\n%s", ipStr, lst)
	return nil
}
