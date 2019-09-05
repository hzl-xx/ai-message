package utils

import (
	"messageserver/utils/configure"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetPage(c *gin.Context) (pageNum, pagesize int) {

	pageNum,_ = strconv.Atoi(c.Query("page"))
	pagesize,_ = strconv.Atoi(c.Query("pagesize"))
	if pagesize <= 0 {
		pagesize = configure.PageSize
	}
	if pageNum > 0 {
		pageNum = (pageNum-1) * pagesize
	}

	return
}
