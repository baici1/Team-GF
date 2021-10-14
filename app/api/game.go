package api

import (
	"errors"
	"team-gf/app/model"
	"team-gf/app/service"
	"team-gf/library/code"
	"team-gf/library/response"

	"github.com/gogf/gf/net/ghttp"
)

// Game 管理Game相关的服务函数
var Game gameApi

type gameApi struct{}

// GetAllGamesDetail 获取比赛所有信息
func (*gameApi) GetAllGamesDetail(r *ghttp.Request) {
	var (
		apiRes []*model.Game
	)
	apiRes, err := service.Game.GetAllGamesDetail()
	if err != nil {
		if errors.Is(err, model.ErrorQueryDataEmpty) {
			response.ResponseError(r, code.CodeQueryDataEmpty)
		}
		response.ResponseError(r, code.CodeServerBusy)
	}
	response.ResponseSuccess(r, code.CodeSuccess, apiRes)
}
