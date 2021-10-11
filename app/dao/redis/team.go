package redis

import (
	"strconv"

	"github.com/gogf/gf/frame/g"
)

var Team = teamDao{}

type teamDao struct{}

// CreateOwnTeam 创建队伍，选择比赛，同时有队伍id。默认创建者是队长.list第一个是队长
func (*teamDao) CreateOwnTeam(teamId, userId int64) error {
	//生成创建队伍的key //由比赛和队伍编号组成
	key := getRedisKey(KeyCreateTeamListPrefix) + strconv.FormatInt(teamId, 10)
	g.Log().Debug("redis-key", key)
	_, err := g.Redis().Do("lpush", key, userId)
	return err
}
