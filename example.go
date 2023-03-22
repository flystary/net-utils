package main

import (
	"fmt"
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

	fmt.Printf("IP地址: %s\n", ipStr)
	fmt.Printf("子网掩码: %s\n", maskStr)
	fmt.Printf("网络地址: %s\n", ipUtils.IPString(network))
}
