package main

import "fmt"

func main() {
	// スライス、可変長配列
	// var a [10]int に近い
	s := make([]int, 10)
	fmt.Println(s)

	s2 := make([]float64, 3)
	s2[0] = 3.14
	// ランタイムパニック
	//s2[4] = 5.1
	fmt.Println(s2)
	// len関数で動的に変化する要素数を調べることができる
	fmt.Println(len(s2))

	// make関数は2つめの引数で要素数の確保、3つめの引数で容量を確保する
	// 3つめは指定しなければ、要素数分だけ容量が確保される
	// 要素数の確保で実際の初期値が入る、容量の確保で初期値は入らないがメモリ上に領域が確保される
	// スライスは確保されていた容量を超えると別のメモリ領域に値をコピーして、容量を超えた分の値を追加するため、比較的コストが高い処理になる
	// 予めスライスに入る容量が分かる場合はなるべく指定したほうが良い
	// 容量を超えて値を入れられないわけではない
	s3 := make([]int, 5, 10)
	fmt.Println(s3) // 標準出力しても容量は確認できない
	fmt.Println(cap(s3))

	// make関数を使わずにスライスを定義することもできる
	// 容量は指定できない
	// 配列に近い
	s4 := []int{1, 2, 3, 4, 5}
	fmt.Println(len(s4), cap(s4))

	// 配列またはスライスからスライスを生成することができる
	a := [5]int{1, 2, 3, 4, 5}
	// [n:m]という指定でnからm-1までの範囲で値を抜き出すことができる
	// m-1というところが要注意
	// [:4] == 0から3のインデックスまで抜き出す
	// [2:] == 2から最後まで
	// [:] == すべて
	// [len(a)-2:] == 後ろの2つを抜き出す
	s5 := a[0:2]
	fmt.Println(s5)

	// 配列と異なりappendで拡張できるのがsliceの最大の特徴
	s6 := []int{1, 2, 3}
	s6 = append(s6, 4)
	fmt.Println(s6)
	// 複数の値を一度に入れることもできる
	s6 = append(s6, 5, 6, 7)
	fmt.Println(s6)

	s7 := []int{1, 2, 3}
	s8 := []int{4, 5, 6}
	// JSのスプレット演算子のような書き方もできる
	s7 = append(s7, s8...)
	fmt.Println(s7, s8)

	// 容量を超えると倍々で容量が確保されていく
	// 環境によって拡張幅が変わるので、常に倍々とは限らない
	s9 := make([]int, 0, 0)
	fmt.Printf("len=%d, cap=%d\n", len(s9), cap(s9))

	s9 = append(s9, 1)
	fmt.Printf("len=%d, cap=%d\n", len(s9), cap(s9))

	s9 = append(s9, []int{2, 3, 4}...)
	fmt.Printf("len=%d, cap=%d\n", len(s9), cap(s9))

	s9 = append(s9, 5)
	fmt.Printf("len=%d, cap=%d\n", len(s9), cap(s9))

	s9 = append(s9, 6, 7, 8, 9, 0)
	fmt.Printf("len=%d, cap=%d\n", len(s9), cap(s9))

	s10 := []int{1, 2, 3, 4, 5}
	s11 := []int{10, 11, 12, 13, 14, 15, 16}
	// 1つめの引数に対して塗りつぶすようにコピーする
	// 拡張まではされない
	n := copy(s10, s11)
	fmt.Println(n, s10, s11)

	a2 := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// 容量を指定しないと最後インデックス分容量を確保
	s12 := a2[2:4]
	fmt.Println(s12, len(s12), cap(s12))

	// 要素数と同じ分だけの容量を確保する場合
	s13 := a2[2:4:4]
	fmt.Println(s13, len(s13), cap(s13))

	// 容量を要素数より多く調整する場合
	s14 := a2[2:4:6]
	fmt.Println(s14, len(s14), cap(s14))

	// 配列と同じようにループを回せる
	s15 := []string{"Apple", "Banana", "Cherry"}
	for i, v := range s15 {
		fmt.Printf("[%d] => %s\n", i, v)

		// rangeでループを回すとappendした値はループの評価式に加味されなくなる
		// このループで"Melon"は登場しない
		// for文の中で動的に変化するスライスを扱いたいのであれば、for i:=0; i < len(s15); len++ 等の別の方法を取ることが必要
		// この例でそれをやると無限ループなので気をつけること
		s15 = append(s15, "Melon")
	}
	fmt.Println(s15)

	fmt.Println(sum(1, 2, 3, 4, 5, 6))

	// 配列は値渡しのためもとの変数に変化なし
	a3 := [3]int{1, 2, 3}
	powA(a3)
	fmt.Println(a3)

	// スライス型は参照渡しのためもとの変数に変化あり
	s16 := []int{1, 2, 3}
	powS(s16)
	fmt.Println(s16)

	var (
		a4 [3]int
		s17 []int
	)
	// 初期値も配列とことなるので要注意
	fmt.Println(a4, s17, s17 == nil)

	a5 := [3]int{1, 2, 3}
	s18 := a5[:]
	fmt.Println(a5, s18)

	// スライスも値が変わる、この時点ではメモリを共有している
	s18[0] = 9
	// こちらでも結果は同じ、append等をしないとメモリは共有されたまま
	//a5[0] = 9
	fmt.Println(a5, s18)

	// appendなどをするとコピーが行われ、別領域になるため値はそれぞれ別々になる
	s18 = append(s18, 4)
	a5[1] = 10
	fmt.Println(a5, s18)
}

// ...で可変長引数を定義でき、それはスライス型となる
func sum(s ...int) int {
	n := 0
	for _, v := range s {
		n += v
	}

	return n
}

func powA(a [3]int)  {
	for i, v := range a {
		a[i] = v * v
	}

	return
}

func powS(s []int)  {
	for i, v := range s {
		s[i] = v * v
	}

	return
}