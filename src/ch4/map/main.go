package main

import (
	"fmt"
)

func main() {
	// mapはいわゆる連想配列
	m := make(map[int]string)

	m[1] = "US"
	m[81] = "Japan"
	m[86] = "China"

	fmt.Println(m)

	m2 := make(map[string]string)
	m2["Yamada"] = "Taro"
	m2["Sato"] = "Hanako"
	m2["Yamada"] = "Jiro"
	fmt.Println(m2)

	// make関数を使わずリテラルで定義するとしたら以下
	m3 := map[int]string{1: "Taro", 2: "Hanako", 3: "Jiro"}
	fmt.Println(m3)

	// map型の中にスライス型を入れるなど複雑な定義も可能
	m4 := map[int][]int{
		1: []int{1},
		2: []int{1, 2},
		3: []int{1, 2, 3},
	}
	fmt.Println(m4)

	// マップの中のスライスのリテラルは省略することができる
	m5 := map[int][]int{
		1: {1},
		2: {1, 2},
		3: {1, 2, 3},
	}
	fmt.Println(m5)

	// マップの中にマップを書いたりと更に複雑な定義もできる
	m6 := map[int]map[float64]string{
		1: {3.14: "円周率"},
	}
	fmt.Println(m6)

	// map型は配列、スライス型のように存在しないキーにアクセスしてもランタイムパニックにならず、
	// 存在しないキーにアクセスした場合は型で定義される初期値が取り出せる
	m7 := map[int]string{1: "A", 2: "B", 3: "C"}
	fmt.Println(m7[9] == "")

	// この例だと0が取り出せるため、間違えやすい
	m8 := map[int]int{1:0}
	fmt.Println(m8[9]) // == 0

	m9:=map[int]string{1:"A", 2:"B", 3:"C"}
	// 存在している場合は2つめの変数にtrueが入る
	// okがよく使われるため、特別な理由がない限りokという変数名を使う
	s, ok := m[1]
	fmt.Println(m9, s, ok)
	// 存在していない場合はfalse
	s, ok = m[9]
	fmt.Println(m9, s, ok)
	// このようにして存在せず初期値が取れたのか、もともとの値が取れたのかを安全に制御できる
	if _, ok := m[1]; ok {
		fmt.Println("exist!")
	}

	// スライスや配列と違い順番は保証されないため注意が必要
	m10:=map[int]string{
		1: "Apple",
		2: "Banana",
		3: "Cherry",
	}
	for k, v := range m10 {
		fmt.Printf("%d => %s\n", k, v)
	}

	fmt.Println(m10, len(m10))
	m10[4] = "Melon"
	m10[5] = "Lemon"
	fmt.Println(m10, len(m10))

	// キーを指定して要素の削除を行うこともできる
	delete(m10, 5)
	fmt.Println(m10)

	// スライスのcapとは意味合いが異なるが、それに近い形
	// メモリ領域を確保してくれる
	// cap関数では確認できない
	// 大量のmapを保持する場合パフォーマンスの向上につながるかもしれない
	//m11 := make(map[int]string, 100)
}
