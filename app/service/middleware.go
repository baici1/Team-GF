package service

import (
	"strings"
	"team-gf/library/code"
	"team-gf/library/jwt"
	"team-gf/library/response"

	"github.com/gogf/gf/net/ghttp"
)

//鉴权中间件，解析token，获取用户id供后续使用
func JWTAuthMiddleware(r *ghttp.Request) {
	//获取头部的token值
	authHeader := r.Header.Get("Authorization")
	//当头部token请求为空时候
	if authHeader == "" {
		//返回错误信息（需要登录）
		response.ResponseError(r, code.CodeNeedLogin)
	}
	//对token进行分割
	parts := strings.SplitN(authHeader, " ", 2)
	//当token格式不正确
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		response.ResponseError(r, code.CodeInvalidToken)
	}
	//获取之前生成好的jwt，解析jwt
	mc, err := jwt.ParseToken(parts[1])
	//当解析发生错误，返回token失效的结果
	if err != nil {
		response.ResponseError(r, code.CodeInvalidToken)
	}
	// 将请求中的id信息保存到请求的上下文c上
	r.SetParam("ID", mc.ID)
	// 执行下一步请求逻辑
	r.Middleware.Next()
}
