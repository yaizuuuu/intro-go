package main

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
	"io"
	"sync"
)

//func main() {
//	mutex = new(sync.Mutex)
//
//	for i := 0; i < 5; i++ {
//		go func() {
//			for i := 0; i < 10; i++ {
//				UpdateAndPrint(i)
//			}
//		}()
//	}
//	for {}
//}
//
//var st struct{A, B, C int}
//
//// ミューテックスを保持する
//var mutex *sync.Mutex
//
//func UpdateAndPrint(n int)  {
//	// ロック
//	// これにより処理を同期化できる
//	mutex.Lock()
//
//	st.A = n
//	time.Sleep(time.Microsecond)
//	st.B = n
//	time.Sleep(time.Microsecond)
//	st.C = n
//	time.Sleep(time.Microsecond)
//
//	fmt.Println(st.A, st.B, st.C)
//
//	// アンロック
//	mutex.Unlock()
//}

func main() {
	wg := new(sync.WaitGroup)
	wg.Add(3)

	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("1st Goroutine")
		}
		wg.Done()
	}()
	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("2nd Goroutine")
		}
		wg.Done()
	}()
	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("3rd Goroutine")
		}
		wg.Done()
	}()

	// WaitGroupのすべての終了を検知するまで待機
	wg.Wait()

	h := md5.New()
	io.WriteString(h, "ABCDE")
	fmt.Println(h.Sum(nil))

	s1 := sha1.New()
	io.WriteString(s1, "ABCDE")
	fmt.Printf("%x\n", s1.Sum(nil))
}
