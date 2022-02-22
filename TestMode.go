package MyModule

import "testing"

type ForTest int

func TestUnit(t *testing.T, index ForTest, function func()) {
	for i := 0; i < int(index); i++ {
		function()
	}
}

func TestBench(b *testing.B, function func()) {
	for i := 0; i < b.N; i++ {
		function()
	}
}
