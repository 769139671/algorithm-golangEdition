package main

import "fmt"

type UnionFind4 struct {
	parent []int
	rank []int
	count int
}
func (u *UnionFind4) Initial(n int) {
	u.count = n
	for i := 0; i < n; i ++ {
		u.parent = append(u.parent, i)
		u.rank = append(u.rank,1)
	}
}
func (u *UnionFind4) Find(p int) (root,ran int){
	if p <0 || p >u.count {
		panic("invalid")
	}
	var rank int
	rank  = 1
	for i := 0; i < u.count; i ++ {
		if u.parent[p] == p {

			break
		} else {
			p = u.parent[p]
		}
		rank ++

	}
	return p, rank
}
func (u *UnionFind4) UnionElement(p, q int) {
	pRoot, pRank := u.Find(p)
	qRoot, qRank := u.Find(q)
	if pRoot == qRoot {
		return
	}
	if pRank < qRank  {
		u.parent[pRoot] = qRoot

	} else {
		if pRank > qRank {
			u.parent[qRoot] = pRoot
		} else {
			u.parent[pRoot] = qRoot
			u.rank[qRoot] +=1
		}
	}
}
func (u *UnionFind4) IsConnected(p, q int) bool {
	Proot, _ := u.Find(p)
	Qroot, _ := u.Find(q)
	fmt.Println("proot",Proot)
	fmt.Println("qroot", Qroot)
	return Proot == Qroot
}
/*
func main() {
	var uf4 UnionFind4
	uf4.Initial(10)
	fmt.Println(uf4.rank)
	uf4.UnionElement(2,3)
	fmt.Println(uf4.parent)
	fmt.Println(uf4.IsConnected(2,3))



}
*/