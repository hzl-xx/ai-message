package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/gomodule/redigo/redis"
	"messageserver/grpc/message/configure"
	"messageserver/grpc/message/protos"
	"messageserver/utils"
	"messageserver/utils/func"
	"messageserver/utils/log"
	"io/ioutil"
	"net/http"
	"time"
	"gopkg.in/gomail.v2"
)

type resJson struct {
	Errcode int64 `json:"errcode"`
	Errmsg string `json:"errmsg"`
	AccessToken string `json:"access_token"`
	ExpiresIn int64 `json:"expires_in"`
}

var (
	redisConn = utils.RedisClient.Get()
)

/*
企业微信消息推送服务
 */
func SendCompanyWechatMessage(req *protos.Message) (string, error) {
	//redisConn1 := utils.RedisClient.Get()
	// 消息转MD5
	reqMd5 := _func.Md5(*req)
	log.Info("本次消息加密结果：",reqMd5)
	// 判断redis是否存在该数据
	is_exists, err := redis.Bool(redisConn.Do("EXISTS", reqMd5))
	log.Info(is_exists, err)
	if err != nil {
		log.Debug(err)
	}
	//defer utils.CloseRedis()
	//defer redisConn.Close()
	if is_exists {
		return "发送失败",errors.New("发送失败")
	}
	// 获取token
	token := GetToken().AccessToken
	res := SendMsg(token, *req)
	//redisConn1.Close()
	return res,nil
}
/*
发送邮件服务
*/
func SendMailMessage(req *protos.Message) (string, error)  {

	// 消息转MD5
	reqMd5 := _func.Md5(*req)

	// 判断redis是否存在该数据
	is_exists, err := redis.Bool(redisConn.Do("EXISTS", reqMd5))
	if err != nil {
		log.Debug(err)
	}

	if is_exists {
		return "发送失败",errors.New("发送失败")
	}



	res := SendMail(*req)
	return res,nil
}
/*
获取token
 */
func GetToken() resJson {
	var token resJson
	// 判断token是否过期
	accToken, err := redis.String(redisConn.Do("GET", "company_wechat_token"))
	if err == nil {
		token.AccessToken = accToken
		return token
	}
	url := "https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid="+ configure.Corpid+"&corpsecret="+ configure.Secret
	rep,err := http.Get(url)	// 发送get请求
	if err != nil {
		log.Error("get access token fail")
	}
	defer rep.Body.Close()
	// 读取body
	res,err := ioutil.ReadAll(rep.Body)
	if err != nil {
		log.Error("read get handle error")
	}


	// 转json
	if err=json.Unmarshal([]byte(string(res)), &token); err != nil {
		log.Error("json str to struct faild")
	}
	// 存储token到redis
	redisConn.Do("SET", "company_wechat_token", token.AccessToken, "EX", token.ExpiresIn)
	return token
}

/*
发送消息
 */
func SendMsg(token string, message protos.Message) string {
	url := "https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=" + token
	log.Info(url)
	var data= make(map[string]interface{})
	var content = make(map[string]string)
	var sendPeople = sendPeople(message.Sentry.ProjectName)
	data["touser"] = sendPeople["touser"]
	data["toparty"] = sendPeople["toparty"]
	data["totag"] = sendPeople["totag"]
	data["msgtype"] = message.Sentry.Type
	data["agentid"] = configure.Agentid
	data["safe"] = 0
	data["enable_id_trans"] = 0
	content["content"] = formatMessage(*message.Sentry)
	if (message.Sentry.Type == "text") {
		data["text"] = content
	} else {
		data["markdown"] = content
	}
	dataJson, err := json.Marshal(data)
	if err != nil {
		log.Debug("msg to json faild")
	}

	readData := bytes.NewReader(dataJson)

	request, err := http.NewRequest("POST", url, readData)
	if err != nil {
		log.Debug("send msg faild1", err)
	}
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	client := http.Client{}
	resp, err := client.Do(request)	// 发送请求
	if err != nil {
		log.Debug("send msg faild2", err)
	}

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Debug("read return msg faild", err)
	}
	str := string(respBytes)
	log.Info(str)
	mesMd5 := _func.Md5(message)
	//log.Info(mesMd5)
	_, err = redisConn.Do("SET", mesMd5, "sentry", "EX", "300")
	if err != nil {
		log.Debug(err)
	}
	//log.Info(str)
	return "发送成功"
}
/**
发送邮件
 */
func SendMail(message protos.Message) string {
	m := gomail.NewMessage()
	from := configure.MailFrom
	password := configure.MailPassword
	var mailPeople = mailPeople(message.Mail.Type)
	//to := message.Mail.To
	m.SetHeader("From",from)//发送人
	m.SetHeader("To",mailPeople["to"]...)  //接收人
	//m.SetHeader("Cc", mailPeople["cc"],"抄送") //抄送

	m.SetHeader("Subject",message.Mail.Title)//邮件主题
	m.SetBody("text/html",message.Mail.Message)//邮件内容

	d := gomail.NewDialer("smtp.exmail.qq.com", 465, from, password) // 发送邮件服务器、端口、发件人账号、发件人密码
	if err := d.DialAndSend(m); err != nil {
		log.Debug(err)
		return "发送失败!"
	} else {
		log.Info("邮件发送结果", err)
	}

	return "发送成功!"
}

/*
拼装消息体
 */
func formatMessage(sentry protos.Sentry) string {
	if (sentry.Type == "text") {
		str := "标题:"+sentry.ProjectName+"\n"
		str += "Level:"+sentry.Level+"\n"
		str += "Time:"+time.Now().Format("2006-01-02 15:04:05")+"\n"
		str += "Info:"+sentry.Message+"\n"
		str += "Href:"+sentry.Href+"\n"
		return str
	} else {
		str := "> #### 标题:"+sentry.ProjectName+"\n"
		str += "##### Level:<font color=\"warning\">"+sentry.Level+"</font> \n"
		str += "##### Time:<font color=\"comment\">"+sentry.Time+"</font> \n"
		str += "##### Info:"+sentry.Message+"\n"
		str += "##### Href:"+sentry.Href+"\n"
		return str
	}
}

/*
推送策略
*/
func sendPeople(project string) map[string]string {

	projectName := map[string][]string{
		"backend":[]string{"ai-x","aiadmin","aiboss","aigeneral","aiplugin","aitemplate","aix","apiboss","apigeneral","dev-aiadmin","dev-aiboss","dev-aigeneral","dev-aix","dev-wxapi","django","front-aigeneral","internal","santi-dev-aiapi","santi-dev-boss","santi-dev-wxapi","st_aiboss","st_general","st-general","stplugin","test-aiadmin","test-aiboss","test-aigeneral","test-aix","test-wxapi","wxapi",},
		"frontend":[]string{},
	}
	var strat string
	for key,val := range projectName{
		for _,v := range val {
			if v == project {
				strat = key
			}
		}
	}
	switch strat {
		case "backend":
			return map[string]string{
				"touser":"",
				"toparty":"",
				"totag":"1",
			}
		case "frontend":
			return map[string]string{
				"touser":"",
				"toparty":"",
				"totag":"2",
			}
		default:
			return map[string]string{
				"touser":"",
				"toparty":"11",
				"totag":"",
			}
	}
}
/**
邮件接收人员分组
 */
func mailPeople(project string) map[string][]string {
	switch project {
	case "aix-wxpp-audit":
		return map[string][]string{
			"to":[]string{"2360655955@qq.com"},//接收人
			//"cc":[]string{},//抄送人
		}
	default:
		return map[string][]string{
			"to":[]string{"2360655955@qq.com","576952208@qq.com"},//接收人
			//"cc":[]string{"2360655955@qq.com","2016331837@qq.com"},//抄送人

		}
	}
}