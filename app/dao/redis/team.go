package redis

import (
	"strconv"
	"team-gf/app/model"

	"github.com/gogf/gf/util/gconv"

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

// GetTeamMembers 获取队伍中的队员id集合
func (*teamDao) GetTeamMembers(teamId int64) ([]int64, error) {
	//获取相关key
	key := getRedisKey(KeyCreateTeamListPrefix) + strconv.FormatInt(teamId, 10)
	//获取除队长除外的成员id集合
	value, err := g.Redis().DoVar("lrange", key, 1, -1)
	if err != nil {
		return nil, err
	}
	return value.Int64s(), err
}

// PushStuInTeam 添加队员进入队伍
func (*teamDao) PushStuInTeam(req *model.TeamApiPushUserInTeamReq) error {
	key := getRedisKey(KeyCreateTeamListPrefix) + gconv.String(req.Team)
	_, err := g.Redis().DoVar("RPUSHX", key, req.Student)
	if err != nil {
		return err
	}
	return err
}
