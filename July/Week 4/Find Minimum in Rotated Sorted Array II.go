// Link - https://leetcode.com/problems/find-minimum-in-rotated-sorted-array-ii

func findMin(nums []int) int {
	// l := 0;
	// r := len(nums) - 1;
	// var mid int;
	// for l < r {
	//     mid = (l+r)/2;
	//     if nums[mid] > nums[r] {
	//         l = mid + 1
	//     } else if nums[mid] < nums[r] {
	//         r = mid;
	//     } else {
	//         r--;
	//     }
	// }
	// return nums[l];

	// Above solution time complexity - O(logn) but performs very poorly
	min := nums[0]
	for _, num := range nums {
		if num < min {
			min = num
			break
		}
	}
	return min
}

// Time - 4ms(beats 88.89%)
// Memory - 3.1MB(beats 60%)