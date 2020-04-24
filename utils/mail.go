package utils

import (
	"encoding/json"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"gopkg.in/gomail.v2"
	"log"
	"strconv"
)

//发送邮箱验证码
func SendEMail(mail string, code string) bool {
	//定义收件人
	mailTo := []string{
		mail,
	}
	subject := "AEasy Framework Account"
	// 邮件正文
	body := "Hello，" + mail + "<br><br><br>" +
		"     You are in AEasy Framework It was registered，" +
		"If it is confirmed to be registered by you。 <br>" +
		"your verification code is：<br><br><br>" +
		code + " <br><br><br>" +
		"If you are not registered, please ignore this email。"

	err := sendMail(mailTo, subject, body)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

func sendMail(mailTo []string, subject string, body string) error {

	mailConn := map[string]string{
		"user": "admin@aeasy.io",
		"pass": "Lx125634897",
		"host": "smtp.exmail.qq.com",
		"port": "465",
	}

	port, _ := strconv.Atoi(mailConn["port"]) //转换端口类型为int

	m := gomail.NewMessage()

	m.SetHeader("From", m.FormatAddress(mailConn["user"], "AEasy Framework")) //这种方式可以添加别名，即“XX官方”
	//说明：如果是用网易邮箱账号发送，以下方法别名可以是中文，如果是qq企业邮箱，以下方法用中文别名，会报错，需要用上面此方法转码
	//m.SetHeader("From", "FB Sample"+"<"+mailConn["user"]+">") //这种方式可以添加别名，即“FB Sample”， 也可以直接用<code>m.SetHeader("From",mailConn["user"])</code> 读者可以自行实验下效果
	//m.SetHeader("From", mailConn["user"])
	m.SetHeader("To", mailTo...)    //发送给多个用户
	m.SetHeader("Subject", subject) //设置邮件主题
	m.SetBody("text/html", body)    //设置邮件正文

	d := gomail.NewDialer(mailConn["host"], port, mailConn["user"], mailConn["pass"])

	err := d.DialAndSend(m)
	return err

}

func SendRegisterSms(phoneNumber string, data string) (res *dysmsapi.SendSmsResponse, e error) {
	client, err := dysmsapi.NewClientWithAccessKey("cn-hangzhou", "LTAI4FhFrGzcs39RnXn3NhZu", "Op4sUb3Uj2rLu5z1aIRWMWBWEz94jr")

	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"

	request.PhoneNumbers = phoneNumber
	request.SignName = "aeasy"
	request.TemplateCode = "SMS_187261217"
	m := map[string]string{"code": data}
	mjson, _ := json.Marshal(m)
	mString := string(mjson)
	request.TemplateParam = mString

	response, err := client.SendSms(request)
	if err != nil {
		fmt.Print(err.Error())
		return response, err
	}
	fmt.Printf("response is %#v\n", response)
	return response, err

}
