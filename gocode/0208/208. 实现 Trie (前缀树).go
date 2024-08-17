package l0208

type try map[byte]try

type Trie struct {
	head try
}

func Constructor() Trie {
	ret := Trie{}
	ret.head = make(map[byte]try)

	return ret
}

func (this *Trie) Insert(word string) {
	wordBytes := []byte(word)
	head := this.head
	for _, b := range wordBytes {
		next, ok := head[b]
		if !ok {
			next = make(map[byte]try)
			head[b] = next
		}
		head = next
	}
	head['0'] = make(map[byte]try)
}

func (this *Trie) Search(word string) bool {
	return this.StartsWith(word + "0")
}

func (this *Trie) StartsWith(prefix string) bool {
	wordBytes := []byte(prefix)
	head := this.head
	for _, b := range wordBytes {
		next, ok := head[b]
		if !ok {
			return false
		}
		head = next
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
