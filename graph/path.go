package main

type path struct {
	graph Graph
	s int
	visited []int
	from []int
	//pathCount int
	rePath []int
	path []int
}

func (gr *Graph) HasPath(s, w int) bool {

	var path path
	path.graph.sparseGraph = gr.sparseGraph
	path.initial(s)
	if path.graph.IsConnected(s, w) == true {
		return true
	} else {
		return false
	}
}
func (gr *Graph) Path(s, w int) {
	var path path
	path.graph.sparseGraph = gr.sparseGraph
	if !path.assert(s, w) {
		panic("out of range")
	}
	path.initial(s)
	path.dfs(s, w)

	var cur int
	cur = w
	path.rePath = append(path.rePath,w)
	//fmt.Println("")
	for {

		if path.from[cur] != s {
			path.rePath = append(path.rePath, path.from[cur])
			cur = path.from[cur]
			continue
		} else {
			path.rePath = append(path.rePath,s)
			break
		}
	}

	for i := len(path.rePath) - 1; i >= 0; i -- {
		path.path = append(path.path, path.rePath[i])
	}


}
func (p *path) assert(s, w int) bool {
	//fmt.Println("p.graph.sparseGraph.n:",p.graph.sparseGraph.n)

	if s < 0 || s >= p.graph.sparseGraph.n || w < 0 || w >= p.graph.sparseGraph.n {
		return false
	} else {
		return true
	}

}
func (p *path) initial(s int) {
    p.s = s
    //p.pathCount = 0
    for i := 0; i < p.graph.sparseGraph.n; i ++ {
    	p.visited = append(p.visited, 0)
    	p.from = append(p.from,-1)
	}

}
func (p *path) dfs(cur int, target int) {
	p.visited[cur] = 1
	var arr []int
	for _, val := range p.graph.sparseGraph.g.value[cur] {
		//发现目标w
		arr = append(arr, val)
	}


	for i := 0; i < len(arr); i ++ {
		if arr[i] == target {

			p.visited[target] = 1

			p.from[target] = cur
			return
		} else {
			if p.visited[arr[i]] == 1 {
				continue
			} else {
				p.from[arr[i]] = cur
				p.dfs(arr[i],target)
			}
		}
	}
}
