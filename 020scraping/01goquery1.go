package main
import (
 "fmt"
 "github.com/PuerkitoBio/goquery"
)

func main() {
 doc, err := goquery.NewDocument("https://rooter.jp/blog/") 
 if err != nil {
  panic("htmlの取得に失敗しました") 
 } 
 title := doc.Find("h3.article-title") 
 title.Each(func(i int, s *goquery.Selection){ 
  fmt.Print(s.Text()) 
 }) 
}
