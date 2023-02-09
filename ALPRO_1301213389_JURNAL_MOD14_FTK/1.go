package main

import "fmt"

var (
	pita  string
	CC    rune
	EOP   bool
	index int
)

func START() {
	index = 0
	CC = rune(pita[index])
	EOP = CC == '.'
}

func ADV() {
	index = index + 1
	CC = rune(pita[index])
	EOP = CC == '.'

}

func DIGIT() bool {
	if CC >= '0' && CC <= '9' {
		return true
	}
	return false
}

func STRIP() {
	for CC == '-' && !EOP {
		ADV()
	}
}

func main() {
	var valid, xAA, xBB, xCC bool
	fmt.Scanln(&pita)
	START()
	valid = false
	for CC != '-' && !EOP {
		ADV()
	}
	if CC == '-' {
		valid = ((pita[0] == 'I' && pita[1] == 'F') || (pita[0] == 'I' && pita[1] == 'F') || (pita[0] == 'S' && pita[1] == 'E'))
		xAA = valid && index == 2
		ADV()
	}
	valid = true
	if CC == '-' {
		valid = false
	} else if !EOP {
		for CC != '-' && !EOP && valid {
			valid = DIGIT()
			ADV()
		}
	}
	if CC == '-' {
		xBB = valid
		ADV()
	}
	valid = true
	if !EOP {
		for !EOP && valid {
			valid = DIGIT()
			ADV()
		}
		xCC = valid
	}
	valid = xAA && xBB && xCC
	fmt.Println(valid)
}
