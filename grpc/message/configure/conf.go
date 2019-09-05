package configure

import (
	"messageserver/utils/func"
	"messageserver/utils/log"
	"gopkg.in/ini.v1"
)

var (
	Cfg *ini.File
	HOST string
	SERVER_NAME string
	Agentid int64
	Secret string
	Corpid string
	RedisHost string
	JWT_KEY string
	RABBMITMQ_HOST string
	RABBMITMQ_NAME string
	MailFrom string
	MailPassword string
)

func init() {
	var err error
	Cfg, err = ini.Load(_func.GetProjectBasePath()+"/conf/message.ini")
	if err != nil {
		log.Fatal("Fail to parse 'conf/message.ini': ", err)
	}

	loadServer()
	loadWechat()
	loadRedis()
	loadKey()
	loadQueue()
	loadBase()
}

func loadServer() {
	sec, err := Cfg.GetSection("qywx_server")
	if err != nil {
		log.Debug("Fail to get section 'grpc_server': ", err)
	}
	HOST = sec.Key("HOST").String()
	SERVER_NAME = sec.Key("SERVER_NAME").String()
}

func loadWechat() {
	sec, err := Cfg.GetSection("company_wechat")
	if err != nil {
		log.Debug("Fail to get section 'company_wechat': ", err)
	}
	Agentid = sec.Key("AGENTID").MustInt64()
	Secret = sec.Key("SECRET").String()
	Corpid = sec.Key("CORPID").String()
}

func loadRedis() {
	sec, err := Cfg.GetSection("redis")
	if err != nil {
		log.Debug("Fail to get section 'redis': %v", err)
	}

	RedisHost = sec.Key("HOST").String()
}

func loadKey() {
	sec, err := Cfg.GetSection("jwt_key")
	if err != nil {
		log.Debug("Fail to get section 'redis': %v", err)
	}

	JWT_KEY = sec.Key("KEY").String()
}

func loadQueue(){
	sec, err := Cfg.GetSection("rabbit_mq")
	if err != nil {
		log.Debug("Fail to get section 'rabbit_mq': %v", err)
	}

	RABBMITMQ_HOST = sec.Key("RABBMITMQ_HOST").String()
	RABBMITMQ_NAME = sec.Key("RABBMITMQ_NAME").String()
}

func loadBase(){
	sec, err := Cfg.GetSection("mail")
	if err != nil {
		log.Debug("Fail to get section 'mail': %v", err)
	}

	MailFrom = sec.Key("FROM").String()
	MailPassword = sec.Key("PASSWORD").String()
}
