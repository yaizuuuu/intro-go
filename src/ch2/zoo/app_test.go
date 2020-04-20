package main

import (
	"testing"
)

func TestAppName(t *testing.T) {
	expect := "Zoo Application"
	actual := AppName()

	if expect != actual {
		t.Errorf("%s =! %s", expect, actual)
	}
}

// `go test` は拡張子が.go, ファイル名が `*_test.go` という条件にマッチしたものに対してテスト実行する