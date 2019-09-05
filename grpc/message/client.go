package main

import (
	"messageserver/grpc/message/configure"
	"messageserver/grpc/message/protos"
	"messageserver/utils/log"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/metadata"
	"golang.org/x/net/context"
)


func main() {
	service := micro.NewService(
		micro.Name(configure.SERVER_NAME),
	)
	message := protos.NewSendMessageService(configure.SERVER_NAME, service.Client())
	sentry := &protos.Sentry{
		ProjectName:"aix",
		Level:"error",
		Time:"2018-01-01",
		Message:"GRPC测试消息!!!",
		Href:"www.baidu.com",
		Type:"markdown",
	}
	commonMsg := &protos.Common{
		Type:"sentry",
		Message:"测试消息",
	}
	wechat := protos.Message{
		Type:"sentry",
		Sentry:sentry,
		Common:commonMsg,
	}
	head := map[string]string{
		"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJLZXkiOnsia2V5IjoiZjk0N2JlODAwOWIzNDJjYTBmNWEzMDA1ODdjYzYzNzkifSwiZXhwIjoxNTY3OTIxODg5LCJpc3MiOiJzcnYud2VjaGF0Lm1lc3NhZ2UifQ.MI-nA7iCWqM89wT2Ppa4EmL6X3b5jtXsjAiGQzDrO74",
	}
	ctx := metadata.NewContext(context.Background(), head)
	rep,err := message.SendMessage(ctx, &wechat)
	if err != nil {
		log.Debug(err)
	}
	log.Info(rep)
}
