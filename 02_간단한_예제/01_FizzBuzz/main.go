package main

import "fmt"

// fmt 패키지 : 기본적인 입출력을 담당하는 패키지

func main() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		}
		fmt.Println(i)
	}
}
