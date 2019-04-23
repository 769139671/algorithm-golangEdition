package main

import (
	"math/rand"
	"testing"
)

func BenchmarkUnionFind5(b *testing.B) {
	b.StopTimer()
	var uf5 UnionFind5
	n := 1000000
	uf5.Initial(n)
	b.StartTimer()
	for i := 0; i < len(uf5.parent) ; i ++{
		a := rand.Intn(n)
		b := rand.Intn(n)
		uf5.UnionElement(a, b)
	}
	for i := 0; i < len(uf5.parent); i ++ {
		a := rand.Intn(n)
		b := rand.Intn(n)
		uf5.IsConnected(a, b)
	}
	b.StopTimer()
}
