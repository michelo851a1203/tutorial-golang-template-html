# 關於 golang template : 由 server 去使用 html 模板去產生前端頁面

### 前言
前幾天好奇如何從後端渲染 html 在前端顯示，於是稍微做了研究。

> 了解如何使用 `html/template` 的使用

`html/template` 提供了一些方法，可以把後端資料預處理帶入到我要使用的模板裡，
所以使用 template 是不錯的選擇。

以下是如何使用 template

```go
tmp, err := template.ParseFile("index.html")
if err != nil {
  log.Fatalf("template error : %s\n", err.Error())
}
tmp.Execute(<io.Writer>, <我們要帶入的資料>)
```

這裡要使用的模板為 `index.html`，但其實為必要 html 的 mime type 可以是 `template.tpl` 等副檔名
> 了解 `net/http` 的使用

官方自己有提供網頁開啟的 `net/http` 包做使用，以下我基本範例

```go
func MainHandler(responseWriter http.ResponseWriter, request *http.Request) {
  fmt.Fprintf(responseWriter, "hello world")
}

func main() {
  http.HandleFunc("/", MainHandler)
  log.Fatal(http.ListenAndServe(":8080", nil))
}
```

#### 開始實作
- 在 `terminal` 輸入以下指令
```sh
go mod init demo-template
```
可以產生`go module` 

- 建立檔案 `main.go` 和建立 `index.html` 作為我們的模板檔案
```sh
touch main.go && index.html
```
- 此時資料夾結構會長成這樣
```
.
├── go.mod
├── index.html
└── main.go
```
- 這時我們可以開始編輯 `index.html` 檔案了

```html
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta http-equiv="X-UA-Compatible" content="IE=edge">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>Document</title>
</head>
<body>
  {{ .input }}
</body>
</html>
```
`{{ .input }}` 作為我們模板用的參數帶入使用
- 編輯 `main.go`
```go

package main

import (
	"log"
	"net/http"
	"text/template"
)

func MainHandler(responseWriter http.ResponseWriter, _ *http.Request) {
	myTemplate, err := template.ParseFiles("index.html")
	if err != nil {
		log.Fatalf("template error : %s", err.Error())
	}
	myTemplate.Execute(responseWriter, map[string]interface{}{
		"input": "testing hello",
	})
}

func main() {
  fmt.Printnl("serve on http://localhost:8080")
	http.HandleFunc("/", MainHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

```
- 這時候去 `terminal` 輸入 `go run main.go`
```sh
go run main.go
```
這時候回應為，就大功告成了
```sh
serve on http://localhost:8080
```
這時我們去瀏覽器去輸入網址 `http://localhost:8080` 就會有預想的結果了。
