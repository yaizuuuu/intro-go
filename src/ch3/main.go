package main

import (
	"./foo"
	"fmt"
)

func main() {
	// Printlnは改行付きで標準出力する
	fmt.Println("Hello, Golang(改行あり)")

	// %vは様々な型のデータを埋め込むことができる
	fmt.Printf("数値=%v 文字列=%v 配列=%v(改行なし)", 5, "Golang", [...]int{1, 2, 3})

	fmt.Print("標準エラー出力(改行なし)")

	// varで変数の明示的な型定義をする
	//var n int

	// 複数の変数の型定義もできる
	//var x, y, z int

	// ()で囲うことで型の違う複数の変数の定義ができる
	//var (
	//	x, y int
	//	name string
	//)

	// = で代入できる
	var n int
	n = 5
	fmt.Println(n)
	// 異なる型は代入できない
	//n = "string"

	// 複数の変数に対する代入もできる
	var x, y int
	x, y = 1, 2
	fmt.Println(x, y)
	// 左辺と右辺の数が異なる場合はコンパイルエラー
	//x, y = 1, 2, 3

	// 暗黙的に型定義をする場合 `:=` を使用する
	i := 1
	fmt.Println(i)
	// 再度暗黙的な定義をするとコンパイルエラー
	//i := 2
	// varでも同じ
	//var i int
	// 通常の代入であればコンパイルは通る
	i = 2
	fmt.Println(i)

	// var()で囲って複数の変数を暗黙的な型付けを行うこともできる
	// `:=` で一つ一つ定義するより目立ち可読性が上がるというメリットがある
	var (
		n2 = 1
		s  = "string"
		b  = true
	)
	fmt.Println(n2, s, b)

	var (
		n3 int
		n4 int64
	)
	n3 = 1
	// コンパイルエラーになる, 同じintでも厳密に型チェックされ、暗黙的な型変換は行うことができない
	//n4 = n3
	n4 = 1
	fmt.Println(n3, n4)

	// 明示的な型変換であれば可能
	n5 := uint(17)
	b2 := byte(n5) // uint8 == byte
	i64 := int64(n5)
	u32 := uint32(n5)
	fmt.Println(n5, b2, i64, u32)

	// コンパイルエラー, byte = uint8 は0 ~ 255までを代入できる
	//b3 := byte(256)
	// こちらは成功する
	n6 := 256
	b3 := byte(n6)
	// オーバーフローにより0になる
	// 直接定義するとコンパイルエラーになるが、
	// 明示的な型変換や演算を通すとコンパイルエラーにはならないがオーバーフローが発生する
	fmt.Println(b3) // 0

	s2 := `Goの
Raw文字列リテラルによる
複数行に渡る
文字列
\nは使えないので注意`
	fmt.Println(s2)

	a := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("%v\n", a)
	fmt.Printf("%v\n", a[0])
	fmt.Printf("%v\n", a[1])
	fmt.Printf("%v\n", a[2])
	fmt.Printf("%v\n", a[3])
	fmt.Printf("%v\n", a[4])
	// コンパイルエラー
	//fmt.Printf("%v\n", a[5])

	// コンパイルエラー
	//a2 := [5]int{1, 2, 3, 4, 5, 6}

	// 配列型をvarで初期化したとき指定した型の初期値が入る
	// int = 0, string = "", bool = false
	var a2 [5]int // [0,0,0,0,0]
	fmt.Printf("%v\n", a2)

	// ...で要素数を指定するのを省略できる
	a3 := [...]int{1, 2, 3}
	fmt.Printf("%v\n", a3)

	// コンパイルエラー
	// 厳密な型チェックが行われる
	//var (
	//	a4 [3]int
	//	a5 [5]int
	//)
	//a4 = a5

	// 配列は値渡しになる
	a4 := [3]int{1, 2, 3}
	a5 := [3]int{4, 5, 6}
	a4 = a5
	a4[0] = 0
	a4[2] = 0
	fmt.Printf("%v\n", a4)

	// interface型はすべての型と互換性がある
	var x2 interface{}
	x2 = 1
	x2 = "文字列"
	fmt.Printf("%v\n", x2)
	var x3, y3 interface{}
	x3, y3 = 1, 2
	fmt.Printf("%v, %v\n", x3, y3)
	// コンパイルエラー
	// すべての型と互換性をもたせる代わりに演算はできなくなる
	//z := x + y

	// 複数返り値がある関数から値を受け取る場合
	q, r := div(19, 7)
	fmt.Println(q, r)
	// 片方の返り値を破棄する場合は_を使って受け取ることで破棄できる
	_, r2 := div(19, 7)
	fmt.Println(r2)

	fmt.Println(doSomething())

	f := returnFunc()
	f()

	f2 := callFunction(func() {
		fmt.Println("I'm function 2")
	})
	f2()

	f3 := later()
	i1 := f3("first")
	i2 := f3("second")
	i3 := f3("third")
	fmt.Println(i1, i2, i3)

	f4 := integers()
	i4 := f4()
	i5 := f4()
	i6 := f4()
	fmt.Println(i4, i5, i6)

	f5 := integers()
	i7 := funcName(f5)
	fmt.Println(i7)

	const ONE = 1
	one := func() (int, int) {
		const TWO = 2
		return ONE, TWO
	}
	fmt.Println(one())

	// var同様複数定義できる
	//const (
	//	X = 1
	//	Y = 2
	//)

	// 複数の定数を定義する際、値が一緒であれば省略もできる
	//const (
	//	X = 1
	//	Y // == 1
	//	Z // == 1
	//	S1 = "あ"
	//	S2 // == "あ"
	//)

	// 何も定義しない場合はコンパイルエラー
	//const (
	//	X
	//	Y
	//	Z
	//)

	const (
		X = 2
		Y = 7
		Z = X + Y

		S1 = "今日"
		S2 = "晴れ"
		S  = S1 + "は" + S2
	)

	fmt.Println(Z, S)

	// iotaというconst定義で増える度にインクリメントされるものを使用してenumに近いものを実装できる
	// constブロックごとにリセットされる
	const (
		A = iota // == 0
		B        // == 1
		C        // == 2
		D = 100
		E = iota // == 4
	)
	fmt.Println(A, B, C, D, E)

	// 最初の一文字目が大文字のもののみ外部に公開ができる
	fmt.Println(foo.FooFunc(foo.MAX))

	printMessage("別ファイルからの呼び出しだよ〜")

	s3 := doSomething2("no use")
	fmt.Println(s3)
}

func funcName(f5 func() int) int {
	i7 := f5()
	return i7
}

// 戻り値のない関数には何も型定義をしない `void` 使わない
func hello() {
	fmt.Println("hello, World!!")
}

// 関数の戻り値を複数にできる, 複数ある場合は()で囲い、それぞれの型を定義する
func div(a, b int) (int, int) {
	q := a / b
	r := a % b

	return q, r
}

// goに例外機構がないため、このようにハンドリングを行うことがほとんど
//result, err := doSomething()
//if (err != nil) {
//	// エラー処理
//}

func doSomething() (a int) {
	return
}

// 上記は以下と同義になる
//func doSomething() int {
//	var a int
//	return a
//}

// 複数の返り値の場合でも同じことができる
//func doSomething() (x, y int) {
//	y = 5
//	return // x == 0, y == 5
//}

func returnFunc() func() {
	// 無名関数を返り値として指定できる
	return func() {
		fmt.Println("I'm a function")
	}
}

// 関数を引数とすることもできる
func callFunction(f func()) func() {
	return f
}

func later() func(string) string {
	var store string

	return func(next string) string {
		s := store
		store = next
		return s
	}
}

func integers() func() int {
	i := 0

	return func() int {
		i += 1
		return i
	}
}

func doSomething2(a string) (b string) {
	var block string
	// 定義済みの変数名と同じ名前を使いたい場合はブロックを定義することで強いようできる
	{
		var a string
		const b = "string"

		a = "test"
		block = a + b
	}

	return block
}