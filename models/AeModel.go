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
	"github.com/tyler-smith/go-bip39"
	"io/ioutil"
	"strconv"
)

//var nodeURL = "https://mainnet.aeternal.io"
var nodeURL = "http://node.aechina.io:3013"

//var nodeURL = nodeURL
//根据助记词返回用户
func MnemonicAccount(mnemonic string) (*account.Account, error) {

	seed, err := account.ParseMnemonic(mnemonic)
	if err != nil {
		return nil, err
	}
	_, err = bip39.EntropyFromMnemonic(mnemonic)

	if err != nil {
		return nil, err
	}
	// Derive the subaccount m/44'/457'/3'/0'/1'
	key, err := account.DerivePathFromSeed(seed, 0, 0)
	if err != nil {
		return nil, err
	}

	// Deriving the aeternity Account from a BIP32 Key is a destructive process
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

//随机创建用户
func CreateAccountUtils() (mnemonic string, signingKey string, address string) {

	//cerate mnemonic
	entropy, _ := bip39.NewEntropy(128)
	mne, _ := bip39.NewMnemonic(entropy)

	//mnemonic := "tail disagree oven fit state cube rule test economy claw nice stable"
	seed, _ := account.ParseMnemonic(mne)

	_, _ = bip39.EntropyFromMnemonic(mne)
	// Derive the subaccount m/44'/457'/3'/0'/1'
	key, _ := account.DerivePathFromSeed(seed, 0, 0)

	// Deriving the aeternity Account from a BIP32 Key is a destructive process
	alice, _ := account.BIP32KeyToAeKey(key)
	return mne, alice.SigningKeyToHexString(), alice.Address
}

//返回最新区块高度
func ApiBlocksTop() (height uint64) {
	client := naet.NewNode(nodeURL, false)
	h, _ := client.GetHeight()
	return h
}

//地址信息返回用户信息
func ApiGetAccount(address string) (account *models.Account, e error) {
	client := naet.NewNode(nodeURL, false)
	acc, e := client.GetAccount(address)
	return acc, e
}

//发起转账
func ApiSpend(account *account.Account, recipientId string, amount float64, data string) (*aeternity.TxReceipt, error) {

	accountNet, e := ApiGetAccount(account.Address)
	if e != nil {
		return nil, e
	}
	tokens, err := strconv.ParseFloat(accountNet.Balance.String(), 64)
	if err == nil {
		if tokens/1000000000000000000 > amount {
			node := naet.NewNode(nodeURL, false)
			//_, _, ttlnoncer := transactions.GenerateTTLNoncer(node)
			ttlnoncer := transactions.NewTTLNoncer(node)

			spendTx, err := transactions.NewSpendTx(account.Address, recipientId, utils.GetRealAebalanceBigInt(amount), []byte(data), ttlnoncer)
			if err != nil {
				return nil, err
			}
			hash, err := aeternity.SignBroadcast(spendTx, account, node, "ae_mainnet")
			return hash, err
		} else {
			return nil, errors.New("tokens number insufficient")
		}
	} else {
		return nil, err
	}
}

//获取Sophia vm 当前编译版本
func ApiVersion() (v string) {
	c := naet.NewCompiler("https://compiler.aepps.com", false)
	v, _ = c.APIVersion()
	return v
}

//返回tx详细信息
func ApiThHash(th string) (tx *models.GenericSignedTx) {
	client := naet.NewNode(nodeURL, false)
	t, _ := client.GetTransactionByHash(th)
	return t
}

//获取Sophia vm 当前编译版本
func CompilerVersion() (v string) {
	c := naet.NewCompiler("https://compiler.aepps.com", false)
	v, _ = c.APIVersion()
	return v
}

//创建AEX9代币
func CompileContractInit(account *account.Account, name string, number string) (s string, e error) {
	n := naet.NewNode(nodeURL, false)
	c := naet.NewCompiler("https://compiler.aepps.com", true)
	ctx := aeternity.NewContext(account, n)
	ctx.SetCompiler(c)
	contract := aeternity.NewContract(ctx)
	expected, _ := ioutil.ReadFile("contract/fungible-token.aes")
	ctID, _, err := contract.Deploy(string(expected), "init", []string{"\"" + name + "\"", "18", "\"" + name + "\"", "Some(" + number + ")"}, config.CompilerBackendFATE)
	if err != nil {
		return "", err
	}

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

//是否大于1ae
func Is1AE(address string) bool {
	accountNet, err := ApiGetAccount(address)
	if err != nil {
		return false
	}
	tokens, err := strconv.ParseFloat(accountNet.Balance.String(), 64)
	if err != nil {
		return false
	}
	if tokens/1000000000000000000 < 1 {
		return false
	}
	return true
}

//调用aex9 合约方法
func CallContractFunction(account *account.Account, ctID string, function string, args []string) (s interface{}, e error) {
	n := naet.NewNode(nodeURL, false)
	//c := naet.NewCompiler(nodeURL, true)
	c := naet.NewCompiler("https://compiler.aepps.com", false)
	ctx := aeternity.NewContext(account, n)
	ctx.SetCompiler(c)
	contract := aeternity.NewContract(ctx)
	expected, _ := ioutil.ReadFile("contract/fungible-token.aes")
	callReceipt, err := contract.Call(ctID, string(expected), function, args, config.CompilerBackendFATE)
	if err != nil {
		return nil, err
	}
	response := utils.Get(nodeURL + "/v2/transactions/" + callReceipt.Hash + "/info")
	var callInfoResult CallInfoResult
	err = json.Unmarshal([]byte(response), &callInfoResult)
	if err != nil {
		return nil, err
	}
	//tx, e := n.GetTransactionByHash(callReceipt.Hash)
	//tx.Tx()
	//genericTx, e := models.UnmarshalGenericTx(bytes.NewBuffer(i.), runtime.())
	decodeResult, err := c.DecodeCallResult(callInfoResult.CallInfo.ReturnType, callInfoResult.CallInfo.ReturnValue, function, string(expected), config.Compiler.Backend)
	if err != nil {
		return nil, err
	}
	//fmt.Println(genericTx)

	return decodeResult, err
}

func TransferAENS(account *account.Account, recipientAddress string, name string) ( *aeternity.TxReceipt, error) {
	client := naet.NewNode(nodeURL, false)
	ctxAlice := aeternity.NewContext(account, client)
	// Transfer the name to a recipient
	transferTx, err := transactions.NewNameTransferTx(account.Address, name, recipientAddress, ctxAlice.TTLNoncer())
	if err != nil {
		return nil, err
	}
	fmt.Println("Transfer")
	txReceipt, err := ctxAlice.SignBroadcastWait(transferTx, config.Client.WaitBlocks)
	if err != nil {
		return nil, err
	}
	return txReceipt, err
}
