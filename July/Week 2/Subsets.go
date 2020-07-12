// Link - https://leetcode.com/problems/subsets/

import "math"

func subsets(nums []int) [][]int {
	setSize := int64(math.Pow(2, float64(len(nums))))
	result := make([][]int, 0, setSize)
	init := make([]int, 0, 0)
	result = append(result, init)
	for _, value := range nums {
		n := len(result)
		for i := 0; i < n; i++ {
			init = make([]int, len(result[i])+1)
			copy(init, result[i])
			init[len(result[i])] = value
			// Following above procedure as append doubles the capacity of main array and messes up the result as it changes a lot of existing slices in it;
			// result = append(result,append(result,value));

			result = append(result, init)
		}
	}
	return result
}

// Time - 0ms(beats 100%)
// Memory - 2.2MB(beats 100%)