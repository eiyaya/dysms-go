# 阿里云短信 dysms api 的go实现

## 示例

```
import "github.com/dysms-go"

func	main() {
	dysms.AccessKeyId = "Your AccessKeyId"
	dysms.AccessSecret = "Your AccessSecret"
	dysms.SendWithTemplate("TemplateCode", "TemplateParam")
}
```
