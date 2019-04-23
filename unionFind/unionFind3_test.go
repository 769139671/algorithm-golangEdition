package main

import (
	"math/rand"
	"testing"
)

func BenchmarkUnionFind3(b *testing.B) {
	b.StopTimer()
	var uf3 UnionFind3
	n := 1000000
	uf3.Initial(n)
	b.StartTimer()
	for i := 0; i < len(uf3.parent) ; i ++{
		a := rand.Intn(n)
		b := rand.Intn(n)
		uf3.UnionElement(a, b)
	}
	for i := 0; i < len(uf3.parent); i ++ {
		a := rand.Intn(n)
		b := rand.Intn(n)
		uf3.IsConnected(a, b)
	}
	b.StopTimer()

}