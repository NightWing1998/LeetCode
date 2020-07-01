// Link - https://leetcode.com/problems/queue-reconstruction-by-height/

import (
	"sort"
)

func reconstructQueue(people [][]int) [][]int {
	sort.SliceStable(people, func(i, j int) bool {
		if people[i][0] == people[j][0] {
			return people[i][1] < people[j][1]
		} else {
			return people[i][0] > people[j][0]
		}
	})
	newPeople := make([][]int, 0, len(people))
	for _, curr := range people {
		newPeople = append(newPeople[:curr[1]], append([][]int{curr}, newPeople[curr[1]:]...)...)
	}
	return newPeople
}

// Time:
//  Usage : 36ms
//  Beats : 22.99%
// Memory:
//  Usage : 6.9MB
//  Beats : 7.69%