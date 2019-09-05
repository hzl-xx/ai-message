package controller

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"messageserver/common"
	"messageserver/utils"
	"messageserver/utils/func"
	"net/http"
)

type TokenController struct {
	BaseController
}

func (t *TokenController) GetKey(c *gin.Context) {
	key := _func.GetKey()
	data := map[string]string{
		"key":key,
	}
	c.JSON(http.StatusOK, data)
}

func (t *TokenController) GetToken(c *gin.Context) {
	key := c.Query("key")

	valid := validation.Validation{}
	valid.Required(key, "key").Message("不能为空")
	if valid.HasErrors() {
		common.Errors(valid.Errors)
		common.Reponse(c,http.StatusOK, common.INVALID_PARAMS, nil)
		return
	}

	// 验证key是否有效
	if validRes := _func.ValidKey(key); validRes == false {
		common.Reponse(c,http.StatusOK, common.ERROR_AUTH_CHECK_KEY_FAIL, nil)
		return
	}

	token , err := utils.GenerateToken(key)
	if err != nil {
		common.Reponse(c,http.StatusInternalServerError, common.ERROR, nil)
		return
	}
	res := map[string]string{
		"access_token":token,
	}
	c.JSON(http.StatusOK, res)
}
