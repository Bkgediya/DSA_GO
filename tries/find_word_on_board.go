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

func FindAllWord(board [][]byte, words []string) []string {

	return []string{}
}
