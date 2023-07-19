package util

import (
	"google.golang.org/protobuf/encoding/protojson"
)

func NewProtoMsgMarshaler() *protojson.MarshalOptions {
	return &protojson.MarshalOptions{
		EmitUnpopulated: true,
		UseEnumNumbers:  true,
		UseProtoNames:   true,
		AllowPartial:    true,
	}
}

/*
func ProtoMsgMarshalerTest() {
	m := NewProtoMsgMarshaler()
	r := user.CountryInfoResp{}
	b, err := m.Marshal(&r)
	fmt.Println(string(b), err)

	r2 := user.CountryReq{}
	b2, err := m.Marshal(&r2)
	fmt.Println(string(b2), err)
}
*/
