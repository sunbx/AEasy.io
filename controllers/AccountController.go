package controllers

import (
	"ae/models"
	"ae/utils"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"github.com/aeternity/aepp-sdk-go/binary"
	"github.com/aeternity/aepp-sdk-go/naet"
	"github.com/aeternity/aepp-sdk-go/transactions"
	"github.com/beego/i18n"
	"github.com/shopspring/decimal"
	_ "github.com/typa01/go-utils"
)

type AccountInfoController struct {
	BaseController
}
type BroadcastTxController struct {
	BaseController
}

//用户信息
func (c *AccountInfoController) Post() {
	if c.verifyAppId() {
		address := c.GetString("address")
		if address == "" {
			c.ErrorJson(-301, i18n.Tr(c.getHeaderLanguage(), "parameter is nul"), JsonData{})
			return
		}
		accountNet, e := models.ApiGetAccount(address)
		if e != nil {
			if e.Error() == "Error: Account not found" {
				c.SuccessJson(map[string]string{
					"balance": "0.00000",
					"address": address,
				})
			} else {
				c.ErrorJson(-500, e.Error(), JsonData{})
			}
			return
		}
		decimalValue, _ := decimal.NewFromString(accountNet.Balance.String())
		f, _ := decimalValue.Float64()
		c.SuccessJson(map[string]string{
			"balance": utils.FormatTokens(f),
			"address": address,
		})
	} else {
		c.ErrorJson(-100, i18n.Tr(c.getHeaderLanguage(), "appId or secret verify error"), JsonData{})
	}
}

//广播
func (c *BroadcastTxController) Post() {
	if c.verifyAppId() {
		signature := c.GetString("signature")
		tx := c.GetString("tx")
		t := c.GetString("type")

		uEnc, _ := base64.URLEncoding.DecodeString(tx)



		if t == "SpendTx"{
			var txObj transactions.SpendTx
			_ = json.Unmarshal(uEnc, &txObj)
			serializeTx, _ := transactions.SerializeTx(&txObj)
			tx = serializeTx
		}
		if t == "ContractCreateTx"{
			var txObj transactions.ContractCreateTx
			_ = json.Unmarshal(uEnc, &txObj)
			serializeTx, _ := transactions.SerializeTx(&txObj)
			tx = serializeTx
		}

		if t == "ContractCallTx"{
			var txObj transactions.ContractCallTx
			_ = json.Unmarshal(uEnc, &txObj)
			serializeTx, _ := transactions.SerializeTx(&txObj)
			tx = serializeTx
			println(string(signature))
			println(string(uEnc))
		}
		if t == "NameClaimTx"{
			var txObj transactions.NameClaimTx
			_ = json.Unmarshal(uEnc, &txObj)
			serializeTx, _ := transactions.SerializeTx(&txObj)
			tx = serializeTx
		}
		if t == "NamePreclaimTx"{
			var txObj transactions.NamePreclaimTx
			_ = json.Unmarshal(uEnc, &txObj)
			serializeTx, _ := transactions.SerializeTx(&txObj)
			tx = serializeTx
		}
		if t == "NamePointer"{
			var txObj transactions.NamePointer
			_ = json.Unmarshal(uEnc, &txObj)
			serializeTx, _ := transactions.SerializeTx(&txObj)
			tx = serializeTx
		}
		if t == "NameRevokeTx"{
			var txObj transactions.NameRevokeTx
			_ = json.Unmarshal(uEnc, &txObj)
			serializeTx, _ := transactions.SerializeTx(&txObj)
			tx = serializeTx
		}
		if t == "NameTransferTx"{
			var txObj transactions.NameTransferTx
			_ = json.Unmarshal(uEnc, &txObj)
			serializeTx, _ := transactions.SerializeTx(&txObj)
			tx = serializeTx
		}
		if t == "NameUpdateTx"{
			var txObj transactions.NameUpdateTx
			_ = json.Unmarshal(uEnc, &txObj)
			serializeTx, _ := transactions.SerializeTx(&txObj)
			tx = serializeTx
		}




		//获取节点信息
		node := naet.NewNode(models.NodeURL, true)
		signatureByte, err := hex.DecodeString(signature)
		if err != nil {
			c.ErrorJson(-101, err.Error(), JsonData{})
			return
		}
		deSerializeTx, err := transactions.DeserializeTxStr(tx)
		if err != nil {
			c.ErrorJson(-102, err.Error(), JsonData{})
			return
		}
		var signedTx = transactions.NewSignedTx([][]byte{}, deSerializeTx)
		signedTx.Signatures = append(signedTx.Signatures, signatureByte)
		var _ = binary.Encode(binary.PrefixSignature, signatureByte)
		txHash, err := transactions.Hash(signedTx)
		signedTxStr, errSerializeTx := transactions.SerializeTx(signedTx)
		if errSerializeTx != nil {
			c.ErrorJson(-100, errSerializeTx.Error(), JsonData{})
			return
		}

		println(signedTxStr)
		println(txHash)
		err = node.PostTransaction(signedTxStr, txHash)
		if err != nil {
			c.ErrorJson(-103, err.Error(), JsonData{})
			return
		}
		c.SuccessJson(map[string]string{
			"hash": txHash,
		})
	} else {
		c.ErrorJson(-100, i18n.Tr(c.getHeaderLanguage(), "appId or secret verify error"), JsonData{})
	}
}
