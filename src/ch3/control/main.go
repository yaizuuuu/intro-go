package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	x := 1

	if x == 1 {
		fmt.Println("true")
	}

	// コンパイルエラー
	// 多言語で見られるようなキャストによる真偽判定はサポートされていない
	//b := 1
	//if b {
	//	fmt.Println("true")
	//}

	if x, y := 1, 3; x < y {
		fmt.Printf("x(%d) is less than y(%d)\n", x, y)
	}

	// if文内で定義された変数はスコープがことなるため、別々の名前が一致しても別々の変数として定義できる
	x2 := 5
	if x2 := 2; true {
		fmt.Println(x2)
	}
	fmt.Println(x2)

	x3, y3 := 3, 5
	if n := x3 * y3; n%2 == 0 {
		fmt.Printf("n(%d) is even\n", n)
	} else {
		fmt.Printf("n(%d) is odd\n", n)
	}

	// 例外処理でこのように定義することが頻出
	// errはif文内で定義しているため変数のスコープ外に影響を及ぼすことがない
	//if _,err := doSomething; err != nil {
	//	// 何らかの例外処理
	//}

	// for文で定義を省略すると無限ループとなる
	i := 0
	for {
		fmt.Println(i)
		i++
		if i == 100 {
			break
		}
	}

	i2 := 0
	for i2 < 100 {
		fmt.Println(i2)
		i2++
	}

	for i := 0; i < 100; i++ {
		if i%2 == 1 {
			continue
		}

		fmt.Println(i)
		i++
	}

	fruits := [3]string{"Apple", "Banana", "Cherry"}

	// forで配列の中身にアクセスできる
	// rangeは予約語、配列等のindexを返す
	for i, s := range fruits {
		fmt.Printf("fruits[%d]=%s\n", i, s)
	}

	// 文字列をforで回すとルーン型が返ってくる
	for i, r := range "ABC" {
		fmt.Printf("[%d] -> %d\n", i, r)
	}

	// 文字列の場合rangeで返ってくるのは何番目の文字というものではない
	// UTF-8でエンコードされた文字列のコードポイントが返ってくる
	// 何バイト目の文字かということになるので、配列等と大きく異なるが要注意
	for i, r := range "あいうえお" {
		fmt.Printf("[%d] -> %d\n", i, r)
	}

	// 多言語とは異なりbreakを定義しなくても次のcaseが実行されないのがポイント
	n := 3
	switch n {
	case 1, 2:
		fmt.Println("1 ot 2")
	case 3, 4:
		fmt.Println("3 or 4")
	default:
		fmt.Println("unknown")
	}

	// 多言語のようなbreakしない場合は次のcaseを実行したい場合はfallthroughを指定することで次のcaseを実行できる
	s := "A"
	switch s {
	case "A":
		s += "B"
		fallthrough
	case "B":
		s += "C"
		fallthrough
	case "C":
		s += "D"
		fallthrough
	default:
		s += "E"
	}
	fmt.Println(s)

	// コンパイルエラー
	// 型が異なるものをcaseに指定するとコンパイルエラーになる
	//s2:= 1
	//switch s2 {
	//case 1:
	//fmt.Println(s2)
	//case "2":
	//	fmt.Println("two")
	//}

	// switch文もifと同じように変数が定義できるため、関数の返り値を受け取ってswitch文の実行などができる
	switch n := 2; n {
	case 1, 3, 5, 7, 9:
		fmt.Printf("%d is odd\n", n)
	case 2, 4, 6, 8:
		fmt.Printf("%d is even\n", n)
	}

	// switch文で条件式を書くこともできる
	// その際はswitch文の宣言の後に変数をしていしない
	n2 := 4
	switch {
	case n2 > 0 && n2 < 3:
		fmt.Println("0 < n < 3")
	case n2 > 3 && 6 > n2:
		fmt.Println("3 < n < 6")
	}

	// interface型を使用することですべての型を受け取れる
	anything := func(a interface{}) {
		fmt.Println(a)
	}
	anything(1)
	anything("海")

	// 型アサーション
	var x4 interface{} = 3
	i3 := x4.(int)
	fmt.Println(i3)
	// 型アサーションに失敗している場合、プログラムエラーになってしまう
	//f := x4.(float32)

	// 型アサーションをする際に変数を2つにして受け取ることで、プログラムエラーにならない
	// 型アサーションに失敗した場合、2つ目の変数はfalseを受け取る
	var x5 interface{} = 3.14
	i4, isInt := x5.(int)
	fmt.Println(i4, isInt)
	f, isFloat := x5.(float64)
	// 型アサーションに成功するときのみ1つめの変数に値を受け取ることができる、失敗した場合は型の初期値が入る(0, "", false)
	fmt.Println(f, isFloat)
	// 型アサーションによる真偽判定のみを行うことが目的であれば、1つめの変数は_を使う
	_, isString := x5.(string)
	fmt.Println(isString)

	// if文で型アサーションを使うことができる
	var x6 interface{} = 3.14
	if x6 == nil {
		fmt.Println("x is nil")
	} else if i, isInt := x6.(int); isInt {
		fmt.Printf("x is integer : %d\n", i)
	} else if s, isString := x6.(string); isString {
		fmt.Println(s)
	} else {
		fmt.Println("unsupported type!")
	}

	// 変数名.(type)で型を受け取ることができる
	// JSでいうtypeofに近い感じ??
	switch x6.(type) {
	case bool:
		fmt.Println("bool")
	case int, uint:
		fmt.Println("integer or unsigned integer")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("don't know")
	}

	// switch文で型アサーションをしながらも、変数で実際の値を受け取ることができる
	// v, isInt := x6.(int)に近い
	switch v := x6.(type) {
	case bool:
		fmt.Println("bool: ", v)
	case int:
		fmt.Println(v * v)
	case float64:
		fmt.Println(v * v)
		// コンパイルエラー
		// 下記の場合はどちらも演算できる方ではあるが、
		// string,intとなった場合演算できないものを列挙される可能性があるので、
		// 複数型を列挙すると演算はできない
		//case int64, int32:
		//	fmt.Println(v * v)
	}

	// gotoで処理をスキップさせることができる
	// 関数内でジャンプするもので、関数外にジャンプはできない
	fmt.Println("A")
	goto L
	// 処理されない
	// 変数定義をジャンプしようとするとコンパイルエラーになる
	//n3 := 1
	fmt.Println("B")
L:
	fmt.Println("C")
	// fot文などの制御構文の中にジャンプすることはできない
	//for {
	//	L:
	//		fmt.Println("Hello")
	//}

	// 制御構文に対してラベルをつけることができる
LOOP:
	for {
		for {
			for {
				fmt.Println("開始")
				// 普通のbreakであれば階層が上のfor文を終了させることはできないが、
				// ラベル名を指定することでまとめて終了させることができる
				break LOOP
			}
		}
	}
	fmt.Println("完了")

	// 上と同様でラベルが付いているfor文をまとめてcontinueできる
M:
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			if j > 1 {
				continue M
			}
			fmt.Printf("%d * %d = %d\n", i, j, i*j)
		}
	}

	// deferがついている行は関数の終了時に実行することを強制できる
	runDefer := func() {
		defer fmt.Println("defer")
		fmt.Println("done")
	}
	runDefer()
	// os.Openのようなファイルを読み書きする処理において、リソースの開放漏れを防ぐと行った場面において有効
	// file, err := os.Open("/path/to/file")
	// if err != nil {
	//   return
	// }
	// defer file.Close

	// defer func() {
	//
	// }()
	// 即時関数にもdeferは有効

	runDefer = func() {
		// LIFOになるため、定義した順に実行されないことは要注意
		defer fmt.Println(1)
		defer fmt.Println(2)
		defer fmt.Println(3)
		fmt.Println("done")
	}
	runDefer()

	// panicを実行してもdeferは実行される
	defer fmt.Println("Hello World!")

	// recoverはdeferと組み合わせて使うのが原則
	// panicを補足するとpanicの引数で受け取った値を補足する(!= nil)
	defer func() {
		if x := recover(); x != nil {
			fmt.Println(x)
		}
	}()

	testRecover(3.14)

	// ランタイムエラーを起こす
	// shellでいうexit 1~255みたいなもの
	// 強力な機能なためよほどのことがないと実装しない
	// 実装しているプロダクトも多くはないらしい
	panic("runtime error!")
	// こっちは実行されない
	fmt.Println("Hello World!!...")


}

func testRecover(src interface{}) {
	defer func() {
		if x := recover(); x != nil {
			switch v := x.(type) {
			case int:
				fmt.Printf("panic: int=%v\n", v)
			case string:
				fmt.Printf("panic: string=%v\n", v)
			default:
				fmt.Println("panic: unknown")
			}
		}

	}()

	panic(src)
	return
}
