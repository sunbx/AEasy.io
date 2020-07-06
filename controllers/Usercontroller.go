package controllers

import (
	"ae/models"
	"github.com/beego/i18n"
)

type UserRegisterController struct {
	BaseController
}
type UserLoginController struct {
	BaseController
}
type UserLogoutController struct {
	BaseController
}

func (c *UserRegisterController) Post() {
	email := c.GetString("email")
	captcha := c.GetString("captcha")
	password := c.GetString("password")
	addr := c.Ctx.Request.RemoteAddr

	if email == "" || captcha == "" || password == "" || addr == "" {
		c.ErrorJson(-301, i18n.Tr(c.getHeaderLanguage(),"parameter is nul"), JsonData{})
		return
	}

	_, e := models.VerifyEmail(email, captcha, 1)

	if e == nil {
		_, e := models.FindUserEmail(email)
		if e == nil {
			c.ErrorJson(-302, "user exist", JsonData{})
			return
		}

		id, _ := models.InstUser(email, password, addr)
		if id > 0 {
			c.SetSession("user_id", int(id))
			user, e := models.FindUserID(id)
			if e == nil {
				c.SuccessJson(user)
			} else {
				c.ErrorJson(-303, "register user error", JsonData{})
			}
		} else {
			c.ErrorJson(-303, "register user error", JsonData{})
		}
	} else {
		c.ErrorJson(-304, "captcha error", JsonData{})
	}
}

func (c *UserLoginController) Post() {
	email := c.GetString("email")
	password := c.GetString("password")
	addr := c.Ctx.Request.RemoteAddr

	if email == "" || password == "" || addr == "" {
		c.ErrorJson(-301, i18n.Tr(c.getHeaderLanguage(),"parameter is nul"), JsonData{})
		return
	}

	user, e := models.FindUserEmailPassword(email, password)
	if e == nil {
		c.SetSession("user_id", int(user.ID))
		c.SuccessJson(user)
	} else {
		c.ErrorJson(-303, "username or password error", JsonData{})
	}

}

func (c *UserLogoutController) Post() {
	c.SetSession("user_id", "")
	c.SuccessJson(JsonData{})
}
