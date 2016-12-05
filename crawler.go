package crawler

import (
	"log"
	"net/http"
	"strings"

	"time"

	"github.com/yhat/scrape"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

type Article struct {
	Title       string   `json:"title"`
	Author      string   `json:"author"`
	Url         string   `json:"url"`
	Summary     string   `json:"summary"`
	Photos      []string `json:"photos"`
	Readtime    int64    `json:"readtime"`
	Publishtime int64    `json:"publishtime"`
}

func ParseArticle(url string) *Article {
	resp, err := http.Get(url)
	if err != nil {
		log.Println("Cannot fetch content from url", url, "with error message", err)
		return nil
	}

	root, err := html.Parse(resp.Body)
	if err != nil {
		log.Println("Cannot parse the html page. Error:", err)
	}

	title, _ := scrape.Find(root, titleMatcher)
	richMediaList, _ := scrape.Find(root, richMediaListMatcher)
	richMediaContent, _ := scrape.Find(root, richMediaContentMatcher)

	postDate, _ := scrape.Find(richMediaList, createTimeMatcher)
	author, _ := scrape.Find(richMediaList, authorMatcher)

	article := new(Article)
	// set title
	article.Title = scrape.Text(title)
	// set url
	article.Url = url

	// process publish time
	publishTime, err := time.Parse("2006-01-02", scrape.Text(postDate))
	if err == nil {
		article.Publishtime = publishTime.Unix() // use short format to parse the date
	} else {
		log.Println("Cannot get the publish time")
	}

	// set read time to current timestamp
	article.Readtime = time.Now().Unix()

	// set author
	article.Author = scrape.Text(author)

	// process content text
	contentNodes := scrape.FindAll(richMediaContent, contentText)
	summaryText := []string{}
	for _, node := range contentNodes {
		summaryText = append(summaryText, scrape.Text(node))
	}
	article.Summary = strings.Join(summaryText, "\n")

	imageNodes := scrape.FindAll(richMediaContent, contentImage)
	images := []string{}
	for _, node := range imageNodes {
		images = append(images, scrape.Attr(node, "data-src"))
	}
	article.Photos = images

	//log.Println("Parsed article", article)
	return article
}

// define title matcher
func titleMatcher(n *html.Node) bool {
	if n.DataAtom == atom.H2 {
		return scrape.Attr(n, "class") == "rich_media_title"
	}
	return false
}

// define rich_media_meta_list matcher. The node should contain publish date and author info
func richMediaListMatcher(n *html.Node) bool {
	if n.DataAtom == atom.Div {
		return scrape.Attr(n, "class") == "rich_media_meta_list"
	}
	return false
}

// define rich_media_content matcher. The node contains the article content text and photos
func richMediaContentMatcher(n *html.Node) bool {
	return scrape.ById("js_content")(n)
}

// define article create time matcher
func createTimeMatcher(n *html.Node) bool {
	return scrape.ById("post-date")(n)
}

// define article author matcher
func authorMatcher(n *html.Node) bool {
	return scrape.ById("post-user")(n)
}

// define article content text nodes matcher
func contentText(n *html.Node) bool {
	return n.DataAtom == atom.Section || n.DataAtom == atom.P
}

// define article images
func contentImage(n *html.Node) bool {
	return n.DataAtom == atom.Img
}
