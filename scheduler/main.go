package main

import (
	"admin/scheduler/taskrunner"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// scheduler 调度器（任务-普通restfullapi 无法完成的任务） {定时触发、延迟触发异步的任务}

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()
	router.GET("/video-del-record/:vid", VidDelRecHandler)
	return router
}

func main() {
	go taskrunner.Start()
	r := RegisterHandlers()
	http.ListenAndServe(":9001", r)
}