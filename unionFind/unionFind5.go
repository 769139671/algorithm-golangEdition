package main

//path compression
type UnionFind5 struct {
	parent []int
	rank []int
	count int
}
func (u *UnionFind5) Initial(n int) {
	u.count = n
	for i := 0; i < n; i ++ {
		u.parent = append(u.parent, i)
		u.rank = append(u.rank,1)
	}
}
func (u *UnionFind5) Find(p int) (root,sz int){
	if p <0 || p >u.count {
		panic("invalid")
	}
	var rank int
	rank  = 1
	for i := 0; i < u.count; i ++ {
		if u.parent[p] == p {

			break
		} else {
			//path compression
			u.parent[p] = u.parent[u.parent[p]]
			p = u.parent[p]
		}
		rank ++

	}
	return p, rank
}
func (u *UnionFind5) UnionElement(p, q int) {
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
func (u *UnionFind5) IsConnected(p, q int) bool {
	Proot, _ := u.Find(p)
	Qroot, _ := u.Find(q)
	return Proot == Qroot
}
/*
func main() {
	var uf5 UnionFind5
	uf5.Initial(10)
	fmt.Println(uf5.rank)
	uf5.UnionElement(2,3)
	fmt.Println(uf5.rank)
	fmt.Println(uf5.IsConnected(2,3))


}
*/
