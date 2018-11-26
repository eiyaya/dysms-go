package dysms

import (
	"strings"
	// "net/url"
)

func specialUrlEncode(value string) string {
	// rstValue := url.QueryEscape(value)
	rstValue := "GET&/&"+value
	rstValue = strings.Replace(rstValue, "+", "%20", -1)
	rstValue = strings.Replace(rstValue, "*", "%2A", -1)
	rstValue = strings.Replace(rstValue, "%7E", "~", -1)
	return rstValue
}
