// Link - https://leetcode.com/problems/two-city-scheduling/

import "sort"

func twoCitySchedCost(costs [][]int) int {
	// Greedy move: sort by diff with 0 first having the least cost
	sort.Slice(costs, func(i, j int) bool {
		return (costs[i][0] - costs[i][1]) < (costs[j][0] - costs[j][1])
	})
	sum := 0
	n := len(costs)
	// First all cities with least cost for a
	for i := 0; i < n/2; i++ {
		sum += costs[i][0]
	}
	for j := n / 2; j < n; j++ {
		sum += costs[j][1]
	}
	return sum
}

// Help from redit was taken

// Time:
//  Usage : 0ms
//  Beats : 100%
// Memory:
//  Usage : 2.5MB
//  Beats : 86.67%