package main

import (
    "fmt"
    "os"
    "github.com/sclevine/agouti"
)

func main() {
    // chrome起動
    // driver := agouti.ChromeDriver(agouti.Browser("chrome"))
    driver := agouti.ChromeDriver(
        agouti.ChromeOptions("args", []string{
            "--headless",
            "--window-size=1200,1200",
            "--blink-settings=imagesEnabled=false", // don't load images
            "--disable-gpu",                        // ref: https://developers.google.com/web/updates/2017/04/headless-chrome#cli
            "no-sandbox",                           // ref: https://github.com/theintern/intern/issues/878
            "disable-dev-shm-usage",                // ref: https://qiita.com/yoshi10321/items/8b7e6ed2c2c15c3344c6
        }),
    )

    fmt.Fprintf(os.Stderr, "%s\n", "start=====")

    if err := driver.Start(); err != nil {
        fmt.Fprintf(os.Stderr, "%s\n", err)
        return
    }
    defer driver.Stop()

    page, err := driver.NewPage()
    if err != nil {
        fmt.Fprintf(os.Stderr, "%s\n", err)
        return
    }

    // // ID, Passの要素を取得し、値を設定
    //  identity := page.FindByID("user_login")//ユーザー名欄のhtmlのidの部分を入力
    //  password := page.FindByID("user_pass")//パスワード欄のhtmlのidの部分を入力
    //  identity.Fill("ログイン時のユーザー名")
    //  password.Fill("ログイン時のパスワード")
    //  // formをサブミット
    //  if err := page.FindByClass("button-large").Submit(); err != nil {
    //      log.Fatalf("Failed to login:%v", err)
    //  }

    // googleにアクセス
    if err := page.Navigate("https://www.google.com/"); err != nil {
        fmt.Fprintf(os.Stderr, "%s\n", err)
        return
    }

    // qというnameを持つ要素を取得する
    q := page.FindByName("q")
    // 検索文字列を入力
    q.Fill("mebee")

    // 検索を実行
    if err := q.Submit(); err != nil {
        fmt.Fprintf(os.Stderr, "%s\n", err)
        return
    }

    if err := page.Screenshot("./shot.jpg"); err != nil {
        fmt.Fprintf(os.Stderr, "%s\n", err)
    }

  //   doc.Find("div.classname > div > div > a.item").Each(func(i int, s *\
  // goquery.Selection){
  //    title := s.Text()
  //    url, exists := s.Attr("href")
  //    if exists == true{
  //      fmt.Printf("Title: %s\n", title)
  //      fmt.Printf("URL: %s\n", url)
  //    }
  //  })
    // 掲載イベントURL一覧を取得
    doc.Find("#rso div").Each(func(i int, s *goquery.Selection) {
        href, _ := s.Attr("href")
        burl, _ := url.Parse(sc_url)

        //相対パスから絶対URLに変換
        var full_url = toAbsUrl(burl, href)
        println("ページ内リンクURL:" + full_url)

    })


    fmt.Fprintf(os.Stderr, "%s\n", "end=====")

}