

```
go get github.com/PuerkitoBio/goquery
go get github.com/sclevine/agouti
brew upgrade chromedriver
```


https://sites.google.com/a/chromium.org/chromedriver/downloads

```
/usr/local/bin/chromedriver
に配置。

“chromedriver”は、開発元を検証できないため開けません。
-> 「システム環境設定」->「セキュリティとプライバシー」で許可する
```


chromeコンソール


```
//$x('.//*[@class="item-name"]')
$x('.//*[@class="item-name"]')[0].textContent

$x('.//*[@class="owl-lazy"]')[0]
$x('.//*[@class="owl-lazy"]')[0].src

$x('.//*[@class="owl-lazy"]')[1]
〜４まである

$x('.//*[@class="owl-stage"]')


.//*[@value=“Google Search”]

.//*[@name=“btnK”]

.//*[@jsaction=“sf.chk”]

.//*[@class=“jsb”]/center/input[1]

.//*[@style=“padding-top:18px”]/center/input[1]

.//*[@class=“tsf-p”]/div[3]/center/input[1]

.//*[@class=“tsf”]/div[2]/div[3]/center/input[1]
```


```
str_select := "input[type='radio'][name='bf'][value='off']"
item := page.Find(str_select)
item.Click()

str_select = "input[type='checkbox'][id='plan_b']"
item = page.Find(str_select)
item.Click()
```


