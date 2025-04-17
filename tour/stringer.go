package main

import "fmt"

type IPAddr [4]byte

func (ipAddr IPAddr) String() string {
	var ret string
	for i, val := range ipAddr {
		ret += fmt.Sprintf("%d", val)
		if i < len(ipAddr)-1 {
			ret += "."
		}
	}
	return ret
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
