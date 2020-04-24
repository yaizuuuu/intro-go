package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

func main() {
	// 構造体から文字列にする
	u := new(User)
	u.Id = 1
	u.Name = "山田太郎"
	u.Email = "test@test.com"
	u.Created = time.Now()

	bs, err := json.Marshal(u)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(bs))

	// 文字列から構造体にする
	src := `
{
  "Id": 12,
  "Name": "田中花子",
  "Email": "test+1@test.com",
  "Created": "2016-12-02T10:00:00.000000000+09:00"
}
`
	u = new(User)
	err = json.Unmarshal([]byte(src), u)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v\n", u)

	// URLをパースする
	r, err := url.Parse("https://www.example.com/search?a=1&b=test#top")
	if err != err {
		log.Fatal(err)
	}
	fmt.Println(r.Scheme)
	fmt.Println(r.Host)
	fmt.Println(r.Path)
	fmt.Println(r.RawQuery)
	fmt.Println(r.Fragment)

	// URLを生成する
	u2 := &url.URL{}
	u2.Scheme = "https"
	u2.Host = "google.com"
	query := u2.Query()
	query.Set("q", "Go言語")
	u2.RawQuery = query.Encode()
	fmt.Println(u2)

	// GETメソッドでuにアクセス
	res, err := http.Get("https://www.google.com/")
	if err != nil {
		log.Fatal(err)
	}
	// レスポンスボディは別途読み込みが必要
	fmt.Println(res)
	fmt.Println(res.StatusCode)
	fmt.Println(res.Header["Date"])
	fmt.Println(res.Request.Method)

	// レスポンスボディを読み込んだ後はCloseする
	defer res.Body.Close()

	// ioutilでレスポンスボディの読み込みを行う
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	// []byteで返ってくる
	fmt.Println(string(body))

	// フォームをPOSTする
	vs := url.Values{}
	vs.Add("id", "1")
	vs.Add("message", "メッセージ")
	fmt.Println(vs.Encode())
	res, err = http.PostForm("https://example.com/comments/post", vs)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(res)
	}

	//// ファイルをPOSTする
	//f, err := os.Open("foo.jpg")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//res, err = http.Post("https://example.com/upload", "image/jpeg", f)
	//if err != nil {
	//	log.Fatal(err)
	//} else {
	//	fmt.Println(res)
	//}

	// webサーバを起動できる
	http.HandleFunc("/info", infoHandler)
	http.ListenAndServe(":8080", nil)

}

type User struct {
	Id      int
	Name    string
	Email   string
	Created time.Time
}

func infoHandler(w http.ResponseWriter, r *http.Request)  {
	io.WriteString(w, `
<!DOCTYPE html>
<html lang="ja">
<head>
<meta charset="UTF-8">
<title>インフォメーション</title>
<body>
<h1>ようこそ！</h1>
</body>
</html>
`)
}