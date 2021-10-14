package mysql

import "team-gf/app/dao"

// MysqlUtils 这里都是一些常用的处理机制（处理数据）
var MysqlUtils mysqlUtils

type mysqlUtils struct{}

// GetGender 处理性别
func (*mysqlUtils) GetGender(gender string) string {
	if gender == "0" {
		return "女"
	} else {
		return "男"
	}
}

// GetGame 根据比赛id获取比赛名
func (*mysqlUtils) GetGame(game string) (string, error) {
	v, err := dao.Game.DB().Model("game").Fields("name").Where("id", game).Value()
	if err != nil {
		return "", err
	}
	return v.String(), nil
}

// GetTeacher 根据老师id获取老师名字
func (*mysqlUtils) GetTeacher(teacher string) (string, error) {
	v, err := dao.Teacher.DB().Model("teacher").Fields("name").Where("id", teacher).Value()
	if err != nil {
		return "", err
	}
	return v.String(), nil
}

// GetLeader 根据leaderId获取队长名字
func (*mysqlUtils) GetLeader(leader string) (string, error) {
	v, err := dao.Student.DB().Model("student").Fields("name").Where("id", leader).Value()
	if err != nil {
		return "", err
	}
	return v.String(), nil
}
