package main



type UnionFind struct {
	id []int
	count int
}
func (u *UnionFind) Initial(n int) {
	for i := 0; i < n ; i ++ {
		u.id = append(u.id,i)
	}
	u.count = n
}
func (u *UnionFind)Find(p int) interface{} {
	if p <0 || p >len(u.id) {
		return "invalid p"
	}
	return u.id[p]
}
func (u *UnionFind) IsConnected(p, q int) bool {
	return u.Find(p) == u.Find(q)
}
func (u *UnionFind) UnionElement(p, q int) {
/*
	go func() {
		pID := u.id[p]
		qID := u.id[q]
		if pID == qID {
			return
		}
		for i := 0; i < u.count; i ++ {
			if u.id[i] == qID {
				u.id[i] = pID
				break
			}

		}
		return
	}()
*/


	pID := u.id[p]
	qID := u.id[q]
	if pID == qID {
		return
	}
	for i := 0; i < u.count; i ++ {
		if u.id[i] == qID {
			u.id[i] = pID
			break
		}

	}
	return


}
/*
func main() {
	var uf UnionFind

	uf.Initial(10)

	fmt.Println(uf.Find(8))
	fmt.Println(uf.IsConnected(2,3))
	fmt.Println(uf.id[2],uf.id[3])
	uf.UnionElement(2,3)
	fmt.Println(uf.id[2],uf.id[3])
	fmt.Println(uf.IsConnected(2,3))
	time.Sleep(1*time.Second)

}
*/