package controllers

import (
	"ae/models"
	"ae/utils"
	"bytes"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"github.com/aeternity/aepp-sdk-go/config"
	"github.com/aeternity/aepp-sdk-go/naet"
	"github.com/aeternity/aepp-sdk-go/transactions"
	"github.com/beego/i18n"
	rlp "github.com/randomshinichi/rlpae"
	"github.com/shopspring/decimal"
	"math/big"
	"strconv"
	"strings"
)

type NamesAuctionsActiveController struct {
	BaseController
}

type NamesPriceController struct {
	BaseController
}

type NamesMyRegisterController struct {
	BaseController
}

type NamesMyOverController struct {
	BaseController
}

type NamesOverController struct {
	BaseController
}

type NamesUpdateController struct {
	BaseController
}

type NamesInfoController struct {
	BaseController
}

type NamesClaimController struct {
	BaseController
}
type PreclaimController struct {
	BaseController
}

type NamesTransferController struct {
	BaseController
}

type NamesBaseController struct {
	BaseController
}

type Names struct {
	Name           string `json:"name"`
	Expiration     int    `json:"expiration"`
	ExpirationTime string `json:"expiration_time"`
	WinningBid     string `json:"winning_bid"`
	WinningBidder  string `json:"winning_bidder"`
}

type NamesMy struct {
	Name             string `json:"name"`
	NameHash         string `json:"name_hash"`
	TxHash           string `json:"tx_hash"`
	CreatedAtHeight  int    `json:"created_at_height"`
	AuctionEndHeight int    `json:"auction_end_height"`
	Owner            string `json:"owner"`
	ExpiresAt        int    `json:"expires_at"`
	ExpirationTime   string `json:"expires_time"`
}

type NamesMyBid struct {
	NameAuctionEntry Names `json:"name_auction_entry"`
}

func (c *NamesAuctionsActiveController) Post() {
	if c.verifyAppId() {
		page, _ := c.GetInt("page", 1)
		height := int(models.ApiBlocksTop())
		namesDb, e := models.FindNameAuctionOver(page, height)
		if e != nil {
			c.ErrorJson(-500, e.Error(), []JsonData{})
		}
		var names []map[string]interface{}
		for i := 0; i < len(namesDb); i++ {
			var name = map[string]interface{}{}
			name["name"] = namesDb[i].Name
			name["length"] = namesDb[i].Length
			name["start_height"] = namesDb[i].StartHeight
			name["end_height"] = namesDb[i].EndHeight
			name["over_height"] = namesDb[i].OverHeight
			name["current_height"] = height
			name["owner"] = namesDb[i].Owner
			name["current_price"] = utils.FormatTokens(namesDb[i].CurrentPrice)
			name["price"] = utils.FormatTokens(namesDb[i].Price)
			name["th_hash"] = namesDb[i].ThHash

			names = append(names, name)
		}
		if names == nil {
			c.SuccessJson([]JsonData{})
			return
		}
		c.SuccessJson(names)
	} else {
		c.ErrorJson(-100, i18n.Tr(c.getHeaderLanguage(), "appId or secret verify error"), []JsonData{})
	}
}

func (c *NamesPriceController) Post() {
	if c.verifyAppId() {
		page, _ := c.GetInt("page", 1)
		height := int(models.ApiBlocksTop())
		namesDb, e := models.FindNameAuctionPrice(page, height)
		if e != nil {
			c.ErrorJson(-500, e.Error(), []JsonData{})
		}
		var names []map[string]interface{}
		for i := 0; i < len(namesDb); i++ {
			var name = map[string]interface{}{}
			name["name"] = namesDb[i].Name
			name["length"] = namesDb[i].Length
			name["start_height"] = namesDb[i].StartHeight
			name["end_height"] = namesDb[i].EndHeight
			name["over_height"] = namesDb[i].OverHeight
			name["current_height"] = height
			name["owner"] = namesDb[i].Owner
			name["current_price"] = utils.FormatTokens(namesDb[i].CurrentPrice)
			name["price"] = utils.FormatTokens(namesDb[i].Price)
			name["th_hash"] = namesDb[i].ThHash

			names = append(names, name)
		}
		if names == nil {
			c.SuccessJson([]JsonData{})
			return
		}
		c.SuccessJson(names)
	} else {
		c.ErrorJson(-100, i18n.Tr(c.getHeaderLanguage(), "appId or secret verify error"), []JsonData{})
	}
}

func (c *NamesOverController) Post() {
	if c.verifyAppId() {
		page, _ := c.GetInt("page", 1)
		height := int(models.ApiBlocksTop())
		namesDb, e := models.FindNameOver(page, height)
		if e != nil {
			c.ErrorJson(-500, e.Error(), []JsonData{})
		}
		var names []map[string]interface{}
		for i := 0; i < len(namesDb); i++ {
			var name = map[string]interface{}{}
			name["name"] = namesDb[i].Name
			name["length"] = namesDb[i].Length
			name["start_height"] = namesDb[i].StartHeight
			name["end_height"] = namesDb[i].EndHeight
			name["over_height"] = namesDb[i].OverHeight
			name["current_height"] = height
			name["owner"] = namesDb[i].Owner
			name["current_price"] = utils.FormatTokens(namesDb[i].CurrentPrice)
			name["price"] = utils.FormatTokens(namesDb[i].Price)
			name["th_hash"] = namesDb[i].ThHash

			names = append(names, name)
		}
		if names == nil {
			c.SuccessJson([]JsonData{})
			return
		}
		c.SuccessJson(names)
	} else {
		c.ErrorJson(-100, i18n.Tr(c.getHeaderLanguage(), "appId or secret verify error"), []JsonData{})
	}
}

func (c *NamesMyRegisterController) Post() {
	if c.verifyAppId() {

		address := c.GetString("address")
		page, _ := c.GetInt("page", 1)
		if address == "" {
			c.ErrorJson(-500, i18n.Tr(c.getHeaderLanguage(), "parameter is nul"), []JsonData{})
			return
		}

		height := int(models.ApiBlocksTop())
		namesDb, e := models.FindNameMyRegister(address, page, height)

		if e != nil {
			c.ErrorJson(-500, e.Error(), []JsonData{})
		}
		var names []map[string]interface{}
		for i := 0; i < len(namesDb); i++ {
			var name = map[string]interface{}{}
			name["name"] = namesDb[i].Name
			name["length"] = namesDb[i].Length
			name["start_height"] = namesDb[i].StartHeight
			name["end_height"] = namesDb[i].EndHeight
			name["over_height"] = namesDb[i].OverHeight
			name["current_height"] = height
			name["owner"] = namesDb[i].Owner
			name["current_price"] = utils.FormatTokens(namesDb[i].CurrentPrice)
			name["price"] = utils.FormatTokens(namesDb[i].Price)
			name["th_hash"] = namesDb[i].ThHash

			names = append(names, name)
		}
		if names == nil {
			c.SuccessJson([]JsonData{})
			return
		}
		c.SuccessJson(names)
	} else {
		c.ErrorJson(-100, i18n.Tr(c.getHeaderLanguage(), "appId or secret verify error"), JsonData{})
	}
}

func (c *NamesMyOverController) Post() {
	if c.verifyAppId() {
		address := c.GetString("address")
		page, _ := c.GetInt("page", 1)
		if address == "" {
			c.ErrorJson(-500, i18n.Tr(c.getHeaderLanguage(), "parameter is nul"), []JsonData{})
			return
		}

		height := int(models.ApiBlocksTop())
		namesDb, e := models.FindNameMyRegisterIng(address, page, height)
		if e != nil {
			c.ErrorJson(-500, e.Error(), []JsonData{})
		}
		var names []map[string]interface{}
		for i := 0; i < len(namesDb); i++ {
			var name = map[string]interface{}{}
			name["name"] = namesDb[i].Name
			name["length"] = namesDb[i].Length
			name["start_height"] = namesDb[i].StartHeight
			name["end_height"] = namesDb[i].EndHeight
			name["over_height"] = namesDb[i].OverHeight
			name["current_height"] = height
			name["owner"] = namesDb[i].Owner
			name["current_price"] = utils.FormatTokens(namesDb[i].CurrentPrice)
			name["price"] = utils.FormatTokens(namesDb[i].Price)
			name["th_hash"] = namesDb[i].ThHash

			names = append(names, name)
		}
		if names == nil {
			c.SuccessJson([]JsonData{})
			return
		}
		c.SuccessJson(names)
	} else {
		c.ErrorJson(-100, i18n.Tr(c.getHeaderLanguage(), "appId or secret verify error"), []JsonData{})
	}
}

func (c *NamesBaseController) Post() {
	if c.verifyAppId() {
		data, e := models.FindNameBase()
		if e != nil {
			c.ErrorJson(-500, e.Error(), JsonData{})
		}
		c.SuccessJson(data)
	} else {
		c.ErrorJson(-100, i18n.Tr(c.getHeaderLanguage(), "appId or secret verify error"), JsonData{})
	}
}

func (c *NamesUpdateController) Post() {
	if c.verifyAppId() {
		name := c.GetString("name")
		address := c.GetString("address")
		if name == "" {
			c.ErrorJson(-500, i18n.Tr(c.getHeaderLanguage(), "parameter is nul"), JsonData{})
			return
		}
		n := strings.Split(name, ".")
		n[0] = strings.Replace(n[0], " ", "", -1)
		if n[0] == "" {
			c.ErrorJson(-500, i18n.Tr(c.getHeaderLanguage(), "name chian error"), JsonData{})
			return
		}

		var balance string
		accountNet, err := models.ApiGetAccount(address)
		if err != nil {
			balance = "0"
		} else {
			balance = accountNet.Balance.String()
		}
		tokens, _ := strconv.ParseFloat(balance, 64)

		if tokens/1000000000000000000 < 1 {
			c.ErrorJson(-500, i18n.Tr(c.getHeaderLanguage(), "The balance is insufficient please keep the number of ae greater than 1"), JsonData{})
			return
		}

		node := naet.NewNode(models.NodeURL, false)
		ttler := transactions.CreateTTLer(node)
		noncer := transactions.CreateNoncer(node)
		ttlNoncer := transactions.CreateTTLNoncer(ttler, noncer)

		pointerAddress, err := transactions.NewNamePointer("account_pubkey", address)
		if err != nil {
			c.ErrorJson(-500, err.Error(), JsonData{})
			return
		}

		updateTx, err := transactions.NewNameUpdateTx(address, name, []*transactions.NamePointer{pointerAddress}, config.Client.Names.ClientTTL+10000, ttlNoncer)
		if err != nil {
			c.ErrorJson(-500, err.Error(), JsonData{})
			return
		}
		updateTx.NameTTL = 50000

		txRaw, _ := rlp.EncodeToBytes(updateTx)
		msg := append([]byte("ae_mainnet"), txRaw...)
		decodeMsg := hex.EncodeToString(msg)

		txJson, _ := json.Marshal(updateTx)
		uEnc := base64.URLEncoding.EncodeToString([]byte(txJson))
		c.SuccessJson(map[string]interface{}{
			"tx":  uEnc,
			"msg": decodeMsg})

	} else {
		c.ErrorJson(-100, i18n.Tr(c.getHeaderLanguage(), "appId or secret verify error"), JsonData{})
	}
}

func (c *NamesInfoController) Post() {
	if c.verifyAppId() {

		name := c.GetString("name")
		if name == "" {
			c.ErrorJson(-500, i18n.Tr(c.getHeaderLanguage(), "parameter is nul"), JsonData{})
			return
		}

		nameDb, err := models.FindNameName(name)
		if err != nil {
			c.ErrorJson(-500, err.Error(), JsonData{})
			return
		}

		if nameDb.CurrentPrice == 0 {
			c.ErrorJson(201, "name is no register", JsonData{})
			return
		}
		height := int(models.ApiBlocksTop())
		var nameMap = map[string]interface{}{}
		nameMap["name"] = nameDb.Name
		nameMap["length"] = nameDb.Length
		nameMap["start_height"] = nameDb.StartHeight
		nameMap["end_height"] = nameDb.EndHeight
		nameMap["over_height"] = nameDb.OverHeight
		nameMap["current_height"] = height
		nameMap["owner"] = nameDb.Owner

		nameMap["current_price"] = utils.FormatTokens(nameDb.CurrentPrice)
		nameMap["price"] = utils.FormatTokens(nameDb.Price)
		nameMap["th_hash"] = nameDb.ThHash

		blocksDb, err := models.FindMicroBlockBlockNameorDatas(name)
		if err != nil {
			c.ErrorJson(-500, err.Error(), JsonData{})
			return
		}
		var txs []map[string]interface{}
		for i := 0; i < len(blocksDb); i++ {
			mapObj := make(map[string]interface{})

			// body是后端的http返回结果
			d := json.NewDecoder(bytes.NewReader([]byte(blocksDb[i].Tx)))
			d.UseNumber()
			err = d.Decode(&mapObj)

			f, _ := mapObj["name_fee"].(json.Number).Float64()
			mapObj["name_fee"] = utils.FormatTokens(f)
			mapObj["time"] = blocksDb[i].Time
			mapObj["block_height"] = blocksDb[i].BlockHeight
			mapObj["block_height"] = blocksDb[i].BlockHeight
			mapObj["hash"] = blocksDb[i].Hash
			txs = append(txs, mapObj)
		}

		nameMap["claim"] = txs
		c.SuccessJson(nameMap)
	} else {
		c.ErrorJson(-100, i18n.Tr(c.getHeaderLanguage(), "appId or secret verify error"), JsonData{})
	}
}

func (c *NamesTransferController) Post() {
	if c.verifyAppId() {
		name := c.GetString("name")
		senderID := c.GetString("senderID")
		recipientID := c.GetString("recipientID")
		if name == "" || senderID == "" || recipientID == "" {
			c.ErrorJson(-500, i18n.Tr(c.getHeaderLanguage(), "parameter is nul"), JsonData{})
			return
		}
		n := strings.Split(name, ".")
		n[0] = strings.Replace(n[0], " ", "", -1)
		if n[0] == "" {
			c.ErrorJson(-500, i18n.Tr(c.getHeaderLanguage(), "name chian error"), JsonData{})
			return
		}

		var balance string
		accountNet, err := models.ApiGetAccount(senderID)
		if err != nil {
			balance = "0"
		} else {
			balance = accountNet.Balance.String()
		}
		tokens, _ := strconv.ParseFloat(balance, 64)
		if tokens/1000000000000000000 < 1 {
			c.ErrorJson(-500, i18n.Tr(c.getHeaderLanguage(), "The balance is insufficient please keep the number of ae greater than 1"), JsonData{})
			return
		}

		node := naet.NewNode(models.NodeURL, false)
		ttler := transactions.CreateTTLer(node)
		noncer := transactions.CreateNoncer(node)
		ttlNoncer := transactions.CreateTTLNoncer(ttler, noncer)
		transferTx, _ := transactions.NewNameTransferTx(senderID, name, recipientID, ttlNoncer)
		txRaw, _ := rlp.EncodeToBytes(transferTx)
		msg := append([]byte("ae_mainnet"), txRaw...)
		decodeMsg := hex.EncodeToString(msg)

		txJsons, _ := json.Marshal(transferTx)
		uEnc := base64.URLEncoding.EncodeToString([]byte(txJsons))

		c.SuccessJson(map[string]interface{}{
			"tx":  uEnc,
			"msg": decodeMsg})

	} else {
		c.ErrorJson(-100, i18n.Tr(c.getHeaderLanguage(), "appId or secret verify error"), JsonData{})
	}
}

func (c *NamesClaimController) Post() {
	if c.verifyAppId() {

		name := c.GetString("name")
		address := c.GetString("address")
		//nameSalt ,_:= c.GetInt64("nameSalt",0)
		println("nameSalt->", nameSalt)
		if name == "" || address == "" {
			c.ErrorJson(-500, i18n.Tr(c.getHeaderLanguage(), "parameter is nul"), JsonData{})
			return
		}
		n := strings.Split(name, ".")
		n[0] = strings.Replace(n[0], " ", "", -1)
		if n[0] == "" {
			c.ErrorJson(-500, i18n.Tr(c.getHeaderLanguage(), "name chian error"), JsonData{})
			return
		}

		if !strings.Contains(name, ".chain") {
			c.ErrorJson(-500, i18n.Tr(c.getHeaderLanguage(), "Please keep the chain end"), JsonData{})
			return
		}

		nameDb, err := models.FindNameName(name)
		if err != nil {
			c.ErrorJson(-500, err.Error(), JsonData{})
			return
		}
		aeHeight, _ := strconv.Atoi(strconv.FormatUint(models.ApiBlocksTop(), 10))
		if nameDb.CurrentPrice > 0 && nameDb.EndHeight < aeHeight {
			c.ErrorJson(-500, i18n.Tr(c.getHeaderLanguage(), "names Has been registered"), JsonData{})
			return
		}

		var nameFee *big.Int
		if nameDb.CurrentPrice == 0 {
			palce := utils.GetAENSPalce(name)
			decimalValue, _ := decimal.NewFromString(palce)
			nameFee = decimalValue.BigInt()
		} else {
			nameFee = decimal.NewFromFloat(nameDb.CurrentPrice).BigInt()
		}

		decimalValue := decimal.NewFromBigInt(nameFee, 0)
		decimalValueMul := decimalValue.Mul(decimal.NewFromFloat(0.05))
		decimalValue = decimalValue.Add(decimalValueMul)

		accountNet, e := models.ApiGetAccount(address)
		var tokens float64
		if e != nil {
			tokens = 0
			return
		} else {
			tokens, err = strconv.ParseFloat(accountNet.Balance.String(), 64)
			if err != nil {
				c.ErrorJson(-500, err.Error(), JsonData{})
				return
			}
		}

		f, _ := decimalValue.Float64()

		if utils.FormatTokensInt(tokens) <= utils.FormatTokensInt(f) {
			c.ErrorJson(-500, i18n.Tr(c.getHeaderLanguage(), "Lack of balance greater than")+utils.FormatTokens(f)+"AE", JsonData{})
			return
		}
		var isUpdate bool
		if nameDb.CurrentPrice > 0 && nameDb.EndHeight > aeHeight {
			isUpdate = true
		}
		node := naet.NewNode(models.NodeURL, false)
		ttler := transactions.CreateTTLer(node)
		noncer := transactions.CreateNoncer(node)
		ttlNoncer := transactions.CreateTTLNoncer(ttler, noncer)

		var salt *big.Int
		if v, ok := nameSalt[name+"#"+address]; ok {
			salt = v
		}
		claimTx, _ := transactions.NewNameClaimTx(address, name, salt, decimalValue.BigInt(), ttlNoncer)
		if nameDb.CurrentPrice != 0 {
			if isUpdate {
				claimTx.NameSalt = new(big.Int)
				claimTx.TTL = 0
			}
		}

		txRaw, _ := rlp.EncodeToBytes(claimTx)
		msg := append([]byte("ae_mainnet"), txRaw...)
		decodeMsg := hex.EncodeToString(msg)

		txJson, _ := json.Marshal(claimTx)
		uEnc := base64.URLEncoding.EncodeToString([]byte(txJson))

		c.SuccessJson(map[string]interface{}{
			"tx":  uEnc,
			"msg": decodeMsg})

	} else {
		c.ErrorJson(-100, i18n.Tr(c.getHeaderLanguage(), "appId or secret verify error"), JsonData{})
	}
}

var nameSalt = make(map[string]*big.Int)

func (c *PreclaimController) Post() {
	if c.verifyAppId() {

		name := c.GetString("name")
		address := c.GetString("address")
		if name == "" || address == "" {
			c.ErrorJson(-500, i18n.Tr(c.getHeaderLanguage(), "parameter is nul"), JsonData{})
			return
		}
		n := strings.Split(name, ".")
		n[0] = strings.Replace(n[0], " ", "", -1)
		if n[0] == "" {
			c.ErrorJson(-500, i18n.Tr(c.getHeaderLanguage(), "name chian error"), JsonData{})
			return
		}

		if !strings.Contains(name, ".chain") {
			c.ErrorJson(-500, i18n.Tr(c.getHeaderLanguage(), "Please keep the chain end"), JsonData{})
			return
		}

		nameDb, err := models.FindNameName(name)
		if err != nil {
			c.ErrorJson(-500, err.Error(), JsonData{})
			return
		}
		aeHeight, _ := strconv.Atoi(strconv.FormatUint(models.ApiBlocksTop(), 10))
		if nameDb.CurrentPrice > 0 && nameDb.EndHeight < aeHeight {
			c.ErrorJson(-500, i18n.Tr(c.getHeaderLanguage(), "names Has been registered"), JsonData{})
			return
		}

		var nameFee *big.Int
		if nameDb.CurrentPrice == 0 {
			palce := utils.GetAENSPalce(name)
			decimalValue, _ := decimal.NewFromString(palce)
			nameFee = decimalValue.BigInt()
		} else {
			nameFee = decimal.NewFromFloat(nameDb.CurrentPrice).BigInt()
		}

		decimalValue := decimal.NewFromBigInt(nameFee, 0)
		decimalValueMul := decimalValue.Mul(decimal.NewFromFloat(0.05))
		decimalValue = decimalValue.Add(decimalValueMul)

		accountNet, e := models.ApiGetAccount(address)
		var tokens float64
		if e != nil {
			tokens = 0
		} else {
			tokens, err = strconv.ParseFloat(accountNet.Balance.String(), 64)
			if err != nil {
				c.ErrorJson(-500, err.Error(), JsonData{})
				return
			}
		}

		f, _ := decimalValue.Float64()

		if utils.FormatTokensInt(tokens) <= utils.FormatTokensInt(f) {
			c.ErrorJson(-500, i18n.Tr(c.getHeaderLanguage(), "Lack of balance greater than")+utils.FormatTokens(f)+"AE", JsonData{})
			return
		}

		node := naet.NewNode(models.NodeURL, false)
		ttler := transactions.CreateTTLer(node)
		noncer := transactions.CreateNoncer(node)
		ttlNoncer := transactions.CreateTTLNoncer(ttler, noncer)

		preclaimTx, salt, err := transactions.NewNamePreclaimTx(address, name, ttlNoncer)
		nameSalt[name+"#"+address] = salt
		txRaw, _ := rlp.EncodeToBytes(preclaimTx)
		msg := append([]byte("ae_mainnet"), txRaw...)
		decodeMsg := hex.EncodeToString(msg)

		txJson, _ := json.Marshal(preclaimTx)
		uEnc := base64.URLEncoding.EncodeToString([]byte(txJson))

		c.SuccessJson(map[string]interface{}{
			"tx":   uEnc,
			"salt": salt,
			"msg":  decodeMsg})

	} else {
		c.ErrorJson(-100, i18n.Tr(c.getHeaderLanguage(), "appId or secret verify error"), JsonData{})
	}
}
