package service

import (
	"database/sql"
	"team-gf/app/dao"
	"team-gf/app/dao/redis"
	"team-gf/app/model"

	"github.com/gogf/gf/frame/g"
)

// TeamCommon 管理学生相关user服务
var TeamCommon = teamService{}

type teamService struct{}

// TeamLeader 身份是leader
var TeamLeader = leaderService{}

type leaderService struct{}

// TeamMember 身份是member
var TeamMember = memberService{}

type memberService struct{}

// TeamVisitor 身份是visitor
var TeamVisitor = visitorService{}

type visitorService struct{}

// GetTeamAllDetail （任何人）查询队伍详细信息包括队员的信息，比赛信息,老师信息
func (*teamService) GetTeamAllDetail(teamId int64) (data *model.TeamApiTeamAllDetailRes, err error) {
	var (
		team    *model.Team
		leader  *model.StuApiGetDetailRes     //leader 返回队长的信息
		members []*model.StuApiGetDetailRes   //member 返回队员信息
		game    *model.Game                   //相关的比赛信息
		teacher *model.TeacherApiGetDetailRes //获取知道老师信息
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
	if err := dao.Student.DB().Model("student").Where("id", memberIds).Scan(&members); err != nil {
		return nil, err
	}
	//for _, memberid := range memberIds {
	//	member, err := User.GetUserData(memberid)
	//	if err != nil {
	//		return nil, err
	//	}
	//	members = append(members, member)
	//}
	//获取详细比赛信息
	game, err = Game.GetGameDetail(team.Game)
	if err != nil {
		return
	}
	//获取指导老师信息
	teacher, err = Teacher.GetTeacherDetail(team.Teacher)
	if err != nil {
		return
	}
	data = &model.TeamApiTeamAllDetailRes{
		Name:      team.Name,
		Introduce: team.Introduce,
		Game:      game,
		Leader:    leader,
		Members:   members,
		Teacher:   teacher,
	}
	return
}

// GetTeamsDetail 获取队伍信息
func (*teamService) GetTeamsDetail(gameId int64, stuId int64) (data []*model.TeamApiTeamsDetailRes, err error) {
	if gameId == 0 {
		if err := dao.Team.DB().Model("team").With(model.Student{}).Where("creator", stuId).Scan(&data); err != nil {
			if err == sql.ErrNoRows {
				return nil, model.ErrorQueryDataEmpty
			}
			return nil, err
		}
	} else {
		if err := dao.Team.DB().Model("team").With(model.Student{}).Where("creator=? and game=?", stuId, gameId).Scan(&data); err != nil {
			if err == sql.ErrNoRows {
				return nil, model.ErrorQueryDataEmpty
			}
			return nil, err
		}
	}
	return data, nil
}

// CreateOwnTeam 用户创建team
func (*leaderService) CreateOwnTeam(t *model.Team) error {
	//数据库添加队伍信息
	if _, err := dao.Team.DB().Model("team").Save(t); err != nil {
		return err
	}
	//创建队伍，redis增加队伍信息
	if err := redis.Team.CreateOwnTeam(t.Id, t.Creator); err != nil {
		return err
	}
	//为用户的队伍表添加队伍编号
	if err := redis.Team.UserInTeams(t.Id, t.Creator); err != nil {
		return err
	}
	return nil
}

// AppendStuInTeam 添加队员进入队伍，防止有重复队员进入队伍
func (*leaderService) AppendStuInTeam(req *model.TeamApiAppendStuInTeamReq) error {
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
	//用户的队伍表添加队伍编号
	if err := redis.Team.UserInTeams(req.Team, req.Student); err != nil {
		return err
	}
	return nil
}

// RemoveStuAtTeam leader删除队伍中的队员
func (*leaderService) RemoveStuAtTeam(req *model.TeamApiRemoveStuAtTeamReq) error {
	if err := redis.Team.RemoveStuAtTeam(req); err != nil {
		return err
	}
	return nil
}

// DeleteOwnTeam leader删除队伍
func (*leaderService) DeleteOwnTeam(teamid int64) error {
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

// LeaveToTeam 成员退出队伍
func (*memberService) LeaveToTeam(teamid int64, stuid int64) error {
	if err := redis.Team.LeaveToTeam(teamid, stuid); err != nil {
		return err
	}
	return nil
}
