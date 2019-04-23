package main

type UnionFind3 struct {
	parent []int
	count int
	sz []int
}
func (u *UnionFind3) Initial(n int) {
	u.count = n
	for i := 0; i < n; i ++ {
		u.parent = append(u.parent,i)
		u.sz = append(u.sz, 1)
	}
}

func (u *UnionFind3) Find(p int) (root,sz int){
	if p <0 || p > u.count {
		panic("invalid")
	}
	var size int
	size  = 1
	for i := 0; i < u.count; i ++ {
		if u.parent[p] == p {

			break
		} else {
			p = u.parent[p]
		}
		size ++

	}
	return p, size
}
func (u *UnionFind3) IsConnected(p, q int) bool {
	Proot, _ := u.Find(p)
	Qroot, _ := u.Find(q)
	return Proot == Qroot
}
func (u *UnionFind3) UnionElement(p, q int) {
	pRoot, pSize := u.Find(p)
	qRoot, qSize := u.Find(q)
	if pRoot == qRoot {
		return
	}
	if pSize < qSize  {
		u.parent[pRoot] = qRoot
		u.sz[qRoot] += u.sz[pRoot]
	} else {
		u.parent[qRoot] = pRoot
		u.sz[pRoot] += u.sz[qRoot]
	}



}

/*
func main() {
	var uf3 UnionFind3
	uf3.Initial(10)
	fmt.Println(uf3.parent)
	fmt.Println(uf3.IsConnected(2,3))
	uf3.UnionElement(2,3)
	fmt.Println(uf3.IsConnected(2,3))
	fmt.Println(uf3.parent)

	fmt.Println(uf3.Find(2))
	fmt.Println(uf3.Find(3))


}
*/