package core

import (
	"hamster/log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// 服务接口
type Runable interface {
	Start()
	Stop()
	Done() <-chan struct{}
}

// 运行服务
func Run(r Runable) {
	csignal := make(chan os.Signal, 1)
	signal.Notify(csignal, syscall.SIGINT, syscall.SIGHUP, syscall.SIGTERM, syscall.SIGILL, syscall.SIGQUIT)
	for {
		go r.Start()
		var s os.Signal
		var doneQuit bool
		select {
		case s = <-csignal:
		case <-r.Done():
			s = syscall.SIGQUIT
			doneQuit = true
		}
		switch s {
		case syscall.SIGINT:
			log.Warnf("event:exist SIGINT")
		case syscall.SIGQUIT:
			if doneQuit {
				log.Warnf("event:exist RUNABLE")
			} else {
				log.Warnf("event:exist SIGQUIT")
			}
		case syscall.SIGHUP:
			time.Sleep(time.Second)
			r.Stop()
			log.Warnf("event:exist SIGHUP")
			continue
		case syscall.SIGKILL:
			log.Warnf("event:exist SIGKILL")
		case syscall.SIGTERM:
			log.Warnf("event:exist SIGTERM")
		default:
			log.Warnf("event:exist UNKNOWN")
		}
		r.Stop()
		break
	}
}
