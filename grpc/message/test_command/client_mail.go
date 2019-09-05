package test_command

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
		File:"index.php",
		Message:"代码出BUG了6!!!",
		Href:"www.baidu.com",
		Type:"markdown",
	}
	commonMsg := &protos.Common{
		Type:"sentry",
		Message:"测试消息",
	}
	mail := &protos.Mail{
		Title:"审核通知测试",
		From:"xiaoxiongwei@xinchanedu.com",
		To:"2360655955@qq.com",
		Password:"qbqKqcEzvgeRAAcv",
		Message:"恭喜你审核通过",
	}
	wechat := protos.Message{
		Type:"sentry",
		Sentry:sentry,
		Common:commonMsg,
		Mail:mail,
	}
	head := map[string]string{
		"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJLZXkiOnsicGhvbmUiOiIxNTAxMTE3MjI2OSJ9LCJleHAiOjE1NjczMjUzMzgsImlzcyI6InNydi53ZWNoYXQubWVzc2FnZSJ9.zlDWOIVhtAUx1_Egv6f-RJSF5DByOphEWR-WAgzTugQ",
	}
	ctx := metadata.NewContext(context.Background(), head)
	rep,err := message.SendMessage(ctx, &wechat)
	if err != nil {
		log.Debug(err)
	}
	log.Info(rep)
}
