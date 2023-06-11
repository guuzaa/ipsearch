# IP Search

Fork of [haoel/ipsearch](https://github.com/haoel/ipsearch) with IPv6 support.

A simple library to search for IP(v4 & v6) addresses for one IP database: IP CIDR List.

## 1. IP Database

The library use the cidr list of [herrbischoff/country-ip-blocks](https://github.com/herrbischoff/country-ip-blocks).

## 2. Usage

### 2.1 Check an IP address is in the CIDR list

```go
package main

import (
	"fmt"
	"github.com/guuzaa/ipsearch"
)

// ipv4 version
func main() {
	search, err := ipsearch.NewIPv4SearchWithFile("./data/ipv4/cn.cidr")
	if err != nil {
		panic(err)
	}

	ipStr := "114.114.114.114"
	ip := search.Search(ipStr)
	if ip != nil {
		fmt.Printf("IP [%s] is in China\n", ipStr)
	} else {
		fmt.Printf("IP [%s] is not China\n", ipStr)
	}
}

// ipv6 version
func main() {
	search, err := ipsearch.NewIPv6SearchWithFile("./data/ipv6/cn.cidr")
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

```

### 2.2 Check an IP address is in the country

```go
package main

import (
	"fmt"
	"github.com/guuzaa/ipsearch"
)

// ipv4 version
func main() {
	search, err := ipsearch.NewIPv4SearchWithCountry("cn")
	if err != nil {
		panic(err)
	}

	ipStr := "114.114.114.114"
	ip := search.Search(ipStr)
	if ip != nil {
		fmt.Printf("IP [%s] is in China\n", ipStr)
	} else {
		fmt.Printf("IP [%s] is not China\n", ipStr)
	}
}

// ipv6 version
func main() {
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
```

## 3. Implementation

As haoel did, the ipsearch package is using the Hash Table and Binary Search algorithm.
Apart from that, we use the net/ip package to parse IP address.

## 4. Licence

This project is MIT licensed.
See the [LICENSE](LICENSE) file for details.