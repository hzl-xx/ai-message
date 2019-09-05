package common

import (
	"messageserver/utils/log"
	"github.com/astaxie/beego/validation"
)

func Errors(errors []*validation.Error) {
	for _, err := range errors {
		log.Info(err.Key, err.Message)
	}
	return
}
