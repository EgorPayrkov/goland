package main

import (
	"fmt"
	"strings"
)

func IsPalindrome(input string) bool {

	input2 := strings.ToLower(strings.Join(strings.Split(input, " "), ""))
	lenInput := len(input2)
	if len(input2)%2 != 0 {
		lenInput = len(input2) - 1
	}
	for i := 0; i < lenInput/2; i++ {
		startRune := []rune(input2)[i]
		lastRune := []rune(input2)[lenInput/2-i-1]
		if string(startRune) != string(lastRune) {
			return false
		}
	}
	return true
}

func main() {
	input := "А роза упаала на лапу Азора"
	fmt.Println(IsPalindrome(input))
}
