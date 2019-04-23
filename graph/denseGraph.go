package main

import (

	"fmt"
)
type DenseGraph struct {
	n, m int
	directed bool
	g vector
}
type vector struct {
	vector [][]int
	value []bool
}
func (d *DenseGraph) Initial(n int, directed bool) {

	d.directed = directed
	d.n = n
	d.m = 0
	for i := 0; i < n; i ++ {
		for j := 0; j < n; j ++ {
			d.g.vector = append(d.g.vector,[]int{i, j})
			d.g.value = append(d.g.value, false)

		}
	}
}
func (d *DenseGraph) Assert(v, w int) bool{
	if v < 0 || v >= d.n || w < 0 || w >= d.n {
		return false
	}
	return true
}
func (d *DenseGraph) hasEdge(v, w int) bool{
	if d.Assert(v,w) == false {
		panic("index out of range")
	}
	var x int
	for i, val := range d.g.vector {
		if val[0] == v && val[1] == w {
			fmt.Println("find")
			x = i
			break
		}
	}
	if d.g.value[x] == true {
		return true
	}
	return false
}
func (d *DenseGraph) AddEdge(v, w int) {
	var x int
	if d.Assert(v, w) == false {
		fmt.Println("v, w out of rang")
		return
	}
	for i, val := range d.g.vector {
		//fmt.Println("i",i)
		//fmt.Println("val:", val)
		tx := val[0]
		ty := val[1]
		//fmt.Println("tx", tx)
		//fmt.Println("ty", ty)
		if tx == v && ty == w {
			fmt.Println("found",val)
			x = i
			break
		}
	}
	if d.hasEdge(v, w) == true {
		fmt.Println("already has edge")
		return
	} else {
		if !d.directed {
			d.g.value[d.n*d.n - x - 1] = true
			//d.m ++

		}
		d.g.value[x] = true
		d.m ++
	}
}
//返回节点个数
func (d *DenseGraph) V() int {
	return d.n
}
//返回边的数目
func (d *DenseGraph) E() int {
	return d.m
}
/*
/*
func main() {
	var a DenseGraph
	a.Initial(5,false)
	fmt.Println(a)
	a.AddEdge(1,2)
	a.AddEdge(1,2)
	a.AddEdge(3,4)
	fmt.Println(a.g.value)
	fmt.Println("has edge",a.hasEdge(0,1))
	fmt.Println("has edge",a.hasEdge(1,2))
	fmt.Println("edge:", a.m)

	
}
*/