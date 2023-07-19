package util

import (
	"sync"
)

type ConcurrentQueue[T any] struct {
	data *List[T]
	sync.Mutex
}

func NewConcurrentQueue[T any]() *ConcurrentQueue[T] {
	return &ConcurrentQueue[T]{
		data: NewList[T](),
	}
}

func (this *ConcurrentQueue[T]) Push(value T) {
	this.Lock()
	defer this.Unlock()
	this.data.PushBack(value)
}

func (this *ConcurrentQueue[T]) Front() *Element[T] {
	this.Lock()
	defer this.Unlock()
	return this.data.Front()
}

func (this *ConcurrentQueue[T]) Pop() *Element[T] {
	this.Lock()
	defer this.Unlock()
	e := this.data.Front()
	if e != nil {
		this.data.Remove(e)
	}
	return e
}

func (this *ConcurrentQueue[T]) ExportWithoutLock(q *ConcurrentQueue[T]) {
	this.data.PushBackList(q.data)
}

func (this *ConcurrentQueue[T]) Export(q *ConcurrentQueue[T]) {
	this.Lock()
	data := this.data
	this.data = NewList[T]()
	this.Unlock()
	data.PushBackList(q.data)
}

func (this *ConcurrentQueue[T]) ExportQueue(q *Queue[T]) {
	this.Lock()
	data := this.data
	this.data = NewList[T]()
	this.Unlock()
	q.data.PushBackList(data)
}

func (this *ConcurrentQueue[T]) IsEmpty() bool {
	this.Lock()
	defer this.Unlock()
	return this.data.Len() == 0
}

func (this *ConcurrentQueue[T]) Len() int {
	this.Lock()
	defer this.Unlock()
	return this.data.Len()
}
