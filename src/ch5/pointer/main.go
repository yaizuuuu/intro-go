package main

import "fmt"

func main() {
	// 型の前に*をつけることでポインタ型を定義できる
	var p *int
	fmt.Println(p == nil)

	// &を使って任意の型からポインタ型を生成することができる
	var i int
	p2 := &i
	fmt.Printf("%T\n", p2)
	// ポインタ型からポインタ型を取得もできる
	pp2 := &p2
	fmt.Printf("%T\n", pp2)

	var i2 int
	p3 := &i2
	i2 = 5
	// ポインタ型の前に*を置くことでデリファレンスができる
	fmt.Println(*p3)
	// デリファレンスからの代入も可能
	*p3 = 10
	fmt.Println(*p3)

	i3 := 1
	// ポインタ型を生成して渡すことで参照渡しが可能になる
	inc(&i3)
	inc(&i3)
	inc(&i3)
	fmt.Println(i3)

	p4 := &[3]int{1, 2, 3}
	pow(p4)
	// len, cap, スライス式の場合もデリファレンスを省略できる
	fmt.Println(p4, len(p4), cap(p4), p4[:])
	// for文でrangeを使用する場合もデリファレンスを省略できる
	for i, v := range p4 {
		fmt.Println(i, v)
	}

	// ポインタがnilの場合にデリファレンスしようとするとランタイムパニック
	var p5 *int
	fmt.Println(p5)
	//fmt.Println(*p5)

	i4 := 5
	ip4 := &i4
	fmt.Printf("type=%T, address=%p\n", ip4, ip4)

	// 配列の要素に対してデリファレンスをしても正常に動作する
	ia := [3]int{1, 2, 3}
	iap := &ia[1]
	*iap = 9
	fmt.Println(ia)
	// スライスの要素に対してもデリファレンスをしても正常に動作する
	fs := []float64{1.1, 2.2, 3.3}
	fsp := &fs[2]
	*fsp = 9.9
	fmt.Println(fs)

	// 文字列の部分参照はコンパイルエラーが発生する
	// Goにおいて文字列型は不変であるため、参照渡しでメモリ上の値を操作できないようになっている
	//s := "ABC"
	//sp := &s
	//fmt.Println(sp[0])
	// 文字列型は不変であるため、文字列結合を行うと新しいメモリ領域に値をコピーしつつ結合をおこなうため効率が悪いプログラムになる
	s := ""
	for _, v := range []string{"A", "B", "C"} {
		s += v
	}
	fmt.Println(s)
	// このような場合はbytesパッケージで用意された処理で対応するとよい TODO: bytesパッケージによる文字列結合を調べる


}

// ポインタ型を受け取る
func inc(p *int) {
	// デリファレンスしつつ加算
	*p++
}

// ポインタ型を受け取る
func pow(p *[3]int) {
	i := 0
	for i < 3 {
		// 配列型のポインタをデリファレンスしてインデックスにアクセスする場合は()で囲う必要がある
		// C由来の書き方
		// Goではp[i]という書き方でOK, デリファレンスは省略できる
		(*p)[i] = (*p)[i] * (*p)[i]
		i++
	}
}
