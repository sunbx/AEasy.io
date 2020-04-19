package controllers

import (
	"ae/models"
	"ae/utils"
	"sort"
	"time"
)

type NamesAuctionsActiveController struct {
	BaseController
}

type NamesOverdueController struct {
	BaseController
}

type NamesMyRegisterController struct {
	BaseController
}
type NamesMyActivityController struct {
	BaseController
}
type NamesNewController struct {
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

func (c *NamesAuctionsActiveController) Get() {
	page, _ := c.GetInt("page", 0)


	height := models.ApiBlocksTop()
	names, err := models.GetNamesOver(int(height), page)
	if err != nil {
		c.ErrorJson(-200, err.Error(), JsonData{})
		return
	}

	var namesJsons []models.NamesJson
	for i := 0; i < len(names); i++ {
		var namesJson models.NamesJson
		namesJson.Id = names[i].Id
		namesJson.AuctionEndHeight = names[i].AuctionEndHeight
		namesJson.CreatedAtHeight = names[i].CreatedAtHeight
		namesJson.ExpiresAt = names[i].ExpiresAt
		namesJson.Name = names[i].Name
		namesJson.NameHash = names[i].NameHash
		namesJson.Owner = names[i].Owner
		namesJson.TxHash = names[i].TxHash

		now := time.Now().Unix()
		namesJson.ExpirationTime = utils.StrTime(now - int64(60*3*(int(names[i].AuctionEndHeight)-int(height))))
		//tokens, _ := strconv.ParseFloat(names[i].WinningBid, 64)
		//content := utils.FormatTokens(tokens, 5)
		//names[i].WinningBid = content + " AE"
		//
		namesJsons = append(namesJsons, namesJson)
		//slice := strings.Split(names[i].WinningBidder, "")
		//str := strings.Join(slice[0:3], "") + "****" + strings.Join(slice[len(names[i].WinningBidder)-4:], "")
		//names[i].WinningBidder = str
	}
	c.SuccessJson(namesJsons)

	//response := utils.Get("https://mainnet.aeternal.io/middleware/names/auctions/active?length=NaN&limit=" + limit + "&page=" + page + "&sort=expiration")
	//var names []Names
	//err := json.Unmarshal([]byte(response), &names)
	//if err != nil {
	//	c.ErrorJson(500, err.Error(), JsonData{})
	//	return
	//}
	//height := models.ApiBlocksTop()
	//for i := 0; i < len(names); i++ {
	//	//names[i].Height = utils.StrTime(int64((names[i].Expiration-int(height))*3*60) * 1000)
	//	now := time.Now().Unix()
	//	names[i].ExpirationTime = utils.StrTime(now - int64(60*3*(names[i].Expiration-int(height))))
	//	tokens, _ := strconv.ParseFloat(names[i].WinningBid, 64)
	//	content := utils.FormatTokensP(tokens, 5)
	//	names[i].WinningBid = content + " AE"
	//
	//	slice := strings.Split(names[i].WinningBidder, "")
	//	str := strings.Join(slice[0:3], "") + "****" + strings.Join(slice[len(names[i].WinningBidder)-4:], "")
	//	names[i].WinningBidder = str
	//}
	//c.SuccessJson(names)
}

func (c *NamesNewController) Get() {

	page, _ := c.GetInt("page", 0)
	height := models.ApiBlocksTop()
	names, err := models.GetNamesAllTime(page)
	if err != nil {
		c.ErrorJson(-200, err.Error(), JsonData{})
		return
	}
	var namesJsons []models.NamesJson
	for i := 0; i < len(names); i++ {
		var namesJson models.NamesJson
		namesJson.Id = names[i].Id
		namesJson.AuctionEndHeight = names[i].AuctionEndHeight
		namesJson.CreatedAtHeight = names[i].CreatedAtHeight
		namesJson.ExpiresAt = names[i].ExpiresAt
		namesJson.Name = names[i].Name
		namesJson.NameHash = names[i].NameHash
		namesJson.Owner = names[i].Owner
		namesJson.TxHash = names[i].TxHash

		now := time.Now().Unix()
		namesJson.ExpirationTime = utils.StrTime(now - int64(60*3*(int(names[i].AuctionEndHeight)-int(height))))
		//tokens, _ := strconv.ParseFloat(names[i].WinningBid, 64)
		//content := utils.FormatTokens(tokens, 5)
		//names[i].WinningBid = content + " AE"
		//
		namesJsons = append(namesJsons, namesJson)
		//slice := strings.Split(names[i].WinningBidder, "")
		//str := strings.Join(slice[0:3], "") + "****" + strings.Join(slice[len(names[i].WinningBidder)-4:], "")
		//names[i].WinningBidder = str
	}
	c.SuccessJson(namesJsons)
}

func (c *NamesOverdueController) Get() {
	page, _ := c.GetInt("page", 0)
	height := models.ApiBlocksTop()
	names, err := models.GetNamesAllHeight(page, int(height+1))
	if err != nil {
		c.ErrorJson(-200, err.Error(), JsonData{})
		return
	}
	var namesJsons []models.NamesJson
	for i := 0; i < len(names); i++ {
		var namesJson models.NamesJson
		namesJson.Id = names[i].Id
		namesJson.AuctionEndHeight = names[i].AuctionEndHeight
		namesJson.CreatedAtHeight = names[i].CreatedAtHeight
		namesJson.ExpiresAt = names[i].ExpiresAt
		namesJson.Name = names[i].Name
		namesJson.NameHash = names[i].NameHash
		namesJson.Owner = names[i].Owner
		namesJson.TxHash = names[i].TxHash

		now := time.Now().Unix()
		namesJson.ExpirationTime = utils.StrTime(now - int64(60*3*(int(names[i].ExpiresAt)-int(height))))
		//tokens, _ := strconv.ParseFloat(names[i].WinningBid, 64)
		//content := utils.FormatTokens(tokens, 5)
		//names[i].WinningBid = content + " AE"
		//
		namesJsons = append(namesJsons, namesJson)
		//slice := strings.Split(names[i].WinningBidder, "")
		//str := strings.Join(slice[0:3], "") + "****" + strings.Join(slice[len(names[i].WinningBidder)-4:], "")
		//names[i].WinningBidder = str
	}
	c.SuccessJson(namesJsons)
}

//func (c *NamesMyController) Get() {
//	address := c.GetString("address")
//	if address == "" {
//		c.ErrorJson(-301, "parameter is nul", JsonData{})
//		return
//	}
//
//	response := utils.Get("https://mainnet.aeternal.io/middleware/names/active?owner=" + address)
//	var names []NamesMy
//	err := json.Unmarshal([]byte(response), &names)
//	if err != nil {
//		c.ErrorJson(500, err.Error(), JsonData{})
//		return
//	}
//	height := models.ApiBlocksTop()
//	for i := 0; i < len(names); i++ {
//		//names[i].Height = utils.StrTime(int64((names[i].Expiration-int(height))*3*60) * 1000)
//		now := time.Now().Unix()
//		names[i].ExpirationTime = utils.StrTime(now - int64(60*3*(names[i].ExpiresAt-int(height))))
//
//		slice := strings.Split(names[i].Owner, "")
//		str := strings.Join(slice[0:3], "") + "****" + strings.Join(slice[len(names[i].Owner)-4:], "")
//		names[i].Owner = str
//	}
//
//	//翻转
//	length := len(names)
//	for i := 0; i < length/2; i++ {
//		temp := names[length-1-i]
//		names[length-1-i] = names[i]
//		names[i] = temp
//	}
//
//	response = utils.Get("https://mainnet.aeternal.io/middleware/names/auctions/bids/account/" + address)
//	var namesBid []NamesMyBid
//	err = json.Unmarshal([]byte(response), &namesBid)
//	if err != nil {
//		c.ErrorJson(500, err.Error(), JsonData{})
//		return
//	}
//	var nameBidMap = make(map[string]Names)
//	for i := 0; i < len(namesBid); i++ {
//		now := time.Now().Unix()
//		namesBid[i].NameAuctionEntry.ExpirationTime = utils.StrTime(now - int64(60*3*(namesBid[i].NameAuctionEntry.Expiration-int(height))))
//		tokens, _ := strconv.ParseFloat(namesBid[i].NameAuctionEntry.WinningBid, 64)
//		content := utils.FormatTokens(tokens, 5)
//		namesBid[i].NameAuctionEntry.WinningBid = content + " AE"
//
//		slice := strings.Split(namesBid[i].NameAuctionEntry.WinningBidder, "")
//		str := strings.Join(slice[0:3], "") + "****" + strings.Join(slice[len(namesBid[i].NameAuctionEntry.WinningBidder)-4:], "")
//		namesBid[i].NameAuctionEntry.WinningBidder = str
//
//		if namesBid[i].NameAuctionEntry.Expiration > int(height) {
//			nameBidMap[namesBid[i].NameAuctionEntry.Name] = namesBid[i].NameAuctionEntry
//		}
//	}
//
//	var nameAuctionEntry = []Names{}
//	// 遍历map
//	for _, v := range nameBidMap {
//		nameAuctionEntry = append(nameAuctionEntry, v)
//	}
//
//	c.SuccessJson(map[string]interface{}{
//		"register": names,
//		"auction":  nameAuctionEntry,
//	})
//}

//func (c *NamesMyController) Get() {
//	address := c.GetString("address")
//	if address == "" {
//		c.ErrorJson(-301, "parameter is nul", JsonData{})
//		return
//	}
//
//	response := utils.Get("https://mainnet.aeternal.io/middleware/names/active?owner=" + address)
//	var names []NamesMy
//	err := json.Unmarshal([]byte(response), &names)
//	if err != nil {
//		c.ErrorJson(500, err.Error(), JsonData{})
//		return
//	}
//	height := models.ApiBlocksTop()
//	for i := 0; i < len(names); i++ {
//		//names[i].Height = utils.StrTime(int64((names[i].Expiration-int(height))*3*60) * 1000)
//		now := time.Now().Unix()
//		names[i].ExpirationTime = utils.StrTime(now - int64(60*3*(names[i].ExpiresAt-int(height))))
//
//		slice := strings.Split(names[i].Owner, "")
//		str := strings.Join(slice[0:3], "") + "****" + strings.Join(slice[len(names[i].Owner)-4:], "")
//		names[i].Owner = str
//	}
//
//	//翻转
//	length := len(names)
//	for i := 0; i < length/2; i++ {
//		temp := names[length-1-i]
//		names[length-1-i] = names[i]
//		names[i] = temp
//	}
//
//	owner, err := models.GetNamesAllOwner(address)
//
//	var nameBidMap = make(map[string]NamesMy)
//	for i := 0; i < len(owner); i++ {
//		now := time.Now().Unix()
//		slice := strings.Split(owner[i].Owner, "")
//		str := strings.Join(slice[0:3], "") + "****" + strings.Join(slice[len(owner[i].Owner)-4:], "")
//		nameMy := NamesMy{
//			Name:             owner[i].Name,
//			NameHash:         owner[i].NameHash,
//			TxHash:           owner[i].TxHash,
//			CreatedAtHeight:  int(owner[i].CreatedAtHeight),
//			AuctionEndHeight: int(owner[i].AuctionEndHeight),
//			Owner:            str,
//			ExpiresAt:        int(owner[i].ExpiresAt),
//			ExpirationTime:   utils.StrTime(now - int64(60*3*(int(owner[i].AuctionEndHeight)-int(height)))),
//		}
//		nameBidMap[owner[i].Name] = nameMy
//	}
//
//	for i := 0; i < len(names); i++ {
//		_, ok := nameBidMap[names[i].Name]
//		if ok {
//			delete(nameBidMap, names[i].Name)
//		}
//	}
//
//	var nameAuctionEntry = []NamesMy{}
//	// 遍历map
//	for _, v := range nameBidMap {
//		nameAuctionEntry = append(nameAuctionEntry, v)
//	}
//
//	c.SuccessJson(map[string]interface{}{
//		"register": names,
//		"auction":  nameAuctionEntry,
//	})
//
//	//response = utils.Get("https://mainnet.aeternal.io/middleware/names/auctions/bids/account/" + address)
//	//var namesBid []NamesMyBid
//	//err = json.Unmarshal([]byte(response), &namesBid)
//	//if err != nil {
//	//	c.ErrorJson(500, err.Error(), JsonData{})
//	//	return
//	//}
//	//var nameBidMap = make(map[string]Names)
//	//for i := 0; i < len(namesBid); i++ {
//	//	now := time.Now().Unix()
//	//	namesBid[i].NameAuctionEntry.ExpirationTime = utils.StrTime(now - int64(60*3*(namesBid[i].NameAuctionEntry.Expiration-int(height))))
//	//	tokens, _ := strconv.ParseFloat(namesBid[i].NameAuctionEntry.WinningBid, 64)
//	//	content := utils.FormatTokens(tokens, 5)
//	//	namesBid[i].NameAuctionEntry.WinningBid = content + " AE"
//	//
//	//	slice := strings.Split(namesBid[i].NameAuctionEntry.WinningBidder, "")
//	//	str := strings.Join(slice[0:3], "") + "****" + strings.Join(slice[len(namesBid[i].NameAuctionEntry.WinningBidder)-4:], "")
//	//	namesBid[i].NameAuctionEntry.WinningBidder = str
//	//
//	//	if namesBid[i].NameAuctionEntry.Expiration > int(height) {
//	//		nameBidMap[namesBid[i].NameAuctionEntry.Name] = namesBid[i].NameAuctionEntry
//	//	}
//	//}
//	//
//	//var nameAuctionEntry = []Names{}
//	//// 遍历map
//	//for _, v := range nameBidMap {
//	//	nameAuctionEntry = append(nameAuctionEntry, v)
//	//}
//	//
//	//c.SuccessJson(map[string]interface{}{
//	//	"register": names,
//	//	"auction":  nameAuctionEntry,
//	//})
//}

func (c *NamesMyRegisterController) Get() {
	address := c.GetString("address")
	if address == "" {
		c.ErrorJson(-301, "parameter is nul", JsonData{})
		return
	}

	height := models.ApiBlocksTop()
	names, _ := models.GetNamesAllMyRegister(int(height), address)

	var namesJsons []models.NamesJson
	for i := 0; i < len(names); i++ {

		if names[i].ExpiresAt-names[i].AuctionEndHeight < 48000 {
			names[i].ExpiresAt = names[i].AuctionEndHeight + 48000
		}

		var namesJson models.NamesJson
		namesJson.Id = names[i].Id
		namesJson.AuctionEndHeight = names[i].AuctionEndHeight
		namesJson.CreatedAtHeight = names[i].CreatedAtHeight
		namesJson.ExpiresAt = names[i].ExpiresAt
		namesJson.Name = names[i].Name
		namesJson.NameHash = names[i].NameHash
		namesJson.Owner = names[i].Owner
		namesJson.TxHash = names[i].TxHash

		now := time.Now().Unix()
		namesJson.ExpirationTime = utils.StrTime(now - int64(60*3*(int(names[i].ExpiresAt)-int(height))))
		namesJsons = append(namesJsons, namesJson)
	}
	sort.Sort(models.NamesJsonRegister(namesJsons))
	c.SuccessJson(namesJsons)
}
func (c *NamesMyActivityController) Get() {
	address := c.GetString("address")
	if address == "" {
		c.ErrorJson(-301, "parameter is nul", JsonData{})
		return
	}

	height := models.ApiBlocksTop()
	names, _ := models.GetNamesAllMyActivity(int(height), address)

	var namesJsons []models.NamesJson
	for i := 0; i < len(names); i++ {

		var namesJson models.NamesJson
		namesJson.Id = names[i].Id
		namesJson.AuctionEndHeight = names[i].AuctionEndHeight
		namesJson.CreatedAtHeight = names[i].CreatedAtHeight
		namesJson.ExpiresAt = names[i].ExpiresAt
		namesJson.Name = names[i].Name
		namesJson.NameHash = names[i].NameHash
		namesJson.Owner = names[i].Owner
		namesJson.TxHash = names[i].TxHash

		now := time.Now().Unix()
		namesJson.ExpirationTime = utils.StrTime(now - int64(60*3*(int(names[i].AuctionEndHeight)-int(height))))
		namesJsons = append(namesJsons, namesJson)
	}
	sort.Sort(models.NamesJsonActivity(namesJsons))
	c.SuccessJson(namesJsons)
}
