package core

import (
	"hamster/log"
	"net/http"
	"net/http/pprof"

	"github.com/gorilla/mux"
)

// 监控服务
func PprofMonitor(listen string) {

	defer Recover()

	if len(listen) == 0 {
		log.Infof("pprof port not enabled")
		return
	}

	router := mux.NewRouter()
	router.Handle("/debug/pprof/", http.HandlerFunc(pprof.Index))
	router.Handle("/debug/pprof/cmdline", http.HandlerFunc(pprof.Cmdline))
	router.Handle("/debug/pprof/profile", http.HandlerFunc(pprof.Profile))
	router.Handle("/debug/pprof/symbol", http.HandlerFunc(pprof.Symbol))
	router.Handle("/debug/pprof/trace", http.HandlerFunc(pprof.Trace))
	router.Handle("/debug/pprof/{cmd}", http.HandlerFunc(pprof.Index))

	log.Infof("pprof start to listening:%s", listen)

	log.Errorf("pprof error: %v", http.ListenAndServe(listen, nil))
}
