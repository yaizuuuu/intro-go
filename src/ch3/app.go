package main

import f "fmt"

// 同じパッケージ内であれば、大文字でなくても別ファイルの定義を使用することができる
func printMessage(s string) {
	f.Println(s)
}
