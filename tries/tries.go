package tries

type TrieNode struct {
	children map[rune]*TrieNode
	is_end   bool
}

type Trie struct {
	root *TrieNode
}

func NewTrie() *Trie {
	return &Trie{
		root: &TrieNode{
			children: make(map[rune]*TrieNode),
		},
	}
}

// Insert a word into the trie
// Time Complexity: O(n)
// Space Complexity: O(n)
func (t *Trie) Insert(word string) {
	node := t.root
	for _, ch := range word {
		if _, ok := node.children[ch]; !ok {
			node.children[ch] = &TrieNode{children: make(map[rune]*TrieNode)}
		}
		node = node.children[ch]
	}
	node.is_end = true
}

// Search a word in the trie
// Time Complexity: O(n)
func (t *Trie) Search(word string) bool {
	node := t.root
	for _, ch := range word {
		if _, ok := node.children[ch]; !ok {
			return false
		}
		node = node.children[ch]
	}
	return node.is_end
}

func (t *Trie) StartsWith(prefix string) bool {
	node := t.root
	for _, ch := range prefix {
		if _, ok := node.children[ch]; !ok {
			return false
		}
		node = node.children[ch]
	}
	return true
}

func (t *Trie) SearchWithWildCard(word string) bool {
	return dfs_search(0, t.root, word)
}

func dfs_search(index int, node *TrieNode, word string) bool {
	for i := index; i < len(word); i++ {
		ch := rune(word[i])

		if ch == '.' {
			for _, child := range node.children {
				if dfs_search(i+1, child, word) {
					return true
				}
			}
			return false
		} else if _, ok := node.children[ch]; ok {
			node = node.children[ch]
		} else {
			return false
		}
	}

	return node.is_end
}
