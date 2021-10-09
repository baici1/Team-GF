package boot

import (
	"fmt"
	"team-gf/library/snowflake"
	_ "team-gf/packed"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/swagger"
)

func init() {
	s := g.Server()
	cfg := g.Cfg()
	//雪花算法生成用户ID
	if err := snowflake.Init(cfg.GetString("app.startTime"), cfg.GetInt64("app.machineId")); err != nil {
		fmt.Printf("init snowflake failed err:%v\n", err)
		return
	}
	s.Plugin(&swagger.Swagger{})
}
