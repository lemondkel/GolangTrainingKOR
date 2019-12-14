package main

import "fmt"

// fmt 패키지 : 기본적인 입출력을 담당하는 패키지
// 99 Bottles of Beer
func main() {
	for bottles := 99; bottles >= 0; bottles-- {
		if bottles >= 1 {
			fmt.Printf("%d bottles of beer on the wall, %d bottles of beer\n", bottles, bottles)
			if bottles > 1 {
				nextStr := bottles - 1
				fmt.Printf("Take one down, pass it around, %d bottles of beer on the wall.\n", nextStr)
			} else {
				nextStr := "No more"
				fmt.Printf("Take one down, pass it around, %s bottles of beer on the wall.\n", nextStr)
			}
		} else {
			fmt.Printf("No more bottles of beer on the wall, no more bottles of beer\n")
			fmt.Printf("Go to the store and buy some more, 99 bottles of beer on the wall.\n")
		}
	}
}
