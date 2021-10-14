package api

import (
	"errors"
	"team-gf/app/model"
	"team-gf/app/service"
	"team-gf/library/code"
	"team-gf/library/response"

	"github.com/gogf/gf/net/ghttp"
)

// Teacher teacher管理相关老师服务
var Teacher teacherApi

type teacherApi struct{}

// GetAllTeachersDetail 获取全部老师信息
func (*teacherApi) GetAllTeachersDetail(r *ghttp.Request) {
	var (
		apiRes []*model.TeacherApiGetDetailRes
	)
	apiRes, err := service.Teacher.GetAllTeachersDetail()
	if err != nil {
		if errors.Is(err, model.ErrorQueryDataEmpty) {
			response.ResponseError(r, code.CodeQueryDataEmpty)
		}
		response.ResponseError(r, code.CodeServerBusy)
	}
	response.ResponseSuccess(r, code.CodeSuccess, apiRes)
}
