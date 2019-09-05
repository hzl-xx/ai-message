package main

import (
	"messageserver/grpc/message/configure"
	"messageserver/grpc/message/protos"
	"github.com/micro/go-micro"
	"context"
	"messageserver/utils/log"
	"fmt"
	//"github.com/micro/go-micro/metadata"
	"os"
)

func main() {
	service := micro.NewService(
		micro.Name(configure.SERVER_NAME),
	)
	service.Init()
	client := protos.NewAuthService(configure.SERVER_NAME, service.Client())

	var token = &protos.Token{
		Token:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJLZXkiOnsicGhvbmUiOiIxNTAxMTE3MjI2OSJ9LCJleHAiOjE1NjcxNDczNTksImlzcyI6InNydi53ZWNoYXQubWVzc2FnZSJ9.7Zhu2vw7xGLdU69QlWdAy-fOHSwHzYv2Qi0gdqi-QA4",
	}

	res, err := client.ValidateToken(context.TODO(), token)
	if err != nil {
		log.Debug("token faild")
	}
	fmt.Println(res)
	os.Exit(0)
}
