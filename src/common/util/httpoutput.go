package util

import (
	"common/defs"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"io"
	"net/http"
	"strconv"
	"strings"
)

type HttpOutput struct {
	marshaler *protojson.MarshalOptions
}

func NewHttpOutput() *HttpOutput {
	marshaler := NewProtoMsgMarshaler()
	return &HttpOutput{marshaler: marshaler}
}

type outResult struct {
	Code int32       `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func (this *HttpOutput) OutputSuccess(c *gin.Context, data interface{}) {
	this._outputObjResult(c, outResult{
		Code: 0,
		Msg:  "OK",
		Data: data,
	})
}

func (this *HttpOutput) OutputFailure(c *gin.Context, code int32, msg string) {
	this._outputObjResult(c, outResult{
		Code: code,
		Msg:  msg,
	})
}

func (this *HttpOutput) OutputGrpcError(c *gin.Context, err error) {
	code, msg := FromGrpcError(err)
	this._outputObjResult(c, outResult{
		Code: code,
		Msg:  msg,
	})
}

func (this *HttpOutput) OutputProtoMsg(c *gin.Context, msg proto.Message) {
	var b strings.Builder

	b.WriteString(`{"code":0,"msg":"OK","data":`)

	if msg == nil {
		b.WriteString("{}")
	} else {
		bstr, err := this.marshaler.Marshal(msg)
		if err != nil {
			this.OutputFailure(c, defs.ErrInternal, err.Error())
			return
		}
		b.Write(bstr)
	}
	b.WriteString("}")

	this._outputResult(c, b.String())
}

func (this *HttpOutput) OutputObjString(c *gin.Context, str string) {
	this._outputResult(c, `{"code":0,"msg":"OK","data":`+str+"}")
}

func (this *HttpOutput) _outputObjResult(c *gin.Context, res outResult) {
	d, _ := json.Marshal(res)
	this._outputResult(c, string(d))
}

func (this *HttpOutput) _outputResult(c *gin.Context, res string) {
	key := c.GetString(defs.CtxKeyCryptoKey)
	if key == "" {
		c.Render(http.StatusOK, httpRenderJson{data: res})
	} else {
		d, err := Encrypt([]byte(res), key)
		if err == nil {
			c.Render(http.StatusOK, httpRenderText{data: d})
		} else {
			c.Render(http.StatusOK, httpRenderJson{data: `{"code":` +
				strconv.FormatInt(defs.ErrNetworkAuthRequired, 10) + `,"msg":"response encrypt error"`})
		}
	}
}

// ----------------------------------------------------------------
var jsonContentType = []string{"application/json; charset=utf-8"}
var textContentType = []string{"text/plain; charset=utf-8"}

type httpRenderJson struct {
	data string
}

func (r httpRenderJson) Render(w http.ResponseWriter) error {
	r.WriteContentType(w)
	_, err := io.WriteString(w, r.data)
	return err
}

func (r httpRenderJson) WriteContentType(w http.ResponseWriter) {
	header := w.Header()
	header["Content-Type"] = jsonContentType
}

type httpRenderText struct {
	data string
}

func (r httpRenderText) Render(w http.ResponseWriter) error {
	r.WriteContentType(w)
	_, err := io.WriteString(w, r.data)
	return err
}

func (r httpRenderText) WriteContentType(w http.ResponseWriter) {
	header := w.Header()
	header["Content-Type"] = textContentType
}
