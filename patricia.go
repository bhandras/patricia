package patricia

import (
//"fmt"
)

type Node struct {
	skip  int
	value []byte
	l     *Node
	r     *Node
}

type Trie struct {
	root Node
}

func NewTrie() *Trie {
	t := &Trie{}
	t.root.l = &t.root
	t.root.skip = -1

	return t
}

func ByteMSB(x byte) int {
	var r int

	if x&0xF0 != 0 {
		x >>= 4
		r += 4
	}

	if x&0xC != 0 {
		x >>= 2
		r += 2
	}

	if x&0x2 != 0 {
		r++
	}

	return r
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func mismatch(a []byte, b []byte) (bool, int) {
	if len(a) == 0 || len(b) == 0 {
		return true, 0
	}

	limit := max(len(a), len(b))

	for i := 0; i < limit; i++ {
		var b1, b2 byte

		if i < len(a) {
			b1 = a[i]
		}

		if i < len(b) {
			b2 = b[i]
		}

		if b1 != b2 {
			return true, i*8 + (8 - 1 - ByteMSB(b1^b2))
		}
	}

	return false, 0
}

func bit(array []byte, n int) bool {
	if len(array) <= (n / 8) {
		return false
	}
	return ((array[n/8] >> uint(n%8)) & 1) == 1
}

func (t *Trie) Search(value []byte) bool {
	node := search(value, t.root.l, -1)

	if len(node.value) != len(value) {
		return false
	}

	for i := 0; i < len(node.value); i++ {
		if node.value[i] != value[i] {
			return false
		}
	}

	return true
}

func search(value []byte, node *Node, skip int) *Node {
	if node.skip <= skip {
		return node
	}

	if !bit(value, node.skip) {
		return search(value, node.l, node.skip)
	} else {
		return search(value, node.r, node.skip)
	}
}

func (t *Trie) Insert(value []byte) {
	node := search(value, t.root.l, 0)
	if res, skip := mismatch(value, node.value); res {
		newNode := &Node{
			value: value,
			skip:  skip,
		}
		t.root.l = insert(t.root.l, newNode, &t.root)
	}
}

func insert(node *Node, newNode *Node, parent *Node) *Node {
	if newNode.skip <= node.skip || node.skip <= parent.skip {
		if !bit(newNode.value, newNode.skip) {
			newNode.l = newNode
			newNode.r = node

			return newNode
		} else {
			newNode.l = node
			newNode.r = newNode

			return newNode
		}
	}

	if !bit(newNode.value, node.skip) {
		node.l = insert(node.l, newNode, node)
	} else {
		node.r = insert(node.r, newNode, node)
	}

	return node
}
