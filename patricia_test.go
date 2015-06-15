package patricia

import (
	"fmt"
	"testing"
)

func TestByteMSB(t *testing.T) {
	var n int

	n = ByteMSB(1)
	if n != 0 {
		t.Errorf("msb(1)=%d != 0", n)
	}

	n = ByteMSB(2)
	if n != 1 {
		t.Errorf("msb(2)=%d != 1", n)
	}

	n = ByteMSB(3)
	if n != 1 {
		t.Errorf("msb(3)=%d != 1", n)
	}
}

func printBits(array []byte) string {
	str := "["
	for i := 0; i < len(array); i++ {
		str += fmt.Sprintf(" %08b", array[i])
	}

	str += " ]"
	return str
}

func mismatchTest(a []byte, b []byte, res bool, n int, t *testing.T) {
	_res, _n := mismatch(a, b)

	if _res != res || _n != n {
		t.Errorf("mismatch error: a=%s, b=%s, res=(e: %t; r: %t), n=(e: %d; r: %d)",
			printBits(a), printBits(b), res, _res, n, _n)
	}
}

func TestMismatch(t *testing.T) {
	mismatchTest([]byte{0}, []byte{0}, false, 0, t)
	mismatchTest([]byte{0, 1}, []byte{0, 1}, false, 0, t)

	mismatchTest([]byte{0}, []byte{1}, true, 7, t)
	mismatchTest([]byte{1, 0}, []byte{0, 1}, true, 7, t)
	mismatchTest([]byte{123}, []byte{233}, true, 0, t)
}

func TestPatTrue(t *testing.T) {
	trie := NewTrie()

	trie.Insert([]byte("alma"))
	trie.Insert([]byte("almafa"))

	if !trie.Search([]byte("alma")) {
		t.Error("\"alma\" not found")
	}

	if !trie.Search([]byte("almafa")) {
		t.Error("\"almafa\" not found")
	}

	if trie.Search([]byte("alm")) {
		t.Error("\"almaf\" found")
	}

	if trie.Search([]byte("almafaa")) {
		t.Error("\"almafaa\" found")
	}
}
