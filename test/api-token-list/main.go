package main

import (
	"hamster/core"
	"hamster/log"
	settleServer "hamster/servers/settle"
)

func main() {

	core.InitConfig()

	s := settleServer.GetSettleServer()
	if s == nil {
		return
	}

	s.Init()

	log.Infof("单元测试")

}
