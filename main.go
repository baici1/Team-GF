package main

import (
	_ "team-gf/boot"
	_ "team-gf/router"

	"github.com/gogf/gf/frame/g"
)

func main() {
	g.Server().Run()
}
