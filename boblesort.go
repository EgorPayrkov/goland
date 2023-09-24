package main

import "fmt"

func Boblesort(s []int) []int {
	ls := len(s)

	for i := 0; i < ls-1; i++ {
		for j := 0; j < (ls - 1 - i); j++ {
			if s[j] > s[j+1] {
				s[j], s[j+1] = s[j+1], s[j]
			}
		}
	}
	return s
}

func main() {
	arr := []int{64, 34, 25, 12, 22, 11, 90}
	fmt.Println(Boblesort(arr))

}
