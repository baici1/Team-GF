package service

import (
	"database/sql"
	"team-gf/app/dao"
	"team-gf/app/model"
)

// Game 管理学生相关user服务
var Game = gameService{}

type gameService struct{}

// GetGameDetail 获取比赛详细信息
func (*gameService) GetGameDetail(gameId int64) (data *model.Game, err error) {
	if err = dao.Game.DB().Model("game").Where("id", gameId).Scan(&data); err != nil {
		return
	}
	return
}

// GetAllGamesDetail 获取比赛全部详细信息
func (*gameService) GetAllGamesDetail() (data []*model.Game, err error) {
	if err = dao.Game.DB().Model("game").Scan(&data); err != nil {
		if err == sql.ErrNoRows {
			return
		}
		return
	}
	return data, nil
}
