package main

import (
	"fmt"
	"time"
)

func main() {
	// チャネルはFIFOという性質を持つ
	// キューの概念に近いデータ構造
	// 送受信可能
	var ch chan int
	// 受信専用のチャネル
	var ch1 <-chan int
	// 送信専用のチャネル
	var ch2 chan<- int

	// まとめて書くと以下
	//var (
	//	ch0 chan int
	//	ch1 <-chan int
	//	ch2 chan<- int
	//)

	// コンパイルが通る
	// 送受信可能であれば、受信用、送信用のチャネルに代入できる
	ch1 = ch
	ch2 = ch
	// コンパイルエラー
	//ch = ch1
	//ch = ch2
	//ch1 = ch2
	//ch2 = ch1
	fmt.Println(ch1, ch2)

	// 第2引数でバッファサイズを指定できる
	ch3 := make(chan int, 10)
	// 送信
	ch3 <- 5
	// 受信
	i := <-ch3
	fmt.Println(i)
	// ランタイムパニック
	// 他のゴルーチンから値が送信されずにチャネルが眠ってしまっている、デッドロックであると認識される
	//fmt.Println(<-ch3)

	ch4 := make(chan int)
	go receiver(ch4)
	i2 := 0
	for i2 < 1000 {
		ch4 <- i2
		i2++
	}

	ch5 := make(chan rune, 3)
	ch5 <- 'A'
	ch5 <- 'B'
	ch5 <- 'C'
	// デッドロック
	//ch5 <- 'D'
	// 下記のような状態でデッドロックが発生する
	// バッファ内が空のチャネルから受信 (variable <- ch)
	// バッファ内に空きがないチャネルへの送信 (ch <- something)

	// lenでチャネルにある要素数を調べる事ができる
	fmt.Println(len(ch5))

	// スライス同様capで容量(バッファ)を調べることができる
	ch6 := make(chan string, 6)
	fmt.Println(cap(ch6))
	ch6 <- "A"
	ch6 <- "B"
	ch6 <- "C"
	close(ch6)
	// ランタイムパニック
	// クローズしたチャネルに対して送信を行うとエラーになる
	//ch6 <- "D"

	// チャネルがクローズしている状態で受信をしてもランタイムパニックにはならない
	s, ok := <-ch6
	fmt.Println(s, ok)
	s, ok = <-ch6
	fmt.Println(s, ok)
	s, ok = <-ch6
	fmt.Println(s, ok)
	// チャネルがクローズしている場合、空になった状態のチャネルから受信してもランタイムパニックが発生しない
	// ゆえに↑のclose関数をコメントアウトすると以下の処理はランタイムパニックが発生してしまう
	// チャネルが**クローズ**かつ**空**のとき2つめの変数がfalseとなる
	s, ok = <-ch6
	fmt.Println(s, ok)
	s, ok = <-ch6
	fmt.Println(s, ok)

	ch7 := make(chan int, 20)
	go receive2("1st goroutine", ch7)
	go receive2("2nd goroutine", ch7)
	go receive2("3rd goroutine", ch7)

	i3 := 0
	for i3 < 100 {
		ch7 <- i3
		i3++
	}
	// 値が送信し終わったらcloseする
	close(ch7)
	// 通常はすべてゴルーチンが終了したことを検知するべきだが、ここでは簡易的にスリープ
	// ゴルーチンの同期化?? TODO: 調べる
	time.Sleep(3 * time.Second)

	// チャネル型をfor文で回すこともできるが、チャネルのクローズを検出するタイミングが得られないというデメリットがある
	// チャネルからひたすら受信し続ける常駐バッチのようなものであれば使うかもしれない??
	//ch8 := make(chan int)
	//ch8 <- 1
	//ch8 <- 2
	//ch8 <- 3
	//
	//for i := range ch8 {
	//	fmt.Println(i)
	//}

	ch11 := make(chan int)
	ch12 := make(chan int)
	ch13 := make(chan int)

	go func() {
		for {
			i := <-ch11
			ch12 <- i * 2
		}
	}()

	go func() {
		for {
			i := <-ch12
			ch13 <- i - 1
		}
	}()

	n := 1

LOOP:
	for {
		select {
		case ch11 <- n:
			n++
		case i := <-ch13:
			fmt.Println("received", i)
		default:
			fmt.Println("why", n)
			if n > 100 {
				break LOOP
			}
		}

	}
}

// 受信専用チャネル
func receiver(ch <-chan int) {
	for {
		// 他のゴルーチンから送信された値を受信する
		i := <-ch
		fmt.Println(i)
	}
}

func receive2(name string, ch <-chan int) {
	for {
		i, ok := <-ch
		// 受信できなくなったら終了 (クローズかつ空)
		if ok == false {
			break
		}

		fmt.Println(name, i)
	}

}
