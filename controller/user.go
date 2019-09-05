package controller

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"messageserver/common"
	"messageserver/model"
	"messageserver/utils"
	"messageserver/utils/func"
	"messageserver/utils/log"
	"math"
	"net/http"
	"strconv"
)

type UserController struct{
	BaseController
}

func (uc *UserController) GetAuth(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	valid := validation.Validation{}
	u := &model.User{Username: username, Password: password}
	ok, _ := valid.Valid(u)

	data := make(map[string]interface{})
	code := common.INVALID_PARAMS
	if ok {
		isExist := u.CheckAuth(username, password)
		if isExist {
			token, err := utils.GenerateToken(username)
			if err != nil {
				code = common.ERROR_AUTH_TOKEN
			} else {
				data["token"] = token

				code = common.SUCCESS
			}

		} else {
			code = common.ERROR_AUTH
		}
	} else {
		for _, err := range valid.Errors {
			log.Info(err.Key, err.Message)
		}
	}
	log.Info(data)

	common.Reponse(c,http.StatusInternalServerError, code, data)
}

func (uc *UserController) GetUser(c *gin.Context) {
	username := c.Query("username")
	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	if username != "" {
		maps["username"] = username
	}
	pagenum,pagesize := utils.GetPage(c)
	user := &model.User{}
	data["list"] = user.GetUser(pagenum, pagesize, maps)
	total := user.GetUserTotal(maps)
	data["total"] = total
	data["page"] = c.Query("page")
	data["page"] = int(math.Ceil(float64(total/pagesize)))
	common.Reponse(c,http.StatusInternalServerError, common.SUCCESS, data)
}

func (uc *UserController) CreateUser(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	phone := c.PostForm("phone")
	name := c.PostForm("name")
	age := _func.QueryUint8(c.PostForm("age"))

	valid := validation.Validation{}
	valid.Required(username, "username").Message("名称不能为空")
	valid.MaxSize(username, 50, "username").Message("名称不能超过50字符")
	valid.Phone(phone, "phone").Message("手机号格式错误")

	if valid.HasErrors() {
		common.Errors(valid.Errors)
		common.Reponse(c,http.StatusOK, common.INVALID_PARAMS, nil)
		return
	}

	user := &model.User{Username:username,Password:password,Phone:phone,Name:name,Age:age}
	err := user.InsertUser()
	if err != nil {
		common.Reponse(c,http.StatusInternalServerError, common.ERROR, nil)
		return
	}
	common.Reponse(c,http.StatusInternalServerError, common.SUCCESS, nil)
}

func (uc *UserController) UpdateUser(c *gin.Context) {
	id,_ := strconv.Atoi(c.Param("id"))
	username := c.PostForm("username")
	password := c.PostForm("password")
	phone := c.PostForm("phone")
	name := c.PostForm("name")
	age := _func.QueryUint8(c.PostForm("age"))

	valid := validation.Validation{}
	valid.Required(id, "id").Message("ID不能为空")
	valid.Min(id, 1,"id").Message("ID格式错误")

	if valid.HasErrors() {
		common.Errors(valid.Errors)
		common.Reponse(c,http.StatusOK, common.INVALID_PARAMS, nil)
		return
	}

	user := &model.User{}

	data := make(map[string]interface{})
	if name != "" {
		valid.MaxSize(name, 50, "name").Message("名称不能超过50字符")
		data["name"] = name
	}
	if username != "" {
		valid.MaxSize(username, 50, "username").Message("用户名不能超过50字符")
		data["username"] = username
	}
	if password != "" {
		valid.MaxSize(password, 50, "password").Message("密码不能超过50字符")
		data["password"] = password
	}
	if phone != "" {
		valid.Phone(phone, "phone").Message("手机号格式错误")
		data["phone"] = phone
	}
	if age > 0 {
		data["age"] = age
	}
	user.EditUser(id, data)

	common.Reponse(c,http.StatusOK, common.SUCCESS, nil)
}

func (uc *UserController) DeleteUser(c *gin.Context) {
	id,_ := strconv.Atoi(c.Param("id"))
	valid := validation.Validation{}
	valid.Required(id, "id").Message("ID不能为空")
	valid.Min(id, 1,"id").Message("ID格式错误")

	if valid.HasErrors() {
		common.Errors(valid.Errors)
		common.Reponse(c,http.StatusOK, common.INVALID_PARAMS, nil)
		return
	}
	user := &model.User{}
	user.DeleteUser(id)
	common.Reponse(c,http.StatusOK, common.SUCCESS, nil)
}
