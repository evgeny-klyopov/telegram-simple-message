# telegram-simple-message

### Install
```bash
go get "github.com/evgeny-klyopov/telegram-simple-message"
```

### Usage
```go
import "github.com/evgeny-klyopov/telegram-simple-message"

func main()  {
	client, _ := telegram.NewClient(telegram.Config{
		UseProxy: true,
		Proxy: "user:password@ip:port",
		Token: "DDDDDDDDDD:AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA",
		ChatId: -0000000000000,
	})

	_ = client.Send("*test*", telegram.MarkdownTypeMessage)
	_ = client.Send("<b>test</b>", telegram.HtmlTypeMessage)
}
```
### Help
* DDDDDDDDDD:AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA - bot token
* -0000000000000 - chatId, link for get chatId https://api.telegram.org/botDDDDDDDDDD:AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA/getChat?chat_id=@NameChannelOrNameClient. example response:
{"ok":true,"result":{"id":-0000000000000,"title":"NameChannelOrNameClient","username":"NameChannelOrNameClient","type":"channel"}}






