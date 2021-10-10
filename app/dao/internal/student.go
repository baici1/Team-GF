// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
)

// StudentDao is the manager for logic model data accessing and custom defined data operations functions management.
type StudentDao struct {
	Table   string         // Table is the underlying table name of the DAO.
	Group   string         // Group is the database configuration group name of current DAO.
	Columns StudentColumns // Columns is the short type for Columns, which contains all the column names of Table for convenient usage.
}

// StudentColumns defines and stores column names for table student.
type StudentColumns struct {
	Id        string // 学生编号
	Name      string // 学生姓名
	Gender    string // 学生性别 0是女生 1是男生 默认为女生
	Stuid     string // 学生学号
	Password  string // 学生密码
	Email     string // 学生邮箱
	Introduce string // 学生简介
}

//  studentColumns holds the columns for table student.
var studentColumns = StudentColumns{
	Id:        "id",
	Name:      "name",
	Gender:    "gender",
	Stuid:     "stuid",
	Password:  "password",
	Email:     "email",
	Introduce: "introduce",
}

// NewStudentDao creates and returns a new DAO object for table data access.
func NewStudentDao() *StudentDao {
	return &StudentDao{
		Group:   "default",
		Table:   "student",
		Columns: studentColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *StudentDao) DB() gdb.DB {
	return g.DB(dao.Group)
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *StudentDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.Table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *StudentDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
