// Link - https://leetcode.com/problems/reconstruct-itinerary/
import "sort"

func findItinerary(tickets [][]string) []string {
	graph := make(map[string][]string)
	for i := 0; i < len(tickets); i++ {
		src, dst := tickets[i][0], tickets[i][1]
		_, exists := graph[src]
		if !exists {
			graph[src] = append(make([]string, 0), dst)
		} else {
			graph[src] = append(graph[src], dst)
		}
	}
	for key, _ := range graph {
		dsts, _ := graph[key]
		sort.Strings(dsts)
		graph[key] = dsts
	}
	path := make([]string, 0)
	stack := append(make([]string, 0), "JFK")
	n := 1
	for n > 0 {
		for dsts, exists := graph[stack[n-1]]; exists && len(dsts) > 0; dsts, exists = graph[stack[n-1]] {
			stack = append(stack, dsts[0])
			graph[stack[n-1]] = dsts[1:]
			dsts = dsts[1:]
			n++
		}
		path = append([]string{stack[n-1]}, path...)
		stack = stack[:n-1]
		n--
	}
	return path
}

// Time:
//  Usage: 16ms
// 	Beats: 37.55%
// Memory:
//  Usage: 6.5MB
// 	Beats: 30.3%