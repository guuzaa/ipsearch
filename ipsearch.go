package ipsearch

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	easy "github.com/t-tomalak/logrus-easy-formatter"
	"os"
	"strings"
)

func init() {
	log.SetFormatter(&easy.Formatter{
		TimestampFormat: "2006-01-02 15:04:05",
		LogFormat:       "[%lvl%]: %time% - %msg%\n",
	})
}

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

// IPv6Search is a struct that contains a map of IPv6 ranges.
type IPv6Search struct {
	container IPv6RangeMapList
}

// NewIPv6Search creates a new IPv6Search struct.
func NewIPv6Search(lines []string) *IPv6Search {
	m := NewIPv6RangeMapList()
	ipRanges := NewIPv6RangeSlice(lines)
	m.AppendBatch(ipRanges)
	m.Sort()
	return &IPv6Search{m}
}

// NewIPv6SearchWithCountry creates a new IPv6Search struct from a country code.
func NewIPv6SearchWithCountry(country string) (*IPv6Search, error) {
	path := fmt.Sprintf("./data/ipv6/%s.cidr", strings.ToLower(country))
	if _, err := os.Stat(path); err != nil {
		return nil, fmt.Errorf("%s doesn't exist, check it carefully", country)
	}

	return NewIPv6SearchWithFile(path)
}

// NewIPv6SearchWithFile creates a new IPv6Search struct from a file.
func NewIPv6SearchWithFile(path string) (*IPv6Search, error) {
	lines, err := ReadFile(path)
	if err != nil {
		return nil, err
	}
	return NewIPv6Search(lines), nil
}

// NewIPv6SearchWithFileFromURL creates a new IPv6Search struct from a URL.
func NewIPv6SearchWithFileFromURL(url string) (*IPv6Search, error) {
	lines, err := ReadFileFromURL(url)
	if err != nil {
		return nil, err
	}
	return NewIPv6Search(lines), nil
}

// Search searches if an IPv6 address is in the map of lists of IPv6 ranges.
func (s *IPv6Search) Search(ip string) *IPv6Range {
	return s.container.Search(ip)
}
