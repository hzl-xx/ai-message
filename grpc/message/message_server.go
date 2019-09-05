package main

import (
	"context"
	"errors"
	"flag"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/metadata"
	"github.com/micro/go-micro/server"
	"messageserver/grpc/message/configure"
	"messageserver/grpc/message/handle"
	"messageserver/grpc/message/protos"
	"messageserver/utils/log"
	"syscall"
)

var (
	// 命令行参数-host，默认服务监听端口
	addr = flag.String("host", configure.HOST, "")
)


func main() {
	// 拦截输出到stderr的内容到日志文件
	syscall.Dup2(int(log.F.Fd()), 2)

	service := micro.NewService(
		micro.Name(configure.SERVER_NAME),
		micro.Version("latest"),
		micro.Address(configure.HOST),
		micro.WrapHandler(AuthWrapper),
	)
	service.Init()
	protos.RegisterSendMessageServiceHandler(service.Server(), handle.NewService())
	protos.RegisterAuthHandler(service.Server(), handle.NewAuthService())

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

func AuthWrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, resp interface{}) error {
		meta, ok := metadata.FromContext(ctx)
		log.Debug("请求头：",meta,ok)
		if !ok {
			return errors.New("no auth meta-data found in request")
		}

		if meta["Micro-Method"] == "Auth.GetToken" || meta["Micro-Method"] == "Auth.ValidateToken" {
			err := fn(ctx, req, resp)
			return err
		}
		// Note this is now uppercase (not entirely sure why this is...)
		token := meta["Token"]
		if token == "" {
			return errors.New("no auth meta-data found in token")
		}

		// Auth here
		authClient := protos.NewAuthService(configure.SERVER_NAME, client.DefaultClient)
		_, err := authClient.ValidateToken(context.Background(), &protos.Token{
			Token: token,
		})
		if err != nil {
			return err
		}
		err = fn(ctx, req, resp)
		return err
	}
}