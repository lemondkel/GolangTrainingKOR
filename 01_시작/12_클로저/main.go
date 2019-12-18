package main

import "fmt"

// fmt 패키지 : 기본적인 입출력을 담당하는 패키지
// Golang은 기본적으로 세미콜론(;) 사용 안함

func main() {
	a, b := 3, 5

	// 익명 함수 바깥의 변수 사용 가능
	f := func(x int) int {
		return a*x + b
	}

	y := f(5)
	fmt.Println(y) // 20

	// 클로저 사용 - 함수가 선언될 때의 환경을 계속 유지함.
	// 함수형 언어의 큰 특징.
	f2 := calc()
	fmt.Println(f2(1)) // 8
	fmt.Println(f2(2)) // 11
}

func calc() func(x int) int {
	a, b := 3, 5
	return func(x int) int {
		return a*x + b
	}
}
