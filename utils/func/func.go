package _func

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"messageserver/common"
	"math"
	"os"
	"strconv"
	"messageserver/utils/log"
	"strings"
	"time"
)

func InterfaceConvertInt(i interface{}) (res int,err error) {
	switch i.(type) {
		case string:
			res,_ := strconv.Atoi(i.(string))
			return res,nil
		case int64:
			return i.(int),nil
		case float64:
			return int(math.Round(i.(float64))), nil
		default:
			return 0, fmt.Errorf("不支持的数据类型")
	}
}

func QueryUint8(s string) (parm uint8) {
	u,err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		log.Debug(err)
	}

	return uint8(u)
}

func QueryUint16(s string) (parm uint16) {
	u,err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		log.Debug(err)
	}

	return uint16(u)
}

func QueryUint32(s string) (parm uint32) {
	u,err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		log.Debug(err)
	}

	return uint32(u)
}

func QueryUint64(s string) (parm uint64) {
	u,err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		log.Debug(err)
	}

	return u
}

func QueryString(i int) (s string) {
	s = strconv.Itoa(i)
	return
}

func Md5(v interface{}) string {
	reqString, err := json.Marshal(v)
	if err != nil {
		log.Debug(err)
	}
	// 将消息转为MD5格式
	reqMd5 := fmt.Sprintf("%x", md5.Sum(reqString))
	return reqMd5
}

func GetKey() string {
	return Md5(time.Now().Format("2006-01-02"))
}

func ValidKey(key string) bool {
	sign := Md5(time.Now().Format("2006-01-02"))
	if sign != key {
		return false
	} else {
		return true
	}
}

func GetProjectBasePath() string {
	basePath,err := os.Getwd()
	if err != nil {
		log.Fatal("get project base path faild")
	}
	filePath := strings.Split(basePath, common.PROJECT_NAME)
	return filePath[0]+common.PROJECT_NAME
}