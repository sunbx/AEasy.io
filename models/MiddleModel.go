package models

import (
	"github.com/astaxie/beego/orm"
)

type AeaMiddleBlock struct {
	Id          int64  `orm:"column(id)" json:"id"`
	Beneficiary string `orm:"column(beneficiary)" json:"beneficiary"`
	Hash        string `orm:"column(hash)" json:"hash"`
	Height      int64  `orm:"column(height)" json:"height"`
	MicroBlocks string `orm:"column(micro_blocks)" json:"micro_blocks"`
	Miner       string `orm:"column(miner)" json:"miner"`
	PrevHash    string `orm:"column(prev_hash)" json:"prev_hash"`
	PrevKeyHash string `orm:"column(prev_key_hash)" json:"prev_key_hash"`
	StateHash   string `orm:"column(state_hash)" json:"state_hash"`
	Target      int64  `orm:"column(target)" json:"target"`
	Time        int64  `orm:"column(time)" json:"time"`
	Version     int64  `orm:"column(version)" json:"version"`
}

// TableName sets the insert table name for this struct type
func (a *AeaMiddleBlock) TableName() string {
	return "aea_middle_block"
}

type AeaMiddleMicroBlock struct {
	Id          int64                  `orm:"column(id)" json:"-"`
	BlockHash   string                 `orm:"column(block_hash)" json:"block_hash"`
	BlockHeight int64                  `orm:"column(block_height)" json:"block_height"`
	Hash        string                 `orm:"column(hash)" json:"hash"`
	RecipientId string                 `orm:"column(recipient_id)" json:"recipient_id"`
	Signatures  string                 `orm:"column(signatures)" json:"signatures"`
	Tx          string                 `orm:"column(tx)" json:"-"`
	Time        int64                  `orm:"column(time)" json:"time"`
	NameId      string                 `orm:"column(name_id)" json:"name_id"`
	Name        string                 `orm:"column(name)" json:"name"`
	AccountId   string                 `orm:"column(account_id)" json:"account_id"`
}

// TableName sets the insert table name for this struct type
func (a *AeaMiddleMicroBlock) TableName() string {
	return "aea_middle_micro_block"
}

//主块库插入信息
func InsertAeaMiddleBlock(beneficiary string,
	hash string,
	height int64,
	microBlocks string,
	miner string,
	prevHash string,
	prevKeyHash string,
	stateHash string,
	target int64,
	time int64,
	version int64) (AeaMiddleBlock, error) {
	middleBlock := AeaMiddleBlock{
		Beneficiary: beneficiary,
		Hash:        hash,
		Height:      height,
		MicroBlocks: microBlocks,
		Miner:       miner,
		PrevHash:    prevHash,
		PrevKeyHash: prevKeyHash,
		StateHash:   stateHash,
		Target:      target,
		Time:        time,
		Version:     version,
	}
	_, err := orm.NewOrm().Insert(&middleBlock)
	return middleBlock, err
}

func InsertAeaMiddleMicroBlockBlock(
	blockHash string,
	blockHeight int64,
	hash string,
	signatures string,
	tx string, time int64) (AeaMiddleMicroBlock, error) {
	middleBlock := AeaMiddleMicroBlock{
		BlockHash:   blockHash,
		BlockHeight: blockHeight,
		Hash:        hash,
		Signatures:  signatures,
		Tx:          tx,
		Time:        time,
	}
	//_, err := orm.NewOrm().Insert(&middleBlock)

	o := orm.NewOrm()
	_, err := o.Raw("insert into aea_middle_micro_block (`block_hash`,`block_height`,`hash`,`signatures`,`tx`,`time`) values(?,?,?,?,?,?)", blockHash, blockHeight, hash, signatures, tx, time).Exec()

	//if err == nil {
	//	fmt.Println(result.RowsAffected())
	//}
	return middleBlock, err
}

//通过id 获取用户信息
func InsertAeaMiddleMicroBlockBlockObj(hash string, height string, tx string) error {
	o := orm.NewOrm()
	_, err := o.QueryTable("aea_middle_micro_block").Filter("hash", hash).Filter("block_height", height).Update(orm.Params{
		"tx": tx,
	})
	return err
}

//通过id 获取用户信息
func FindBlockHeight() (int64, error) {
	var aeaMiddleBlock []AeaMiddleBlock
	o := orm.NewOrm()
	_, err := o.Raw("select * from aea_middle_block order by height desc limit 1").QueryRows(&aeaMiddleBlock)
	return aeaMiddleBlock[0].Height, err
}

//通过id 获取用户信息
func FindMicroBlockBlockTimeUpdate() ([]AeaMiddleMicroBlock, error) {
	var aeaMiddleMicroBlock []AeaMiddleMicroBlock
	o := orm.NewOrm()
	_, err := o.Raw("SELECT hash , block_height FROM `aea_middle_micro_block` LIMIT 1000").QueryRows(&aeaMiddleMicroBlock)
	return aeaMiddleMicroBlock, err
}

//通过nameid 获取name 注册成功后的信息并且是转移域名的信息
func FindMicroBlockNameIdData(nameId string) (AeaMiddleMicroBlock, error) {
	var aeaMiddleMicroBlock AeaMiddleMicroBlock
	o := orm.NewOrm()
	err := o.Raw("SELECT * FROM `aea_middle_micro_block` where  `recipient_id` !='' and type = 'NameTransferTx' and `name_id` =? order by `time` desc", nameId).QueryRow(&aeaMiddleMicroBlock)
	return aeaMiddleMicroBlock, err
}

//通过nameid 获取name 注册成功后的信息并且是更新域名的信息
func FindMicroBlockNameIdUpdate(nameId string) (AeaMiddleMicroBlock, error) {
	var aeaMiddleMicroBlock AeaMiddleMicroBlock
	o := orm.NewOrm()
	err := o.Raw("SELECT * FROM `aea_middle_micro_block` where  type = 'NameUpdateTx' and `name_id` =? order by `time` desc", nameId).QueryRow(&aeaMiddleMicroBlock)
	return aeaMiddleMicroBlock, err
}

//通过id 获取用户信息
func FindMicroBlockBlockNames() ([]AeaMiddleMicroBlock, error) {
	var aeaMiddleMicroBlock []AeaMiddleMicroBlock
	o := orm.NewOrm()
	_, err := o.Raw("SELECT name FROM `aea_middle_micro_block` WHERE `name` !='null' and `type`  = 'NameClaimTx' group by name ").QueryRows(&aeaMiddleMicroBlock)
	return aeaMiddleMicroBlock, err
}

//通过id 获取用户信息
func FindMicroBlockBlockNameorData(name string) (AeaMiddleMicroBlock, error) {
	var aeaMiddleMicroBlock []AeaMiddleMicroBlock
	o := orm.NewOrm()
	_, err := o.Raw("select block_height,hash,account_id,tx from `aea_middle_micro_block` WHERE `name` ='" + name + "' order by block_height desc limit 1").QueryRows(&aeaMiddleMicroBlock)
	return aeaMiddleMicroBlock[0], err
}

func FindMicroBlockBlockNameorDatas(name string) ([]AeaMiddleMicroBlock, error) {
	var aeaMiddleMicroBlock []AeaMiddleMicroBlock
	o := orm.NewOrm()
	_, err := o.Raw("select * from `aea_middle_micro_block` WHERE `name` ='" + name + "' order by block_height desc limit 100").QueryRows(&aeaMiddleMicroBlock)
	return aeaMiddleMicroBlock, err
}

func FindMicroBlockBlockList(address string, page int, t string) ([]AeaMiddleMicroBlock, error) {
	var aeaMiddleMicroBlock []AeaMiddleMicroBlock
	o := orm.NewOrm()
	if t == "all" {
		_, err := o.Raw("select * from aea_middle_micro_block where (account_id = ? or sender_id=? or recipient_id = ? or caller_id = ?) order by time desc limit ?,?", address, address, address,address, (page-1)*20, 20).QueryRows(&aeaMiddleMicroBlock)
		return aeaMiddleMicroBlock, err
	} else {
		_, err := o.Raw("select * from aea_middle_micro_block where (account_id = '?' or sender_id='?' or recipient_id = '?' or caller_id = ?) and type=? order by time desc limit ?,?", address, address, address,address, t, (page-1)*20, 20).QueryRows(&aeaMiddleMicroBlock)
		return aeaMiddleMicroBlock, err
	}
}

func RegisterAeaMiddleBlockDB() {
	orm.RegisterModel(new(AeaMiddleBlock))
}

func RegisterAeaMiddleMicroBlockDB() {
	orm.RegisterModel(new(AeaMiddleMicroBlock))
}
