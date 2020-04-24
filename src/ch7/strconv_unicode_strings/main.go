package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	b := true
	s := strconv.FormatBool(b)
	// bool型をstring型に変換
	fmt.Printf("[%T]%v\n", s, s)

	// int型を10進数でstring型に変換
	s = strconv.FormatInt(-1234, 10)
	fmt.Printf("[%T]%v\n", s, s)

	// string型をbool型に変換
	b, _ = strconv.ParseBool("true")
	b, _ = strconv.ParseBool("1")
	b, _ = strconv.ParseBool("t")
	b, _ = strconv.ParseBool("T")
	b, _ = strconv.ParseBool("True")
	b, _ = strconv.ParseBool("TRUE")
	fmt.Printf("[%T]%v\n", b, b)

	// 変換できない文字列を受け取ると2つめの変数にエラーとして値を返す
	_, err := strconv.ParseBool("foo")
	fmt.Println(err)

	// 文字列からint型に変換, 10進数で表現
	i, _ := strconv.ParseInt("12345", 10, 0)
	fmt.Printf("[%T]%v\n", i, i)

	// ルーン型の引数が数字かどうか判定する
	fmt.Println(unicode.IsDigit('X'))
	fmt.Println(unicode.IsDigit('2'))
	// ルーン型の引数が文字かどうか判定する
	fmt.Println(unicode.IsLetter('A'))
	fmt.Println(unicode.IsLetter('3'))
	// ルーン型の引数がスペースかどうか判定する
	fmt.Println(unicode.IsSpace(' '))
	fmt.Println(unicode.IsSpace('\t'))
	fmt.Println(unicode.IsSpace('　'))

	// 文字列を結合する
	fmt.Println(strings.Join([]string{"A", "B", "C"}, ", "))

	// 指定した文字の位置をインデックスで取得する
	fmt.Println(strings.Index("ABCDE", "C"))
	// 存在しない場合は-1が返す
	fmt.Println(strings.Index("ABCDE", "Z"))
	// 複数存在する場合は最後のインデックスを返す
	// Index()は最初のインデックスを返す
	fmt.Println(strings.LastIndex("ABCDEABCDE", "ABC"))

	// プレフィックスが一致するかどうか
	fmt.Println(strings.HasPrefix("Go言語", "Go"))
	// サフィックスが一致するかどうか
	fmt.Println(strings.HasSuffix("Go言語", "言語"))

	// 指定した文字列が含まれるかどうか
	fmt.Println(strings.Contains("ABCDE", "AB"))
	fmt.Println(strings.Contains("ABCDE", "X"))
	// 空文字列はtrue判定になる
	fmt.Println(strings.Contains("ABCDE", ""))

	// 指定した文字列のいずれかは含まれるか
	fmt.Println(strings.ContainsAny("ABCDE", "AZ"))

	// 指定した文字列が何回登場するか
	fmt.Println(strings.Count("ABCDEABCDE", "A"))

	// 文字列を繰り返して結合する
	fmt.Println(strings.Repeat("ABC", 3))

	// 文字列の置換, 最後の数字は何回置き換えるかを指定する, -1が指定された場合はすべて置換する
	fmt.Println(strings.Replace("AAAAA", "A", "Z", 1))
	fmt.Println(strings.Replace("AAAAA", "A", "Z", -1))

	// 文字列の分割
	s1 := strings.Split("A,B,C,D,E", ",")
	fmt.Printf("[%T]%s\n", s1, s1)

	// 小文字に変換
	fmt.Println(strings.ToLower("ABCDE"))
	// 大文字に変換
	fmt.Println(strings.ToUpper("abcde"))

	// 文字列からスペースを取り除く
	fmt.Println(strings.TrimSpace("  - Hello, World! -  "))
	// 全角スペースや改行やタブコードも取り除かれる
	fmt.Println(strings.TrimSpace("　　- Hello, World! -\t\n\r"))

	// 1つ以上のスペースを区切りに分割する
	fmt.Println(strings.Fields("a b c"))

}
