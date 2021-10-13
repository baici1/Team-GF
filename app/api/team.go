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

// TeamCommon 公共组队相关的api
var TeamCommon = commonApi{}

type commonApi struct{}

// TeamLeader 身份是leader
var TeamLeader = leaderApi{}

type leaderApi struct{}

// TeamMember 身份是member
var TeamMember = memberApi{}

type memberApi struct{}

// TeamVisitor 身份是visitor
var TeamVisitor = visitorApi{}

type visitorApi struct{}

// GetTeamAllDetail （任何人）查询队伍信息,包括队员的信息，比赛信息
func (*commonApi) GetTeamAllDetail(r *ghttp.Request) {
	var teamId int64
	teamId = gconv.Int64(r.Get("teamId"))
	if teamId == 0 {
		g.Log().Debug("参数获取 teamId failed")
		response.ResponseError(r, code.CodeInvalidParam)
	}
	Res, err := service.TeamCommon.GetTeamAllDetail(teamId)
	if err != nil {
		if errors.Is(err, model.ErrorQueryDataEmpty) {
			response.ResponseError(r, code.CodeTeamNotExist)
		}
		response.ResponseError(r, code.CodeServerBusy)
	}
	response.ResponseError(r, code.CodeSuccess, Res)
}

// GetOwnTeams 用户查询自己的参加的队伍信息
func (*commonApi) GetOwnTeams(r *ghttp.Request) {
	var (
		ApiRes []*model.TeamApiTeamsDetailRes
	)
	//如果没有获取到gameId参数，则返回全部参加的队伍信息，如果有则返回参加的队伍信息
	//参数为0则代表全部队伍信息，如果>0则返回相对应的比赛队伍信息
	gameId := gconv.Int64(r.Get("game"))
	stuId := r.GetParam(service.ContextUserIDKey)
	ApiRes, err := service.TeamCommon.GetTeamsDetail(gameId, gconv.Int64(stuId))
	if err != nil {
		g.Log().Error(err.Error())
		response.ResponseError(r, code.CodeServerBusy)
	}
	response.ResponseSuccess(r, code.CodeSuccess, ApiRes)
}

// GetOtherGameTeams 显示不同比赛的队伍
func (*commonApi) GetOtherGameTeams(r *ghttp.Request) {
	//gameId := gconv.Int64(r.Get("game"))

}

// CreateOwnTeam 用户创建Team
// @summary 用户给比赛创建Team
// @tags    用户服务
// @produce json
// @param   entity  body model.StuApiCreateTeam true "提交信息请求"
// @router  /stu/create/team [POST]
// @success 200 {object} response.JsonResponse "执行结果"
func (*leaderApi) CreateOwnTeam(r *ghttp.Request) {
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
	if err := service.TeamLeader.CreateOwnTeam(serviceReq); err != nil {
		g.Log().Error("service.User.CreatOwnTeam failed", err.Error())
		response.ResponseError(r, code.CodeServerBusy)
	}
	response.ResponseSuccess(r, code.CodeSuccess)
}

// PushStuInTeam leader让用户加入队伍
func (*leaderApi) PushStuInTeam(r *ghttp.Request) {
	//创建参数对象
	var (
		apiReq *model.TeamApiAppendStuInTeamReq
	)
	//获取参数
	if err := r.Parse(&apiReq); err != nil {
		g.Log().Error("获取参数 model.TeamApiPushUserInTeamReq failed", gerror.Current(err).Error())
		response.ResponseError(r, code.CodeInvalidParam, gerror.Current(err).Error())
	}
	//不能添加队长进入队伍
	ownid := r.GetParam(service.ContextUserIDKey)
	if ownid == apiReq.Student {
		response.ResponseError(r, code.CodeRepeatUser)
	}
	//业务逻辑
	if err := service.TeamLeader.AppendStuInTeam(apiReq); err != nil {
		if errors.Is(err, model.ErrorRepeatUser) {
			response.ResponseError(r, code.CodeRepeatUser)
		}
		if errors.Is(err, model.ErrorTeamNotExist) {
			response.ResponseError(r, code.CodeTeamNotExist)
		}
		response.ResponseError(r, code.CodeServerBusy)
	}
	response.ResponseSuccess(r, code.CodeSuccess)
}

// RemoveUserAtTeam leader删除队伍中的用户
func (*leaderApi) RemoveUserAtTeam(r *ghttp.Request) {
	var (
		apiReq *model.TeamApiRemoveStuAtTeamReq
	)
	if err := r.Parse(&apiReq); err != nil {
		g.Log().Error("获取参数 model.TeamApiRemoveUserAtTeamReq  failed", gerror.Current(err).Error())
		response.ResponseError(r, code.CodeInvalidParam, gerror.Current(err).Error())
	}
	//业务逻辑函数
	if err := service.TeamLeader.RemoveStuAtTeam(apiReq); err != nil {
		g.Log().Error("service.Team.RemoveStuAtTeam", err.Error())
		response.ResponseError(r, code.CodeServerBusy)
	}
	response.ResponseSuccess(r, code.CodeSuccess)
}

// DeleteOwnTeam leader删除队伍
func (*leaderApi) DeleteOwnTeam(r *ghttp.Request) {
	var teamId int64
	teamId = gconv.Int64(r.Get("team"))
	if teamId == 0 {
		g.Log().Debug("参数获取 teamId failed")
		response.ResponseError(r, code.CodeInvalidParam)
	}
	if err := service.TeamLeader.DeleteOwnTeam(teamId); err != nil {
		g.Log().Error("service.Team.DeleteOwnTeam failed", err.Error())
		response.ResponseError(r, code.CodeServerBusy)
	}
	response.ResponseSuccess(r, code.CodeSuccess)
}

// LeaveToTeam 用户离开队伍
func (*memberApi) LeaveToTeam(r *ghttp.Request) {
	var teamId int64
	//获取请求参数
	teamId = gconv.Int64(r.Get("team"))
	if teamId == 0 {
		response.ResponseError(r, code.CodeInvalidParam)
	}
	stuId := gconv.Int64(r.GetParam(service.ContextUserIDKey))
	g.Log().Debug(stuId)
	if err := service.TeamMember.LeaveToTeam(teamId, stuId); err != nil {
		response.ResponseError(r, code.CodeServerBusy)
	}
	response.ResponseSuccess(r, code.CodeSuccess)
}
