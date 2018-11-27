package dysms

import (
	"time"
	"strings"
	"errors"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
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
	smsSenders = make(map[string]*SmsSender)
	// 默认的SmsSender
	Default *SmsSender
)

// 初始化默认 SmsSender
func init() {
	Default = &SmsSender{}
	smsSenders["default"] = Default
}

// 根据给出的名称获取SmsSender
// 可选参数 AccessKeyId AccessSecret SignName TemplateCode
func GetSmsSender(name string, param ...string) (smsSender *SmsSender) {
	// 如果smsSenders中保存了以name为键的SmsSender,
	// 返回对应的SmsSender,
	// 如果没有，创建一个新的SmsSender并返回
	_, ok := smsSenders[name]
	if !ok {
		sender := NewSmsSender()
		if len(param) == 3 {

		}
		smsSenders[name] = sender
	}
	return smsSenders[name]
}

// 实现短信请求的组装和发送
type SmsSender struct {
	// AccessKeyId 和 AccessSecret
	AccessKeyId string
	AccessSecret string
	// 签名
	SignName string
	// 模板码
	TemplateCode string
	Values Values
	GetSignatureNonce func() string
}

// 创建SmsSender
func NewSmsSender() *SmsSender {
	s := &SmsSender{}
	s.Values.Add("Action", "SendSms")
	s.Values.Add("Version", "2017-05-25")
	s.Values.Add("RegionId", "cn-hangzhou")
	return s
}

func (s *SmsSender) timestamp() string {
	return time.Now().UTC().Format("2006-01-02T15:04:05Z")
}

// 设置 SmsSender 的参数
func (s *SmsSender) Set(accessKeyId, accessSecret, signName, templateCode string) {
	s.AccessKeyId = accessKeyId
	s.AccessSecret = accessSecret
	s.SignName = signName
	s.TemplateCode = templateCode
}

// 根据给出的 phoneNumbers templateParam发送短信
func (s *SmsSender) Send(phoneNumbers, templateParam string, outId ...string) (successed bool, err error) {
	s.Values.Add("AccessKeyId", s.AccessKeyId)
	s.Values.Add("Timestamp", s.timestamp())
	// 业务参数
	s.Values.Add("SignName", s.SignName)
	s.Values.Add("TemplateCode", s.TemplateCode)
	s.Values.Add("PhoneNumbers", phoneNumbers)
	s.Values.Add("TemplateCode", s.TemplateCode)
	s.Values.Add("TemplateParam", templateParam)
	if (len(outId) > 0) {
		s.Values.Add("OutId", outId[0])
	} else {
		s.Values.Del("OutId")
	}
	queryString := s.signature()
	Client.Get(Path+"?"+queryString)
	// TODO: 结果检查
	return
}

// 使用给出的 accessSecret 进行签名
func (s *SmsSender) signature() (queryString string) {
	// 签名方法
	s.Values.Add("SignatureMethod", "HMAC-SHA1")
	// 签名版本
	s.Values.Add("SignatureVersion", "1.0")
	// 随机数
	// s.Values.Add("SignatureNonce", s.GetSignatureNonce())
	s.Values.Add("SignatureNonce", uuid.Must(uuid.NewV4()).String())
	// 去除签名字段
	s.Values.Del("Signature")
	sortedQueryString := s.Values.Encode()
	var buf strings.Builder
	buf.WriteString("GET")
	buf.WriteString("&")
	buf.WriteString(SpecialEncode("/"))
	buf.WriteString("&")
	buf.WriteString(SpecialEncode(sortedQueryString))
	stringToSign := buf.String()
	// 进行签名
	h := hmac.New(sha1.New, []byte(s.AccessSecret+"&"))
	h.Write([]byte(stringToSign))
	signedString := base64.StdEncoding.EncodeToString(h.Sum(nil))
	return "Signature="+signedString+"&"+sortedQueryString
}

// 解析返回值
func parseResp(resp *http.Response) (err error) {
	return nil
}
