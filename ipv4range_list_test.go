package ipsearch_test

import (
	"github.com/guuzaa/ipsearch"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"testing"
)

const expectedPart1 = `1.0.1.0/24
1.0.2.0/23
1.4.1.0/24
36.0.16.0/20
43.224.242.0/24
45.119.116.0/22
`

const expectedPart2 = `59.83.0.0/18
101.236.0.0/14
103.196.64.0/22
`

const expected = expectedPart1 + expectedPart2

var testCIDRDataList = []struct {
	ip       string
	cidr     string
	rangeStr string
	find     bool
}{
	{"1.0.1.24", "1.0.1.0/24", "1.0.1.0 - 1.0.1.255", true},
	{"1.4.1.1", "1.4.1.0/24", "1.4.1.0 - 1.4.1.255", true},
	{"43.224.242.100", "43.224.242.0/24", "43.224.242.0 - 43.224.242.255", true},
	{"101.236.0.1", "101.236.0.0/14", "101.236.0.0 - 101.239.255.255", true},
	{"101.240.0.1", "", "", false},
	{"5.5.5.5", "", "", false},
}

func TestIPv4CIDRList(t *testing.T) {
	log.SetLevel(log.DebugLevel)

	ipCIDRList := ipsearch.NewIPv4RangeList(cidrs)
	assert.Equal(t, ipCIDRList.Len(), len(cidrs))

	ipCIDRList.Sort()
	assert.Equal(t, ipCIDRList.String(), expected)

	newIP := "58.83.0.0/16"
	ipCIDRList.InsertSorted(ipsearch.NewIPv4Range(newIP))

	newExpected := expectedPart1 + newIP + "\n" + expectedPart2
	assert.Equal(t, ipCIDRList.String(), newExpected)

	for _, data := range testCIDRDataList {
		ip := ipCIDRList.Search(data.ip)
		assert.Equal(t, ip != nil, data.find)
		if ip != nil {
			assert.Equal(t, ip.String(), data.cidr)
			assert.Equal(t, ip.Range(), data.rangeStr)
		}
	}
}
