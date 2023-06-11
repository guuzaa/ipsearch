package ipsearch_test

import (
	"github.com/guuzaa/ipsearch"
	"github.com/stretchr/testify/assert"
	"testing"
)

// todo
const v6CIDRs1 = `2001:b28::/32
2605:9e40::/32
2620:6f:2000::/48
2620:85:a000::/48
2620:95:a000::/48
2620:105:a000::/48
`

const v6CIDRs2 = `2a00:11d8::/32
2a02:1e0::/29
2a03:5f80::/32
`

const expectedV6 = v6CIDRs1 + v6CIDRs2

func TestIPv6CIDRList(t *testing.T) {
	ipCIDRList := ipsearch.NewIPv6RangeList(ipv6CIDRs)
	assert.Equal(t, ipCIDRList.Len(), len(ipv6CIDRs))

	ipCIDRList.Sort()
	assert.Equal(t, ipCIDRList.String(), expectedV6)

	newIP := "2840:95::ae"
	ipCIDRList.InsertSorted(ipsearch.NewIPv6Range(newIP))

	newExpected := v6CIDRs1 + newIP + "\n" + v6CIDRs2
	assert.Equal(t, ipCIDRList.String(), newExpected)

	for _, data := range testcasesV6 {
		ip := ipCIDRList.Search(data.ipStr)
		assert.Equal(t, ip != nil, data.flag)
		if ip != nil {
			assert.Equal(t, ip.String(), data.cidr)
			assert.Equal(t, ip.Range(), data.rangeStr)
		}
	}
}
