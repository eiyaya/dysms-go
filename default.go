package dysms

import (
)
var (
	// 用于 Default 的 AccessKeyId
	AccessKeyId string
	// 用于 Default 的 AccessSecret
	AccessSecret string
	// 用于 Default 的 TemplateCode
	TemplateCode string
	// 默认的SmsSender
	Default *SmsSender
)

// 初始化默认 SmsSender
func init() {
	Default, _ = NewSmsSender("Default")
	Default.GetAccessKeyId = func() string {
		return AccessKeyId
	}
	Default.GetAccessSecret = func() string {
		return AccessSecret
	}
	Default.GetTeplateCode = func() string {
		return TemplateCode
	}
	smsSenders["default"] = Default
}

// 设置AccessKeyId
func SetAccessKeyId(accessKeyId string) {
	AccessKeyId = accessKeyId
	Default.AccessKeyId = accessKeyId
}

// 设置AccessSecret
func SetAccessSecret(accessSecret string) {
	AccessSecret = accessSecret
	Default.AccessSecret = accessSecret
}

// 设置TemplateCode
func SetTemplateCode(templateCode string)  {
	TemplateCode = templateCode
	Default.TemplateCode = templateCode
}
