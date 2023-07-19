package util

import (
	"common/defs"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/proto"
	"net/http"
)

type HttpProtobufOutput struct {
}

func NewHttpProtobufOutput() *HttpProtobufOutput {
	return &HttpProtobufOutput{}
}

func (this *HttpProtobufOutput) OutputSuccess(c *gin.Context, data proto.Message) {
	w := NewPacketWriter()
	w.WriteByte(0) // flag
	w.WriteU16(0)  // errCode
	if data != nil {
		buf, _ := proto.Marshal(data)
		w.WriteRawBytes(buf)
	}
	this._outputResult(c, w.Data())
}

func (this *HttpProtobufOutput) OutputFailure(c *gin.Context, code int16, msg string) {
	if code < 1 {
		code = defs.ErrInternal
	}
	w := NewPacketWriter()
	w.WriteByte(0) // flag
	w.WriteS16(code)
	w.WriteString(msg)
	this._outputResult(c, w.Data())
}

func (this *HttpProtobufOutput) OutputGrpcError(c *gin.Context, err error) {
	code, msg := FromGrpcError(err)
	this.OutputFailure(c, int16(code), msg)
}

func (this *HttpProtobufOutput) _outputResult(c *gin.Context, data []byte) {
	c.Render(http.StatusOK, httpRenderProtobuf{data: data})
}

// ----------------------------------------------------------------
var protobufContentType = []string{"application/octet-stream"}

type httpRenderProtobuf struct {
	data []byte
}

func (r httpRenderProtobuf) Render(w http.ResponseWriter) error {
	r.WriteContentType(w)
	_, err := w.Write(r.data)
	return err
}

func (r httpRenderProtobuf) WriteContentType(w http.ResponseWriter) {
	header := w.Header()
	header["Content-Type"] = jsonContentType
}
