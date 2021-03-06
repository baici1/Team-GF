package model

import "errors"

//关于mysql
var (
	ErrorUserExist           = errors.New("用户已存在！")
	ErrorUserNotExist        = errors.New("用户不存在")
	ErrorQueryFailedUser     = errors.New("查询用户失败")
	ErrorInvalidUser         = errors.New("学号或密码错误")
	ErrorInvalidID           = errors.New("无效的ID")
	ErrorGenerateTokenFailed = errors.New("生成token失败")
	ErrorWriteData           = errors.New("写入信息出错")
	ErrorRepeatUser          = errors.New("有重复用户")
	ErrorQueryDataEmpty      = errors.New("查询信息不存在")
	ErrorTeamNotExist        = errors.New("队伍不存在")
)
