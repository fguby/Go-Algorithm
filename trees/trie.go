package trees

type Trie struct {
	size int
	root *trie
}

type trie struct {
	isWord bool
	next   map[rune]*trie
}

func NewTrie() *Trie {
	return &Trie{
		size: 0,
		root: &trie{
			isWord: false,
			next:   make(map[rune]*trie, 26),
		},
	}
}

func (t *Trie) GetSize() int {
	return t.size
}

func (t *Trie) Insert(words string) {
	b := t.root.insert(words)
	if !b {
		t.size++
	}
}

func (t *trie) insert(words string) bool {

	tree := t

	for _, v := range words {
		_, b := tree.next[v]
		if !b {
			tree.next[v] = &trie{
				isWord: false,
				next:   make(map[rune]*trie, 26),
			}
		}
		tree = tree.next[v]
	}
	b := tree.isWord
	tree.isWord = true
	return b
}

func (t *Trie) Search(words string) bool {
	return t.root.search(words)
}

func (t *trie) search(words string) bool {

	tree := t

	for _, v := range words {
		node, b := tree.next[v]
		if b {
			tree = node
		} else {
			return false
		}
	}

	return tree.isWord
}

func (t *Trie) Prefix(words string) bool {
	tree := t.root

	for _, v := range words {

		if _, b := tree.next[v]; !b {
			return false
		}
		tree = tree.next[v]
	}
	return true
}

func (t *Trie) PrefixWithRegular(words string) bool {
	return t.root.prefix(words)
}

func (t *trie) prefix(words string) bool {

	tree := t

	for k, v := range words {
		//rune '.'
		if v == 46 {
			for _, node := range tree.next {
				return node.prefix(words[k+1:])
			}
		}

		if _, b := tree.next[v]; !b {
			return false
		}
		tree = tree.next[v]
	}

	return true
}
