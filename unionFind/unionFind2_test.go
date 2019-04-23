package main

import (
	"math/rand"
	"testing"
)

func BenchmarkUnionFind2(b *testing.B) {
	b.StopTimer()
	var uf2 UnionFind2
	n := 100000
	uf2.Initial(n)
	b.StartTimer()
	for i := 0; i < len(uf2.parent) ; i ++{
		a := rand.Intn(n)
		b := rand.Intn(n)
		uf2.UnionElement(a, b)
	}
	for i := 0; i < len(uf2.parent); i ++ {
		a := rand.Intn(n)
		b := rand.Intn(n)
		uf2.IsConnected(a, b)
	}
	b.StopTimer()

}