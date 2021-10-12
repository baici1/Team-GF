package api

import (
	"errors"
	"team-gf/app/model"
	"team-gf/app/service"
	"team-gf/library/code"
	"team-gf/library/response"
	"team-gf/library/snowflake"

	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
)

var Team = teamApi{}

type teamApi struct{}

// CreateOwnTeam 用户创建Team
// @summary 用户给比赛创建Team
// @tags    用户服务
// @produce json
// @param   entity  body model.StuApiCreateTeam true "提交信息请求"
// @router  /stu/create/team [POST]
// @success 200 {object} response.JsonResponse "执行结果"
func (*teamApi) CreateOwnTeam(r *ghttp.Request) {
	//创建请求得参数
	var (
		apiReq     *model.TeamApiCreateTeamReq
		serviceReq *model.Team
	)
	//获取请求中得参数
	if err := r.Parse(&apiReq); err != nil {
		g.Log().Error("获取参数 model.TeamApiCreateTeam failed", err.Error())
		response.ResponseError(r, code.CodeInvalidParam, gerror.Current(err).Error())
	}
	//类型转换-方便后续业务函数使用
	if err := gconv.Struct(apiReq, &serviceReq); err != nil {
		g.Log().Error("gconv.Struct failed", err.Error())
		response.ResponseError(r, code.CodeServerBusy)
	}
	//业务逻辑函数
	serviceReq.Id = snowflake.GenID()
	serviceReq.Creator = r.GetParam(service.ContextUserIDKey).(int64)
	if err := service.Team.CreateOwnTeam(serviceReq); err != nil {
		g.Log().Error("service.User.CreatOwnTeam failed", err.Error())
		response.ResponseError(r, code.CodeServerBusy)
	}
	response.ResponseSuccess(r, code.CodeSuccess)
}

// GetTeamAllDetail （任何人）查询队伍信息,包括队员的信息，比赛信息
func (*teamApi) GetTeamAllDetail(r *ghttp.Request) {
	var teamId int64
	teamstr := r.Get("teamId").(string)
	teamId = gconv.Int64(teamstr)
	//if err := r.GetParam(); err != nil {
	//	g.Log().Error("获取参数 teamId failed", err.Error())
	//	response.ResponseError(r, code.CodeInvalidParam)
	//}
	g.Log().Debug(teamstr)
	Res, err := service.Team.GetTeamAllDetail(teamId)
	if err != nil {
		response.ResponseError(r, code.CodeServerBusy)
	}
	response.ResponseError(r, code.CodeSuccess, Res)
}

// PushStuInTeam 用户加入队伍
func (*teamApi) PushStuInTeam(r *ghttp.Request) {
	//创建参数对象
	var (
		apiReq *model.TeamApiPushUserInTeamReq
	)
	//获取参数
	if err := r.Parse(&apiReq); err != nil {
		g.Log().Error("获取参数 model.TeamApiPushUserInTeamReq failed", err.Error())
		response.ResponseError(r, code.CodeInvalidParam, gerror.Current(err).Error())
	}
	//不能添加队长进入队伍
	ownid := r.GetParam(service.ContextUserIDKey)
	if ownid == apiReq.Student {
		response.ResponseError(r, code.CodeRepeatUser)
	}
	//业务逻辑
	if err := service.Team.PushStuInTeam(apiReq); err != nil {
		if errors.Is(err, model.ErrorRepeatUser) {
			response.ResponseError(r, code.CodeRepeatUser)
		}
		response.ResponseError(r, code.CodeServerBusy)
	}
	response.ResponseSuccess(r, code.CodeSuccess)
}

// RemoveUserAtTeam leader删除队伍中的用户
func (*teamApi) RemoveUserAtTeam(r *ghttp.Request) {

}

//删除队伍
