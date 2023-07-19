package util

import (
	"fmt"
	"testing"
)

func checkOrder(t *testing.T, nodes []*SortedSetNode, expectedOrder []SKEY) {
	if len(expectedOrder) != len(nodes) {
		t.Errorf("nodes does not contain %d elements", len(expectedOrder))
	}
	for i := 0; i < len(expectedOrder); i++ {
		if nodes[i].Key() != expectedOrder[i] {
			t.Errorf("nodes[%d] is %q, but the expected key is %q", i, nodes[i].Key(), expectedOrder[i])
		}
	}
}

func checkIterateRangeByRank(t *testing.T, sortedset *SortedSet, start int, end int, expectedOrder []SKEY) {
	var keys []SKEY

	// check nil callback should do nothing
	sortedset.IterateRangeByRank(start, end, nil)

	sortedset.IterateRangeByRank(start, end, func(key SKEY, _ any) bool {
		keys = append(keys, key)
		return true
	})
	if len(expectedOrder) != len(keys) {
		t.Errorf("keys does not contain %d elements", len(expectedOrder))
	}
	for i := 0; i < len(expectedOrder); i++ {
		if keys[i] != expectedOrder[i] {
			t.Errorf("keys[%d] is %q, but the expected key is %q", i, keys[i], expectedOrder[i])
		}
	}

	// check return early
	if len(expectedOrder) < 1 {
		return
	}
	// reset data
	keys = []SKEY{}
	var i int
	sortedset.IterateRangeByRank(start, end, func(key SKEY, _ any) bool {
		keys = append(keys, key)
		i++
		// return early
		return i < len(expectedOrder)-1
	})
	if len(expectedOrder)-1 != len(keys) {
		t.Errorf("keys does not contain %d elements", len(expectedOrder)-1)
	}
	for i := 0; i < len(expectedOrder)-1; i++ {
		if keys[i] != expectedOrder[i] {
			t.Errorf("keys[%d] is %q, but the expected key is %q", i, keys[i], expectedOrder[i])
		}
	}

}

func checkRankRangeIterAndOrder(t *testing.T, sortedset *SortedSet, start int, end int, remove bool, expectedOrder []SKEY) {
	checkIterateRangeByRank(t, sortedset, start, end, expectedOrder)
	nodes := sortedset.RangeByRank(start, end, remove)
	checkOrder(t, nodes, expectedOrder)
}

func TestCase1(t *testing.T) {
	sortedset := NewSortedSet()

	sortedset.Add('a', 89, "Kelly")
	sortedset.Add('b', 100, "Staley")
	sortedset.Add('c', 100, "Jordon")
	sortedset.Add('d', -321, "Park")
	sortedset.Add('e', 101, "Albert")
	sortedset.Add('f', 99, "Lyman")
	sortedset.Add('g', 99, "Singleton")
	sortedset.Add('h', 70, "Audrey")

	sortedset.Add('e', 99, "ntrnrt")

	sortedset.Remove('b')

	node := sortedset.GetByRank(3, false)
	if node == nil || node.Key() != 'a' {
		t.Error("GetByRank() does not return expected value `a`")
	}

	node = sortedset.GetByRank(-3, false)
	if node == nil || node.Key() != 'f' {
		t.Error("GetByRank() does not return expected value `f`")
	}

	// get all nodes since the first one to last one
	checkRankRangeIterAndOrder(t, sortedset, 1, -1, false, []SKEY{'d', 'h', 'a', 'e', 'f', 'g', 'c'})

	// get & remove the 2nd/3rd nodes in reserve order
	checkRankRangeIterAndOrder(t, sortedset, -2, -3, true, []SKEY{'g', 'f'})

	// get all nodes since the last one to first one
	checkRankRangeIterAndOrder(t, sortedset, -1, 1, false, []SKEY{'c', 'e', 'a', 'h', 'd'})
}

func TestCase2(t *testing.T) {

	// create a new set
	sortedset := NewSortedSet()

	// fill in new node
	sortedset.Add('a', 89, "Kelly")
	sortedset.Add('b', 100, "Staley")
	sortedset.Add('c', 100, "Jordon")
	sortedset.Add('d', -321, "Park")
	sortedset.Add('e', 101, "Albert")
	sortedset.Add('f', 99, "Lyman")
	sortedset.Add('g', 99, "Singleton")
	sortedset.Add('h', 70, "Audrey")

	// update an existing node
	sortedset.Add('e', 99, "ntrnrt")

	// remove node
	sortedset.Remove('b')

	nodes := sortedset.RangeByScore(-500, 500, false, false, 0)
	checkOrder(t, nodes, []SKEY{'d', 'h', 'a', 'e', 'f', 'g', 'c'})

	nodes = sortedset.RangeByScore(500, -500, false, false, 0)
	//t.Logf("%v", nodes)
	checkOrder(t, nodes, []SKEY{'c', 'g', 'f', 'e', 'a', 'h', 'd'})

	nodes = sortedset.RangeByScore(600, 500, false, false, 0)
	checkOrder(t, nodes, []SKEY{})

	nodes = sortedset.RangeByScore(500, 600, false, false, 0)
	checkOrder(t, nodes, []SKEY{})

	rank := sortedset.Rank('f')
	if rank != 5 {
		t.Error("FindRank() does not return expected value `5`")
	}

	rank = sortedset.Rank('d')
	if rank != 1 {
		t.Error("FindRank() does not return expected value `1`")
	}

	nodes = sortedset.RangeByScore(99, 100, false, false, 0)
	checkOrder(t, nodes, []SKEY{'e', 'f', 'g', 'c'})

	nodes = sortedset.RangeByScore(90, 50, false, false, 0)
	checkOrder(t, nodes, []SKEY{'a', 'h'})

	nodes = sortedset.RangeByScore(99, 100, true, false, 0)
	checkOrder(t, nodes, []SKEY{'c'})

	nodes = sortedset.RangeByScore(100, 99, true, false, 0)
	checkOrder(t, nodes, []SKEY{'g', 'f', 'e'})

	nodes = sortedset.RangeByScore(99, 100, false, true, 0)
	checkOrder(t, nodes, []SKEY{'e', 'f', 'g'})

	nodes = sortedset.RangeByScore(100, 99, false, true, 0)
	checkOrder(t, nodes, []SKEY{'c'})

	nodes = sortedset.RangeByScore(50, 100, false, false, 2)
	checkOrder(t, nodes, []SKEY{'h', 'a'})

	nodes = sortedset.RangeByScore(100, 50, false, false, 2)
	checkOrder(t, nodes, []SKEY{'c', 'g'})

	minNode := sortedset.GetMin()
	if minNode == nil || minNode.Key() != 'd' {
		t.Error("PeekMin() does not return expected value `d`")
	}

	minNode = sortedset.PopMin()
	if minNode == nil || minNode.Key() != 'd' {
		t.Error("PopMin() does not return expected value `d`")
	}

	nodes = sortedset.RangeByScore(-500, 500, false, false, 0)
	checkOrder(t, nodes, []SKEY{'h', 'a', 'e', 'f', 'g', 'c'})

	maxNode := sortedset.GetMax()
	if maxNode == nil || maxNode.Key() != 'c' {
		t.Error("PeekMax() does not return expected value `c`")
	}

	maxNode = sortedset.PopMax()
	if maxNode == nil || maxNode.Key() != 'c' {
		t.Error("PopMax() does not return expected value `c`")
	}

	nodes = sortedset.RangeByScore(500, -500, false, false, 0)
	checkOrder(t, nodes, []SKEY{'g', 'f', 'e', 'a', 'h'})
}

func TestCase3(t *testing.T) {
	// fill in new node
	var m = map[int64]int64{}
	m['a'] = 100
	m['b'] = 200
	m['c'] = 300
	sortedset := NewSortedSet()
	sortedset.SetSameScoreCompareFn(func(k1, k2 SKEY) int {
		if k1 == k2 {
			return 0
		}

		kv1 := m[int64(k1)]
		kv2 := m[int64(k2)]
		if kv1 != kv2 {
			return If(kv1 > kv2, 1, -1)
		}

		return If(k1 > k2, 1, -1)
	})

	sortedset.Add('e', 99, "test")
	sortedset.Add('a', 100, "Kelly")
	sortedset.Add('b', 100, "Staley")
	sortedset.Add('c', 100, "Jordon")

	sortedset.IterateRangeByRank(0, sortedset.Count(), func(key SKEY, value any) bool {
		fmt.Println(key)
		return true
	})
}
