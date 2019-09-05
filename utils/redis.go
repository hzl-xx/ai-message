package utils

import (
	"github.com/gomodule/redigo/redis"
	"messageserver/utils/configure"
	"time"
)

var (
	RedisConn redis.Conn
	RedisClient *redis.Pool
	err error
)

func init() {
	// 最大空闲连接数
	maxIdle := 3
	//if v, ok := conf["MaxIdle"]; ok {
	//	maxIdle = int(v.(int64))
	//}
	// 最大连接数
	maxActive := 3
	//if v, ok := conf["MaxActive"]; ok {
	//	maxActive = int(v.(int64))
	//}

	// 建立连接池
	RedisClient = &redis.Pool{
		MaxIdle:     maxIdle,
		MaxActive:   maxActive,
		IdleTimeout: 20 * time.Second, // 空闲连接超时时间
		Wait:        true,
		Dial: func() (redis.Conn, error) {
			con, err := redis.Dial("tcp", configure.RedisHost,
				redis.DialPassword(configure.RedisPassword),
				redis.DialDatabase(configure.RedisSelect),
				redis.DialConnectTimeout(30*time.Second),
				redis.DialReadTimeout(30*time.Second),
				redis.DialWriteTimeout(30*time.Second))
			if err != nil {
				return nil, err
			}
			return con, nil
		},
	}

	//RedisConn, err = redis.Dial("tcp",
	//	configure.RedisHost,
	//	redis.DialPassword(configure.RedisPassword),
	//)

	//if err != nil {
	//	log.Debug("Connect to redis error", err)
	//}
}

func CloseRedis() {
	defer RedisConn.Close()
}
