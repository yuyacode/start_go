package bar

import (
	"testing"
)

// 名前がTestから始まる関数はテスト用関数であることを示す
func TestIsOne(t *testing.T) {
	n := 1
	b := IsOne(n)
	if b != true {
		t.Errorf("%d is not one", n)
	}
}

// go test bar

// 成功
// ok bar 0.002s

// 失敗
// 0 is not one
// FAIL bar 0.002s
