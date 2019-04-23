package main

import (
	"math/rand"
	"testing"
)

func BenchmarkUnionFind_IsConnected(b *testing.B) {
	b.StopTimer()
	n := 100000
	var uf1 UnionFind
	uf1.Initial(n)
	b.StartTimer()
	for i := 0; i < n; i ++ {
		a := rand.Intn(n)
		b := rand.Intn(n)
		uf1.UnionElement(a, b)
	}
	for i := 0; i < n; i ++ {
		a := rand.Intn(n)
		b := rand.Intn(n)
		uf1.IsConnected(a, b)
	}
	b.StopTimer()
}
