package router

import (
	"team-gf/app/api"
	"team-gf/app/service"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func init() {
	s := g.Server()
	s.Use(service.Middleware.CORS)
	s.Group("/stu", func(group *ghttp.RouterGroup) {
		//学生注册
		group.POST("/signup", api.Student.SignUp)
		//学生登录
		group.POST("/signin", api.Student.SignIn)
		group.Group("/", func(group *ghttp.RouterGroup) {
			group.Middleware(service.Middleware.JWTAuthMiddleware)
			group.ALL("/ping", api.Student.Ping)
		})
	})
}
