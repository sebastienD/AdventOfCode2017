package main

import (
	"os"
	"bufio"
	"strings"
	"log"
	"strconv"
	"fmt"
	"sync"
	"time"
)


type registers struct {
	reg map[string]int64
	sent int64
	id int64
}

func newRegisters(pValue int64) *registers {
	r := &registers{make(map[string]int64, 0), 0, pValue}
	r.reg["p"] = pValue
	return r
}

func (r *registers) value(name string) int64 {
	v, err := strconv.ParseInt(name, 10, 64)
	if err == nil {
		return int64(v)
	}
	return r.reg[name]
}

func (r *registers) snd(varVal string, c chan int64) {
	v := r.value(varVal)
	//if v != 0 {
		//fmt.Printf("set last to %v\n", v)
		//r.last = v
		r.sent++
		fmt.Printf("[%v] sent %v %v\n", r.id, varVal, v)
		c <- v
		//fmt.Printf("[%v] end sent %v %v\n", r.id, varVal, v)
	//}
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

func (r *registers) rcv(name string, c chan int64) bool {
	fmt.Printf("[%v] wait rcv %v\n", r.id, name)
	select {
		case res := <- c:
			fmt.Printf("[%v] end wait rcv %v %v \n", r.id, name, res)
			r.reg[name] = res
		case <-time.After(1 * time.Second):
			return false
	}
	return true
}

func (r *registers) jgz(x string, y string) int64 {
	v := r.value(x)
	if v > 0 {
		return r.value(y)
	}
	return 1
}

func main() {
	file, err := os.Open("input-test2.txt")
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

	instr1 := make([]string, len(instr))
	for i,v := range instr {
		instr1[i] = v
	}

	var wg sync.WaitGroup
	wg.Add(2)

	c0 := make(chan int64, 1000)
	c1 := make(chan int64, 1000)

	go launch(instr, 0, c1, c0, wg)
	go launch(instr1, 1, c0, c1, wg)

	wg.Wait()

	//fmt.Printf("Nb sent is %v\n", result)
}

func launch(instr []string, pValue int64, sendChan chan int64, rcvChan chan int64,  wg sync.WaitGroup) {
	regist := newRegisters(pValue)
	index := 0
	loop: for {
		line := instr[index]
		words := strings.Split(line, " ")
		fmt.Printf("[%v] Line %v\n", pValue, words)
		switch words[0] {
		case "snd":
			regist.snd(words[1], sendChan)
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
			fmt.Printf("[%v] line recv and sent %v\n", pValue, regist.sent)
			wg.Done()
			if !regist.rcv(words[1], rcvChan) {
				break loop
			}
			wg.Add(1)
			fmt.Printf("[%v] line recv release\n", pValue)
			index ++
		case "jgz":
			index += int(regist.jgz(words[1], words[2]))
		}
	}
	fmt.Printf("[%v] Done\n", pValue)
}