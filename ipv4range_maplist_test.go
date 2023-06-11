package ipsearch_test

import (
	"github.com/guuzaa/ipsearch"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIPv4CIDRMapList(t *testing.T) {
	ipCIDRMapList := ipsearch.NewIPv4RangeMapList()
	ipRangeList := ipsearch.NewIPv4RangeList(cidrs)
	ipCIDRMapList.AppendBatch(ipRangeList)

	// without sorting, it would be a bug
	ip := ipCIDRMapList.Search("1.4.1.2")
	assert.Nil(t, ip)

	// with sorting, it would be ok
	ipCIDRMapList.Sort()
	ip = ipCIDRMapList.Search("1.4.1.3")
	assert.NotNil(t, ip)
	assert.Equal(t, ip.String(), "1.4.1.0/24")

	ipCIDRMapList = ipsearch.NewIPv4RangeMapList()
	ipRangeList = ipsearch.NewIPv4RangeList(cidrs)
	ipCIDRMapList.InsertSortedCIDRs(ipRangeList)

	for _, data := range testCIDRDataList {
		ip := ipCIDRMapList.Search(data.ip)
		assert.Equal(t, ip != nil, data.find)
		if ip != nil {
			assert.Equal(t, ip.String(), data.cidr)
		}
	}

	ipCIDRMapList.InsertSorted(ipsearch.NewIPv4Range("1.3.0.0/16"))
	ip = ipCIDRMapList.Search("1.3.1.1")
	assert.NotNil(t, ip)
	assert.Equal(t, ip.String(), "1.3.0.0/16")
}
