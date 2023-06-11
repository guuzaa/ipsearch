package ipsearch_test

import (
	"github.com/guuzaa/ipsearch"
	"github.com/stretchr/testify/assert"
	"testing"
)

var cidrs = []string{
	"1.4.1.0/24",
	"1.0.1.0/24",
	"1.0.2.0/23",
	"36.0.16.0/20",
	"43.224.242.0/24",
	"59.83.0.0/18",
	"103.196.64.0/22",
	"101.236.0.0/14",
	"45.119.116.0/22",
}

func TestCIDRIPRange(t *testing.T) {
	cidrIPs := ipsearch.NewIPv4RangeSlice(cidrs)
	assert.Equal(t, len(cidrIPs), len(cidrs))
	for i, cidrIP := range cidrIPs {
		assert.Equal(t, cidrs[i], cidrIP.CIDR())
		assert.Equal(t, cidrs[i], cidrIP.String())
	}
}
