package world

import (
	utilc "client0/util"
	"common/util"
	"fmt"
)

type FrameComponent struct {
	queue         *utilc.FrameQueue
	tailSeqId     int
	headSeqId     int
	assert        bool
	nextSendSeqId int
	fn            func(headE, dailE utilc.IFrameElement) bool
}

type Option func(component *FrameComponent)

func WithAssert() Option {
	return func(component *FrameComponent) {
		component.assert = true
	}
}

func NewFrameComponent(options ...Option) *FrameComponent {
	component := &FrameComponent{}
	for _, o := range options {
		o(component)
	}

	if component.queue == nil {
		component.queue = utilc.NewFrameQueue(256)
	}
	return component
}

func (this *FrameComponent) HeadSeqId() int {
	return this.headSeqId
}

func (this *FrameComponent) TailSeqId() int {
	return this.tailSeqId
}

func (this *FrameComponent) Capacity() int {
	return this.queue.Capacity()
}

func (this *FrameComponent) NextSendSeqId() int {
	return this.nextSendSeqId
}

func (this *FrameComponent) SetNextSendSeqId(seqId int) {
	this.nextSendSeqId = seqId
}

func (this *FrameComponent) ConfirmSeqId(seqId int) bool {
	if this.queue.CanReWriter(seqId) {
		this.headSeqId = seqId + 1
		this.queue.ResetHeadIndex(this.headSeqId)
		return true
	}

	if this.assert {
		util.AssertTrue(false, "ConfirmSeqId :%d", seqId)
	}
	return false
}

// 回拨confirm seq id
func (this *FrameComponent) DailBack(seqId int) bool {
	element := this.queue.GetElement(seqId)
	if element == nil {
		if this.assert {
			util.AssertTrue(
				false,
				fmt.Sprintf("ResetConfirmSeqId :%d not allowed, because seqId data is nil", seqId))
		}
		return false
	}
	headElement := this.queue.GetElement(this.headSeqId)
	if this.fn != nil {
		var headFrameData utilc.IFrameElement
		if headElement != nil {
			headFrameData = headElement.(utilc.IFrameElement)
		}
		if !this.fn(headFrameData, element.(utilc.IFrameElement)) {
			return false
		}
	}
	this.headSeqId = seqId
	this.queue.ResetHeadIndex(seqId)
	return true
}

func (this *FrameComponent) GetElement(index int) utilc.IFrameElement {
	if index < 0 {
		return nil
	}
	elem := this.queue.GetElement(index)
	if elem != nil {
		return elem
	}
	return nil
}

func (this *FrameComponent) Len() int {
	return this.tailSeqId - this.headSeqId
}

func (this *FrameComponent) TraceStart(index int, fn func(element utilc.IFrameElement) bool) {
	this.queue.TraceStart(index, fn)
}

// 初始化队列index, 断线重连后使用
func (this *FrameComponent) ResetQueue(nextConfirmSeqId int) {
	this.headSeqId = nextConfirmSeqId
	this.tailSeqId = nextConfirmSeqId
	this.nextSendSeqId = nextConfirmSeqId
	this.queue.InitIndex(nextConfirmSeqId)
}

func (this *FrameComponent) ReplaceOrEnqueue(data utilc.IFrameElement) bool {
	if !this.Replace(data) {
		return this.queue.Enqueue(data)
	}
	return true
}

func (this *FrameComponent) Replace(data utilc.IFrameElement) bool {
	return this.queue.Replace(data.GetSeqId(), data)
}

func (this *FrameComponent) TailElement() utilc.IFrameElement {
	if this.queue.IsEmpty() {
		return nil
	}
	return this.queue.TailElement().(utilc.IFrameElement)
}

func (this *FrameComponent) HeadElement() utilc.IFrameElement {
	if this.queue.IsEmpty() {
		return nil
	}
	return this.queue.HeadElement().(utilc.IFrameElement)
}

func (this *FrameComponent) EnQueue(data utilc.IFrameElement) bool {
	r := this.queue.Enqueue(data)
	if !r {
		if this.assert {
			util.AssertTrue(false, "EnQueue")
		}
		return false
	}
	this.tailSeqId++
	return true
}

func (this *FrameComponent) IsFull() bool {
	return this.queue.IsFull()
}

func (this *FrameComponent) IsEmpty() bool {
	return this.queue.IsEmpty()
}

func (this *FrameComponent) Clear() {
	this.queue.Clear()
}
