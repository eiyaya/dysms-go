package dysms

import (
	"time"
	"errors"
	"net/url"
	"net/http"
	// "crypto/rand"
)

var (
	// 已经存在一个同名的SmsSender
	ErrAlreadyExisted = errors.New("SmsSender already existed with the same name!")
	// SmsSender.TemplateCode没有设置
	ErrTemplateCodeNotSet = errors.New("TemplateCode not set!")
)

var (
	// 用于发送http请求的Client
	Client = &http.Client{}
	// 记录所有创建的SmsSender
	smsSenders = make(map[string]*SmsSender)
)

// 根据给出的名称获取SmsSender
func GetSmsSender(names ...string) (smsSender *SmsSender) {
	// 如果没有参数，返回默认smsSender，
	// 如果有参数，根据name[0] 返回smsSender
	if len(names) <=0 {
		return Default
	} else {
		name := names[0]
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

func NewSmsSender(name string) (smsSender *SmsSender, err error) {
	// 检查smsSenders 中是否存在以name为键的SmsSender，
	// 如果已经存在，返回已存在错误
	smsSender, ok := smsSenders[name]
	if ok {
		return nil, ErrAlreadyExisted
	}
	// 创建新的SmsSender
	smsSender = &SmsSender{
		Format: "json",
		SigntureMethod: "HMAC-SHA1",
	}
	smsSender.GetAccessKeyId = func() string {
		return smsSender.AccessKeyId
	}
	smsSender.GetAccessSecret = func() string {
		return smsSender.AccessSecret
	}
	smsSender.GetTimestamp = func() string {
		return time.Now().Format("2006-01-02T15:04:05Z")
	}
	smsSender.GetSignatureNonce = func() string {
		// 实现返回随机字符串
		// TODO: 参考uuidV4
		return ""
	}
	// 以name为键，向smsSenders中添加新创建的SmsSender
	smsSenders[name] = smsSender
	return smsSender, nil
}

type Request struct {
	// 系统参数
	Timestamp string
	Format string
	SigntureMethod string
	SignatureVersion string
	SignatureNonce func() string
	Signature string
	// 业务参数
	Action string
	Version string
	RegionId string
	PhoneNumbers string
	SignName string
	TemplateCode string
	TemplateParam string
	OutId string
}

// 实现短信请求的组装和发送
type SmsSender struct {
	// AccessKeyId 和 AccessSecret
	AccessKeyId string
	AccessSecret string
	// 模板码
	TemplateCode string
	// 格式
	Format string
	SigntureMethod string

	// 获取AccessKeyId 和 AccessSecret的函数
	GetAccessKeyId func() string
	GetAccessSecret func() string
	GetTeplateCode func() string
	GetSignatureNonce func() string
	// 获取时间戳的函数
	GetTimestamp func() string
}

// 根据给出的TemplateParam发送短信
func (s *SmsSender) Send(templateParam string) (successed bool, err error) {
	// 检查s.TemplateCode
	if s.TemplateCode == "" {
		return false, ErrTemplateCodeNotSet
	}
	return s.SendWithTemplate(s.TemplateCode, templateParam)
}

// 按照给出的 TemplateCode 和 TemplateParam 发送短信
func (s *SmsSender) SendWithTemplate(templateCode, templateParam string) (successed bool, err error) {
	// request := NewRequest()
	locator, err := url.Parse("dysmsapi.aliyuncs.com")
	resp, err := Client.Get(locator.String())
	// 如果请求失败返回失败信息
	if err != nil {
		return false, err
	}
	// 解析响应
	err = parseResp(resp)
	if err != nil {
		return false, err
	}
	return true, nil
}

// 解析返回值
func parseResp(resp *http.Response) (err error) {
	return nil
}
