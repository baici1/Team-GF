package response

import (
	"team-gf/library/code"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

// JsonResponse 数据返回的的JSON数据结构
type JsonResponse struct {
	Code code.ResCode `json:"code"`           //错误码
	Msg  interface{}  `json:"msg"`            //提示信息
	Data interface{}  `json:"data,omitempty"` //返回数据
}

// ResponseError 返回自定义错误信息（可以选择带信息，或者利用code自带的信息）
func ResponseError(r *ghttp.Request, code code.ResCode, msg ...interface{}) {
	rd := &JsonResponse{
		Code: code,
		Msg:  code.Msg(),
		Data: nil,
	}
	if len(msg) > 0 {
		rd.Msg = msg[0]
	}
	if err := r.Response.WriteJson(rd); err != nil {
		g.Log().Debug("请求输出数据返回失败", err.Error())
	}
	r.Exit()
}

// ResponseSuccess 返回请求成功的数据
func ResponseSuccess(r *ghttp.Request, code code.ResCode, data ...interface{}) {

	rd := &JsonResponse{
		Code: code,
		Msg:  code.Msg(),
		Data: data,
	}
	if len(data) > 0 {
		rd.Data = data[0]
	}
	if err := r.Response.WriteJson(rd); err != nil {
		g.Log().Debug("请求输出数据返回失败", err.Error())
	}
	r.Exit()
}
