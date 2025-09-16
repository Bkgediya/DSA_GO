package tries

type TrieNodeBoard struct {
	children map[rune]*TrieNodeBoard
	word     string
}

type TrieBoard struct {
	root *TrieNodeBoard
}

func NewTrieBoard() *TrieBoard {
	return &TrieBoard{
		root: &TrieNodeBoard{
			children: make(map[rune]*TrieNodeBoard),
		},
	}
}

func (t *TrieBoard) Insert(word string) {
	node := t.root

	for _, ch := range word {
		if _, ok := node.children[ch]; !ok {
			node.children[ch] = &TrieNodeBoard{
				children: make(map[rune]*TrieNodeBoard),
			}
		}
		node = node.children[ch]
	}
	node.word = word
}

func FindAllWord(board [][]rune, words []string) []string {
	tr := NewTrieBoard()

	for _, word := range words {
		tr.Insert(word)
	}

	res := []string{}

	for row := range board {
		for col := range board[0] {
			if child, ok := tr.root.children[board[row][col]]; ok {
				dfs(board, row, col, child, &res)
			}
		}
	}

	return res
}

func dfs(board [][]rune, row, col int, node *TrieNodeBoard, res *[]string) {
	if node.word != "" {
		*res = append(*res, node.word)
		node.word = ""
	}

	ch := board[row][col]
	board[row][col] = '#'

	// explore all the adjucent cells
	dirs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for _, dir := range dirs {
		next_row, next_col := row+dir[0], col+dir[1]
		if is_within_bound(next_row, next_col, board) {
			if child, ok := node.children[board[next_row][next_col]]; ok {
				dfs(board, next_row, next_col, child, res)
			}
		}
	}

	board[row][col] = ch

}

func is_within_bound(row, col int, board [][]rune) bool {
	return row >= 0 && row < len(board) && col >= 0 && col < len(board[0])
}
