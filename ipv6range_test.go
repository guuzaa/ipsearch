package ipsearch_test

import (
	"github.com/guuzaa/ipsearch"
	"github.com/stretchr/testify/assert"
	"testing"
)

var ipv6CIDRs = []string{
	"2a03:5f80::/32",
	"2001:b28::/32",
	"2605:9e40::/32",
	"2620:105:a000::/48",
	"2620:6f:2000::/48",
	"2620:85:a000::/48",
	"2620:95:a000::/48",
	"2a00:11d8::/32",
	"2a02:1e0::/29",
}

func TestIPv6CIDRRange(t *testing.T) {
	ipv6Slice := ipsearch.NewIPv6RangeSlice(ipv6CIDRs)
	assert.Equal(t, len(ipv6CIDRs), len(ipv6Slice))
	for i, r := range ipv6Slice {
		assert.Equal(t, ipv6CIDRs[i], r.CIDR())
		assert.Equal(t, ipv6CIDRs[i], r.String())
		t.Log(r.Range())
	}
}
