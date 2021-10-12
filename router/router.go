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
		//用户相关的，三个身份共同的api
		group.Group("/stu", func(group *ghttp.RouterGroup) {
			group.POST("/submit", api.Student.SubmitUserData)
			group.GET("/check", api.Student.GetUserDate)
		})
		//这里是用户操作队伍相关的
		group.Group("/team", func(group *ghttp.RouterGroup) {
			//三个身份共同的api
			group.GET("/check", api.TeamCommon.GetTeamAllDetail)
			//当用户为leader的api操作
			group.Group("/leader", func(group *ghttp.RouterGroup) {
				group.POST("/create", api.TeamLeader.CreateOwnTeam)
				group.POST("/append", api.TeamLeader.PushStuInTeam)
				group.DELETE("/member", api.TeamLeader.RemoveUserAtTeam)
				group.DELETE("/ownTeam", api.TeamLeader.DeleteOwnTeam)
			})
			//当身份是member的api操作
			group.Group("/member", func(group *ghttp.RouterGroup) {
				group.DELETE("/leave", api.TeamMember.LeaveToTeam)
			})

		})

	})
}
