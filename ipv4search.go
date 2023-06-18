package ipsearch

import (
	"fmt"
	"os"
	"strings"
)

// IPv4Search is a struct that contains a map of IPv4 ranges.
type IPv4Search struct {
	container IPv4RangeMapList
}

// NewIPv4Search creates a new IPv4Search struct.
func NewIPv4Search(lines []string) *IPv4Search {
	m := NewIPv4RangeMapList()
	ipRanges := NewIPv4RangeSlice(lines)
	m.AppendBatch(ipRanges)
	m.Sort()
	return &IPv4Search{m}
}

// NewIPv4SearchWithCountry creates a new IPv4Search struct from a country code.
func NewIPv4SearchWithCountry(country string) (*IPv4Search, error) {
	path := fmt.Sprintf("./data/ipv4/%s.cidr", strings.ToLower(country))
	if _, err := os.Stat(path); err != nil {
		return nil, fmt.Errorf("%s doesn't exist, check it carefully", country)
	}

	return NewIPv4SearchWithFile(path)
}

// NewIPv4SearchWithFile creates a new IPv4Search struct from a file.
func NewIPv4SearchWithFile(path string) (*IPv4Search, error) {
	lines, err := ReadFile(path)
	if err != nil {
		return nil, err
	}
	return NewIPv4Search(lines), nil
}

// NewIPv4SearchWithFileFromURL creates a new IPv4Search struct from a URL.
func NewIPv4SearchWithFileFromURL(url string) (*IPv4Search, error) {
	lines, err := ReadFileFromURL(url)
	if err != nil {
		return nil, err
	}
	return NewIPv4Search(lines), nil
}

// Search searches if an IPv4 address is in the map of lists of IPv4 ranges.
func (s *IPv4Search) Search(ip string) *IPv4Range {
	return s.container.Search(ip)
}
