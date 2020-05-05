package main

import (
	"ae/models"
	"ae/utils"
	"github.com/shopspring/decimal"
	"strings"

	//"fmt"
	//"github.com/astaxie/beego/toolbox"
	//"time"
	"encoding/json"
	"fmt"
	"strconv"
)

func SynAeBlock() {
	//查询当前同步的最新高度
	dbHeight, _ := models.FindBlockHeight()
	//dbHeight = 248022
	//获取AE的最新高度
	aeHeight, _ := strconv.Atoi(strconv.FormatUint(models.ApiBlocksTop(), 10))

	if dbHeight >= int64(aeHeight) {
		fmt.Println(aeHeight, "最新区块,没有产生新区块")
		return
	}
	fmt.Println(aeHeight, "---", dbHeight)
	for i := dbHeight; i <= int64(aeHeight); i++ {
		//从 node 当前高度的区块
		response := utils.Get("http://node.aechina.io:3013/v2/generations/height/" + strconv.Itoa(int(i)))
		//解析区块信息为实体
		var block Block
		err := json.Unmarshal([]byte(response), &block)
		if err != nil {
			fmt.Println("aea_middle_block 主块JSON转换失败 => height->" + strconv.Itoa(int(i)) + " " + err.Error())
		}

		//微块信息转移成json
		microJson, err := json.Marshal(block.MicroBlocks)

		//插入区块高度库
		_, err = models.InsertAeaMiddleBlock(block.KeyBlock.Beneficiary, block.KeyBlock.Hash, block.KeyBlock.Height, string(microJson), block.KeyBlock.Miner, block.KeyBlock.PrevHash, block.KeyBlock.PrevKeyHash, block.KeyBlock.StateHash, block.KeyBlock.Target, block.KeyBlock.Time, block.KeyBlock.Version)
		if err != nil {
			fmt.Println("aea_middle_block 主块数据库插入失败 => height->" + strconv.Itoa(int(i)) + " " + err.Error())
		}
		fmt.Println("aea_middle_block 主块SUCESS => height->" + strconv.Itoa(int(i)))

		//获取微块 ID 机型循环
		for j := 0; j < len(block.MicroBlocks); j++ {

			//从 node 获取微块详细信息
			response := utils.Get("http://node.aechina.io:3013/v2/micro-blocks/hash/" + block.MicroBlocks[j] + "/transactions")

			//解析微块信息
			var block MicroBlock
			err := json.Unmarshal([]byte(response), &block)
			if err != nil {
				fmt.Println("aea_middle_micro_block 微块JSON解析失败 => height->" + strconv.Itoa(int(i)) + " MicroBlock -> " + strconv.Itoa(j) + " " + err.Error())
				return
			}

			//解析微块的转账记录
			for k := 0; k < len(block.Transactions); k++ {
				//从note 获取微块信息,主要获取的是time
				response := utils.Get("http://node.aechina.io:3013/v2/micro-blocks/hash/" + block.Transactions[k].BlockHash + "/header")
				var blockHeader BlocksHeader
				err = json.Unmarshal([]byte(response), &blockHeader)
				if err != nil {
					fmt.Println("aea_middle_micro_block 微块转账JSON解析失败 => height->" + strconv.Itoa(int(i)) + " MicroBlock -> " + strconv.Itoa(j) + " " + err.Error())

					return
				}

				//获取 th信息 ,根据类型分别插入库 , 和过滤垃圾信息
				mapObj, err := Obj2map(block.Transactions[k].Tx)
				if err != nil {
					fmt.Println("Obj2map error", err.Error())
					return
				}

				inter := mapObj["type"]

				if inter.(string) == "SpendTx" {
					senderId := mapObj["sender_id"]
					recipientId := mapObj["recipient_id"]
					if senderId.(string) == "ak_zvU8YQLagjcfng7Tg8yCdiZ1rpiWNp1PBn3vtUs44utSvbJVR" ||
						senderId.(string) == "ak_2QkttUgEyPixKzqXkJ4LX7ugbRjwCDWPBT4p4M2r8brjxUxUYd" ||
						senderId.(string) == "ak_wTPFpksUJFjjntonTvwK4LJvDw11DPma7kZBneKbumb8yPeFq" ||
						senderId.(string) == "ak_KHfXhF2J6VBt3sUgFygdbpEkWi6AKBkr9jNKUCHbpwwagzHUs" ||
						senderId.(string) == "ak_kdxBz4kzVot86bcrUMQwisDpA5m1gciycQBqs1Cj7MojxQEGo" {
						//fmt.Println("aea_middle_micro_block 转账账号过滤=>height->", i, "->"+strconv.Itoa(j)+"TH =>"+strconv.Itoa(k)+"-->error:"+"垃圾转账过滤")
						continue
					}
					if recipientId.(string) == "ak_zvU8YQLagjcfng7Tg8yCdiZ1rpiWNp1PBn3vtUs44utSvbJVR" ||
						recipientId.(string) == "ak_2QkttUgEyPixKzqXkJ4LX7ugbRjwCDWPBT4p4M2r8brjxUxUYd" ||
						recipientId.(string) == "ak_wTPFpksUJFjjntonTvwK4LJvDw11DPma7kZBneKbumb8yPeFq" ||
						recipientId.(string) == "ak_KHfXhF2J6VBt3sUgFygdbpEkWi6AKBkr9jNKUCHbpwwagzHUs" ||
						recipientId.(string) == "ak_kdxBz4kzVot86bcrUMQwisDpA5m1gciycQBqs1Cj7MojxQEGo" {
						//fmt.Println("aea_middle_micro_block 转账账号过滤=>height->", i, "->"+strconv.Itoa(j)+"TH =>"+strconv.Itoa(k)+"-->error:"+"垃圾转账过滤")
						continue
					}
					//更新address表,发送和接收的用户都更新一下
					if InsertAddressBlock(senderId.(string), blockHeader) {
						return
					}
					if InsertAddressBlock(recipientId.(string), blockHeader) {
						return
					}

				} else if inter.(string) == "NameClaimTx" {
					if i < 161150 {
						fmt.Println("aea_middle_micro_block AENS 161150测试数据过滤->" + strconv.Itoa(int(i)) + "MicroBlock->" + strconv.Itoa(j) + "TH =>" + strconv.Itoa(k) + " NameClaimTx < 161150")
						continue
					}
					if InsertNameBlock(mapObj, block, k) {
						return
					}

				} else if inter.(string) == "NameUpdateTx" {
					if i < 161150 {
						fmt.Println("aea_middle_micro_block AENS 161150测试数据过滤->" + strconv.Itoa(int(i)) + "MicroBlock->" + strconv.Itoa(j) + "TH =>" + strconv.Itoa(k) + " NameUpdateTx < 161150")
						continue
					}
					if updateNameTimeBlock(i, mapObj) {
						return
					}

				} else if inter.(string) == "NameTransferTx" {
					if i < 161150 {
						fmt.Println("aea_middle_micro_block AENS 161150测试数据过滤->" + strconv.Itoa(int(i)) + "MicroBlock->" + strconv.Itoa(j) + "TH =>" + strconv.Itoa(k) + " NameUpdateTx < NameTransferTx")
						continue
					}
					if updateNameOwnerBlock(mapObj) {
						return
					}

				}

				if len(block.Transactions[k].Signatures) > 0 {
					txJson, err := json.Marshal(block.Transactions[k].Tx)

					_, err = models.InsertAeaMiddleMicroBlockBlock(block.Transactions[k].BlockHash, block.Transactions[k].BlockHeight, block.Transactions[k].Hash, block.Transactions[k].Signatures[0], string(txJson), blockHeader.Time)
					if err != nil {
						fmt.Println("aea_middle_micro_block ERROR 微块转账记录插入失败=>height->", i, strconv.Itoa(j)+"TH =>"+strconv.Itoa(k)+"-->error:"+err.Error())
					} else {
						fmt.Println("aea_middle_micro_block 微块转账记录插入成功 =>height->", i, strconv.Itoa(int(i))+" "+strconv.Itoa(j)+"TH =>"+strconv.Itoa(k))
					}

				}
			}
		}

	}

	if dbHeight == int64(aeHeight) {
		fmt.Println(aeHeight, "最新区块,没有产生新区块")
	}
	fmt.Println("Sucess+" + strconv.Itoa(int(dbHeight)))
}

func updateNameOwnerBlock(mapObj map[string]interface{}) bool {
	nameId := mapObj["name_id"].(string)
	recipientId := mapObj["recipient_id"].(string)
	err := models.UpdateNameOwner(nameId, recipientId)
	if err != nil {
		fmt.Println(err.Error())
		return true
	}
	fmt.Println("aea_middle_micro_block AENS更新成功->", nameId)

	return false
}

func updateNameTimeBlock(i int64, mapObj map[string]interface{}) bool {
	nameId := mapObj["name_id"].(string)
	err := models.UpdateNameHeight(nameId, i+50000)
	if err != nil {
		fmt.Println(err.Error())
		return true
	}
	fmt.Println("aea_middle_micro_block AENS更新成功->", nameId)

	return false
}

func InsertNameBlock(mapObj map[string]interface{}, block MicroBlock, k int) bool {
	name := mapObj["name"].(string)
	response := utils.Get("http://node.aechina.io:3013/v2/names/" + name)
	var v2Name V2Name
	err := json.Unmarshal([]byte(response), &v2Name)
	if err != nil {
		fmt.Println(500, err.Error())
		return true
	}
	var endHeight int
	if len(name)-6 <= 4 {
		endHeight = int(block.Transactions[k].BlockHeight + 29760)
	} else if len(name)-6 >= 5 && len(name)-6 <= 8 {
		endHeight = int(block.Transactions[k].BlockHeight + 14880)
	} else if len(name)-6 >= 9 && len(name)-6 <= 12 {
		endHeight = int(block.Transactions[k].BlockHeight + 480)
	} else {
		endHeight = int(block.Transactions[k].BlockHeight + 0)
	}
	if v2Name.TTL == 0 {
		v2Name.TTL = int64(endHeight + 50000)
	}
	var price string
	if len(name)-6 == 1 {
		price = "570288700000000000000"
	} else if len(name)-6 == 2 {
		price = "352457800000000000000"
	} else if len(name)-6 == 3 {
		price = "217830900000000000000"
	} else if len(name)-6 == 4 {
		price = "134626900000000000000"
	} else if len(name)-6 == 5 {
		price = "83204000000000000000"
	} else if len(name)-6 == 6 {
		price = "51422900000000000000"
	} else if len(name)-6 == 7 {
		price = "31781100000000000000"
	} else if len(name)-6 == 8 {
		price = "19641800000000000000"
	} else if len(name)-6 == 9 {
		price = "12139300000000000000"
	} else if len(name)-6 == 10 {
		price = "7502500000000000000"
	} else if len(name)-6 == 11 {
		price = "4636800000000000000"
	} else if len(name)-6 == 12 {
		price = "2865700000000000000"
	} else if len(name)-6 >= 13 {
		price = "2865700000000000000"
	}
	priceFloat, _ := strconv.ParseFloat(price, 64)
	priceFormat := utils.FormatTokensP(priceFloat, 4)
	f, _ := mapObj["name_fee"].(float64)
	accountId := mapObj["account_id"].(string)
	decimalNum := decimal.NewFromFloat(f)
	priceFloat2, _ := strconv.ParseFloat(decimalNum.String(), 64)
	priceFormat2 := utils.FormatTokensP(priceFloat2, 4)
	fmt.Println("name->" + name)
	fmt.Println("createHeight->" + strconv.Itoa(int(block.Transactions[k].BlockHeight)))
	fmt.Println("endHeight->" + strconv.Itoa(endHeight))
	fmt.Println("orderHeight->" + strconv.Itoa(int(v2Name.TTL)))
	fmt.Println("th->" + block.Transactions[k].Hash)
	fmt.Println("len->" + strconv.Itoa(len(name)-6))
	fmt.Println("name_id->" + v2Name.ID)
	fmt.Println("owner->" + accountId)
	fmt.Println("price->" + priceFormat)
	fmt.Println("priceCurrent->" + priceFormat2)
	//fmt.Println(priceFormat2)
	fmt.Println("=============================")
	var aeaMiddleName models.AeaMiddleNames
	aeaMiddleName.Name = name
	aeaMiddleName.StartHeight = int(block.Transactions[k].BlockHeight)
	aeaMiddleName.EndHeight = endHeight
	aeaMiddleName.OverHeight = int(v2Name.TTL)
	aeaMiddleName.ThHash = block.Transactions[k].Hash
	aeaMiddleName.Length = len(name) - 6
	aeaMiddleName.NameID = v2Name.ID
	aeaMiddleName.Owner = accountId
	aeaMiddleName.Price = priceFloat
	aeaMiddleName.CurrentPrice = priceFloat2
	models.InsertName(aeaMiddleName)
	fmt.Println("aea_middle_names AENS插入成功->", name)

	return false
}

func InsertAddressBlock(senderId string, blockHeader BlocksHeader) bool {
	if strings.Contains(senderId, "nm_") {
		names, e := models.FindNameId(senderId)
		if e != nil {
			fmt.Println("FindNameId - ", e.Error())
			return true
		}
		senderId = names.Owner
	}

	account, e := models.ApiGetAccount(senderId)
	if e != nil {
		fmt.Println(e.Error())
		return true
	}

	var address models.AeaMiddleAddress
	tokens, e := strconv.ParseFloat(account.Balance.String(), 64)
	address.Balance = tokens
	address.Address = senderId
	address.UpdateTime = blockHeader.Time
	models.InsertAddress(address)
	fmt.Println("aea_middle_address address数据库插入成功->", senderId, "->", address.Balance, "->", address.UpdateTime)
	return false
}

func Obj2map(obj interface{}) (mapObj map[string]interface{}, err error) {
	// 结构体转json
	b, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(b, &result); err != nil {
		return nil, err
	}
	return result, nil
}

type Block struct {
	KeyBlock    KeyBlock `json:"key_block"`
	MicroBlocks []string `json:"micro_blocks"`
}

type KeyBlock struct {
	Beneficiary string  `json:"beneficiary"`
	Hash        string  `json:"hash"`
	Height      int64   `json:"height"`
	Info        string  `json:"info"`
	Miner       string  `json:"miner"`
	Pow         []int64 `json:"pow"`
	PrevHash    string  `json:"prev_hash"`
	PrevKeyHash string  `json:"prev_key_hash"`
	StateHash   string  `json:"state_hash"`
	Target      int64   `json:"target"`
	Time        int64   `json:"time"`
	Version     int64   `json:"version"`
}

type MicroBlock struct {
	Transactions []Transactions `json:"transactions"`
}

type Transactions struct {
	BlockHash   string      `json:"block_hash"`
	BlockHeight int64       `json:"block_height"`
	Hash        string      `json:"hash"`
	Signatures  []string    `json:"signatures"`
	Tx          interface{} `json:"tx"`
}

//func TimeUpdate() {
//	aeaMiddleMicroBlocks, e := models.FindMicroBlockBlockTimeUpdate()
//	if e != nil {
//		fmt.Println("error+" + e.Error())
//		return
//	}
//
//	for i := 0; i < len(aeaMiddleMicroBlocks); i++ {
//		//response := utils.Get("http://node.aechina.io:3013/v2/micro-blocks/hash/" + aeaMiddleMicroBlocks[i].BlockHash + "/header")
//		response := utils.Get("http://node.aechina.io:3013/v2/micro-blocks/hash/" + aeaMiddleMicroBlocks[i].BlockHash + "/header")
//		var block BlocksHeader
//		err := json.Unmarshal([]byte(response), &block)
//		if err != nil {
//			fmt.Println(500, err.Error())
//			return
//		}
//		aeaMiddleMicroBlocks[i].Time = block.Time
//		err = models.InsertAeaMiddleMicroBlockBlockObj(aeaMiddleMicroBlocks[i])
//		if err != nil {
//
//			fmt.Println("error+" + err.Error())
//			return
//		}
//		fmt.Println("sucess->" + aeaMiddleMicroBlocks[i].BlockHash + " #### " + strconv.FormatInt(aeaMiddleMicroBlocks[i].Time, 10))
//	}
//
//}

type BlocksHeader struct {
	Hash        string `json:"hash"`
	Height      int64  `json:"height"`
	PofHash     string `json:"pof_hash"`
	PrevHash    string `json:"prev_hash"`
	PrevKeyHash string `json:"prev_key_hash"`
	Signature   string `json:"signature"`
	StateHash   string `json:"state_hash"`
	Time        int64  `json:"time"`
	TxsHash     string `json:"txs_hash"`
	Version     int64  `json:"version"`
}
