package controllers


import (
	"ae/models"
	"ae/utils"
	"fmt"
	"strconv"
)
//发送邮箱验证码
type MailSendController struct {
	BaseController
}

//发送验证码
func (c *MailSendController) Post() {
	email := c.GetString("email")
	t := c.GetString("type")
	//code := c.GetString("code")
	addr := c.Ctx.Request.RemoteAddr
	//captcha := "8888"
	captcha := utils.CreateCaptcha()

	//if true{
	//	c.ErrorJson(-301, "未开放注册", JsonData{})
	//}

	if email == ""  || captcha == "" || t == "" || addr == "" {
		c.ErrorJson(-301, "parameter is nul", JsonData{})
		return
	}

	//if !utils.IsEmail(email) {
	//	c.ErrorJson(-310, "Mailbox format error", JsonData{})
	//	return
	//}

	//转成成int
	tInt, err := strconv.Atoi(t)

	//如果转换失败直接报错
	if err != nil {
		c.ErrorJson(-201, "error type is null", JsonData{})
		return
	}

	//判断是不是规定的type类型
	if tInt > 2 || tInt < 1 {
		c.ErrorJson(-202, "error type is not 1", JsonData{})
		return
	}
	//查询有没有发过验证码
	_, e := models.VerifyIpEmail(email, addr)
	//表示没有查到验证码
	if e != nil {

		if !utils.IsEmail(email) {
			res, _ := utils.SendRegisterSms(email, captcha)
			fmt.Println(email)
			if res.Code != "OK" {
				c.ErrorJson(-203, res.Message, JsonData{})
				return
			}
			//插入邮箱信息到数据库
			instEmail, e := models.InstEmail(email, addr, tInt, captcha)
			if e == nil && instEmail > 0 {
				c.SuccessJson(JsonData{})
				return
			} else {
				c.ErrorJson(-203, "Send captcha error", JsonData{})
				return
			}
		} else {
			bool := utils.SendEMail(email, captcha)
			if !bool {
				c.ErrorJson(-203, "Send Email error", JsonData{})
				return
			}
			//插入邮箱信息到数据库
			instEmail, e := models.InstEmail(email, addr, tInt, captcha)
			if e == nil && instEmail > 0 {
				c.SuccessJson(JsonData{})
				return
			} else {
				c.ErrorJson(-203, "Send captcha error", JsonData{})
				return
			}
		}
	} else {
		c.ErrorJson(-204, "Captcha acquisition is frequent", JsonData{})
		return
	}
}
