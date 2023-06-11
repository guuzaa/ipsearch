package ipsearch_test

import (
	"github.com/guuzaa/ipsearch"
	"github.com/stretchr/testify/assert"
	"net/http"
	"strings"
	"testing"
	"time"
)

const (
	TestDir      = "data/"
	IPv4CIDRFile = "data/ipv4/cn.cidr"
	IPv6CIDRFile = "data/ipv6/kg.cidr"
)

func TestLoadFromFile(t *testing.T) {
	search, err := ipsearch.NewIPv4SearchWithFile("not-exist-file")
	assert.NotNil(t, err)

	search, err = ipsearch.NewIPv4SearchWithFile(IPv4CIDRFile)
	assert.Nil(t, err)
	testCIDRSearch(t, search)

	searchV6, err := ipsearch.NewIPv6SearchWithFile("not-exist-file")
	assert.NotNil(t, err)

	searchV6, err = ipsearch.NewIPv6SearchWithFile(IPv6CIDRFile)
	assert.Nil(t, err)
	testV6CIDRSearch(t, searchV6)
}

func TestLoadWithCountry(t *testing.T) {
	// IPv4
	search, err := ipsearch.NewIPv4SearchWithCountry("cn")
	assert.Nil(t, err)
	testCIDRSearch(t, search)

	search, err = ipsearch.NewIPv4SearchWithCountry("CN")
	assert.Nil(t, err)
	testCIDRSearch(t, search)

	search, err = ipsearch.NewIPv4SearchWithCountry("not-exist")
	assert.NotNil(t, err)

	// IPv6
	searchV6, err := ipsearch.NewIPv6SearchWithCountry("kg")
	assert.Nil(t, err)
	testV6CIDRSearch(t, searchV6)

	searchV6, err = ipsearch.NewIPv6SearchWithCountry("KG")
	assert.Nil(t, err)
	testV6CIDRSearch(t, searchV6)

	searchV6, err = ipsearch.NewIPv6SearchWithCountry("not-exist")
	assert.NotNil(t, err)
}

func TestLoadFromURL(t *testing.T) {
	endpoint := "127.0.0.1:9898"
	protocol := "http://" + endpoint + "/"
	search, err := ipsearch.NewIPv4SearchWithFileFromURL(protocol + "not-exist-file")
	assert.NotNil(t, err)

	// start a http server
	go func() {
		http.ListenAndServe(endpoint, http.FileServer(http.Dir(TestDir)))
	}()

	for {
		search, err = ipsearch.NewIPv4SearchWithFileFromURL(protocol + "ipv4/cn.cidr")
		if err != nil && strings.Contains(err.Error(), "connection refused") {
			time.Sleep(100 * time.Millisecond)
			continue
		}

		assert.Nil(t, err)
		testCIDRSearch(t, search)
		break
	}

	for {
		searchV6, err := ipsearch.NewIPv6SearchWithFileFromURL(protocol + "ipv6/kg.cidr")
		if err != nil && strings.Contains(err.Error(), "connection refused") {
			time.Sleep(100 * time.Millisecond)
			continue
		}

		assert.Nil(t, err)
		testV6CIDRSearch(t, searchV6)
		break
	}
}

func testCIDRSearch(t *testing.T, search *ipsearch.IPv4Search) {
	type testCIDRData struct {
		ip   string
		cidr string
		find bool
	}
	var testCIDRDataList = []testCIDRData{
		{"1.0.1.24", "1.0.1.0/24", true},
		{"8.8.8.8", "", false},
		{"1.1.1.1", "", false},
	}

	for _, data := range testCIDRDataList {
		ip := search.Search(data.ip)
		assert.Equal(t, ip != nil, data.find)
		if ip != nil {
			assert.Equal(t, ip.CIDR(), data.cidr)
		}
	}
}

func testV6CIDRSearch(t *testing.T, search *ipsearch.IPv6Search) {
	type testCIDRData struct {
		ip   string
		cidr string
		find bool
	}
	var testCIDRDataList = []testCIDRData{
		{"2a00:5700::44", "2a00:5700::/32", true},
		{"2001:da8::33", "", false},
		{"2001:da8:3::1", "", false},
	}

	for _, data := range testCIDRDataList {
		ip := search.Search(data.ip)
		assert.Equal(t, ip != nil, data.find)
		if ip != nil {
			assert.Equal(t, ip.CIDR(), data.cidr)
		}
	}
}
