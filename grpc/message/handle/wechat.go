package handle

import (
	"messageserver/common"
	"messageserver/grpc/message/protos"
	"messageserver/grpc/message/services"
	"context"
	"messageserver/utils/log"
)

//对外提供的工厂函数
func NewService() *MessageService {
	return &MessageService{}
}

type MessageService struct {
}

func (w *MessageService) SendMessage(ctx context.Context, req *protos.Message, rep *protos.Reponse) error {
	err := services.NewMessageService().PushMessage(req)
	if err != nil {
		log.Info(err)
		rep.Msg = "发送失败"
		rep.Code = common.ERROR
	} else {
		rep.Msg = "发送成功"
		rep.Code = common.SUCCESS
	}

	return nil
}