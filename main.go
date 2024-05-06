package main

import (
	"github.com/01-edu/z01"
)

func main() {
	sr := []string{"200", "400", "150", "355", "60", "10"}
	r := Max(sr)
	q := Itoa(r)
	PrintRune(q)
}
func Atoi(str string) int {
	// var charInteger int
	// result := 0
	// for _, char := range str {
	// 	charInteger = int(char - '0')
	// 	result = result*10 + charInteger

	// }
	// return result
	var num int
	for _, v := range str {
		if v < '0' || v > '9' {
			return 0
		}
		num = num*10 + (int(v) - 48)
	}
	return num
}
func Max(s []string) int {
	var result []int

	for _, element := range s {
		charInt := Atoi(element)
		result = append(result, charInt)

	}
	max := result[0]
	for i := 0; i <= len(result)-1; i++ {
		if max < result[i] {
			max = result[i]
		}

	}
	return max

}
func Itoa(t int) string {
	var div int
	var mod int
	var st []string
	var run rune
	var finalst string
	div = t
	for div > 0 {
		mod = div % 10
		div = div / 10
		run = int32(mod) + '0'

		st = append(st, string(run))
	}
	for i := len(st) - 1; i >= 0; i-- {
		finalst += st[i]
	}
	return finalst

}
func PrintRune(s string) {
	for _, char := range s {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')

}
