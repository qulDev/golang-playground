package main

import "fmt"

func main() {
	nums := []int{10, 20, 30}
	grow(nums)

}

func grow(s []int) {
	s = append(s, 40)
	s[0] = 99
	fmt.Println("Cap :", cap(s))
}
