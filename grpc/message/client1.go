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

	key := "f947be8009b342ca0f5a300587cc6379"

	var user = protos.Key{
		Key:key,
	}
	//var head = map[string]string{
		//"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VyIjp7InVzZXJuYW1lIjoiemhhbmdzYW4iLCJwYXNzd29yZCI6InpoYW5nc2FuMSJ9LCJleHAiOjE1NjcwNjcyODYsImlzcyI6InNydi53ZWNoYXQubWVzc2FnZSJ9.aVmEMr8XBUITO1s388qUy2Epv1goraMRqkVn0alti-c",
	//}

	//ctx := metadata.NewContext(context.Background(), head)

	authReponse, err := client.GetToken(context.Background(), &user)

	if err != nil {
		log.Fatal("鉴权错误", key, err)
	}
	fmt.Println(authReponse.Token)
	os.Exit(0)
}
