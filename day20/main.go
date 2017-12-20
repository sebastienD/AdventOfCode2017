package main

import (
	"fmt"
	"os"
	"bufio"
	"log"
	"strings"
	"unicode"
	"strconv"
)

type ByManhattan []*particule

func (a ByManhattan) Len() int           { return len(a) }
func (a ByManhattan) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByManhattan) Less(i, j int) bool { return manhattan(a[i]) < manhattan(a[j]) }

type coord struct {
	x,y,z int
}

func (c *coord) equalsTo(o *coord) bool {
	return c.x == o.x && c.y == o.y && c.z == o.z
}

func (c *coord)String() string  {
	return fmt.Sprintf("(x:%v, y:%v, z:%v )", c.x, c.y, c.z)
}

type particule struct {
	p, v, a *coord
	m int64
}

func (p *particule)String() string  {
	return fmt.Sprintf("(p:%v, v:%v, a:%v )", p.p,p.v,p.a)
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Open file failed")
	}
	defer file.Close()

	particles := make([]*particule, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		f := func(c rune) bool {
			return !unicode.IsDigit(c) && c != '-'
		}
		parts := strings.FieldsFunc(scanner.Text(), f)
		partsInt := make([]int, len(parts))
		for i,v := range parts {
			partsInt[i],err = strconv.Atoi(v)
			if err != nil {
				log.Fatalf("Error with parse %v", v)
			}
		}
		particle := &particule{
			&coord{partsInt[0], partsInt[1], partsInt[2]},
			&coord{partsInt[3], partsInt[4], partsInt[5]},
			&coord{partsInt[6], partsInt[7], partsInt[8]},
			0,
			}
		particles = append(particles, particle)
	}

	n := 0
	loop := 1000
	nbCollisions := 0
	for loop > 0 {
		for _, p := range particles {
			incVelocity(p)
			incPosition(p)
		}
		//fmt.Println("avant coll", particles)
		particles, n = collisions(particles)
		//fmt.Println("apres coll", particles)
		nbCollisions += n
		fmt.Printf("[%v] nb collisions %v, len particules %v, total nb collisions %v\n", loop, n, len(particles), nbCollisions)
		loop--
	}

	//sort.Sort(ByManhattan(particles))

	indexMin := 0
	manhattanMin := manhattan(particles[0])
	for i,v := range particles {
		m := manhattan(v)
		if m < manhattanMin {
			manhattanMin = m
			indexMin = i
		}
	}

	fmt.Println("Closest particle is", indexMin)
	fmt.Println("Nb particles left", len(particles))
}

func collisions(particules []*particule) ([]*particule, int) {
	colls := make(map[int]*particule, 0)

	for i:=0; i<len(particules)/2; i++ {
		for j:=i+1; j < len(particules); j++ {
			if particules[i].p.equalsTo(particules[j].p) {
				colls[i] = particules[i]
				colls[j] = particules[j]
			}
		}
	}

	newParticles := make([]*particule, 0)
	for i,p := range particules {
		_,exists := colls[i]
		if !exists {
			newParticles = append(newParticles, p)
		}
	}

	return newParticles, len(colls)
}

func manhattan(p *particule) int64 {
	return abs(p.p.x) + abs(p.p.y) + abs(p.p.z)
}

func abs(v int) int64 {
	if v < 0 {
		return int64(-v)
	}
	return int64(v)
}

func incPosition(p *particule) {
	p.p.x += p.v.x
	p.p.y += p.v.y
	p.p.z += p.v.z
}

func incVelocity(p *particule) {
	p.v.x += p.a.x
	p.v.y += p.a.y
	p.v.z += p.a.z
}
