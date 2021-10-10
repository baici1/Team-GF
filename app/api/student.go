package api

import (
	"errors"
	"strconv"
	"team-gf/app/model"
	"team-gf/app/service"
	"team-gf/library/code"
	"team-gf/library/response"
	"team-gf/library/snowflake"

	"github.com/gogf/gf/frame/g"

	"github.com/gogf/gf/errors/gerror"

	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
)

var Student = studentApi{}

type studentApi struct{}

// SignUp
// @summary 用户注册接口
// @tags    用户服务
// @produce json
// @param   entity  body model.StuApiSignUpReq true "注册请求"
// @router  /stu/signup [POST]
// @success 200 {object} response.JsonResponse "执行结果"
func (*studentApi) SignUp(r *ghttp.Request) {
	//创建注册请求的参数
	var (
		apiReq     *model.StuApiSignUpReq //注册请求参数
		serviceReq *model.StuServiceSignUpReq
	)
	//注册api请求输入
	if err := r.Parse(&apiReq); err != nil {
		g.Log().Error("r.Parse(&apiReq) failed", gerror.Current(err).Error())
		response.ResponseError(r, code.CodeInvalidParam, gerror.Current(err).Error())
	}
	//类型转换--将请求的参数转换成业务参数
	if err := gconv.Struct(apiReq, &serviceReq); err != nil {
		g.Log().Error(" gconv.Struct(apiReq, &serviceReq) failed", err.Error())
		response.ResponseError(r, code.CodeServerBusy)
	}
	//添加编号
	id := snowflake.GenID()
	serviceReq.ID = strconv.FormatInt(id, 10)
	//业务函数
	if err := service.User.SignUp(serviceReq); err != nil {
		g.Log().Error("service.User.SignUp failed", err.Error())
		//如果错误是ErrorUserExist，返回对应的错误信息
		if errors.Is(err, model.ErrorUserExist) {
			response.ResponseError(r, code.CodeUserExist)
		}
		response.ResponseError(r, code.CodeServerBusy)
	} else {
		response.ResponseSuccess(r, code.CodeSuccess, nil)
	}
}

// SignIn  学生用户登录功能
// @summary 用户登录接口
// @tags    用户服务
// @produce json
// @param   entity  body model.StuApiSignInReq true "注册请求"
// @router  /stu/signin [POST]
// @success 200 {object} response.JsonResponse "执行结果"
func (*studentApi) SignIn(r *ghttp.Request) {
	//创建登录请求的参数对象
	var (
		data *model.StuApiSignInReq
	)
	//获取登录请求的参数
	if err := r.Parse(&data); err != nil {
		g.Log().Error("获取参数 failed", err.Error())
		response.ResponseError(r, code.CodeInvalidParam, err.Error())
	}
	//业务处理函数
	token, err := service.User.SignIn(data)
	if err != nil {
		g.Log().Error("service.User.SignIn failed", err.Error())
		response.ResponseError(r, code.CodeInvalidPassword)
	}
	response.ResponseSuccess(r, code.CodeSuccess, g.Map{
		"token": token,
	})
}

func (*studentApi) Ping(r *ghttp.Request) {
	response.ResponseSuccess(r, code.CodeSuccess, r.GetParam(service.ContextUserIDKey))
}
