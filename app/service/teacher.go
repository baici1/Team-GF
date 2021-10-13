package service

import (
	"database/sql"
	"team-gf/app/dao"
	"team-gf/app/model"
)

// Teacher 管理学生相关user服务
var Teacher = teacherService{}

type teacherService struct{}

// GetTeacherDetail 获取teacher详细信息
func (*teacherService) GetTeacherDetail(teacherId int64) (data *model.TeacherApiGetDetailRes, err error) {
	data = new(model.TeacherApiGetDetailRes)
	if err := dao.Teacher.DB().Model("teacher").Where("id", teacherId).Scan(data); err != nil {
		if err == sql.ErrNoRows {
			err = model.ErrorQueryDataEmpty
		}
		return data, err
	}
	return
}
