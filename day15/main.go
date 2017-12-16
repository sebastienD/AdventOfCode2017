package main

import (
	"fmt"
	"strconv"
	"strings"
)

type param struct {
	factor int64
	first  int64
}

func generator(num int64, pA param, pB param) int {
	prevA := pA.first
	prevB := pB.first
	count := 0
	for num > 0 {
		for prevA%4 != 0 {
			prevA = (prevA * pA.factor) % 2147483647
		}
		//fmt.Printf("prev A: %v\n", prevA)
		binA := leftPad2Len(strconv.FormatInt(prevA, 2), "0", 16)
		sA := binA[len(binA)-16:]
		prevA = (prevA * pA.factor) % 2147483647

		for prevB%8 != 0 {
			prevB = (prevB * pB.factor) % 2147483647
		}
		//fmt.Printf("prev B: %v\n", prevB)
		binB := leftPad2Len(strconv.FormatInt(prevB, 2), "0", 16)
		sB := binB[len(binB)-16:]
		prevB = (prevB * pB.factor) % 2147483647

		if sA == sB {
			count++
		}
		if num%1000000 == 0 {
			fmt.Printf("iter %v count %v\n", num, count)
		}

		num--
	}
	return count
}

func leftPad2Len(s string, padStr string, overallLen int) string {
	var padCountInt int
	padCountInt = 1 + ((overallLen - len(padStr)) / len(padStr))
	var retStr = strings.Repeat(padStr, padCountInt) + s
	return retStr[(len(retStr) - overallLen):]
}

func main() {
	result := generator(5000000, param{16807, 634}, param{48271, 301})

	fmt.Printf("result %v\n", result)
}
