package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {

	// os.Exitするとdefer破棄されてしまう
	//defer func() {
	//	fmt.Println("defer")
	//}()
	//
	//os.Exit(1)

	//_, err := os.Open("./foo")
	//if err != nil {
	//	// interface{}型の値を標準出力してos.Exit(1)を実行
	//	log.Fatal(err)
	//}

	// コマンドラインの引数を取得する
	// os.Args[0]は実行コマンドを受け取る
	//fmt.Println(os.Args[0], os.Args[1])

	fmt.Printf("length=%d\n", len(os.Args))

	for _, v := range os.Args {
		fmt.Println(v)
	}

	f, err := os.Open("./foo.txt")
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(f)
	}
	// 関数の終了時に確実にクローズできる
	defer f.Close()

	bs := make([]byte, 128)
	n, err := f.Read(bs)
	fmt.Println(n)

	bs = make([]byte, 128)
	n, err = f.ReadAt(bs, 2)
	fmt.Println(n)

	// カーソルをnバイト目にセット
	offset, err := f.Seek(2, os.SEEK_SET)
	fmt.Println(offset)

	// 現在のカーソルからnバイト目
	offset, err = f.Seek(-1, os.SEEK_CUR)
	fmt.Println(offset)

	// ファイルの末尾からnバイト目
	offset, err = f.Seek(-2, os.SEEK_END)
	fmt.Println(offset)

	fi, err := f.Stat()
	if err == nil && fi != nil {
		fmt.Println(fi)
		fmt.Println(fi.Name())
		fmt.Println(fi.Size())
		fmt.Println(fi.Mode())
		fmt.Println(fi.ModTime())
		fmt.Println(fi.IsDir())
	}

	f, _ = os.Create("bar.txt")

	fi, _ = f.Stat()
	if fi != nil {
		fmt.Println(fi.Name())
		fmt.Println(fi.Size())
		fmt.Println(fi.IsDir())
	}

	f.Write([]byte("Hello, World!\n"))
	// 7バイト目から上書き
	f.WriteAt([]byte("Golang"), 7)
	f.Seek(0, os.SEEK_END)
	f.WriteString("Yeah!")

	// OpenFile関数で様々なオプションでファイルを操作できる
	// 例
	// os.O_CREATE => ファイルが存在しなければ新規作成
	// os.O_TRUNC => 可能であればファイルの内容をオープン時に空にする
	f, err = os.OpenFile("./bar.txt", os.O_RDONLY, 0666)
	fmt.Println(f)

	// ファイルの削除
	//err = os.Remove("./bar.txt")
	//fmt.Println(err)

	// ディレクトリの削除
	// ディレクトリの中にファイルやディレクトリが残っているとエラーになる
	//err = os.Remove("./bar")
	// ディレクトリbarをまとめて削除
	//err = os.RemoveAll("./bar")

	// ファイル名・ディレクトリ変更ができる
	//err = os.Rename("bar.txt", "bar2.txt")

	// カレントディレクトリの取得
	dir, err := os.Getwd()
	fmt.Println(dir)

	// チェンジディレクトリ
	err = os.Chdir("./src")
	dir, err = os.Getwd()
	fmt.Println(dir)

	// カレントディレクトリのオープン
	err = os.Chdir("../")
	f, err = os.Open("./")
	if err != nil {
		log.Fatal(err)
	}
	fis, err := f.Readdir(0)
	for _, fi := range fis {
		if fi.IsDir() {
			fmt.Println(fi.Name())
		}
	}

	// ディレクトリの作成
	err = os.Mkdir("foo", 0775)
	// 要はmkdir -p
	err = os.MkdirAll("foo/bar/baz", 0775)
	// まとめて削除
	err = os.RemoveAll("foo")

	// OSのテンポラリディレクトのパスを取得できる
	fmt.Println(os.TempDir())

	// シンボリックリンクの作成
	err = os.Symlink("./foo.txt", "./foo2.txt")

	// ホスト名の取得
	host, err := os.Hostname()
	fmt.Println(host)

	// 環境変数一覧の取得
	for _, v := range os.Environ() {
		fmt.Println(v)
	}

	fmt.Println(os.Getenv("HOME"))
	// 存在しない環境変数であれば、空文字列が入る
	fmt.Println(os.Getenv("HOME2"))

	os.Setenv("HOME2", "/path/to/home")
	fmt.Println(os.Getenv("HOME2"))
	os.Unsetenv("HOME2")
	fmt.Println(os.Getenv("HOME2"))

	// 環境変数の存在チェックを厳密に行う
	getHome := func(envName string) {
		if home, ok := os.LookupEnv(envName); ok {
			fmt.Println(home)
		} else {
			fmt.Println("no $HOME")
		}
	}
	getHome("HOME")
	getHome("HOME2")

	// プロセスに関する情報をまとめる
	fmt.Println(os.Getpid())
	fmt.Println(os.Getppid())
	fmt.Println(os.Getuid())
	fmt.Println(os.Geteuid())
	fmt.Println(os.Getgid())
	fmt.Println(os.Getegid())

	fmt.Println(time.Now())
	// 自分で情報を入れることもできる
	//t := time.Date(2015, 7, 19, 10, 14, 23, 0, time.Local)
	t := time.Now()
	fmt.Println(t.Year())
	fmt.Println(t.YearDay())
	fmt.Println(t.Month())
	fmt.Println(t.Weekday())
	fmt.Println(t.Day())
	fmt.Println(t.Hour())
	fmt.Println(t.Minute())
	fmt.Println(t.Second())
	fmt.Println(t.Nanosecond())
	fmt.Println(t.Zone())

	// Durationも予め定義されている
	fmt.Println(time.Hour)
	fmt.Println(time.Minute)
	fmt.Println(time.Second)
	fmt.Println(time.Millisecond)
	fmt.Println(time.Microsecond)
	fmt.Println(time.Nanosecond)

	// 任意のDurationを文字列から作成できる
	d, _ := time.ParseDuration("2h30m")
	t = time.Now()
	t = t.Add(d)
	fmt.Println(t)
	// デフォルトのDurationを計算して使うこともできる
	t = t.Add(2*time.Hour + 30*time.Minute)
	fmt.Println(t)

	// 時間の差分計算もできる
	t0 := time.Date(2021, 7, 24, 0, 0, 0, 0, time.Local)
	t1 := time.Now()
	// 差分はDurationで返ってくる
	d = t0.Sub(t1)
	fmt.Println(d)

	// 時刻の比較
	fmt.Println(t1.Before(t0))
	fmt.Println(t1.After(t0))
	fmt.Println(t0.Before(t0))
	fmt.Println(t0.After(t1))

	t = time.Now()
	t0 = t.AddDate(1, 0, 0)
	fmt.Println(t0)
	t1 = t.AddDate(0, -1, 0)
	fmt.Println(t1)

	t, err = time.Parse("2006/01/02", "2020/04/24")
	fmt.Println(t)
	t, err = time.Parse("2006年1月2日15時04分05秒", "2015年11月27日14時30分29秒")
	fmt.Println(t)

	// 時刻を文字列化
	t = time.Now()
	fmt.Println(t.Format(time.RFC3339))

	// タイムゾーンを変更する
	t = time.Now()
	utc := t.UTC()
	fmt.Println(utc)

	// ローカルのタイムゾーンにする
	t = time.Now()
	localTz := t.Local()
	fmt.Println(localTz)

	// UNIX時間
	unixTime := t.Unix()
	fmt.Println(unixTime)

	// UNIX時間をDateTimeに戻すこともできる
	t = time.Unix(unixTime, 0)
	fmt.Println(t)

	// ゴルーチンの停止
	// Durationと組み合わせて使うとわかりやすい
	fmt.Println("start")
	time.Sleep(2 * time.Second)
	fmt.Println("finish")

	// time.Tickは指定した秒数の間隔でDateTimeが送信されるチャネルを生成する
	ch := time.Tick(1 * time.Second)
	counter := 0
	for  {
		// 受信
		t := <-ch
		counter += 1
		if counter > 3 {
			break
		}
		fmt.Println(t)
	}

	// 二秒後に値をチャネルに送信
	ch1 := time.After(2 * time.Second)
	fmt.Println(<-ch1)
}
