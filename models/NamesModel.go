package models

import (
	"database/sql"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

//CREATE TABLE `aea_names` (
//`id` mediumint(8) NOT NULL AUTO_INCREMENT,
//`name` varchar(256) DEFAULT '' COMMENT 'name',
//`name_hash` varchar(256) DEFAULT '' COMMENT 'name_hash',
//`tx_hash` varchar(256) DEFAULT '' COMMENT 'tx_hash',
//`created_at_height` int(10) DEFAULT '0' COMMENT 'created_at_height',
//`auction_end_height` int(10) DEFAULT '0' COMMENT 'auction_end_height',
//`owner` varchar(256) DEFAULT '' COMMENT 'owner',
//`expires_at` int(10) DEFAULT '0' COMMENT 'expires_at',
//PRIMARY KEY (`id`)
//) ENGINE=InnoDB AUTO_INCREMENT=501 DEFAULT CHARSET=utf8 COMMENT='names表';
type Names struct {
	Id               int64  `orm:"column(id)" json:"id"`
	Name             string `orm:"column(name)" json:"name"`
	NameHash         string `orm:"column(name_hash)" json:"name_hash"`
	TxHash           string `orm:"column(tx_hash)" json:"tx_hash"`
	CreatedAtHeight  int64  `orm:"column(created_at_height)" json:"created_at_height"`
	AuctionEndHeight int64  `orm:"column(auction_end_height)" json:"auction_end_height"`
	Owner            string `orm:"column(owner)" json:"owner"`
	ExpiresAt        int64  `orm:"column(expires_at)" json:"expires_at"`
}
type NamesJson struct {
	Id               int64  `orm:"column(id)" json:"id"`
	Name             string `orm:"column(name)" json:"name"`
	NameHash         string `orm:"column(name_hash)" json:"name_hash"`
	TxHash           string `orm:"column(tx_hash)" json:"tx_hash"`
	CreatedAtHeight  int64  `orm:"column(created_at_height)" json:"created_at_height"`
	AuctionEndHeight int64  `orm:"column(auction_end_height)" json:"auction_end_height"`
	Owner            string `orm:"column(owner)" json:"owner"`
	ExpiresAt        int64  `orm:"column(expires_at)" json:"expires_at"`
	ExpirationTime   string ` json:"expires_time"`
}

type NamesJsonRegister []NamesJson

func (s NamesJsonRegister) Len() int { return len(s) }

func (s NamesJsonRegister) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

func (s NamesJsonRegister) Less(i, j int) bool { return s[i].ExpiresAt < s[j].ExpiresAt }


type NamesJsonActivity []NamesJson

func (s NamesJsonActivity) Len() int { return len(s) }

func (s NamesJsonActivity) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

func (s NamesJsonActivity) Less(i, j int) bool { return s[i].AuctionEndHeight < s[j].AuctionEndHeight }

func InsertNames(names []Names) {
	for i := 0; i < len(names); i++ {
		_, _ = orm.NewOrm().Insert(&names[i])
	}
}

//获取即将过期的域名
func GetNamesAllHeight(page int, top int) ([]Names, error) {
	var names []Names
	ids := []int{top, top, 49589, (page - 1) * 20, page * 20}
	o := orm.NewOrm()
	_, err := o.Raw("select * from aea_names where expires_at > ? and auction_end_height < ?  and expires_at - auction_end_height > ? order by expires_at limit ?,?", ids).QueryRows(&names)
	return names, err
}

//获取地址下已注册成功的域名
func GetNamesAllMyRegister(top int, address string) ([]Names, error) {
	var names []Names
	o := orm.NewOrm()
	_, err := o.Raw("select * from aea_names where owner = ? and auction_end_height< ? and (expires_at > ? or auction_end_height+49589 > ?) order by expires_at", address, top, top, top).QueryRows(&names)
	return names, err
}

//获取地址下拍卖的域名
func GetNamesAllMyActivity(top int, address string) ([]Names, error) {
	var names []Names
	o := orm.NewOrm()
	_, err := o.Raw("select * from aea_names where owner = ? and auction_end_height> ? and (expires_at > ? or auction_end_height+49589 > ?) order by expires_at", address, top, top, top).QueryRows(&names)
	return names, err
}

//获取即将结束拍卖的域名int
func GetNamesOver(top int,page int) ([]Names, error) {
	var names []Names
	o := orm.NewOrm()
	_, err := o.Raw("select * from aea_names where created_at_height < ? and auction_end_height > ?  order by auction_end_height  limit ?,?", top, top,(page - 1) * 20, page * 20).QueryRows(&names)
	return names, err
}



func GetNamesAllTime(page int) ([]Names, error) {
	var names []Names
	qs := orm.NewOrm().QueryTable("aea_names")
	_, err := qs.Limit(20).Offset((page - 1) * 20).OrderBy("-created_at_height").All(&names)
	return names, err
}

func GetNamesAllOwner(owner string) ([]Names, error) {
	var names []Names
	qs := orm.NewOrm().QueryTable("aea_names")
	_, err := qs.Filter("owner", owner).All(&names)
	return names, err
}

func DeleteAllNames() {
	host := beego.AppConfig.String("db::host")
	port := beego.AppConfig.String("db::port")
	dbname := beego.AppConfig.String("db::databaseName")
	user := beego.AppConfig.String("db::userName")
	pwd := beego.AppConfig.String("db::password")
	dbconnect := user + ":" + pwd + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8"
	db, _ := sql.Open("mysql", dbconnect)
	////root为数据库用户名，后面为密码，tcp代表tcp协议，test处填写自己的数据库名称
	//o := orm.NewOrm()
	//for i := 0; i < 500; i++ {
	//	_, e := o.Delete(&Top500{Id: int64(i)})
	//	if e != nil {
	//		fmt.Println(e)
	//	}
	//}

	db.Exec("Delete from aea_names")
	db.Exec("ALTER TABLE aea_names auto_increment=1")
}

func (email *Names) TableName() string {
	return "aea_names"
}

func RegisterNamesDB() {
	orm.RegisterModel(new(Names))
}
