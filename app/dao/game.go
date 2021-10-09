// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"team-gf/app/dao/internal"
)

// gameDao is the manager for logic model data accessing and custom defined data operations functions management.
// You can define custom methods on it to extend its functionality as you wish.
type gameDao struct {
	*internal.GameDao
}

var (
	// Game is globally public accessible object for table game operations.
	Game = gameDao{
		internal.NewGameDao(),
	}
)

// Fill with you ideas below.