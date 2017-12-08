package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type node struct {
	name   string
	weight int
	nodes  []*node
	sum    int
	dad    *node
}

func newNode(name string, w int, nodes []*node) *node {
	return &node{
		name:   name,
		weight: w,
		sum:    w,
		nodes:  nodes,
	}
}

func (n *node) refreshSum() {
	n.sum = n.weight
	for _, v := range n.nodes {
		n.sum += v.sum
	}
	if n.dad != nil {
		n.dad.refreshSum()
	}
}

func (n *node) String() string {
	return fmt.Sprintf("%v %v %v %v", n.name, n.weight, n.nodes, n.sum)
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var last *node
	nodes := make(map[string]*node)

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.Replace(line, "(", "", -1)
		line = strings.Replace(line, ")", "", -1)
		line = strings.Replace(line, "->", "", -1)
		line = strings.Replace(line, ",", "", -1)
		scan := bufio.NewScanner(strings.NewReader(line))
		scan.Split(bufio.ScanWords)

		scan.Scan()
		n := scan.Text()
		scan.Scan()
		w, _ := strconv.Atoi(scan.Text())
		r := make([]*node, 0)
		for scan.Scan() {
			name := scan.Text()
			son, found := nodes[name]
			if found {
				r = append(r, son)
			} else {
				son = &node{name: name}
				nodes[name] = son
				r = append(r, son)
			}
		}

		nod, found := nodes[n]
		if found {
			nod.weight = w
			nod.nodes = r
		} else {
			nodes[n] = newNode(n, w, r)
			last = nodes[n]
		}

		nod = nodes[n]
		for _, v := range nod.nodes {
			v.dad = nod
		}
		nod.refreshSum()
	}

	for last.dad != nil {
		last = last.dad
	}

	print(last)
	nmin, nmax := last.nodes[0], last.nodes[0]
	for _, v := range last.nodes {
		if v.sum < nmin.sum {
			nmin = v
		}
		if v.sum > nmax.sum {
			nmax = v
		}
	}

	// uownj
	fmt.Printf("last : %v\n", last.name)
	fmt.Printf("min: %v, max: %v, w:%v\n", nmin.sum, nmax.sum, last.weight)
	fmt.Printf("il faut enlever %v à %v => %v\n", (nmax.sum - nmin.sum), nmax.weight, (nmax.weight - (nmax.sum - nmin.sum)))

}

func print(nod *node) {
	//if len(nod.nodes) > 0 {
	fmt.Printf("[%v] ", nod.name)
	for _, v := range nod.nodes {
		fmt.Printf("n:%v s:%v w:%v ", v.name, v.sum, v.weight)
	}
	fmt.Println("")
	//for _, v := range nod.nodes {
	//	print(v)
	//}
	//}
}
