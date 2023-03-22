package ip

import (
	"errors"
	"strconv"
	"strings"
)

type Uint32 uint32

func (u32 Uint32) ToIPString() string {
	octets := make([]string, 4)
	for i := 0; i < 4; i++ {
		octets[3-i] = strconv.Itoa(int(u32 & 0xff))
		u32 >>= 8
	}
	return strings.Join(octets, ".")

}

type IPUtils interface {
	ParseIP(ipStr string) (uint32, error)
	ParseMask(maskStr string) (uint32, error)
	GetNetworkAddress(ip, mask uint32, maskLen int) uint32
}

type ipUtils struct{}

func NewIPUtils() IPUtils {
	return &ipUtils{}
}

func (i *ipUtils) ParseIP(ipStr string) (uint32, error) {
	ip := uint32(0)
	parts := strings.Split(ipStr, ".")
	if len(parts) != 4 {
		return 0, errors.New("无效的IP地址")
	}
	for j := 0; j < 4; j++ {
		part, err := strconv.Atoi(parts[j])
		if err != nil || part < 0 || part > 255 {
			return 0, errors.New("无效的IP地址")
		}
		ip = (ip << 8) | uint32(part)
	}
	return ip, nil
}

func (i *ipUtils) ParseMask(maskStr string) (uint32, error) {
	mask := uint32(0)
	parts := strings.Split(maskStr, ".")
	if len(parts) != 4 {
		return 0, errors.New("无效的子网掩码")
	}
	for j := 0; j < 4; j++ {
		part, err := strconv.Atoi(parts[j])
		if err != nil || part < 0 || part > 255 {
			return 0, errors.New("无效的子网掩码")
		}
		mask = (mask << 8) | uint32(part)
	}
	return mask, nil
}

func (i *ipUtils) GetNetworkAddress(ip, mask uint32, maskLen int) uint32 {
	network := ip & mask
	if maskLen < 32 {
		network &= i.GetMask(maskLen)
	}
	return network
}

func (i *ipUtils) GetMask(maskLen int) uint32 {
	mask := uint32(0)
	for j := 0; j < maskLen; j++ {
		mask |= (1 << uint(31-j))
	}
	return mask
}
