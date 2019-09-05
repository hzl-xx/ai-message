package configure

import (
	"gopkg.in/ini.v1"
	"log"
	"time"
	"messageserver/utils/func"
)

var (
	Cfg *ini.File
	AppEnv string
	HTTPPort int
	ReadTimeout time.Duration
	WriteTimeout time.Duration
	PageSize int
	JwtSecret string
	JwtExpireTime int
	RedisHost string
	From string
	Password string
	RedisPassword string
	RedisSelect int
)

func init() {
	var err error
	Cfg, err = ini.Load(_func.GetProjectBasePath()+"/conf/app.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
	}

	LoadBase()
	LoadServer()
	LoadApp()
	LoadRedis()
	LoadMail()
}

func LoadBase() {
	AppEnv = Cfg.Section("").Key("APP_ENV").String()
}

func LoadServer() {
	sec, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatalf("Fail to get section 'server': %v", err)
	}

	HTTPPort = sec.Key("HTTP_PORT").MustInt(8000)
	ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeout =  time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
}

func LoadApp() {
	sec, err := Cfg.GetSection("app")
	if err != nil {
		log.Fatalf("Fail to get section 'app': %v", err)
	}

	JwtSecret = sec.Key("JWT_SECRET").MustString("!@)*#)!@U#@*!@!)")
	JwtExpireTime = sec.Key("JWT_EXPIRE_TIME").MustInt(24)
	PageSize = sec.Key("PAGE_SIZE").MustInt(10)
}

func LoadRedis() {
	sec, err := Cfg.GetSection("redis")
	if err != nil {
		log.Fatalf("Fail to get section 'redis': %v", err)
	}

	RedisHost = sec.Key("HOST").String()
	RedisPassword = sec.Key("PASSWORD").String()
	RedisSelect = sec.Key("SELECT").MustInt(15)
}

func LoadMail() {
	sec, err := Cfg.GetSection("mail")
	if err != nil {
		log.Fatalf("Fail to get section 'redis': %v", err)
	}

	From = sec.Key("FROM").String()
	Password = sec.Key("PASSWORD").String()
}


