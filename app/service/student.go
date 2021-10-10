package service

import (
	"team-gf/app/dao"
	"team-gf/app/model"
	"team-gf/library/jwt"

	"github.com/gogf/gf/frame/g"
)

// User 管理学生相关user服务
var User = userService{}

type userService struct{}

// SignUp 学生用户注册功能
func (u *userService) SignUp(s *model.StuServiceSignUpReq) error {
	//学号唯一性检验
	if !u.CheckStuID(s.Stuid) {
		return model.ErrorUserExist
	}
	if _, err := dao.Student.DB().Model("student").Save(s); err != nil {
		g.Log().Error("存储学生信息发生错误", err.Error())
		return err
	}
	return nil
}

// CheckStuID 学号唯一性检验存在返回false,否则true
func (u *userService) CheckStuID(stuid string) bool {
	if i, err := dao.Student.DB().Model("student").FindCount("stuid", stuid); err != nil {
		return false
	} else {
		return i == 0
	}
}

// 学生登录功能
func (u *userService) SignIn(s *model.StuApiSignInReq) (string, error) {
	//创建学生信息对象
	var stu *model.Student
	//查询学生信息
	err := dao.Student.DB().Model("student").Where("stuid=? and password=?", s.Stuid, s.Password).Scan(&stu)
	if err != nil {
		g.Log().Error("查询学生信息发生错误!", err.Error())
		return "", model.ErrorQueryFailedUser
	}
	//如果查询结果为nil，那么账号和密码发生错误
	if stu == nil {
		return "", model.ErrorInvalidUser
	}
	//返回生成的token信息
	token, err := jwt.GenerateToken(stu.Id)
	if err != nil {
		return "", model.ErrorGenerateTokenFailed
	}
	return token, nil
}
