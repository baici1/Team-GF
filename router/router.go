package router

import (
	"team-gf/app/api"

	"github.com/gogf/gf/frame/g"
)

func init() {
	s := g.Server()
	s.BindObject("/user", api.Student)
}
