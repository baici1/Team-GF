package model

type TeacherApiGetDetailRes struct {
	Name      string `json:"name,omitempty"`         // 老师姓名
	Gender    int64  `d:"0" json:"gender,omitempty"` // 老师性别 1是男生 0是女生
	Phone     string `json:"phone,omitempty"`        // 老师手机号
	Email     string `json:"email,omitempty"`        // 老师邮箱
	Introduce string `json:"introduce,omitempty"`    // 老师简介
}
