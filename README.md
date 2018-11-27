# 阿里云短信 dysms api 的go实现

## 示例1

```
import "github.com/dysms-go"

func main() {
	dysms.Set("AccessKeyId", "AccessSecret", "SignName", "TemplateCode")
	dysms.Send("PhoneNumbers", "TemplateParam")
}
```

## 示例2

```
import "github.com/dysms-go"

func main() {
	sender := dysms.GetSmsSender("Mysender")
	sender.Set("AccessKeyId", "AccessSecret", "SignName", "TemplateCode")
	sender.Send("PhoneNumbers", "TemplateParam")
}
```

## 示例3

```
import "github.com/dysms-go"

func main() {
	sender := dysms.NewSmsSender()
	sender.Set("AccessKeyId", "AccessSecret", "SignName", "TemplateCode")
	sender.Send("PhoneNumbers", "TemplateParam")
}
```

### 示例4

```
import "github.com/dysms-go"

func main() {
	dysms.RawSend("AccessKeyId", "AccessSecret", "SignName", "TemplateCode", "PhoneNumbers", "TemplateParam")
}
```
