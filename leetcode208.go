/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-01-27 00:57:48
 * @LastEditors: cm.d
 * @LastEditTime: 2022-01-27 01:22:05
 */
package myleetcode

type Trie struct {
	kv  map[byte]*Trie
	End bool
}

func Constructor() Trie {
	t := Trie{kv: make(map[byte]*Trie)}
	return t
}

func (this *Trie) Insert(word string) {
	z := this.kv
	for i, j := range []byte(word) {
		if t, ok := z[j]; ok {
			z = t.kv
			if i == len(word)-1 {
				t.End = true
			}
		} else {
			z[j] = &Trie{make(map[byte]*Trie), false}
			if i == len(word)-1 {
				z[j].End = true
			}
			z = z[j].kv

		}
	}
}

func (this *Trie) Search(word string) bool {
	z := this.kv
	for i, j := range []byte(word) {
		if t, ok := z[j]; ok {
			z = t.kv
			if i == len(word)-1 && t.End {
				return true
			}
		} else {
			return false
		}
	}

	return false
}

func (this *Trie) StartsWith(prefix string) bool {
	z := this.kv
	for _, j := range []byte(prefix) {
		if t, ok := z[j]; ok {
			z = t.kv
		} else {
			return false
		}
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
