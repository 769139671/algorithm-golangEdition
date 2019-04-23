package main

import (
	"math/rand"
	"testing"
)

func BenchmarkUnionFind4(b *testing.B) {
	b.StopTimer()
	var uf4 UnionFind4
	n := 1000000
	uf4.Initial(n)
	b.StartTimer()
	for i := 0; i < len(uf4.parent) ; i ++{
		a := rand.Intn(n)
		b := rand.Intn(n)
		uf4.UnionElement(a, b)
	}
	for i := 0; i < len(uf4.parent); i ++ {
		a := rand.Intn(n)
		b := rand.Intn(n)
		uf4.IsConnected(a, b)
	}
	b.StopTimer()

}

