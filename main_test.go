package main

import (
	"testing"
	"time"
)

type A struct {
	X    int
	Laen int
}

func (a A) Len() int {
	return 1
}

func NewBig() A {
	time.Sleep(1 * time.Second)
	return A{1, 2}
}
func BenchmarkBigLen(b *testing.B) {
	// if benchmark needs some expensive setup, the
	// timer may be reset
	big := NewBig()
	//reseting timer
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		big.Len()
	}
}

func BenchmarkTemplateParallel(b *testing.B) {
	big := NewBig()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			big.Len()
		}
	})
}

// for covering
// $go test -cover
// go tool can generate coverage profile thay may be interpreted by the cover tool.
// $go test -coverprofile=cover.out
// $go tool cover -func=cover.out
