package logic

import (
	"common/proto/comm"
)

type FrameData struct {
	*comm.StatusFrame
	FrameType int
}

func (frame *FrameData) GetSeqId() int {
	return int(frame.Head.SeqId)
}
