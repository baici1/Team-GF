package student

import (
	"team-gf/app/dao"
	"team-gf/app/model"
)

// User 管理学生相关user服务
var User = userService{}

type userService struct{}

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
