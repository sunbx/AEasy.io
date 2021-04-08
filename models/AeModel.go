package models

import (
	"ae/utils"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/aeternity/aepp-sdk-go/account"
	"github.com/aeternity/aepp-sdk-go/aeternity"
	"github.com/aeternity/aepp-sdk-go/config"
	"github.com/aeternity/aepp-sdk-go/naet"
	"github.com/aeternity/aepp-sdk-go/swagguard/node/models"
	"github.com/aeternity/aepp-sdk-go/transactions"
	"github.com/shopspring/decimal"
	"github.com/tyler-smith/go-bip39"
	"io/ioutil"
	"math/big"
	"net/http"
	"strconv"
)

//var NodeURL = "https://mainnet.aeternal.io"
//var NodeURL = "http://localhost:3013"
//var NodeURL = "http://www.aestore.co:3013"
//var NodeURL = "http://47.108.93.212:3013"
var NodeURL = "https://node.aechina.io"
//var compilerURL = "http://localhost:3080"s
var compilerURL = "https://compiler.aeasy.io"

//===================================================================================================================================================================================================
//|                           															AE-BASE																										 |
///===================================================================================================================================================================================================

//根据助记词返回用户
func MnemonicAccount(mnemonic string, addressIndex uint32) (*account.Account, error) {

	//生成种子
	seed, err := account.ParseMnemonic(mnemonic)
	if err != nil {
		return nil, err
	}

	//验证助记词
	_, err = bip39.EntropyFromMnemonic(mnemonic)

	if err != nil {
		return nil, err
	}

	//获取子账户
	// Derive the subaccount m/44'/457'/3'/0'/1'
	key, err := account.DerivePathFromSeed(seed, 0, addressIndex-1)
	if err != nil {
		return nil, err
	}

	// 生成账户
	alice, err := account.BIP32KeyToAeKey(key)
	if err != nil {
		return nil, err
	}
	return alice, nil
}

//根据私钥返回用户
func SigningKeyHexStringAccount(signingKey string) (*account.Account, error) {
	acc, e := account.FromHexString(signingKey)
	return acc, e
}

//随机创建用户
func CreateAccount() (*account.Account, string) {
	mnemonic, signingKey, _ := CreateAccountUtils()
	acc, _ := account.FromHexString(signingKey)
	return acc, mnemonic
}

//随机创建用户,返回助记词
func CreateAccountUtils() (mnemonic string, signingKey string, address string) {
	//创建助记词
	entropy, _ := bip39.NewEntropy(128)
	//生成助记词
	mne, _ := bip39.NewMnemonic(entropy)
	//生成种子
	seed, _ := account.ParseMnemonic(mne)
	//验证助记词
	_, _ = bip39.EntropyFromMnemonic(mne)
	//生成子账户
	key, _ := account.DerivePathFromSeed(seed, 0, 0)
	//获取账户
	alice, _ := account.BIP32KeyToAeKey(key)
	//返回私钥和信息
	return mne, alice.SigningKeyToHexString(), alice.Address
}

//返回最新区块高度
func ApiBlocksTop() (height uint64) {
	client := naet.NewNode(NodeURL, false)
	h, _ := client.GetHeight()
	return h
}

//地址信息返回用户信息和余额
func ApiGetAccount(address string) (account *models.Account, e error) {
	client := naet.NewNode(NodeURL, false)
	acc, e := client.GetAccount(address)
	return acc, e
}

//发起转账
func ApiSpend(account *account.Account, recipientId string, amount float64, data string) (*aeternity.TxReceipt, error) {
	//获取账户
	accountNet, e := ApiGetAccount(account.Address)
	if e != nil {
		return nil, e
	}
	//格式化账户的tokens
	tokens, err := strconv.ParseFloat(accountNet.Balance.String(), 64)
	if err == nil {

		//判断账户余额是否大于要转账的余额
		if tokens/1000000000000000000 >= amount {
			//获取节点信息
			node := naet.NewNode(NodeURL, false)
			//生成ttl
			ttler := transactions.CreateTTLer(node)
			noncer := transactions.CreateNoncer(node)

			ttlNoncer := transactions.CreateTTLNoncer(ttler, noncer)
			//生成转账tx
			spendTx, err := transactions.NewSpendTx(account.Address, recipientId, utils.GetRealAebalanceBigInt(amount), []byte(data), ttlNoncer)
			spendTxJson, _ := json.Marshal(spendTx)
			println("spendTx->",string(spendTxJson))
			feeTokens, _ := strconv.ParseFloat(spendTx.Fee.String(), 64)

			if (feeTokens/1000000000000000000+amount) >= tokens/1000000000000000000 {
				decimalValue := decimal.NewFromFloat(feeTokens/1000000000000000000+amount)
				return nil, errors.New("fee number insufficient , fee Need to be " +decimalValue.String())
			}

			if err != nil {
				return nil, err
			}
			//广播转账信息
			hash, err := aeternity.SignBroadcast(spendTx, account, node, "ae_mainnet")

			//err = aeternity.WaitSynchronous(hash, config.Client.WaitBlocks, node)

			if err != nil {
				return nil, err
			}
			return hash, err
		} else {
			return nil, errors.New("tokens number insufficient")
		}
	} else {
		return nil, err
	}
}

//返回tx详细信息
func ApiThHash(th string) (tx *models.GenericSignedTx) {
	client := naet.NewNode(NodeURL, false)
	t, _ := client.GetTransactionByHash(th)
	return t
}

////获取Sophia vm 当前编译版本
//func ApiVersion() (v string) {
//	c := naet.NewCompiler("https://compiler.aepps.com", false)
//	v, _ = c.APIVersion()
//	return v
//}
//
////获取Sophia vm 当前编译版本
//func CompilerVersion() (v string) {
//	c := naet.NewCompiler("https://compiler.aepps.com", false)
//	v, _ = c.APIVersion()
//	return v
//}

//===================================================================================================================================================================================================
//|                           															AEX-9																										 |
///===================================================================================================================================================================================================

//创建AEX9代币
func CompileContractInit(account *account.Account, name string, number string) (s string, e error) {
	//创建节点
	n := naet.NewNode(NodeURL, false)
	//设置虚拟机地址
	c := naet.NewCompiler(compilerURL, false)
	//创context
	ctx := aeternity.NewContext(account, n)
	//设置编译器
	ctx.SetCompiler(c)
	//生成合约
	contract := aeternity.NewContract(ctx)
	//获取aex9 代币合约
	expected, _ := ioutil.ReadFile("contract/fungible-token.aes")
	//部署aex9 合约代买
	ctID, _, err := contract.Deploy(string(expected), "init", []string{"\"" + name + "\"", "18", "\"" + name + "\"", "Some(" + number + ")"}, config.CompilerBackendFATE)
	if err != nil {
		return "", err
	}
	//获取合约地址,将其返回
	_, err = n.GetContractByID(ctID)
	if err != nil {
		return "", err
	}
	return ctID, err
}

type CallInfoResult struct {
	CallInfo CallInfo `json:"call_info"`
}

type CallInfo struct {
	ReturnType  string `json:"return_type"`
	ReturnValue string `json:"return_value"`
}

//检查账户是否大于1ae
func Is1AE(address string) bool {
	//获取账户信息
	accountNet, err := ApiGetAccount(address)
	if err != nil {
		return false
	}
	//转换token
	tokens, err := strconv.ParseFloat(accountNet.Balance.String(), 64)
	if err != nil {
		return false
	}
	//判断token 是否大于1
	if tokens/1000000000000000000 < 1 {
		return false
	}
	return true
}

//调用aex9 合约方法
func CallContractFunction(account *account.Account, ctID string, function string, args []string) (s interface{}, e error) {
	//获取节点信息
	n := naet.NewNode(NodeURL, false)
	//获取编译器信息
	c := naet.NewCompiler(compilerURL, false)
	//创建上下文
	ctx := aeternity.NewContext(account, n)
	//关联编译器
	ctx.SetCompiler(c)
	//创建合约
	contract := aeternity.NewContract(ctx)
	//获取合约代码
	expected, _ := ioutil.ReadFile("contract/fungible-token.aes")
	//调用合约代码
	callReceipt, err := contract.Call(ctID, string(expected), function, args, config.CompilerBackendFATE)
	if err != nil {
		return nil, err
	}

	resp, err := http.Get(NodeURL + "/v2/transactions/" + callReceipt.Hash + "/info")
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	//获取合约调用信息
	//response := utils.Get(NodeURL + "/v2/transactions/" + callReceipt.Hash + "/info")
	//解析jSON
	var callInfoResult CallInfoResult
	err = json.Unmarshal(body, &callInfoResult)
	if err != nil {
		return nil, err
	}
	//解析结果
	decodeResult, err := c.DecodeCallResult(callInfoResult.CallInfo.ReturnType, callInfoResult.CallInfo.ReturnValue, function, string(expected), config.Compiler.Backend)
	if err != nil {
		return nil, err
	}
	//返回结果
	return decodeResult, err
}

//===================================================================================================================================================================================================
//|                           															AENS																										 |
///===================================================================================================================================================================================================

//aens 转账
func TransferAENS(account *account.Account, recipientAddress string, name string) (*aeternity.TxReceipt, error) {
	//获取节点
	client := naet.NewNode(NodeURL, false)
	//创建账户信息
	ctxAlice := aeternity.NewContext(account, client)
	// 创建转账aens tx
	transferTx, err := transactions.NewNameTransferTx(account.Address, name, recipientAddress, ctxAlice.TTLNoncer())
	if err != nil {
		return nil, err
	}
	// 广播
	txReceipt, err := ctxAlice.SignBroadcast(transferTx, config.Client.WaitBlocks)
	if err != nil {
		return nil, err
	}
	return txReceipt, err
}

//更新aens
func UpdateAENS(account *account.Account, name string) (*aeternity.TxReceipt, error) {
	//获取节点
	client := naet.NewNode(NodeURL, false)
	//获取账户信息
	ctxAlice := aeternity.NewContext(account, client)
	alicesAddress, err := transactions.NewNamePointer("account_pubkey", account.Address)
	if err != nil {
		return nil, err
	}
	//生成tx
	updateTx, err := transactions.NewNameUpdateTx(account.Address, name, []*transactions.NamePointer{alicesAddress}, config.Client.Names.ClientTTL+10000, ctxAlice.TTLNoncer())
	if err != nil {
		return nil, err
	}
	updateTx.NameTTL = 50000
	fmt.Println("Update")
	txReceipt, err := ctxAlice.SignBroadcast(updateTx, config.Client.WaitBlocks)
	if err != nil {
		return nil, err
	}

	return txReceipt, err
}

//注册域名
func ClaimAENS(account *account.Account, name string, fee *big.Int, isUpdate bool) (*aeternity.TxReceipt, error) {
	client := naet.NewNode(NodeURL, false)
	ctxAlice := aeternity.NewContext(account, client)
	preclaimTx, salt, err := transactions.NewNamePreclaimTx(account.Address, name, ctxAlice.TTLNoncer())
	if err != nil {
		return nil, err
	}

	_, err = ctxAlice.SignBroadcastWait(preclaimTx, config.Client.WaitBlocks)
	if err != nil {
		return nil, err
	}

	claimTx, err := transactions.NewNameClaimTx(account.Address, name, salt, fee, ctxAlice.TTLNoncer())
	if err != nil {
		return nil, err
	}

	if isUpdate {
		claimTx.NameSalt = new(big.Int)
		claimTx.TTL = 0
	}
	txReceipt, err := ctxAlice.SignBroadcast(claimTx, config.Client.WaitBlocks)

	if err != nil {
		return nil, err
	}
	return txReceipt, err
}

//===================================================================================================================================================================================================
//|                           															Oracle																										 |
///===================================================================================================================================================================================================

//注册预言鸡
func OracleRegister(account *account.Account, querySpace string, responseSpec string, queryFee float64) (hash string, oraclePubKey string) {
	client := naet.NewNode(NodeURL, false)
	ctxAlice := aeternity.NewContext(account, client)

	// Register Tx
	register, err := transactions.NewOracleRegisterTx(account.Address, querySpace, responseSpec, utils.GetRealAebalanceBigInt(queryFee), config.OracleTTLTypeDelta, 10000000, config.Client.Oracles.ABIVersion, ctxAlice.TTLNoncer())
	//register, err := transactions.NewOracleRegisterTx(account.Address, querySpace, responseSpec, utils.GetRealAebalanceBigInt(queryFee), config.OracleTTLTypeDelta, config.Client.Oracles.OracleTTLValue, config.Client.Oracles.ABIVersion, ctxAlice.TTLNoncer())
	if err != nil {
		println("err", err)
	}
	//println("account address",register.JSON())
	println("utils.GetRealAebalanceBigInt(queryFee) ", utils.GetRealAebalanceBigInt(queryFee).String())
	txReceipt, err := ctxAlice.SignBroadcast(register, config.Client.WaitBlocks)
	if err != nil {
		println(err)
	}
	oraclePubKey = register.ID()
	return txReceipt.Hash, oraclePubKey
}

func OracleQuery(account *account.Account, oracleID string, querySpec string, queryFee float64) (hash string, opId string) {
	client := naet.NewNode(NodeURL, false)
	ctxAlice := aeternity.NewContext(account, client)

	// Query
	query, err := transactions.NewOracleQueryTx(account.Address, oracleID, querySpec, utils.GetRealAebalanceBigInt(queryFee), 0, 100, 0, 100, ctxAlice.TTLNoncer())
	if err != nil {
		println(err)
	}
	txReceipt, err := ctxAlice.SignBroadcast(query, config.Client.WaitBlocks)
	if err != nil {
		println(err)
	}
	id, _ := query.ID()
	return txReceipt.Hash, id
}

func OracleResponse(account *account.Account, oracleID string, queryID string, response string) (hash string) {
	client := naet.NewNode(NodeURL, false)
	ctxAlice := aeternity.NewContext(account, client)
	// Respond
	respond, err := transactions.NewOracleRespondTx(account.Address, oracleID, queryID, response, config.OracleTTLTypeDelta, config.Client.Oracles.ResponseTTLValue, ctxAlice.TTLNoncer())
	if err != nil {
		println(err)
	}
	txReceipt, err := ctxAlice.SignBroadcast(respond, config.Client.WaitBlocks)
	if err != nil {
		println(err)
	}
	return txReceipt.Hash
}
