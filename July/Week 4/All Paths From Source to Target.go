// Link - https://leetcode.com/problems/all-paths-from-source-to-target

type node struct {
	current int
	trace   []int
}

func allPathsSourceTarget(graph [][]int) [][]int {
	n := len(graph)
	result := make([][]int, 0, n)
	queue := make([]node, 0, n)
	for _, dst := range graph[0] {
		queue = append(queue, node{current: dst, trace: []int{0, dst}})
	}
	for len(queue) > 0 {
		top := queue[0]
		queue = queue[1:]
		if top.current == n-1 {
			result = append(result, top.trace)
		} else {
			for _, dst := range graph[top.current] {
				newTrace := append([]int{}, top.trace...)
				queue = append(queue, node{current: dst, trace: append(newTrace, dst)})
			}
		}
	}
	return result
}

// Time - 16ms(beats 44.83%)
// Memory - 6.9MB(-)