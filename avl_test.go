package avl_test

import (
	"math/rand"
	"testing"

	"github.coom/zyedidia/avl"
)

const (
	opAdd = iota
	opRemove
	opSearch
)

func TestTree(t *testing.T) {
	tree := &avl.Tree{}
	m := make(map[int]int)

	const maxKey = 100
	const nops = 10000
	for i := 0; i < nops; i++ {
		op := rand.Intn(3)
		k := rand.Intn(maxKey)

		switch op {
		case opAdd:
			v := rand.Int()
			tree.Add(k, v)
			m[k] = v
		case opRemove:
			tree.Remove(k)
			delete(m, k)
		case opSearch:
			tv, _ := tree.Search(k)
			mv := m[k]
			if tv != mv {
				t.Errorf("Incorrect value for key %d, want: %d, got: %d", k, mv, tv)
			}
		}
	}
}
