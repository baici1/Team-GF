package code

type ResCode int64

/*
封装所有的请求码以及提示信息
*/

const (
	CodeSuccess         ResCode = 1000 + iota //成功
	CodeInvalidParam                          //请求参数出错
	CodeUserExist                             //用户已存在
	CodeUserNoExist                           //用户不存在
	CodeInvalidPassword                       //学号或密码错误
	CodeServerBusy                            //服务繁忙
	CodeInvalidToken                          //token无效
	CodeNeedLogin                             //需要登录
	CodeOperation                             //错误操作
	CodeRepeatUser                            //重复用户
)

var codeMsgMap = map[ResCode]string{
	CodeSuccess:         "success",
	CodeInvalidParam:    "请求参数出错",
	CodeUserExist:       "用户已存在,",
	CodeUserNoExist:     "用户不存在",
	CodeInvalidPassword: "用户名或密码错误",
	CodeServerBusy:      "服务器繁忙",
	CodeNeedLogin:       "需要登录",
	CodeInvalidToken:    "无效token",
	CodeOperation:       "操作错误",
	CodeRepeatUser:      "请勿添加重复用户",
}

// Msg 返回错误码的对应的提示信息
func (c ResCode) Msg() string {
	msg, ok := codeMsgMap[c]
	if !ok {
		msg = codeMsgMap[CodeServerBusy]
	}
	return msg
}
