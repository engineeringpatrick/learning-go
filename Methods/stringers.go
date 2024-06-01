package main

import (
	"fmt"
	"strconv"
	"strings"
)

type IPAddr [4]byte

// now IPAddr implicitly implements the Stringer interface
func (ip IPAddr) String() string {
	// return fmt.Sprintf("%d.%d.%d.%d", ip[0], ip[1], ip[2], ip[3])
	ipStr := make([]string, 4)
	for i, val := range ip {
		ipStr[i] = strconv.Itoa(int(val))
	}
	return strings.Join(ipStr, ".")
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
