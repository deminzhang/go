package util

type IFrameElement interface {
	GetSeqId() int
}

// 循环队列结构定义
type FrameQueue struct {
	q        []any
	capacity int
	head     int
	tail     int
}

// 新建队列
func NewFrameQueue(n int) *FrameQueue {
	if n == 0 {
		return nil
	}
	return &FrameQueue{
		q:        make([]any, n),
		capacity: n,
		head:     0,
		tail:     0,
	}
}

func (queue *FrameQueue) Capacity() int {
	return queue.capacity
}

// 队列判空
func (queue *FrameQueue) IsEmpty() bool {
	if queue.head == queue.tail {
		return true
	}
	return false
}

// 队列判满
func (queue *FrameQueue) IsFull() bool {
	if queue.head == (queue.tail+1)%queue.capacity {
		return true
	}
	return false
}

func (queue *FrameQueue) Len() int {
	if queue.tail >= queue.head {
		return queue.tail - queue.head
	} else {
		return queue.tail + queue.capacity - queue.head
	}
}

func (queue *FrameQueue) CanReWriter(index int) bool {
	index = index % queue.capacity
	if queue.tail > queue.head {
		if index >= queue.head && index < queue.tail {
			return true
		}
		return false
	} else {
		if index < queue.tail || index >= queue.head {
			return true
		}
		return false
	}
}

func (queue *FrameQueue) Replace(index int, v IFrameElement) bool {
	index = index % queue.capacity
	if queue.CanReWriter(index) {
		queue.q[index] = v
		return true
	}
	return false
}

func (queue *FrameQueue) Enqueue(v IFrameElement) bool {
	if queue.IsFull() {
		return false
	}

	if v.GetSeqId()%queue.capacity != queue.tail {
		return false
	}

	queue.q[queue.tail] = v
	queue.tail = (queue.tail + 1) % queue.capacity
	return true
}

func (queue *FrameQueue) Trace(fn func(IFrameElement) bool) {
	size := queue.Len()
	start := queue.head
	for i := 0; i < size; i++ {
		data := queue.q[start]
		if fn(data.(IFrameElement)) {
			break
		}
		start = (start + 1) % queue.capacity
	}
}

func (queue *FrameQueue) TraceStart(start int, fn func(IFrameElement) bool) {
	if queue.IsEmpty() {
		return
	}
	start = start % queue.capacity
	size := -1
	if queue.tail > queue.head {
		if start >= queue.head && start < queue.tail {
			size = queue.tail - start
		}
	} else {
		if start < queue.tail || start >= queue.head {
			if start < queue.tail {
				size = queue.tail - start
			} else {
				size = queue.tail + queue.capacity - start
			}
		}
	}

	for i := 0; i < size; i++ {
		data := queue.q[start]
		if fn(data.(IFrameElement)) {
			break
		}
		start = (start + 1) % queue.capacity
	}
}

func (queue *FrameQueue) HeadIndex() int {
	return queue.head
}

func (queue *FrameQueue) TailIndex() int {
	return queue.tail
}

// 回溯
func (queue *FrameQueue) ResetHeadIndex(index int) {
	head := index % queue.capacity
	queue.head = head
}

func (queue *FrameQueue) GetElement(index int) IFrameElement {
	index = index % queue.capacity
	element := queue.q[index]
	if element != nil {
		return element.(IFrameElement)
	}
	return nil
}

func (queue *FrameQueue) InitIndex(index int) {
	index = index % queue.capacity
	queue.head = index
	queue.tail = index
}

func (queue *FrameQueue) TailElement() IFrameElement {
	tail := (queue.tail - 1 + queue.capacity) % queue.capacity
	element := queue.q[tail]
	if element != nil {
		return element.(IFrameElement)
	}
	return nil
}

func (queue *FrameQueue) HeadElement() IFrameElement {
	element := queue.q[queue.head]
	if element != nil {
		return element.(IFrameElement)
	}
	return nil
}

func (queue *FrameQueue) Clear() {
	queue.head = queue.tail
}
