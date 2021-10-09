// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
)

// GameDao is the manager for logic model data accessing and custom defined data operations functions management.
type GameDao struct {
	Table   string      // Table is the underlying table name of the DAO.
	Group   string      // Group is the database configuration group name of current DAO.
	Columns GameColumns // Columns is the short type for Columns, which contains all the column names of Table for convenient usage.
}

// GameColumns defines and stores column names for table game.
type GameColumns struct {
	Id        string // 编号
	Name      string // 比赛名称
	Introduce string // 比赛简介
}

//  gameColumns holds the columns for table game.
var gameColumns = GameColumns{
	Id:        "id",
	Name:      "name",
	Introduce: "introduce",
}

// NewGameDao creates and returns a new DAO object for table data access.
func NewGameDao() *GameDao {
	return &GameDao{
		Group:   "default",
		Table:   "game",
		Columns: gameColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *GameDao) DB() gdb.DB {
	return g.DB(dao.Group)
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *GameDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.Table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *GameDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
