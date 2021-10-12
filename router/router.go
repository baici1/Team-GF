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
	s.Group("/", func(group *ghttp.RouterGroup) {
		//学生注册
		group.POST("/signup", api.Student.SignUp)
		//学生登录
		group.POST("/signin", api.Student.SignIn)
		group.Middleware(service.Middleware.JWTAuthMiddleware)
		//这里路由只与用户相关的
		group.Group("/stu", func(group *ghttp.RouterGroup) {
			group.ALL("/ping", api.Student.Ping)
			group.POST("/submit", api.Student.SubmitUserData)
			group.GET("/get", api.Student.GetUserDate)
		})
		//这里是用户操作队伍相关的
		group.Group("/team", func(group *ghttp.RouterGroup) {
			group.POST("/create", api.Team.CreateOwnTeam)
			group.GET("/get", api.Team.GetTeamAllDetail)
			group.POST("/push", api.Team.PushStuInTeam)
			group.DELETE("/member", api.Team.RemoveUserAtTeam)
			group.DELETE("/ownteam", api.Team.DeleteOwnTeam)
		})

	})
}
