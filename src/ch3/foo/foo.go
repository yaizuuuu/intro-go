package foo

// インポートするパッケージに別名をつけることもできる
import f "fmt"
// . をつけるとパッケージ名が省略できる
import . "math"

const (
	MAX           = 100
	internalConst = 1
)

func FooFunc(n int) int {
	// Piはmathに属する変数, . をつけてインポートしたため直接使用できる
	f.Println(n, Pi)

	return n + 1
}

func internalFunc(n int) int {
	return n + 1
}
