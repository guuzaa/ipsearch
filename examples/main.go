package main

import (
	"fmt"
	"github.com/guuzaa/ipsearch"
)

func checkChinaIPv4() {
	search, err := ipsearch.NewIPv4SearchWithFile("./data/ipv4/cn.cidr")
	if err != nil {
		panic(err)
	}

	ipStr := "114.114.114.114"
	ip := search.Search(ipStr)
	if ip != nil {
		fmt.Printf("IP [%s] is in China\n", ipStr)
	} else {
		fmt.Printf("IP [%s] is not in China\n", ipStr)
	}
}

func checkChinaIPv6() {
	search, err := ipsearch.NewIPv6SearchWithCountry("cn")
	if err != nil {
		panic(err)
	}

	ipStr := "2001:da8::2"
	ip := search.Search(ipStr)
	if ip != nil {
		fmt.Printf("IP [%s] is in China\n", ipStr)
	} else {
		fmt.Printf("IP [%s] is not in China\n", ipStr)
	}
}

func checkUSIPv4() {
	search, err := ipsearch.NewIPv4SearchWithCountry("us")
	if err != nil {
		panic(err)
	}

	ipStr := "8.8.8.8"
	ip := search.Search(ipStr)
	if ip != nil {
		fmt.Printf("IP [%s] is in US\n", ipStr)
	} else {
		fmt.Printf("IP [%s] is not in US\n", ipStr)
	}

	ipStr = "1.1.1.1"
	ip = search.Search(ipStr)
	if ip != nil {
		fmt.Printf("IP [%s] is in US\n", ipStr)
	} else {
		fmt.Printf("IP [%s] is not in US\n", ipStr)
	}
}

func main() {
	checkChinaIPv4()
	checkChinaIPv6()
	checkUSIPv4()
}
