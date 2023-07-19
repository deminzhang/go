package util

import "github.com/golang/snappy"

func GetProtoMsg(msg []byte) (int16, []byte) {
	if len(msg) < 5 {
		return -1, nil
	}

	var flag = msg[0]
	if flag&0x01 != 0 { //压缩包解压
		dmsg, err := snappy.Decode(nil, msg[1:])
		if err != nil {
			return -1, []byte(err.Error())
		}
		msg = dmsg
	} else {
		msg = msg[1:]
	}
	var op int16 //协议号
	op = int16(msg[0])<<8 + int16(msg[1])

	var errCode int16 //错误码
	errCode = int16(msg[2]<<8) + int16(msg[3])
	if errCode != 0 {
		return -errCode, append(make([]byte, 0), msg[4:]...)
	}
	var protoMsg []byte //proto数据结构
	protoMsg = append(protoMsg, msg[4:]...)
	return op, protoMsg
}
