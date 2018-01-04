package main

import (
	"strconv"
	"strings"
	"bytes"
)

func Hash(tab []string) string {
	if len(tab) == 2 {
		return hash2(tab)
	}
	return hash3(tab)
}

func hash2(tab []string) string {
	return strconv.Itoa(strings.Count(tab[0], "#")) + strconv.Itoa(strings.Count(tab[1], "#"))
}

func hash3(tab []string) string {
	var buffer bytes.Buffer
	buffer.WriteString(strconv.Itoa(countAround(tab[0])))
	buffer.WriteString(strconv.Itoa(countMiddle(tab[0])))

	buffer.WriteString(strconv.Itoa(count(tab[0][2]) + count(tab[2][2])))
	buffer.WriteString(strconv.Itoa(count(tab[1][2])))

	buffer.WriteString(strconv.Itoa(countAround(tab[2])))
	buffer.WriteString(strconv.Itoa(countMiddle(tab[2])))

	buffer.WriteString(strconv.Itoa(count(tab[0][0]) + count(tab[2][0])))
	buffer.WriteString(strconv.Itoa(count(tab[1][0])))

	buffer.WriteString(strconv.Itoa(countMiddle(tab[1])))
	return buffer.String()
}

func countAround(v string) int {
	count := 0
	if v[0] == '#' {
		count++
	}
	if v[2] == '#' {
		count++
	}
	return count
}

func countMiddle(v string) int {
	if v[1] == '#' {
		return 1
	}
	return 0
}

func count(b byte) int {
	if b == '#' {
		return 1
	}
	return 0
}

func Rotate(tab []string) []string {
	if len(tab) == 2 {
		return rotate2(tab)
	}
	return rotate3(tab)
}

func rotate2(tab []string) []string {
	rot := make([]string, len(tab))
	rot[0] = string(tab[1][0]) + string(tab[0][0])
	rot[1] = string(tab[1][1]) + string(tab[0][1])
	return rot
}

func rotate3(tab []string) []string {
	rot := make([]string, len(tab))
	rot[0] = string(tab[2][0]) + string(tab[1][0]) + string(tab[0][0])
	rot[1] = string(tab[2][1]) + string(tab[1][1]) + string(tab[0][1])
	rot[2] = string(tab[2][2]) + string(tab[1][2]) + string(tab[0][2])
	return rot
}
