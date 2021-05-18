package models

import (
	"github.com/astaxie/beego/orm"
)

type Contract struct {
	Id                 int64   `orm:"column(id)" json:"-"`
	BlockHash          string  `orm:"column(block_hash)" json:"block_hash"`
	BlockHeight        int64   `orm:"column(block_height)" json:"block_height"`
	Hash               string  `orm:"column(hash)" json:"hash"`
	Function           string  `orm:"column(function)" json:"function"`
	DecodeJson         string  `orm:"column(decode_json)" json:"-"`
	ContractId         string  `orm:"column(contract_id)" json:"contract_id"`
	CallAddress        string  `orm:"column(call_address)" json:"call_address"`
	Aex9Amount         float64 `orm:"column(aex9_amount)" json:"-"`
	Aex9AmountString   string  `orm:"-"json:"aex9_amount"`
	Amount             float64 `orm:"column(amount)" json:"-"`
	AmountString       string  `orm:"-"json:"amount"`
	Fee                float64 `orm:"column(fee)" json:"-"`
	FeeString          string  `orm:"-" json:"fee"`
	ResultType         string  `orm:"column(result_type)" json:"result_type"`
	Aex9ReceiveAddress string  `orm:"column(aex9_receive_address)" json:"aex9_receive_address"`
	CreateTime         int64   `orm:"column(create_time)" json:"create_time"`
}

func InsertContract(
	blockHash string,
	blockHeight int64,
	hash string,
	function string,
	decodeJson string,
	contractId string,
	callAddress string,
	aex9Amount float64,
	amount float64,
	fee float64,
	resultType string,
	aex9ReceiveAddress string,
	createTime int64) (Contract, error) {

	contract := Contract{
		BlockHash:          blockHash,
		BlockHeight:        blockHeight,
		Hash:               hash,
		Function:           function,
		DecodeJson:         decodeJson,
		ContractId:         contractId,
		CallAddress:        callAddress,
		Aex9Amount:         aex9Amount,
		Amount:             amount,
		Fee:                fee,
		ResultType:         resultType,
		Aex9ReceiveAddress: aex9ReceiveAddress,
		CreateTime:         createTime,
	}
	_, err := orm.NewOrm().Insert(&contract)
	return contract, err
}

//
func GetAEX9RecordAll(page int, address string, contractId string) ([]Contract, error) {
	var contracts []Contract
	function := "transfer"
	resultType := "ok"
	o := orm.NewOrm()
	_, err := o.Raw("select * from `aea_contracts` where(`call_address`= ? or `aex9_receive_address`= ?)  AND `function` = ? AND `contract_id` = ? AND `result_type` = ? order by `create_time` desc limit ?,?", address, address, function, contractId, resultType,(page-1)*20, 20).QueryRows(&contracts)
	return contracts, err
}

//
//func FindArticleId(articleId string) (Article, error) {
//	var article Article
//	qs := orm.NewOrm().QueryTable("aea_article")
//	err := qs.
//		Filter("article_id", articleId).
//		One(&article)
//	return article, err
//}
//
func (contract *Contract) TableName() string {
	return "aea_contracts"
}

func RegisterContractDB() {
	orm.RegisterModel(new(Contract))
}
