# wechat-article-parser
A wechat article parser wirtten in Golang

The library parse the wechat public account article ang generate a data object in the following format.
```go
type Article struct {
	Title       string   `json:"title"`
	Author      string   `json:"author"`
	Url         string   `json:"url"`
	Summary     string   `json:"summary"`
	Photos      []string `json:"photos"`
	Readtime    int64    `json:"readtime"`
	Publishtime int64    `json:"publishtime"`
}
```
## Example
```go
package main

import (
	"log"

	"github.com/xzdev/wechat-article-parser"
)

func main() {
	article := parser.ParseArticle("http://mp.weixin.qq.com/wechat-article-url-here")
	log.Println(article.Title, article.Summary)
}
```

## License
This library is under the [MIT License](http://opensource.org/licenses/MIT)
