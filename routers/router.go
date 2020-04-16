package routers

import (
	"ae/controllers"
	"github.com/astaxie/beego"
)

func init() {
	//====================静态页====================
	//首页
	beego.Router("/", &controllers.MainController{})
	//登录页面
	beego.Router("/login", &controllers.LoginController{})
	//个人页面
	beego.Router("/user", &controllers.UserController{})
	//授权页面
	beego.Router("/accredit", &controllers.AccreditController{})
	//授权页面-绑定手机号
	beego.Router("/accredit_bind", &controllers.AccreditBindController{})
	//授权页面-pay
	beego.Router("/pay", &controllers.PayController{})
	//token页面
	beego.Router("/token", &controllers.TokenController{})
	//token页面
	beego.Router("/token/transfer", &controllers.TransferController{})
	beego.Router("/token/info", &controllers.TokenInfoController{})

	//====================网站相关====================
	//验证码发送
	beego.Router("/mail/send", &controllers.MailSendController{})
	//用户注册
	beego.Router("/user/register", &controllers.UserRegisterController{})
	//用户登录
	beego.Router("/user/login", &controllers.UserLoginController{})
	//用户登出
	beego.Router("/user/logout", &controllers.UserLogoutController{})

	//====================oauth2授权相关====================
	//授权登录
	beego.Router("/accredit/login", &controllers.AccreditLoginController{})
	//授权注册
	beego.Router("/accredit/register", &controllers.AccreditRegisterController{})
	//邮箱绑定
	beego.Router("/accredit/bind", &controllers.AccreditBindEmailController{})
	//通过code获取access_token
	beego.Router("/accredit/access_token", &controllers.AccreditAccessTokenController{})
	//获取account账户余额
	beego.Router("/accredit/info", &controllers.AccreditInfoController{})
	//创建订单
	beego.Router("/accredit/create_order", &controllers.AccreditCreateOrderController{})
	//授权页面-pay
	beego.Router("/pay/buy", &controllers.PayBuyController{})
	//合约token创建
	beego.Router("/token/create", &controllers.TokenCreateController{})
	//====================ae api相关====================
	//查询当前区块高度
	beego.Router("/ae/api/blocks_top", &controllers.ApiBlocksTopController{})
	//查询th_hash
	beego.Router("/ae/api/th_hash", &controllers.ApiThHashController{})
	//数据上链
	beego.Router("/ae/api/block_data", &controllers.ApiTransferController{})
	//账户余额
	beego.Router("/ae/api/balance", &controllers.ApiBalanceController{})
	//创建账户
	beego.Router("/ae/api/create_account", &controllers.ApiCreateAccountController{})

	//创建账户
	beego.Router("/test", &controllers.TestController{})
	beego.Router("/test2", &controllers.TestController2{})
}
