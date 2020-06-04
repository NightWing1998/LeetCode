func reverseString(s []byte) {
	n := len(s)
	for i := 0; i < n/2; i++ {
		// XOR swap
		if s[i] != s[n-1-i] {
			s[i] ^= s[n-1-i]
			s[n-1-i] ^= s[i]
			s[i] ^= s[n-1-i]
		}
	}
}

// faster algo
// for i := 0; i < len(s)/2; i++ {
//     s[i], s[len(s)-i-1] = s[len(s)-i-1], s[i]
// }

// h e l l o
// 0 1 2 3 4 n=5