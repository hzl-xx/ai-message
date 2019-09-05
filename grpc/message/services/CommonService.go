package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/gomodule/redigo/redis"
	"messageserver/grpc/message/protos"
	"messageserver/utils/func"
	"messageserver/utils/log"
	"net/http"
)

var (
	webHook = "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=7589d9ce-0ea3-4881-9a53-43aa991ce088"
)

func SendCompanyWechatWebHook(message *protos.Message) (string, error) {
	// 消息转MD5
	reqMd5 := _func.Md5(*message)
	// 判断redis是否存在该数据
	is_exists, err := redis.Bool(redisConn.Do("EXISTS", reqMd5))
	//defer utils.CloseRedis()
	if err != nil {
		log.Debug(err)
	}
	if is_exists {
		return "发送失败",errors.New("发送失败")
	}

	data := make(map[string]interface{})
	content := make(map[string]interface{})
	methodList := []string{"@all"}
	mobileList := []string{""}
	data["msgtype"] = message.Common.Type
	content["content"] = message.Common.Message
	content["mentioned_list"] = methodList
	content["mentioned_mobile_list"] = mobileList
	data[message.Common.Type] = content

	dataJson, err := json.Marshal(data)
	if err != nil {
		log.Debug("msg to json faild")
	}

	readData := bytes.NewReader(dataJson)

	request, err := http.NewRequest("POST", webHook, readData)

	if err != nil {
		log.Debug("send msg faild1", err)
	}
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	client := http.Client{}
	_, err = client.Do(request)	// 发送请求
	if err != nil {
		log.Debug("send msg faild2", err)
	}
	mesMd5 := _func.Md5(message)
	_, err = redisConn.Do("SET", mesMd5, "common", "EX", "300")
	if err != nil {
		log.Debug(err)
	}
	return "发送成功", nil
}

