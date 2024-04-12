package main

import (
	"fmt"
)

const cipher = `
141291826972
252615211269_211813261324182615_
81876267181213_1192224269181268+14687_
11262_21129_819269228_1813_
116924192682223_25922422918228_152687_
141213719+267_15222687_71992222_
7191268261323_1112613238+1821_21269_
1112972112151812_241261523_16222211_
71922_1426916227_8726251522_8726_
`

func main() {
	currChar := int(0)
	for _, char := range cipher {
		switch {
		case char == '\n':
			printIfNotEmpty(currChar)
			currChar = 0
			fmt.Printf("\n")

		case char == '_':
			printIfNotEmpty(currChar)
			currChar = 0
			fmt.Print(" ")

		case char == '+':
			printIfNotEmpty(currChar)
			currChar = 0
			fmt.Print("+")

		case currChar == 1 || currChar == 2:
			currChar *= 10
			currChar += (asInt(char))

			print(currChar)
			currChar = 0

		case char == '1' || char == '2':
			currChar = asInt(char)

		default:
			print(asInt(char))
		}
	}
}

func printIfNotEmpty(num int) {
	if num != 0 {
		print(num)
	}
}

// print letter of the English alphabet by reversing it's number, z=1, y=2, ..., a=26
func print(num int) {
	if num == 0 {
		fmt.Printf(" ")
		return
	}
	fmt.Printf("%c", 'z'-num+1)
}

func asInt(num rune) int {
	return int(num - '0')
}
