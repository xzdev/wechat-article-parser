package main

import (
	"log"

	"github.com/xzdev/wechat-article-parser"
)

func main() {
	article := parser.ParseArticle("http://mp.weixin.qq.com/s?__biz=MTQzMjE1NjQwMQ==&mid=2655537855&idx=1&sn=6fe48ecaf2d8a367a36e7db252e419e7&chksm=66dfe52151a86c3715228ead2e9ccc099649208cb3e86c947c7c60342b60a047a7656c9f8f73&scene=0#rd")
	log.Println(article.Title, article.Summary)
}
