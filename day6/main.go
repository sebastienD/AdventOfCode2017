package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
)

func read() []int {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = '\t'
	reader.FieldsPerRecord = -1

	data, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	d := make([]int, 0)
	for _, each := range data[0] {
		v, _ := strconv.Atoi(each)
		d = append(d, v)
	}

	return d
}

func main() {
	//input := []int{0, 2, 7, 0}
	input := read()
	moves := addMove([][]int{}, input)
	counter := 0
	firstIndex := 0
Loop:
	for {
		index, value := 0, 0
		for i, v := range input {
			if v > value {
				value = v
				index = i
			}
		}
		input[index] = 0
		for i := 0; i < value; value-- {
			index = (index + 1) % len(input)
			input[index]++
		}
		counter++
		for i, v := range moves {
			if reflect.DeepEqual(v, input) {
				firstIndex = i
				break Loop
			}
		}
		moves = addMove(moves, input)
	}

	fmt.Printf("counter %v, first %v, delta %v\n", counter, firstIndex, (counter - firstIndex))
}

func addMove(moves [][]int, move []int) [][]int {
	tmp := make([]int, len(move))
	copy(tmp, move)
	return append(moves, tmp)
}
