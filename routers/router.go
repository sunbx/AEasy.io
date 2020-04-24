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
	//token创建
	beego.Router("/token", &controllers.TokenController{})
	//token详情
	beego.Router("/token/info", &controllers.TokenInfoController{})
	//token详情
	beego.Router("/article/info", &controllers.ArticleInfoController{})

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
	//token转账 AEX9
	beego.Router("/token/transfer", &controllers.TransferController{})

	//抓取数据
	//文章数据抓取
	beego.Router("/article/data", &controllers.ArticleDataController{})
	//钱包排行榜数据抓取
	beego.Router("/wealth/data", &controllers.WealthDataController{})
	//aens 数据抓取
	beego.Router("/names/data", &controllers.NameshDataController{})

	//api调用
	//文章列表
	beego.Router("/article/list", &controllers.ArticleListController{})
	//钱包排行榜
	beego.Router("/wealth/list", &controllers.WealthListController{})
	//基础数据
	beego.Router("/base/data", &controllers.BaseDataController{})

	//即将结束拍卖
	beego.Router("/names/auctions", &controllers.NamesAuctionsActiveController{})
	//即将过期未续费的域名
	beego.Router("/names/overdue", &controllers.NamesOverdueController{})
	//最新注册的域名
	beego.Router("/names/new", &controllers.NamesNewController{})
	//获取我的已注册域名
	beego.Router("/names/my/register", &controllers.NamesMyRegisterController{})
	//获取我的拍卖中域名
	beego.Router("/names/my/activity", &controllers.NamesMyActivityController{})

}
