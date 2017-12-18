package main

import (
	"os"
	"bufio"
	"strings"
	"log"
	"strconv"
	"fmt"
)


type registers struct {
	reg map[string]int64
	last int64
	receive int64
}

func newRegisters() *registers {
	return &registers{make(map[string]int64, 0), 0, 0}
}

func (r *registers) value(name string) int64 {
	v, err := strconv.ParseInt(name, 10, 64)
	if err == nil {
		return int64(v)
	}
	return r.reg[name]
}

func (r *registers) snd(varVal string) {
	v := r.value(varVal)
	if v != 0 {
		fmt.Printf("set last to %v\n", v)
		r.last = v
	}
}

func (r *registers) set(name string, varVal string) {
	r.reg[name] = r.value(varVal)
}

func (r *registers) add(name string, varVal string) {
	r.reg[name] += r.value(varVal)
}

func (r *registers) mul(name string, varVal string) {
	r.reg[name] *= r.value(varVal)
}

func (r *registers) mod(name string, varVal string) {
	r.reg[name] = r.reg[name] % r.value(varVal)
}

func (r *registers) rcv(varVal string) {
	v := r.value(varVal)
	if v != 0 {
		fmt.Printf("set receive to %v\n", r.last)
		r.receive = r.last
	}
}

func (r *registers) jgz(x string, y string) int64 {
	v := r.value(x)
	if v > 0 {
		return r.value(y)
	}
	return 1
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Open file failed")
	}
	defer file.Close()

	instr := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		instr = append(instr, scanner.Text())
		//fmt.Printf("tab %v\n", instr)
	}

	regist := newRegisters()
	index := 0
	for regist.receive == 0 {
		fmt.Printf("loop %v\n", index)
		line := instr[index]
		words := strings.Split(line, " ")
		switch words[0] {
		case "snd":
			regist.snd(words[1])
			index ++
		case "set":
			regist.set(words[1], words[2])
			index ++
		case "add":
			regist.add(words[1], words[2])
			index ++
		case "mul":
			regist.mul(words[1], words[2])
			index ++
		case "mod":
			regist.mod(words[1], words[2])
			index ++
		case "rcv":
			regist.rcv(words[1])
			index ++
		case "jgz":
			index += int(regist.jgz(words[1], words[2]))
		}
	}

	fmt.Printf("Receive is %v\n", regist.receive)
}
