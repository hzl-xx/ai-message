package middleware

import (
	"time"
	"net/http"

	"github.com/gin-gonic/gin"

	"messageserver/utils"
	"messageserver/common"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = common.SUCCESS
		token := c.GetHeader("token")

		if token == "" {
			code = common.ERROR_AUTH_CHECK_TOKEN_FAIL
		} else {
			claims, err := utils.ParseToken(token)
			if err != nil {
				code = common.ERROR_AUTH_CHECK_TOKEN_FAIL
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = common.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			}
		}

		if code != common.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code" : code,
				"msg" : common.GetMessage(code),
				"data" : data,
			})

			c.Abort()
			return
		}

		c.Next()
	}
}

