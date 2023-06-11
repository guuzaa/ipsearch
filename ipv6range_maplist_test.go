package ipsearch_test

import (
	"github.com/guuzaa/ipsearch"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIPv6CIDRMapList(t *testing.T) {
	ipv6MapList := ipsearch.NewIPv6RangeMapList()
	ipv6RangeList := ipsearch.NewIPv6RangeList(ipv6CIDRs)
	ipv6MapList.AppendBatch(ipv6RangeList)

	ip := ipv6MapList.Search("2620:105:a000::1")
	assert.Nil(t, ip)

	ipv6MapList.Sort()
	ip = ipv6MapList.Search("2620:105:a000::1")
	assert.NotNil(t, ip)
	assert.Equal(t, ip.String(), "2620:105:a000::/48")

	ipv6MapList = ipsearch.NewIPv6RangeMapList()
	ipv6RangeList = ipsearch.NewIPv6RangeList(ipv6CIDRs)
	ipv6MapList.InsertSortedCIDRs(ipv6RangeList)
	for _, data := range testcasesV6 {
		ip := ipv6MapList.Search(data.ipStr)
		assert.Equal(t, ip != nil, data.flag)
		if ip != nil {
			assert.Equal(t, ip.String(), data.cidr)
		}
	}
}
