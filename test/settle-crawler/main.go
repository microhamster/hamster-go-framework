package main

import (
	"hamster/core"
	"hamster/log"
	apiServer "hamster/servers/api"
)

func main() {

	core.InitConfig()

	s := apiServer.GetApiServer()
	if s == nil {
		return
	}

	s.Init()

	log.Infof("单元测试")

}
