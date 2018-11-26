package dysms

import (
	"time"
	// "strings"
	"errors"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"net/url"
	"net/http"

	"github.com/satori/go.uuid"
	// "crypto/rand"
)

var (
	// 已经存在一个同名的SmsSender
	ErrAlreadyExisted = errors.New("SmsSender already existed with the same name!")
	// SmsSender.TemplateCode没有设置
	ErrTemplateCodeNotSet = errors.New("TemplateCode not set!")
)

var (
	// 接口地址
	Path string = "http://dysmsapi.aliyuncs.com/"
	// 用于发送http请求的Client
	Client = &http.Client{}
	// 记录所有创建的SmsSender
	smsSenders = make(map[string]*smsSender)
	// 默认的SmsSender
	Default *smsSender
)

// 初始化默认 SmsSender
func init() {
	Default, _ = &smsSender{}
	smsSenders["default"] = Default
}

// 根据给出的名称获取SmsSender
func GetSmsSender(param ...string) (smsSender *smsSender) {
	// 如果没有参数，返回默认smsSender，
	// 如果有参数，根据name[0] 返回smsSender
	if len(param) <= 0 {
		return Default
	} else {
		name := param[0]
		// 如果smsSenders中保存了以name为键的SmsSender,
		// 返回对应的SmsSender,
		// 如果没有，创建一个新的SmsSender并返回
		_, ok := smsSenders[name]
		if !ok {
			NewSmsSender(name)
		}
		return smsSenders[name]
	}
}



// 实现短信请求的组装和发送
type smsSender struct {
	// AccessKeyId 和 AccessSecret
	AccessKeyId string
	AccessSecret string
	// 签名
	SignName string
	// 模板码
	TemplateCode string
	// 执行签名的函数
	SignatureHandler func(values url.Values, accessSecret string) (rstValues url.Values, sign string)
}

func (s *smsSender) timestamp() string {
	return time.Now().UTC().Format("2006-01-02T15:04:05Z")
}

// 根据给出的 phoneNumbers templateParam发送短信
func (s *smsSender) Send(phoneNumbers, templateParam string) (successed bool, err error) {
	return RawSend(phoneNumbers, s.SignName, s.TemplateCode, templateParam)
}

// 按照给出的 所有参数 发送短信
func RawSend(accessKeyId, accessSecret, phoneNumbers, signName, templateCode, templateParam string) (requestId, code, message, bizId string) {
	values := url.Values{}
	// 权限ID
	values.Add("AccessKeyId", accessKeyId)
	// 电话号码
	values.Add("PhoneNumbers", phoneNumbers)
	// 签名
	values.Add("SignName", signName)
	// 模板码
	values.Add("TemplateCode", templateCode)
	// 模板数据
	values.Add("TemplateParam", templateParam)
	// 时间戳
	values.Add("Timestamp", time.Now().UTC().Format("2006-01-02T15:04:05Z"))
	// 行为
	values.Add("Action", "SendSms")
	values.Add("Version", "2017-05-25")
	values.Add("RegionId", "cn-hangzhou")
	signedArgs = Signature(values)
	// 发送请求
	resp, err := Client.Get(Path+"?"+signedArgs)
	// 处理结果
	return parseResp(resp)
}

// 使用给出的 accessSecret 进行签名
func Signature(values url.Values, accessSecret string) (signedArgs string) {
	// 签名方法
	values.Add("SignatureMethod", "HMAC-SHA1")
	// 签名版本
	values.Add("SignatureVersion", "1.0")
	// 随机数
	values.Add("SignatureNonce", uuid.Must(uuid.NewV4()).String())
	// 去除签名字段
	values.Del("Signature")
	urlEncoded := values.Encode()
	stringToSign := "GET"+"&"+specialUrlEncode("/")+"&"+specialUrlEncode(urlEncoded)
	h := hmac.New(sha1.New, []byte(accessSecret+"&"))
	h.Write([]byte(stringToSign))
	signString := base64.StdEncoding.EncodeToString(h.Sum(nil))
	return values, signString
}

// 解析返回值
func parseResp(resp *http.Response) (err error) {
	return nil
}
