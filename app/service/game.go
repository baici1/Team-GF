package service

import (
	"team-gf/app/dao"
	"team-gf/app/model"
)

// Game 管理学生相关user服务
var Game = gameService{}

type gameService struct{}

// GetGameDetail 获取比赛详细信息
func (*gameService) GetGameDetail(gameid int64) (data *model.Game, err error) {
	if err = dao.Game.DB().Model("game").Where("id", gameid).Scan(&data); err != nil {
		return
	}
	return
}
