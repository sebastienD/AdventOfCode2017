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
		prevA = (prevA * pA.factor) % 2147483647
		binA := leftPad2Len(strconv.FormatInt(prevA, 2), "0", 16)
		sA := binA[len(binA)-16:]

		prevB = (prevB * pB.factor) % 2147483647
		binB := leftPad2Len(strconv.FormatInt(prevB, 2), "0", 16)
		sB := binB[len(binB)-16:]

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
	result := generator(40000000, param{16807, 634}, param{48271, 301})

	fmt.Printf("result %v\n", result)
}
