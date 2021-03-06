// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package model

// Game is the golang structure for table game.
type Game struct {
	Id        int64  `orm:"id,primary" json:"id"`        // 编号
	Name      string `orm:"name"       json:"name"`      // 比赛名称
	Introduce string `orm:"introduce"  json:"introduce"` // 比赛简介
}

// Student is the golang structure for table student.
type Student struct {
	Id        int64  `orm:"id,primary" json:"id"`        // 学生编号
	Name      string `orm:"name"       json:"name"`      // 学生姓名
	Gender    int64  `orm:"gender"     json:"gender"`    // 学生性别 0是女生 1是男生 默认为女生
	Stuid     string `orm:"stuid"      json:"stuid"`     // 学生学号
	Password  string `orm:"password"   json:"password"`  // 学生密码
	Email     string `orm:"email"      json:"email"`     // 学生邮箱
	Introduce string `orm:"introduce"  json:"introduce"` // 学生简介
}

// Teacher is the golang structure for table teacher.
type Teacher struct {
	Id        int64  `orm:"id,primary" json:"id"`        // 老师编号
	Name      string `orm:"name"       json:"name"`      // 老师姓名
	Gender    int64  `orm:"gender"     json:"gender"`    // 老师性别 1是男生 0是女生
	Phone     string `orm:"phone"      json:"phone"`     // 老师手机号
	Email     string `orm:"email"      json:"email"`     // 老师邮箱
	Introduce string `orm:"introduce"  json:"introduce"` // 老师简介
}

// Team is the golang structure for table team.
type Team struct {
	Id        int64  `orm:"id,primary" json:"id"`        // 编号
	Name      string `orm:"name"       json:"name"`      // 队伍名字
	Introduce string `orm:"introduce"  json:"introduce"` // 简介
	Game      int64  `orm:"game"       json:"game"`      // 比赛名
	Creator   int64  `orm:"creator"    json:"creator"`   // 创建者的id
	Teacher   int64  `orm:"teacher"    json:"teacher"`   // 指导老师
}
