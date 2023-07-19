package util

import (
	"common/defs"
	"common/tlog"
	"encoding/binary"
	"google.golang.org/protobuf/proto"
)

func BuildRawSuccessResponse(opCode int16, msg []byte) []byte {
	w := NewPacketWriter()
	w.WriteByte(0) // flags
	w.WriteS16(opCode)
	w.WriteU16(0) // errCode
	if len(msg) > 0 {
		w.WriteRawBytes(msg)
	}
	return w.Data()
}

func BuildSuccessResponse(opCode int16, msg proto.Message) []byte {
	w := NewPacketWriter()
	w.WriteByte(0) // flags
	w.WriteS16(opCode)
	w.WriteU16(0) // errCode
	if msg != nil {
		buf, _ := proto.Marshal(msg)
		w.WriteRawBytes(buf)
	}
	return w.Data()
}

func BuildFailureResponse(opCode int16, errCode int16, errMsg string) []byte {
	tlog.Errorf("fail response: op %d, err: %d, %s", opCode, errCode, errMsg)
	if errCode < 1 {
		errCode = defs.ErrInternal
	}
	w := NewPacketWriter()
	w.WriteByte(0) // flags
	w.WriteS16(opCode)
	w.WriteS16(errCode)
	w.WriteRawBytes(UnsafeStr2Bytes(errMsg))
	return w.Data()
}

func BuildGrpcErrorResponse(opCode int16, err error) []byte {
	errCode, errMsg := FromGrpcError(err)
	return BuildFailureResponse(opCode, int16(errCode), errMsg)
}

func ParseResponse(msg []byte, resp proto.Message) error {
	if len(msg) < 2 {
		return MakeGrpcError(defs.ErrInternal, "parse response failed")
	}

	errCode := binary.BigEndian.Uint16(msg)
	if errCode != 0 {
		return MakeGrpcError(int32(errCode), string(msg[2:]))
	} else if resp != nil {
		err := proto.Unmarshal(msg[2:], resp)
		if err != nil {
			return MakeGrpcError(defs.ErrInternal, err.Error())
		}
	}
	return nil
}
