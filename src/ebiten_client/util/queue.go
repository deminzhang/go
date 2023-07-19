package util

type Queue[T any] struct {
	data *List[T]
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		data: NewList[T](),
	}
}

func (this *Queue[T]) Push(value T) {
	this.data.PushBack(value)
}

func (this *Queue[T]) Pop() *Element[T] {
	e := this.data.Front()
	if e != nil {
		this.data.Remove(e)
	}
	return e
}

func (this *Queue[T]) Export(q *ConcurrentQueue[T]) {
	this.data.PushBackList(q.data)
}
