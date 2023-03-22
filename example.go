package main

import (
	"fmt"
	iputils "github.com/flystary/net-utils/ip"
)

func main() {
	ipStr := "192.168.1.100"
	maskStr := "255.255.255.0"

	ipUtils := iputils.NewIPUtils()

	ip, err := ipUtils.ParseIP(ipStr)
	if err != nil {
		fmt.Println(err)
		return
	}
	mask, err := ipUtils.ParseMask(maskStr)
	if err != nil {
		fmt.Println(err)
		return
	}
	maskLen := 0
	for i := uint32(1); i != 0; i <<= 1 {
		if mask&i != 0 {
			maskLen++
		}
	}
	network := ipUtils.GetNetworkAddress(ip, mask, maskLen)

	fmt.Printf("address: %s\n", ipStr)
	fmt.Printf("netmask: %s\n", maskStr)

	fmt.Printf("network: %s\n", iputils.ToIPString(network))
}
