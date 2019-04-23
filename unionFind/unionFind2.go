package main

type UnionFind2 struct {
	parent []int
	count int
}
func (u *UnionFind2) Initial(n int) {
	u.count = n
	for i := 0; i < n; i ++ {
		u.parent = append(u.parent,i)
	}
}
//find the element p of root
func (u *UnionFind2) Find(p int) int{
	if p <0 || p > u.count {
		panic("invalid")
	}
	for i := 0; i < u.count; i ++ {
		//found the root of p
		if u.parent[p] == p {
			break
		} else {
			//p = the parent of p
			p = u.parent[p]
		}

	}
	return p
}
func (u *UnionFind2) IsConnected(p, q int) bool {
	return u.Find(p) == u.Find(q)
}
func (u *UnionFind2) UnionElement(p, q int) {
	//find the root of p
	pRoot := u.Find(p)
	//find the root of q
	qRoot := u.Find(q)
	if pRoot == qRoot {
		return
	}
	//change the root of p as q's root
	u.parent[pRoot] = qRoot

}
/*
func main() {
	var uf2 UnionFind2
	uf2.Initial(10)
	fmt.Println(uf2.parent)
	fmt.Println(uf2.IsConnected(2,3))
	uf2.UnionElement(2,3)
	fmt.Println(uf2.IsConnected(2,3))
	fmt.Println(uf2.parent)


}
*/