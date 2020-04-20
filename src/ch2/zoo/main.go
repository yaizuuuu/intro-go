package main

import (
	// ディレクトリの相対的なパスを記述してパッケージを読み込むこともできる
	// そのため、パッケージ名とディレクトリ名をあわせるのがよい
	"./animals"
	"fmt"
)

func main()  {
	fmt.Println(AppName())

	fmt.Println(animals.ElephantFeed())
	fmt.Println(animals.MonkeyFeed())
	fmt.Println(animals.RabbitFeed())
}

// このファイルがあるディレクトリで `go build` を実行すると `zoo` という実行ファイルが出来上がる
// 特に指定をしない場合はカレントディレクトリと同じ名前の実行ファイルが出来上がる
// また、カレントディレクトリ内のすべてのgoファイルをコンパイルしてくれる
// このファイルをコンパイルすると内包しているanimalsパッケージのファイルもコンパイルされる