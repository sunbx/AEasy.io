package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	crypt_rand "crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/shopspring/decimal"
	"io"
	"io/ioutil"
	"math"
	"math/big"
	"math/rand"
	"net/http"
	"regexp"
	"strconv"
	"time"
)

type Resp struct {
	Code string      `json:"code"`
	Msg  string      `json:"msg"`
	data interface{} `json:"object"`
}

/**
* @des 时间转换函数
* @param atime string 要转换的时间戳（秒）
* @return string
 */
func StrTime(atime int64) string {
	var byTime = []int64{365 * 24 * 60 * 60, 24 * 60 * 60, 60 * 60, 60, 1}
	var unit = []string{"年", "天", "小时", "分钟", "秒钟"}
	now := time.Now().Unix()
	ct := now - atime
	if ct < 0 {
		return "已结束"
	}
	var res string
	for i := 0; i < len(byTime); i++ {
		if ct < byTime[i] {
			continue
		}
		var temp = math.Floor(float64(ct / byTime[i]))
		ct = ct % byTime[i];
		if temp > 0 {
			var tempStr string
			tempStr = strconv.FormatFloat(temp, 'f', -1, 64)
			res = MergeString(tempStr, unit[i]) //此处调用了一个我自己封装的字符串拼接的函数（你也可以自己实现）
		}
		break //我想要的形式是精确到最大单位，即："2天前"这种形式，如果想要"2天12小时36分钟48秒前"这种形式，把此处break去掉，然后把字符串拼接调整下即可（别问我怎么调整，这如果都不会我也是无语）
	}
	return res
}

/**
* @des 拼接字符串
* @param args ...string 要被拼接的字符串序列
* @return string
 */
func MergeString(args ...string) string {
	buffer := bytes.Buffer{}
	for i := 0; i < len(args); i++ {
		buffer.WriteString(args[i])
	}
	return buffer.String()
}

func CreateCaptcha() string {
	return fmt.Sprintf("%04v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(10000))
}
func Md5V(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

//生成Guid字串
func UniqueId(salt string) string {
	b := make([]byte, 8)

	if _, err := io.ReadFull(crypt_rand.Reader, b); err != nil {
		return ""
	}
	s := base64.URLEncoding.EncodeToString(b) + salt
	return base64.URLEncoding.EncodeToString([]byte(s))
}

func FormatTokens(tokens float64) string {
	if tokens > 0 {
		decimalValue := decimal.NewFromFloat(tokens)
		decimalValue = decimalValue.Div(decimal.NewFromInt(1000000000000000000))
		f, _ := decimalValue.Float64()
		str := strconv.FormatFloat(math.Trunc(f/0.00001) * 0.00001, 'f', 5, 64)
		return str
	} else {
		return "0"
	}
}

func FormatTokensInt(tokens float64) float64 {
	if tokens > 0 {
		decimalValue := decimal.NewFromFloat(tokens)
		decimalValue = decimalValue.Div(decimal.NewFromInt(1000000000000000000))
		f, _ := decimalValue.Float64()
		str := strconv.FormatFloat(f, 'f', 5, 64)
		parseFloat, _ := strconv.ParseFloat(str, 64)
		return parseFloat
	} else {
		return 0
	}
}

func FormatTokensP(tokens float64, p int) string {
	if tokens > 0 {
		decimalValue := decimal.NewFromFloat(tokens)
		decimalValue = decimalValue.Div(decimal.NewFromInt(1000000000000000000))
		f, _ := decimalValue.Float64()
		str := strconv.FormatFloat(f, 'f', p, 64)
		return str
	} else {
		return "0"
	}
}

func GetAENSPalce(name string) string {
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
	return price
}

func GetRealAebalanceBigInt(amount float64) *big.Int {
	newFloat := big.NewFloat(amount)
	basefloat := big.NewFloat(1000000000000000000)
	float1 := big.NewFloat(1)
	float1.Mul(newFloat, basefloat)
	resultAmount := new(big.Int)
	float1.Int(resultAmount)
	return resultAmount
}

func GetRealAebalanceFloat64(amount float64) float64 {
	return amount * 1000000000000000000
}

func AesEncrypt(orig string, key string) string {
	// 转成字节数组
	origData := []byte(orig)
	k := []byte(key)
	// 分组秘钥
	// NewCipher该函数限制了输入k的长度必须为16, 24或者32
	block, _ := aes.NewCipher(k)
	// 获取秘钥块的长度
	blockSize := block.BlockSize()
	// 补全码
	origData = PKCS7Padding(origData, blockSize)
	// 加密模式
	blockMode := cipher.NewCBCEncrypter(block, k[:blockSize])
	// 创建数组
	cryted := make([]byte, len(origData))
	// 加密
	blockMode.CryptBlocks(cryted, origData)
	return base64.StdEncoding.EncodeToString(cryted)
}
func AesDecrypt(cryted string, key string) string {
	// 转成字节数组
	crytedByte, _ := base64.StdEncoding.DecodeString(cryted)
	k := []byte(key)
	// 分组秘钥
	block, _ := aes.NewCipher(k)
	// 获取秘钥块的长度
	blockSize := block.BlockSize()
	// 加密模式
	blockMode := cipher.NewCBCDecrypter(block, k[:blockSize])
	// 创建数组
	orig := make([]byte, len(crytedByte))
	// 解密
	blockMode.CryptBlocks(orig, crytedByte)
	// 去补全码
	orig = PKCS7UnPadding(orig)
	return string(orig)
}

//补码
//AES加密数据块分组长度必须为128bit(byte[16])，密钥长度可以是128bit(byte[16])、192bit(byte[24])、256bit(byte[32])中的任意一个。
func PKCS7Padding(ciphertext []byte, blocksize int) []byte {
	padding := blocksize - len(ciphertext)%blocksize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

//去码
func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func IsEmail(email string) bool {
	// 识别电子邮件地址
	isorno, _ := regexp.MatchString(`\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*`, email)

	if isorno {
		return true
	} else {
		return false
	}
}

//检测是不是手机访问
func IsMobile(userAgent string) bool {
	mobileRe, _ := regexp.Compile("(?i:Mobile|iPod|iPhone|Android|Opera Mini|BlackBerry|webOS|UCWEB|Blazer|PSP)")
	if mobileRe.FindString(userAgent) == "" {
		return false
	}
	return true
}

//发送GET请求
//url:请求地址
//response:请求返回的内容
func Get(url string) (response string) {
	client := http.Client{Timeout: 600 * time.Second}
	resp, error := client.Get(url)
	defer resp.Body.Close()
	if error != nil {
		panic(error)
	}
	var buffer [512]byte
	result := bytes.NewBuffer(nil)
	for {
		n, err := resp.Body.Read(buffer[0:])
		result.Write(buffer[0:n])
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
	}
	response = result.String()
	return

}


//发送POST请求
//url:请求地址		data:POST请求提交的数据		contentType:请求体格式，如：application/json
//content:请求返回的内容
func PostBody(url string, data string, contentType string) (content string) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(data)))
	req.Header.Add("content-type", contentType)
	if err != nil {
		return ""
	}
	defer req.Body.Close()
	client := &http.Client{Timeout: 5 * time.Second}
	resp, error := client.Do(req)
	if error != nil {
		return ""
	}
	defer resp.Body.Close()
	result, _ := ioutil.ReadAll(resp.Body)
	content = string(result)
	return
}
