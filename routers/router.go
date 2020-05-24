package routers

import (
	"ae/controllers"
	"github.com/astaxie/beego"
)

func init() {
	//====================网站静态页====================
	//首页
	beego.Router("/", &controllers.MainController{})

	//切换语言
	beego.Router("/language", &controllers.LanguageController{})

	//登录页面
	beego.Router("/login", &controllers.LoginController{})

	//个人页面
	beego.Router("/user", &controllers.UserController{})
	beego.Router("/show", &controllers.ShowController{})

	//授权页面
	beego.Router("/accredit", &controllers.AccreditController{})

	//token创建和展示
	beego.Router("/token", &controllers.TokenController{})

	//文章详情
	beego.Router("/article/info", &controllers.ArticleInfoController{})

	//====================网站相关接口====================

	//验证码发送
	beego.Router("/mail/send", &controllers.MailSendController{})

	//用户注册
	beego.Router("/user/register", &controllers.UserRegisterController{})

	//用户登录
	beego.Router("/user/login", &controllers.UserLoginController{})

	//用户登出
	beego.Router("/user/logout", &controllers.UserLogoutController{})

	//合约token创建
	beego.Router("/token/create", &controllers.TokenCreateController{})

	//合约token转账 AEX9
	beego.Router("/token/transfer", &controllers.TokenTransferController{})

	//====================API基础接口====================

	//数据上链
	beego.Router("/api/ae/block_data", &controllers.ApiTransferController{})

	//查询当前区块高度
	beego.Router("/api/ae/block_top", &controllers.ApiBlocksTopController{})

	//查询th_hash
	beego.Router("/api/ae/th_hash", &controllers.ApiThHashController{})

	//ae china文章列表
	beego.Router("/api/article/list", &controllers.ArticleListController{})

	//钱包排行榜
	beego.Router("/api/wallet/list", &controllers.WalletListController{})

	//ae 价格等一些基础数据
	beego.Router("/api/base/data", &controllers.BaseDataController{})

	//====================API高级接口====================

	//助记词登录
	beego.Router("/api/user/login", &controllers.AccreditLoginController{})

	//助记词注册
	beego.Router("/api/user/register", &controllers.AccreditRegisterController{})

	//转账
	beego.Router("/api/wallet/transfer", &controllers.WalletTransferController{})

	//转账记录
	beego.Router("/api/wallet/transfer/record", &controllers.WalletTransferRecordController{})

	//获取account账户余额
	beego.Router("/api/user/info", &controllers.AccreditInfoController{})

	//拍卖中 - 即将结束拍卖
	beego.Router("/api/names/auctions", &controllers.NamesAuctionsActiveController{})

	//拍卖中 - 价格最贵的域名
	beego.Router("/api/names/price", &controllers.NamesPriceController{})

	//即将过期 未续费的域名
	beego.Router("/api/names/over", &controllers.NamesOverController{})

	//我的 - 已注册的域名
	beego.Router("/api/names/my/register", &controllers.NamesMyRegisterController{})

	//我的 - 即将过期的域名
	beego.Router("/api/names/my/over", &controllers.NamesMyOverController{})

	//更新域名
	beego.Router("/api/names/update", &controllers.NamesUpdateController{})

	//域名详细信息
	beego.Router("/api/names/info", &controllers.NamesInfoController{})

	//域名注册
	beego.Router("/api/names/add", &controllers.NamesAddController{})

	//域名转移
	beego.Router("/api/names/transfer", &controllers.NamesTransferController{})

	//所有域名基础信息
	beego.Router("/api/names/base", &controllers.NamesBaseController{})

	//====================定时任务====================

	//文章数据抓取
	beego.Router("/article/data", &controllers.ArticleDataController{})

	beego.Router("/test", &controllers.TestController{})
}
