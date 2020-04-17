package controllers

import (
	"ae/models"
	"encoding/json"
	"github.com/aeternity/aepp-sdk-go/account"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	"time"
)

type BaseController struct {
	beego.Controller
}

type ReturnMsg struct {
	Code  int         `json:"code"`
	Msg   string      `json:"msg"`
	Stime int64       `json:"time"`
	Data  interface{} `json:"data"`
}

type JsonData struct {
}

var bm, _ = cache.NewCache("file", `{"CachePath":"./cache","FileSuffix":".cache","DirectoryLevel":"2","EmbedExpiry":"120"}`)

func (c *BaseController) SuccessJson(data interface{}) {
	serviceTime := time.Now().UnixNano() / 1e6
	res := ReturnMsg{
		200, "success", serviceTime, data,
	}
	jsons, _ := json.Marshal(res)

	c.Ctx.WriteString(string(jsons))
}

func (c *BaseController) ErrorJson(code int, msg string, data interface{}) {
	serviceTime := time.Now().UnixNano() / 1e6
	res := ReturnMsg{
		code, msg, serviceTime, data,
	}
	jsons, _ := json.Marshal(res)

	c.Ctx.WriteString(string(jsons))
}

func (c *BaseController) isLogin() bool {
	v := c.GetSession("user_id")
	if value, ok := v.(int); ok == true {
		if value > 0 {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

func (c *BaseController) getCurrentUserId() int {
	v := c.GetSession("user_id")
	if value, ok := v.(int); ok == true {
		if value > 0 {
			return value
		} else {
			return -1
		}
	} else {
		return value
	}
}

func (c *BaseController) getAccessTokenOpenId(token string) string {
	openId := bm.Get(token)
	if value, ok := openId.(string); ok == true {
		if value == "" {
			return ""
		} else {
			//更新时间
			_ = bm.Put(token, value, 30*24*60*60*time.Second)
			return value
		}
	} else {
		return ""
	}
}

func (c *BaseController) GetSecretAeAccount() *account.Account {
	var appId string
	var appSecret string

	appId = c.GetString("app_id")
	appSecret = c.GetString("app_secret")

	if appId == "" || appSecret == "" {
		appId = c.Ctx.Input.Header("aeasy-APP-ID")
		appSecret = c.Ctx.Input.Header("aeasy-APP-SECRET")
	}

	if appId == "" || appSecret == "" {
		return nil
	} else {
		secret, e := models.FindSecretIdSecret(appId, appSecret)
		if e == nil {
			stringAccount, _ := models.SigningKeyHexStringAccount(secret.SigningKey)
			if stringAccount != nil {
				return stringAccount
			}
		}
	}
	return nil
}

func (c *BaseController) verifySecret() bool {
	var appId string
	var appSecret string

	appId = c.GetString("app_id")
	appSecret = c.GetString("app_secret")

	if appId == "" || appSecret == "" {
		appId = c.Ctx.Input.Header("aeasy-APP-ID")
		appSecret = c.Ctx.Input.Header("aeasy-APP-SECRET")
	}

	if appId == "" || appSecret == "" {
		return false
	} else {
		appSecretCache := bm.Get(appId)

		if value, ok := appSecretCache.(string); ok == true {
			if appSecret == value {
				return true
			}
		}
		_, e := models.FindSecretIdSecret(appId, appSecret)
		if e == nil {
			_ = bm.Put(appId, appSecret, 24*time.Hour)
			return true
		} else {
			return false
		}
	}
}

func (c *BaseController) verifyAppId() bool {
	var appId string
	appId = c.GetString("app_id")
	if appId == "" {
		return false
	} else {
		appSecretCache := bm.Get(appId)
		if value, ok := appSecretCache.(string); ok == true {
			if "yes" == value {
				return true
			}
		}
		_, e := models.FindSecretId(appId)
		if e == nil {
			_ = bm.Put(appId, "yes", 24*time.Hour)
			return true
		} else {
			return false
		}
	}
}
