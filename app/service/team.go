package service

import (
	"database/sql"
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
		team    *model.Team
		leader  *model.StuApiGetDetailRes   //leader 返回队长的信息
		members []*model.StuApiGetDetailRes //member 返回队员信息
		game    *model.Game                 //相关的比赛信息
	)
	team = new(model.Team) //对应的team的信息
	//获取队伍所有信息，部分信息供后续使用
	if err = dao.Team.DB().Model("team").Where("id", teamId).Scan(team); err != nil {
		if err == sql.ErrNoRows {
			return nil, model.ErrorQueryDataEmpty
		}
		return
	}
	//查询leader的信息
	leader, err = User.GetUserData(team.Creator)
	if err != nil {
		g.Log().Debug("err", err)
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

// AppendStuInTeam 添加队员进入队伍，防止有重复队员进入队伍
func (*teamService) AppendStuInTeam(req *model.TeamApiAppendStuInTeamReq) error {
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
	if err := redis.Team.AppendStuInTeam(req); err != nil {
		return err
	}
	return nil
}

// RemoveStuAtTeam leader删除队伍中的队员
func (*teamService) RemoveStuAtTeam(req *model.TeamApiRemoveStuAtTeamReq) error {
	if err := redis.Team.RemoveStuAtTeam(req); err != nil {
		return err
	}
	return nil
}

// DeleteOwnTeam leader删除队伍
func (*teamService) DeleteOwnTeam(teamid int64) error {
	//从redis删除队伍信息
	if err := redis.Team.DeleteOwnTeam(teamid); err != nil {
		return err
	}
	//从mysql删除队伍信息
	if _, err := dao.Team.DB().Model("team").Where("id", teamid).Delete(); err != nil {
		return err
	}
	return nil
}
