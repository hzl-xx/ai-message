package v1

import (
	"messageserver/common"
	"messageserver/grpc/message/protos"
	"messageserver/grpc/message/services"
	//"go-api/utils"
	"messageserver/utils/log"
	//"context"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro"
	//"github.com/micro/go-micro/metadata"
	"net/http"
	"messageserver/controller"
)

type WechatController struct {
	controller.BaseController
}

func (w *WechatController) SendMessage(c *gin.Context) {
	var params common.Params
	err := c.BindJSON(&params)
	log.Info(params)
	if err != nil {
		log.Debug("params bind faild")
		common.Reponse(c,http.StatusBadRequest, common.ERROR, "")
		return
	}
	//log.Debug(params)
	//service := getService();
	//key := protos.Key{
	//	Key:utils.GetKey(),
	//}
	//tokenService := &services.TokenService{}
	//token ,err := tokenService.Encode(&key)
	//if err != nil {
	//	log.Debug("get token faild")
	//}
	//ctx := metadata.NewContext(context.TODO(), map[string]string{"token":token,})
	message := formatParams(params)
	err = services.NewMessageService().PushMessage(&message)
	if err != nil {
		common.Reponse(c,http.StatusBadRequest, common.FAILD, "发送失败")
	}


	//rep,err := service.SendMessage(ctx, &message)
	//if err != nil {
	//	log.Debug(err)
	//}
	common.Reponse(c,http.StatusOK, common.SUCCESS, "发送成功")
}

func getService() protos.SendMessageService {
	service := micro.NewService(
		micro.Name("srv.send.message"),
	)
	message := protos.NewSendMessageService("srv.send.message", service.Client())
	return message
}

func formatParams(params common.Params) protos.Message {
	var message protos.Message
	message.Type = params.Type
	if params.Type == "sentry" {
		sentry := protos.Sentry{}
		sentry.ProjectName = params.Sentry.ProjectName
		sentry.Level = params.Sentry.Level
		sentry.Message = params.Sentry.Message
		sentry.Href = params.Sentry.Href
		sentry.Type = params.Sentry.Type
		message.Sentry = &sentry
	}else if params.Type == "mail" {
		mail := protos.Mail{}
		mail.From=params.Mail.From
		mail.To=params.Mail.To
		mail.Title=params.Mail.Title
		mail.Message=params.Mail.Message
		mail.Password=params.Mail.Password
		message.Mail = &mail
	}else {
		commonMsg := protos.Common{}
		commonMsg.Type = params.Common.Type
		commonMsg.Message = params.Common.Message
		message.Common = &commonMsg
	}
	return message
}