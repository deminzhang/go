package util

import (
	"common/defs"
	"encoding/json"
	"net/http"
)

type HttpOutResult struct {
	Code int32           `json:"code"`
	Msg  string          `json:"msg"`
	Data json.RawMessage `json:"data"`
}

func HttpCall(client *http.Client, url string, params interface{}, ret interface{}) error {
	var resp HttpOutResult
	err := HttpPost(client, url, nil, params, &resp)
	if err != nil {
		return MakeGrpcError(defs.ErrInternal, err.Error())
	}
	if resp.Code != 0 {
		return MakeGrpcError(resp.Code, resp.Msg)
	}

	if resp.Data != nil {
		if err = json.Unmarshal(resp.Data, ret); err != nil {
			return MakeGrpcError(defs.ErrInternal, err.Error())
		}
	}
	return nil
}
