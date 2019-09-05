package main

import (
	"messageserver/router"
	"messageserver/utils/configure"
	"fmt"
	"github.com/fvbock/endless"
	"messageserver/utils/log"
	"syscall"
)

func main() {

	//db := model.InitMysql()               
	endless.DefaultReadTimeOut = configure.ReadTimeout
	endless.DefaultWriteTimeOut = configure.WriteTimeout
	endless.DefaultMaxHeaderBytes = 1 << 20
	endPoint := fmt.Sprintf(":%d", configure.HTTPPort)

	server := endless.NewServer(endPoint, router.Router)
	server.BeforeBegin = func(add string) {
		log.Info("Actual pid is %d", syscall.Getpid())
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Info("Server err: %v", err)
	}
}
