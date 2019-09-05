package main

import (
	"encoding/json"
	"fmt"
	"messageserver/grpc/message/protos"
	"messageserver/grpc/message/services"
	"messageserver/utils/log"
)

func main() {
	var (
		message = &protos.Message{}
		res string
	)
	messageList, err := services.NewMessageService().ConsumeMessage()

	if err != nil {
		log.Info(err)
	}
	forever := make(chan bool)
	go func() {
		for d := range messageList  {
			json.Unmarshal(d.Body, message)
			switch message.Type {
				case "sentry":
					res, err = services.SendCompanyWechatMessage(message)
				case "mail":
					res, err = services.SendMailMessage(message)
				default:
					res, err = services.SendCompanyWechatWebHook(message)
			}
			if err != nil {
				log.Info("consume message: %s", err)
			} else {
				fmt.Println(res)
			}

		}
	}()
	<-forever
}
