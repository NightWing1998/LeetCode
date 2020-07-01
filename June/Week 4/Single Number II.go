func singleNumber(nums []int) int {
	count := make(map[int]int)
	for _, num := range nums {
		if _, exist := count[num]; exist {
			count[num]++
		} else {
			count[num] = 1
		}
	}
	for num, cnt := range count {
		if cnt == 1 {
			return num
		}
	}
	return -1
}