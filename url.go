package dysms

import (
	"sort"
	"strings"
	"net/url"
)

type Values interface {
	Get(string) string
	Set(string, string)
	Add(string, string)
	Del(string)
	Encode() string
}
// 用于保存生成请求的参数
type values url.Values

// 参照url.Values URL编码，
// 编写POP协议中的特殊编码函数
func (v values) Encode() string {
	if v == nil {
		return ""
	}
	var buf strings.Builder
	keys := make([]string, 0, len(v))
	for k := range v {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		vs := v[k]
		// keyEscaped := QueryEscape(k)
		// 换用特殊的编码格式
		keyEncoded := SpecialEncode(k)
		for _, v := range vs {
			if buf.Len() > 0 {
				buf.WriteByte('&')
			}
			buf.WriteString(keyEncoded)
			buf.WriteByte('=')
			// buf.WriteString(QueryEscape(v))
			// 换用特殊的编码格式
			buf.WriteString(SpecialEncode(v))
		}
	}
	return buf.String()
}

// 采用POP协议特殊的编码格式编码字符
func SpecialEncode(value string) string {
	rstValue := url.QueryEscape(value)
	rstValue = strings.Replace(rstValue, "+", "%20", -1)
	rstValue = strings.Replace(rstValue, "*", "%2A", -1)
	rstValue = strings.Replace(rstValue, "%7E", "~", -1)
	return rstValue
}
