package service

import (
	"team-gf/app/dao"
	"team-gf/app/dao/redis"
	"team-gf/app/model"

	"github.com/gogf/gf/frame/g"
)

// Team 管理学生相关user服务
var Team = teamService{}

type teamService struct{}

// CreateOwnTeam 用户创建team
func (*teamService) CreateOwnTeam(t *model.Team) error {
	if _, err := dao.Team.DB().Model("team").Save(t); err != nil {
		return err
	}
	if err := redis.Team.CreateOwnTeam(t.Id, t.Creator); err != nil {
		return err
	}
	return nil
}

// GetTeamAllDetail （任何人）查询队伍详细信息包括队员的信息，比赛信息
func (*teamService) GetTeamAllDetail(teamId int64) (data *model.TeamApiTeamAllDetailRes, err error) {
	var (
		team    *model.Team                 //对应的team的信息
		leader  *model.StuApiGetDetailRes   //leader 返回队长的信息
		members []*model.StuApiGetDetailRes //member 返回队员信息
		game    *model.Game                 //相关的比赛信息
	)
	g.Log().Debug(teamId)
	//获取队伍所有信息，部分信息供后续使用
	if err = dao.Team.DB().Model("team").Where("id", teamId).Scan(&team); err != nil {
		g.Log().Error(err)
		return
	}
	g.Log().Debug(team.Creator)
	//查询leader的信息
	leader, err = User.GetUserData(team.Creator)
	if err != nil {
		return
	}
	//批量获取队伍的队员id
	memberIds, err := redis.Team.GetTeamMembers(teamId)
	if err != nil {
		return
	}
	for _, memberid := range memberIds {
		member, err := User.GetUserData(memberid)
		if err != nil {
			return nil, err
		}
		members = append(members, member)
	}
	//获取详细比赛信息
	game, err = Game.GetGameDetail(team.Game)
	if err != nil {
		return
	}
	data = &model.TeamApiTeamAllDetailRes{
		Name:      team.Name,
		Introduce: team.Introduce,
		Game:      game,
		Leader:    leader,
		Members:   members,
	}
	return
}

// PushStuInTeam 添加队员进入队伍，防止有重复队员进入队伍
func (*teamService) PushStuInTeam(req *model.TeamApiPushUserInTeamReq) error {
	//查询队伍中是否有重复的学生
	stus, err := redis.Team.GetTeamMembers(req.Team)
	if err != nil {
		return err
	}
	for _, stu := range stus {
		if stu == req.Student {
			return model.ErrorRepeatUser
		}
	}
	if err := redis.Team.PushStuInTeam(req); err != nil {
		return err
	}
	return nil
}
