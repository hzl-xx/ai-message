package handle

import (
	"context"
	"errors"
	"messageserver/grpc/message/protos"
	"messageserver/grpc/message/services"
	"messageserver/utils/func"
	"messageserver/utils/log"
)

func NewAuthService() *AuthService {
	return &AuthService{}
}

type AuthService struct {

}

var (
	service  = &services.TokenService{}
)

func (auth *AuthService) GetToken(ctx context.Context, key *protos.Key, rep *protos.Response) error {
	// 验证key是否有效
	if validRes := _func.ValidKey(key.Key); validRes == false {
		rep.Code = 500
		rep.Message = "check key faild"
		rep.Valid = false
		rep.Token = &protos.Token{Token:""}
		return errors.New("key验证失败")
	}
	str , err := service.Encode(key)
	log.Info("获取token：",str)
	if err != nil {
		rep.Code = 500
		rep.Message = "get token faild"
		rep.Valid = false
		rep.Token = &protos.Token{Token:""}
		return err
	}

	rep.Code = 200
	rep.Message = "success"
	rep.Valid = true
	rep.Token = &protos.Token{Token:str}
	return nil
}

func (auth *AuthService) ValidateToken(ctx context.Context, token *protos.Token, rep *protos.Response) error {

	res , err := service.Decode(token.Token)
	log.Info(res,err)
	if err != nil {
		rep.Valid = false
		rep.Code = 500
		rep.Message = "token validate faild"
		rep.Token = token
		return err
	}
	rep.Valid = true
	rep.Code = 200
	rep.Message = "success"
	rep.Token = token
	return nil
}
