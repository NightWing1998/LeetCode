// Link - https://leetcode.com/problems/reverse-bits

func reverseBits(num uint32) uint32 {
	// extract half bits, shift for swap, | to complete the ans
	num = ((num & 0xffff0000) >> 16) | ((num & 0x0000ffff) << 16)
	num = ((num & 0xff00ff00) >> 8) | ((num & 0x00ff00ff) << 8)
	num = ((num & 0xf0f0f0f0) >> 4) | ((num & 0x0f0f0f0f) << 4)
	num = ((num & 0xcccccccc) >> 2) | ((num & 0x33333333) << 2)
	num = ((num & 0xaaaaaaaa) >> 1) | ((num & 0x55555555) << 1)
	return num
}

// Time - 0ms(beats 100%)
// Mempry - 2.6MB(beats 14.81%)