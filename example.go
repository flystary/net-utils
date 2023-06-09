package main

import (
	"fmt"
	. "github.com/flystary/net-utils/ip" //不使用别名引入包
)

func main() {
	ipStr := "192.168.1.100"
	maskStr := "255.255.255.0"

	ipUtils := NewIPUtils()

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
	var network uint32
	network = ipUtils.GetNetworkAddress(ip, mask, maskLen)

	fmt.Printf("address: %s\n", ipStr)
	fmt.Printf("netmask: %s\n", maskStr)

	var net Uint32 = Uint32(network)

	fmt.Printf("network: %s\n", net.ToIPString())
}
