package main

import (
	_ "team-gf/boot"
	_ "team-gf/router"

	"github.com/gogf/gf/frame/g"
)

// @title       team-GF服务API
// @version     1.0
// @description 项目team-GF服务API接口文档。
// @schemes     http
func main() {
	g.Server().Run()
}
