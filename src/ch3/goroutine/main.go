package main

import (
	"fmt"
	"runtime"
)

// mainの前に実行できる
// 引数などは指摘できず、指定するとコンパイルエラーになる
func init()  {
	fmt.Println("init")
}

// パッケージ内にinitは何回も定義できる
// 定義した順番に実行される
// 特別な事情がなければ推奨されない書き方である
func init()  {
	fmt.Println("init 2")
}

func main() {
	// ゴルーチンによる非同期処理
	//sub := func() {
	//	for {
	//		fmt.Println("sub loop")
	//	}
	//}
	//go sub()
	//for {
	//	fmt.Println("main loop")
	//}

	go fmt.Println("Yeah!")

	fmt.Printf("NumCPU: %d\n", runtime.NumCPU())
	// 使用しているgoroutineの数を表せる
	fmt.Printf("NumGoroutine: %d\n", runtime.NumGoroutine())
	fmt.Printf("Version: %s\n", runtime.Version())
}
