// Link - https://leetcode.com/problems/random-pick-with-weight/

import (
	"math/rand"
	"time"
)

type Solution struct {
	probability []int
}

func Constructor(w []int) Solution {
	probability := make([]int, len(w))
	if len(w) > 0 {
		probability[0] = w[0]
		for i := 1; i < len(w); i++ {
			probability[i] = w[i] + probability[i-1]
		}
	}
	return Solution{probability}
}

func (this Solution) PickIndex() int {
	time.Now().UnixNano()
	val := rand.Intn(this.probability[len(this.probability)-1] + 1)
	start := 0
	end := len(this.probability) - 1
	for start < end {
		mid := start + (end-start)/2
		if this.probability[mid] < val {
			start = mid + 1
		} else {
			end = mid
		}
	}
	return start
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(w);
 * param_1 := obj.PickIndex();
 */

// GOLANG RANDOM NOT UNIFORM HENCE CODE IN PYTHON