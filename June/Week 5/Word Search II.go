// Link - https://leetcode.com/problems/word-search-ii/

type Trie struct {
	word     string
	children [26]*Trie
}

// take arr of string as parameter and generate a Trie out of it
func GenerateTrie(words []string) *Trie {
	var root Trie = Trie{}
	for _, word := range words {
		current := &root
		for _, ch := range word {
			if current.children[ch-'a'] == nil {
				current.children[ch-'a'] = &Trie{} // if alpha does not have a node the gen one
			}
			current = current.children[ch-'a']
		}
		current.word = word // assign word to last alphabet of word
	}
	return &root
}

func findWords(board [][]byte, words []string) []string {
	result := make([]string, 0, len(words))
	root := GenerateTrie(words)
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			dfs(board, i, j, root, &result)
		}
	}
	return result
}

func dfs(board [][]byte, i, j int, root *Trie, result *[]string) {
	ch := board[i][j]
	if ch == '*' || root.children[ch-'a'] == nil { // * indicates visited cell
		return
	}
	temp := root.children[ch-'a'] // pick up the node coresponding to alpha
	if len(temp.word) != 0 {      // if word  != "" the found match append to result
		*result = append(*result, temp.word)
		temp.word = "" // delete word so that it will be added again in result
	}

	board[i][j] = '*' // cell visited
	if i > 0 {
		dfs(board, i-1, j, temp, result) // up
	}
	if j > 0 {
		dfs(board, i, j-1, temp, result) // left
	}
	if i < len(board)-1 {
		dfs(board, i+1, j, temp, result) // down
	}
	if j < len(board[i])-1 {
		dfs(board, i, j+1, temp, result) // right
	}
	board[i][j] = ch // backtracking
}

// Time :
// 	Usage : 24ms
//	Beats : 94.96%
// Memory:
// 	Usage : 14.8MB
// 	Beats : 32.97%