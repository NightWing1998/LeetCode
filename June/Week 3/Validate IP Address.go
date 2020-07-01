// Link - https://leetcode.com/problems/validate-ip-address/

import (
	"strconv"
	"strings"
)

func validateIPv4(address []string) bool {
	for _, ad := range address {
		if len(ad) == 0 {
			return false
		}
		if intAd, err := strconv.Atoi(ad); err != nil || intAd >= 256 || (intAd < 10 && len(ad) > 1) || (intAd < 100 && len(ad) > 2) {
			return false
		}
	}
	return true
}

func validateIPv6(address []string) bool {
	for _, ad := range address {
		if len(ad) > 4 {
			return false
		}
		if _, err := strconv.ParseUint(ad, 16, 16); err != nil {
			return false
		}
	}
	return true
}

func validIPAddress(IP string) string {
	addressBytes := strings.Split(IP, ".")
	switch len(addressBytes) {
	case 4:
		if validateIPv4(addressBytes) {
			return "IPv4"
		}
	case 1:
		addressBytes = strings.Split(addressBytes[0], ":")
		if len(addressBytes) == 8 && validateIPv6(addressBytes) {
			return "IPv6"
		}
	}
	return "Neither"
}

// Time:
// 	Usage : 0ms
// 	Beats : 100%
// Memory:
// 	Usage : 2MB
// 	Beats : 69.94%