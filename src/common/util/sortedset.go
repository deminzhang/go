package util

import (
	"math"
	"math/rand"
)

//from: https://github.com/wangjia184/sortedset

type SKEY int64
type SSCORE int64 // the type of score

const (
	SkiplistMaxLevel = 32   /* Should be enough for 2^32 elements */
	SkiplistP        = 0.25 /* Skiplist P = 1/4 */
)

type sortedSetLevel struct {
	forward *SortedSetNode
	span    int64
}

// 1 k1score > k2score  k1score < k2score
type SortedSameScoreCompareFn func(k1, k2 SKEY) int

func defaultSortedSameScoreCompare(k1, k2 SKEY) int {
	if k1 == k2 {
		return 0
	}

	return If(k1 > k2, 1, -1)
}

// Node in skip list
type SortedSetNode struct {
	key      SKEY        // unique key of this node
	score    SSCORE      // score to determine the order of this node in the set
	Value    interface{} // associated data
	backward *SortedSetNode
	level    []sortedSetLevel
}

func (this *SortedSetNode) Key() SKEY {
	return this.key
}

func (this *SortedSetNode) Score() SSCORE {
	return this.score
}

type SortedSet struct {
	header         *SortedSetNode
	tail           *SortedSetNode
	length         int64
	level          int
	dict           map[SKEY]*SortedSetNode
	scoreCompareFn SortedSameScoreCompareFn
}

func createNode(level int, score SSCORE, key SKEY, value interface{}) *SortedSetNode {
	/*node := _sortedNodePool.get(level)
	node.Value = value
	node.key = key
	node.score = score
	return node*/

	return &SortedSetNode{
		score: score,
		key:   key,
		Value: value,
		level: make([]sortedSetLevel, level),
	}
}

// Returns a random level for the new skiplist node we are going to create.
// The return value of this function is between 1 and SkiplistMaxLevel
// (both inclusive), with a powerlaw-alike distribution where higher
// levels are less likely to be returned.
func randomLevel() int {
	level := 1
	for float64(rand.Int31()&0xFFFF) < (SkiplistP * 0xFFFF) {
		level++
	}
	if level < SkiplistMaxLevel {
		return level
	}

	return SkiplistMaxLevel
}

func (this *SortedSet) insertNode(score SSCORE, key SKEY, value interface{}) *SortedSetNode {
	var update [SkiplistMaxLevel]*SortedSetNode
	var rank [SkiplistMaxLevel]int64

	sameScoreCompareFn := this.getSameScoreCompare()
	x := this.header
	for i := this.level - 1; i >= 0; i-- {
		/* store rank that is crossed to reach the insert position */
		if this.level-1 == i {
			rank[i] = 0
		} else {
			rank[i] = rank[i+1]
		}

		for x.level[i].forward != nil &&
			(x.level[i].forward.score < score ||
				(x.level[i].forward.score == score && // score is the same but the key is different
					sameScoreCompareFn(x.level[i].forward.key, key) == -1)) {
			rank[i] += x.level[i].span
			x = x.level[i].forward
		}
		update[i] = x
	}

	/* we assume the key is not already inside, since we allow duplicated
	 * scores, and the re-insertion of score and redis object should never
	 * happen since the caller of Insert() should test in the hash cfg
	 * if the element is already inside or not. */
	level := randomLevel()

	if level > this.level { // add a new level
		for i := this.level; i < level; i++ {
			rank[i] = 0
			update[i] = this.header
			update[i].level[i].span = this.length
		}
		this.level = level
	}

	x = createNode(level, score, key, value)
	for i := 0; i < level; i++ {
		x.level[i].forward = update[i].level[i].forward
		update[i].level[i].forward = x

		/* update span covered by update[i] as x is inserted here */
		x.level[i].span = update[i].level[i].span - (rank[0] - rank[i])
		update[i].level[i].span = (rank[0] - rank[i]) + 1
	}

	/* increment span for untouched levels */
	for i := level; i < this.level; i++ {
		update[i].level[i].span++
	}

	if update[0] == this.header {
		x.backward = nil
	} else {
		x.backward = update[0]
	}
	if x.level[0].forward != nil {
		x.level[0].forward.backward = x
	} else {
		this.tail = x
	}
	this.length++
	return x
}

/* Internal function used by delete, DeleteByScore and DeleteByRank */
func (this *SortedSet) deleteNode(x *SortedSetNode, update [SkiplistMaxLevel]*SortedSetNode) {
	for i := 0; i < this.level; i++ {
		if update[i].level[i].forward == x {
			update[i].level[i].span += x.level[i].span - 1
			update[i].level[i].forward = x.level[i].forward
		} else {
			update[i].level[i].span -= 1
		}
	}
	if x.level[0].forward != nil {
		x.level[0].forward.backward = x.backward
	} else {
		this.tail = x.backward
	}
	for this.level > 1 && this.header.level[this.level-1].forward == nil {
		this.level--
	}
	this.length--
	delete(this.dict, x.key)
	//_sortedNodePool.put(x)
}

/* Delete an element with matching score/key from the skiplist. */
func (this *SortedSet) delete(score SSCORE, key SKEY) bool {
	var update [SkiplistMaxLevel]*SortedSetNode

	sameScoreCompareFn := this.getSameScoreCompare()
	x := this.header
	for i := this.level - 1; i >= 0; i-- {
		for x.level[i].forward != nil &&
			(x.level[i].forward.score < score ||
				(x.level[i].forward.score == score &&
					sameScoreCompareFn(x.level[i].forward.key, key) == -1)) {
			x = x.level[i].forward
		}
		update[i] = x
	}
	/* We may have multiple elements with the same score, what we need
	 * is to find the element with both the right score and object. */
	x = x.level[0].forward
	if x != nil && score == x.score && x.key == key {
		this.deleteNode(x, update)
		// free x
		return true
	}
	return false /* not found */
}

// Create a SortedSet
func NewSortedSet() *SortedSet {
	sortedSet := SortedSet{
		level: 1,
		dict:  make(map[SKEY]*SortedSetNode),
	}
	sortedSet.header = createNode(SkiplistMaxLevel, 0, 0, nil)
	return &sortedSet
}

// Get the number of elements
func (this *SortedSet) Count() int {
	return int(this.length)
}

func (this *SortedSet) getSameScoreCompare() SortedSameScoreCompareFn {
	if this.scoreCompareFn != nil {
		return this.scoreCompareFn
	}
	return defaultSortedSameScoreCompare
}

func (this *SortedSet) SetSameScoreCompareFn(fn SortedSameScoreCompareFn) {
	if this.scoreCompareFn == nil {
		if len(this.dict) == 0 {
			this.scoreCompareFn = fn
		}
	}
}

// get the element with minimum score, nil if the set is empty
// Time complexity of this method is : O(log(N))
func (this *SortedSet) GetMin() *SortedSetNode {
	return this.header.level[0].forward
}

// get and remove the element with minimal score, nil if the set is empty
// Time complexity of this method is : O(log(N))
func (this *SortedSet) PopMin() *SortedSetNode {
	x := this.header.level[0].forward
	if x != nil {
		this.Remove(x.key)
	}
	return x
}

// get the element with maximum score, nil if the set is empty
// Time Complexity : O(1)
func (this *SortedSet) GetMax() *SortedSetNode {
	return this.tail
}

// get and remove the element with maximum score, nil if the set is empty
// Time complexity of this method is : O(log(N))
func (this *SortedSet) PopMax() *SortedSetNode {
	x := this.tail
	if x != nil {
		this.Remove(x.key)
	}
	return x
}

// Add an element into the sorted set with specific key / value / score.
// if the element is added, this method returns true; otherwise false means updated
// Time complexity of this method is : O(log(N))
func (this *SortedSet) Add(key SKEY, score SSCORE, value interface{}) bool {
	var newNode *SortedSetNode = nil

	found := this.dict[key]
	if found != nil {
		// score does not change, only update value
		if found.score == score {
			found.Value = value
		} else { // score changes, delete and re-insert
			this.delete(found.score, found.key)
			newNode = this.insertNode(score, key, value)
		}
	} else {
		newNode = this.insertNode(score, key, value)
	}

	if newNode != nil {
		this.dict[key] = newNode
	}
	return found == nil
}

// Delete element specified by key
// Time complexity of this method is : O(log(N))
func (this *SortedSet) Remove(key SKEY) *SortedSetNode {
	found := this.dict[key]
	if found != nil {
		this.delete(found.score, found.key)
		return found
	}
	return nil
}

// Get the nodes whose score within the specific range
// If options is nil, it searchs in interval [start, end] without any limit by default
// Time complexity of this method is : O(log(N))
func (this *SortedSet) RangeByScore(start SSCORE, end SSCORE, excludeStart bool, excludeEnd bool, limit int) []*SortedSetNode {
	if limit <= 0 {
		limit = math.MaxInt
	}

	reverse := start > end
	if reverse {
		start, end = end, start
		excludeStart, excludeEnd = excludeEnd, excludeStart
	}

	var nodes []*SortedSetNode

	//determine if out of range
	if this.length == 0 {
		return nodes
	}

	if reverse { // search from end to start
		x := this.header

		if excludeEnd {
			for i := this.level - 1; i >= 0; i-- {
				for x.level[i].forward != nil &&
					x.level[i].forward.score < end {
					x = x.level[i].forward
				}
			}
		} else {
			for i := this.level - 1; i >= 0; i-- {
				for x.level[i].forward != nil &&
					x.level[i].forward.score <= end {
					x = x.level[i].forward
				}
			}
		}

		for x != nil && limit > 0 {
			if excludeStart {
				if x.score <= start {
					break
				}
			} else {
				if x.score < start {
					break
				}
			}

			next := x.backward

			nodes = append(nodes, x)
			limit--

			x = next
		}
	} else {
		// search from start to end
		x := this.header
		if excludeStart {
			for i := this.level - 1; i >= 0; i-- {
				for x.level[i].forward != nil &&
					x.level[i].forward.score <= start {
					x = x.level[i].forward
				}
			}
		} else {
			for i := this.level - 1; i >= 0; i-- {
				for x.level[i].forward != nil &&
					x.level[i].forward.score < start {
					x = x.level[i].forward
				}
			}
		}

		/* Current node is the last with score < or <= start. */
		x = x.level[0].forward

		for x != nil && limit > 0 {
			if excludeEnd {
				if x.score >= end {
					break
				}
			} else {
				if x.score > end {
					break
				}
			}

			next := x.level[0].forward

			nodes = append(nodes, x)
			limit--

			x = next
		}
	}

	return nodes
}

// sanitizeIndexes return start, end, and reverse flag
func (this *SortedSet) sanitizeIndexes(start int, end int) (int, int, bool) {
	if start < 0 {
		start = int(this.length) + start + 1
	}
	if end < 0 {
		end = int(this.length) + end + 1
	}
	if start <= 0 {
		start = 1
	}
	if end <= 0 {
		end = 1
	}

	reverse := start > end
	if reverse { // swap start and end
		start, end = end, start
	}
	return start, end, reverse
}

func (this *SortedSet) findNodeByRank(start int, remove bool) (traversed int, x *SortedSetNode, update [SkiplistMaxLevel]*SortedSetNode) {
	x = this.header
	for i := this.level - 1; i >= 0; i-- {
		for x.level[i].forward != nil &&
			traversed+int(x.level[i].span) < start {
			traversed += int(x.level[i].span)
			x = x.level[i].forward
		}
		if remove {
			update[i] = x
		} else {
			if traversed+1 == start {
				break
			}
		}
	}
	return
}

// Get nodes within specific rank range [start, end]
// Note that the rank is 1-based integer. Rank 1 means the first node; Rank -1 means the last node;
// If start is greater than end, the returned array is in reserved order
// If remove is true, the returned nodes are removed
// Time complexity of this method is : O(log(N))
func (this *SortedSet) RangeByRank(start int, end int, remove bool) []*SortedSetNode {
	start, end, reverse := this.sanitizeIndexes(start, end)

	var nodes []*SortedSetNode

	traversed, x, update := this.findNodeByRank(start, remove)

	traversed++
	x = x.level[0].forward
	for x != nil && traversed <= end {
		next := x.level[0].forward

		nodes = append(nodes, x)

		if remove {
			this.deleteNode(x, update)
		}

		traversed++
		x = next
	}

	if reverse {
		for i, j := 0, len(nodes)-1; i < j; i, j = i+1, j-1 {
			nodes[i], nodes[j] = nodes[j], nodes[i]
		}
	}
	return nodes
}

// IterateRangeByRank apply fn to node within specific rank range [start, end]
// or until fn return false
// Note that the rank is 1-based integer. Rank 1 means the first node; Rank -1 means the last node;
// If start is greater than end, apply fn in reserved order
// If fn is nil, this function return without doing anything
func (this *SortedSet) IterateRangeByRank(start int, end int, fn func(key SKEY, value any) bool) {
	if fn == nil {
		return
	}

	start, end, reverse := this.sanitizeIndexes(start, end)
	traversed, x, _ := this.findNodeByRank(start, false)
	var nodes []*SortedSetNode

	x = x.level[0].forward
	for x != nil && traversed < end {
		next := x.level[0].forward

		if reverse {
			nodes = append(nodes, x)
		} else if !fn(x.key, x.Value) {
			return
		}

		traversed++
		x = next
	}

	if reverse {
		for i := len(nodes) - 1; i >= 0; i-- {
			if !fn(nodes[i].key, nodes[i].Value) {
				return
			}
		}
	}
}

// Get node by rank.
// Note that the rank is 1-based integer. Rank 1 means the first node; Rank -1 means the last node;
// If remove is true, the returned nodes are removed
// If node is not found at specific rank, nil is returned
// Time complexity of this method is : O(log(N))
func (this *SortedSet) GetByRank(rank int, remove bool) *SortedSetNode {
	nodes := this.RangeByRank(rank, rank, remove)
	if len(nodes) == 1 {
		return nodes[0]
	}
	return nil
}

// Get node by key
// If node is not found, nil is returned
// Time complexity : O(1)
func (this *SortedSet) GetByKey(key SKEY) *SortedSetNode {
	return this.dict[key]
}

// Get the rank of the node specified by key
// Note that the rank is 1-based integer. Rank 1 means the first node
// If the node is not found, 0 is returned. Otherwise rank(> 0) is returned
// Time complexity of this method is : O(log(N))
func (this *SortedSet) Rank(key SKEY) int {
	var rank int = 0
	node := this.dict[key]
	if node != nil {
		x := this.header
		for i := this.level - 1; i >= 0; i-- {
			for x.level[i].forward != nil &&
				(x.level[i].forward.score < node.score ||
					(x.level[i].forward.score == node.score &&
						x.level[i].forward.key <= node.key)) {
				rank += int(x.level[i].span)
				x = x.level[i].forward
			}

			if x.key == key {
				return rank
			}
		}
	}
	return 0
}

/*
//level pool 可以构建为2,4,8,16,32 5中类型的池子, 根据createNode 传递的level,

type noCopy struct{}

func (*noCopy) Lock()   {}
func (*noCopy) Unlock() {}

type sortedPoolGet func(level int) *SortedSetNode
type sortedPoolPut func(node *SortedSetNode)

type sortedNodePool struct {
	noCopy
	get sortedPoolGet
	put sortedPoolPut
}

var _sortedNodePool sortedNodePool

func init() {
	const level = 5
	var table = [level]int{2, 4, 8, 16, 32}
	var nodePool = sync.Pool{
		New: func() any {
			return &SortedSetNode{}
		},
	}

	var levelPool [level]*sync.Pool
	for i := 0; i < len(levelPool); i++ {
		l := table[i]
		levelPool[i] = &sync.Pool{
			New: func() any {
				return make([]sortedSetLevel, l)
			},
		}
	}

	_sortedNodePool.get = func(l int) *SortedSetNode {
		for i := 0; i < level; i++ {
			if l <= table[i] {
				node := nodePool.Get().(*SortedSetNode)
				lp := levelPool[i].Get().([]sortedSetLevel)
				node.level = unsafe.Slice(&lp[0], l)
				return node
			}
		}
		node := nodePool.Get().(*SortedSetNode)
		node.level = make([]sortedSetLevel, l)
		return node
	}

	_sortedNodePool.put = func(node *SortedSetNode) {
		nodeLevel := node.level
		node.level = nil
		node.Value = nil
		nodePool.Put(node)
		for i := 0; i < level; i++ {
			if len(node.level) <= table[i] {
				nodeLevel = unsafe.Slice(&nodeLevel[0], table[i])
				for j := 0; j < len(nodeLevel); j++ {
					nodeLevel[j].forward = nil
					nodeLevel[j].span = 0
				}
				levelPool[i].Put(nodeLevel)
				return
			}
		}
	}
}*/
