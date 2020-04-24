package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {
	//// 標準入力を読み取るスキャナ
	//scanner := bufio.NewScanner(os.Stdin)
	//
	//for scanner.Scan() {
	//	fmt.Println(scanner.Text())
	//}
	//
	//if err := scanner.Err(); err != nil {
	//	fmt.Println(os.Stderr, "読み込みエラー: ", err)
	//}

	s := `XXXXX
YYYYY
ZZZZZ`

	// 文字列から読み取り ↑だとos.Stdinから読み取る
	r := strings.NewReader(s)
	scanner := bufio.NewScanner(r)
	// 一行ずつScan
	scanner.Scan()
	fmt.Println(scanner.Text())
	scanner.Scan()
	fmt.Println(scanner.Text())
	scanner.Scan()
	fmt.Println(scanner.Text())

	s = `ABC DEF
GHI JKL MNO
PQR STU VWX YZ
`
	r = strings.NewReader(s)
	scanner = bufio.NewScanner(r)
	// スペースごとのScan方法に変更(デフォルトは一行ずつ)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	currentDir, _ := os.Getwd()
	f, err := ioutil.TempFile(currentDir, "foo")
	if err != nil {
		log.Fatal(err)
	}
	f.WriteString("Hello, World!\n")
	fmt.Println(f.Name())

	// 正規表現のパターンを登録
	re := regexp.MustCompile(`bc`)
	fmt.Println(re.MatchString("abcdefg"))

	// +?, *?で最短マッチ, +, *は最長マッチ
	re = regexp.MustCompile(`A+?A+?X`)
	fmt.Println(re.MatchString("AAX"))

	re = regexp.MustCompile(`^[0-9A-Za-z]{3}`)
	fmt.Println(re.MatchString("ABC"))
	fmt.Println(re.MatchString("   "))

	// [^XXXX]で否定
	re = regexp.MustCompile(`[^0-9A-Za-z]{3}`)
	fmt.Println(re.MatchString("ABC"))
	fmt.Println(re.MatchString("   "))

	// 正規表現にマッチした文字列の取得
	re = regexp.MustCompile(`\w+`)
	// 一番最初に該当したものを取得
	fmt.Println(re.FindString("abc XYZ 123"))
	// 指定した数だけ取得
	fmt.Println(re.FindAllString("abc XYZ 123", 2))
	// -1ですべてを取得
	fmt.Println(re.FindAllString("abc XYZ 123", -1))

	re = regexp.MustCompile(`\s+`)
	// 先頭から指定した数だけ分割
	fmt.Println(re.Split("A B C D\nE", 3))
	// すべてを分割
	fmt.Println(re.Split("A B C D\nE", -1))

	// 文字列の置換
	re = regexp.MustCompile(`佐藤`)
	fmt.Println(re.ReplaceAllString("佐藤さんと鈴木さん", "田中"))

	// グループによるサブマッチ
	re = regexp.MustCompile(`(\d+)-(\d+)-(\d+)`)
	s = `
00-111-2222
3333-44-55555
666-777-888
0-0-0
`
	// スライスでマッチした文字列とグループでサブマッチした文字列が取得できる
	ms := re.FindAllStringSubmatch(s, -1)
	for _, s := range ms {
		fmt.Println(s)
	}

	// サブマッチを$nで取得してReplaceに活用することもできる
	re = regexp.MustCompile(`(\d+)-(\d+)-(\d+)`)
	fmt.Println(re.ReplaceAllString("Tel: 0120-3333-9066", "$2-1234-$1"))
}
