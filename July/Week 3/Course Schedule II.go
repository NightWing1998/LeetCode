// Link - https://leetcode.com/problems/course-schedule-ii

func findOrder(numCourses int, prerequisites [][]int) []int {
	label := make([]int, numCourses)
	coursesOrder := make([]int, 0, numCourses)
	graph := make(map[int][]int)
	for _, edge := range prerequisites {
		src, dest := edge[0], edge[1]
		neighbours, _ := graph[src]
		graph[src] = append(neighbours, dest)
	}
	for node, _ := range label {
		if !visit(node, graph, label, &coursesOrder) {
			return []int{}
		}
	}
	return coursesOrder
}

// https://en.wikipedia.org/wiki/Topological_sorting#Algorithms --> Topological sort DFS
func visit(node int, graph map[int][]int, label []int, coursesOrder *[]int) bool {
	if label[node] == 1 {
		return true
	} else if label[node] == -1 {
		return false
	}
	label[node] = -1

	if neighbours, exists := graph[node]; exists {
		for _, newNode := range neighbours {
			if !visit(newNode, graph, label, coursesOrder) {
				return false
			}
		}
	}

	label[node] = 1
	(*coursesOrder) = append((*coursesOrder), node)
	return true
}

// Time - 16ms(beats 53.57%)
// Memory - 6.3MB(beats 43.28%)