package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	var (
		max int
		msg string
		x bool
	)

	// オプションの定義
	// 代入する変数, オプション名, デフォルト値, description
	flag.IntVar(&max, "n", 32, "処理数の最大数")
	flag.StringVar(&msg, "m", "", "処理メッセージ")
	flag.BoolVar(&x, "x", false, "拡張オプション")

	// オプション読み取り
	// 定義していないオプションを指定するとエラーをはいてくれる
	flag.Parse()

	fmt.Println("処理数の最大値 = ", max)
	fmt.Println("処理メッセージ = ", msg)
	fmt.Println("拡張オプション = ", x)

	// 5桁にまるめて右詰め
	fmt.Printf("%5d\n", 123)
	// 5桁に丸めて左詰め
	fmt.Printf("%-5d\n", 123)
	// 小数点以下2桁に丸める
	fmt.Printf("%.2f\n", 1.4141356)
	// 小数点以下2桁にして全体を8桁に丸めて右詰め
	fmt.Printf("%8.2f\n", 1.4141356)
	// 10文字にして右詰め
	fmt.Printf("%10s\n", "Go言語")
	// 変数の型を表示
	fmt.Printf("%T\n", []string{"A", "B", "C"})
	// あらゆる型を表示
	fmt.Printf("%v\n", 3.123)

	type User struct {
		Id int
		Email string
	}
	u := User{Id : 123, Email: "test@test.com"}
	fmt.Printf("%v\n", u)
	// 構造体のフィールド名も表示する
	fmt.Printf("%+v\n", u)
	// 構造体のフィールド名と型も表示する
	fmt.Printf("%#v\n", u)

	// ログを標準出力にはきだす
	log.Print("ログの1行目\n")
	log.Println("ログの2行目")
	log.Printf("ログの%d行目\n", 3)

	f, err := os.Create("./test.log")
	if err != nil {
		return
	}
	log.SetOutput(f)
	log.Println("ログのメッセージ")

	// ログのはきだし先を標準出力にする
	log.SetOutput(os.Stdout)
	// デフォルトのLogFormat
	log.SetFlags(log.LstdFlags)
	log.Println("Default Log Format")

	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)
	log.Println("Add MicroSecond")

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Println("Add File Name and Line")

	log.SetPrefix("[LOG] ")
	log.Println("Add Prefix")

	// ログ出力をして、os.Exit(1)
	//log.Fatal()
	// ログ出力をして、ランタイムパニックを発生
	//log.Panic()

	// ログの出力先を細かく分けるなどの制御を行いたいときには新しいロガーを生成する
	logger := log.New(os.Stdout, "[LOG] ", log.LstdFlags | log.Lshortfile)
	logger.Println("デフォルト以外で作成したロガー")
}