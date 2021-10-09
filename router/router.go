package router

import (
	"team-gf/app/api"
	"team-gf/app/service"

	"github.com/gogf/gf/frame/g"
)

func init() {
	s := g.Server()
	s.Use(service.JWTAuthMiddleware)
	s.BindObject("/user", api.Student)
}
