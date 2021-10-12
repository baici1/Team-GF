package model

// TeamApiCreateTeamReq 学生创建队伍的参数请求
type TeamApiCreateTeamReq struct {
	Name      string `json:"name" v:"required#队伍名字不能为空"`
	Game      int64  `json:"game" v:"required#需要选择比赛"`
	Introduce string `json:"introduce" v:"required#队伍介绍不能为空"`
}

// TeamApiTeamAllDetailRes 返回队伍所有的相关信息（队长。队员，比赛等）
type TeamApiTeamAllDetailRes struct {
	Name      string `json:"name,omitempty"`
	Introduce string `json:"introduce,omitempty"`
	*Game     `json:"game,omitempty"`
	Leader    *StuApiGetDetailRes   `json:"leader"`
	Members   []*StuApiGetDetailRes `json:"members"`
}

// TeamApiAppendUserInTeamReq 用户进入队伍
type TeamApiAppendStuInTeamReq struct {
	Team    int64 `json:"team,omitempty" v:"required#请选择队伍"`
	Student int64 `json:"student,omitempty" v:"required#请选择需要邀请的学生"`
}

type TeamApiRemoveStuAtTeamReq struct {
	Team    int64 `json:"team,omitempty" v:"required#请选择队伍"`
	Student int64 `json:"student,omitempty" v:"required#请选择需要删除的学生"`
}
