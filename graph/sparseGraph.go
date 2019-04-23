package main

import "fmt"

type SparseGraph struct {
	n, m int
	directed bool
	g vector2
}
type vector2 struct {
	vector []int
	value [][]int
}

func (s *SparseGraph) Initial(n int, dir bool) {
	s.directed = dir
	//fmt.Println("n :",n)

	s.n = n
	//fmt.Println("s.n: ",s.n)
	s.m = 0
	for i := 0; i < n; i ++ {
		s.g.vector = append(s.g.vector,i)
		s.g.value = append(s.g.value,[]int{})
	}
}
func (s *SparseGraph) assertRange(v, w int) bool{
	if v < 0 || v >= s.n || w < 0 || w >= s.n {
		return false
	}
	return true
}
func (s *SparseGraph) hasEdge(v, w int) bool{
	if s.assertRange(v, w) == false {
		panic("out of range")
	}
	for _, val := range s.g.value[v] {
		if val == w {
			return true
		}
	}

	return false
}
func (s *SparseGraph) AddEdge(v, w int) {
	if !s.assertRange(v, w) {
		fmt.Println("v and w out of range")
		return
	}

	//自环边
	if v == w {
		//fmt.Println("==")
		if !s.hasEdge(v, w) {
			s.g.value[v] = append(s.g.value[v], v)
			s.m ++
		}
		return
	}

	//非自环边
	//false 无向图
	if !s.directed {
		if !s.hasEdge(v,w) {
			//fmt.Println("v,w")
			s.g.value[v] = append(s.g.value[v],w)
			s.m ++
		}
		if !s.hasEdge(w, v) {
			//fmt.Println("w,v")
			s.g.value[w] = append(s.g.value[w],v)
			s.m ++
		}
	}
	//true 有向图
	if s.directed {
		if !s.hasEdge(v,w) {
			s.g.value[v] = append(s.g.value[v],w)
			s.m ++
		}
	}


}

