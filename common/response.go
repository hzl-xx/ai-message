package common

import (
	"github.com/gin-gonic/gin"
)

const (
	SUCCESS = 200
	ERROR = 500
	FAILD = 400
	INVALID_PARAMS = 10001
	ERROR_AUTH_CHECK_TOKEN_FAIL = 10002
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT = 10003
	ERROR_AUTH_TOKEN = 10004
	ERROR_AUTH = 10005
	ERROR_AUTH_CHECK_KEY_FAIL = 10006
)

var Message = map[int]string{
	SUCCESS : "ok",
	ERROR : "error",
	FAILD : "faild",
	INVALID_PARAMS : "params invalid",
	ERROR_AUTH_CHECK_TOKEN_FAIL : "token check faild",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT : "token timeout",
	ERROR_AUTH_TOKEN : "get token faild",
	ERROR_AUTH : "auth faild",
	ERROR_AUTH_CHECK_KEY_FAIL : "key check faild",
}

func GetMessage(code int) string {
	msg, ok := Message[code]

	if ok {
		return msg
	}

	return Message[ERROR]
}

func Reponse(c *gin.Context,httpCode int, errCode int, data interface{}) {
	c.JSON(httpCode, gin.H{
		"code":errCode,
		"msg":GetMessage(errCode),
		"data":data,
	})
	return
}