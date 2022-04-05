package main

import (
	"HelloGo/controller/hello"
	"HelloGo/controller/user"
	"HelloGo/tool"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := tool.ParseConfig("./config/app.json")
	if err != nil {
		panic(err.Error())
	}
	engine := gin.Default()
	user.Router(engine)
	hello.Router(engine)
	engine.Run(cfg.AppHost + ":" + cfg.AppPort)
}
