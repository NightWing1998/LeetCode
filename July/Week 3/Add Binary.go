// Link - https://leetcode.com/problems/add-binary

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func addBinary(a string, b string) string {
	res := make([]byte, 0, max(len(a), len(b))+1)
	carry := 0
	i, j := len(a)-1, len(b)-1
	for ; i >= 0 && j >= 0; i, j = i-1, j-1 {
		ai := int(a[i]) - '0'
		bj := int(b[j]) - '0'
		res = append(res, byte((carry+ai+bj)%2+'0'))
		carry = (carry + ai + bj) / 2
	}
	for ; i >= 0; i-- {
		ai := int(a[i]) - '0'
		res = append(res, byte((carry+ai)%2+'0'))
		carry = (carry + ai) / 2
	}
	for ; j >= 0; j-- {
		bj := int(b[j]) - '0'
		res = append(res, byte((carry+bj)%2+'0'))
		carry = (carry + bj) / 2
	}
	if carry == 1 {
		res = append(res, '1')
	}
	for k := 0; k < len(res)/2; k++ {
		res[k], res[len(res)-1-k] = res[len(res)-1-k], res[k]
	}
	return string(res)
}

// Time - 0ms(beats 100%)
// Memory - 2.2MB(beats 100%)