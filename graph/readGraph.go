package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type Graph struct {
	fileName string
	sparseGraph SparseGraph
	denseGraph DenseGraph

}
func (r *Graph) ReadGraph(fileName string, tp string,directed bool) {
	switch tp {
	case "denseGraph" :
		//TODO: read file as a denseGraph

	case "sparseGraph" :
		//r.sparseGraph.Initial(10,false)
		//fmt.Println(r.sparseGraph.g.value)
		 r.buildSparseGraph(fileName,directed)
	}
}

func (r *Graph) buildSparseGraph(fileName string,dir bool)  {
	//var sparseGraph SparseGraph
	fi, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	br := bufio.NewReader(fi)
	file := make([][]int,0)
	index := make([]int,0)
	count := 0
	for {
		line, _, err := br.ReadLine()

		if err == io.EOF {
			break
		}
		st := string(line)
		r := strings.NewReader(st)
		var b int
		var c []int
		for {
			_, err := fmt.Fscan(r,&b)
			if err == io.EOF {
				break
			}
			if err != nil {
				panic(err)
			}
			c = append(c, b)
		}

		if count == 0 {
			index = append(index,c...)
		} else {
			file = append(file,c)
		}
		count ++
	}
	//fmt.Println("initial:  ", index[0])
	r.sparseGraph.Initial(index[0],dir)
	for i := 0; i < index[0]; i ++ {
		r.sparseGraph.AddEdge(file[i][0],file[i][1])
	}
	//fmt.Println("value",r.sparseGraph.g.value)

}
