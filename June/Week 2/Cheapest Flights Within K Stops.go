// Link - https://leetcode.com/problems/cheapest-flights-within-k-stops/

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
func findCheapestPrice(n int, flights [][]int, src int, dst int, K int) int {
	graph := make(map[int][][]int)
	// fill graph
	for _, nodes := range flights {
		source, dest, cost := nodes[0], nodes[1], nodes[2]
		_, exists := graph[source]
		if exists {
			graph[source] = append(graph[source], []int{dest, cost})
		} else {
			graph[source] = [][]int{[]int{dest, cost}}
		}
	}

	queue := [][]int{[]int{src, 0, 0}}
	mincost := 1000100 //random number larger than sum of costs
	for len(queue) != 0 {
		front := queue[0]
		tempSrc, tempK, tempCost := front[0], front[1], front[2]
		neighbours, exists := graph[tempSrc]
		if exists {
			for _, edge := range neighbours {
				if tempCost <= mincost && edge[0] != dst && tempK < K {
					queue = append(queue, []int{edge[0], tempK + 1, tempCost + edge[1]})
				} else if edge[0] == dst {
					mincost = min(tempCost+edge[1], mincost)
				}
			}
		}
		queue = queue[1:]
	}
	if mincost == 1000100 {
		return -1
	} else {
		return mincost
	}
}

// Time:
// 	Usage : 48ms
// 	Beats : 15.67%
// Memory:
// 	Usage : 8MB
// 	Beats : 23.53%