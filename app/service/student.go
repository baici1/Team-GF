package service

import (
	"team-gf/app/dao"
	"team-gf/app/model"
	"team-gf/library/jwt"
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

// SignIn 学生登录功能
func (u *userService) SignIn(s *model.StuApiSignInReq) (string, error) {
	//创建学生信息对象
	var stu *model.Student
	//查询学生信息
	err := dao.Student.DB().Model("student").Where("stuid=? and password=?", s.Stuid, s.Password).Scan(&stu)
	if err != nil {

		return "", err
	}
	//如果查询结果为nil，那么账号和密码发生错误
	if stu == nil {
		return "", model.ErrorInvalidUser
	}
	//返回生成的token信息
	token, err := jwt.GenerateToken(stu.Id)
	if err != nil {
		return "", err
	}
	return token, nil
}

// SubmitUserData 提交user信息到数据库中
func (u *userService) SubmitUserData(data *model.StuApiSubmitDataReq, id int64) error {
	_, err := dao.Student.DB().Model("student").Where("id", id).Update(data)
	if err != nil {
		return err
	}
	return nil
}

// GetUserData 查询user相关信息，并返回
func (u *userService) GetUserData(id int64) (data *model.StuApiGetDetailRes, err error) {
	err = dao.Student.DB().Model("student").Where("id", id).Scan(&data)
	if err != nil {
		return data, err
	}
	if data == nil {
		err = model.ErrorQueryFailedUser
	}
	return
}
