package parser

import "testing"

func TestParseArticle(t *testing.T) {
	//ParseArticle("http://mp.weixin.qq.com/s?__biz=MTQzMjE1NjQwMQ==&mid=2655537855&idx=1&sn=6fe48ecaf2d8a367a36e7db252e419e7&chksm=66dfe52151a86c3715228ead2e9ccc099649208cb3e86c947c7c60342b60a047a7656c9f8f73&scene=0#rd")
	article := ParseArticle("http://mp.weixin.qq.com/s?__biz=MzA3NzI5MTEzMg==&mid=2658506017&idx=1&sn=43ae2fd172fbf29e752c3d564bc91bbf&chksm=84d57e32b3a2f724d86a337a5aceac6d3f7403e27a5de58142468bea77627f4f49f5ee9e3264&scene=0#rd")
	if article.Author != "油麻菜" {
		t.Error("Failed to parse author")
	}

	if article.Title != "紫薇的翅膀" {
		t.Error("Failed to parse title")
	}
}
