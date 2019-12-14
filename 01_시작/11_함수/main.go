package main

import "fmt"

// fmt 패키지 : 기본적인 입출력을 담당하는 패키지
// Golang은 기본적으로 세미콜론(;) 사용 안함

func main() {
	hello()

	result := sum(5, 15)
	fmt.Println(result)

	// 리턴값이 2개인 함수
	result2, result3 := sumAndDiff(5, 15)
	fmt.Println(result2, result3)

	// 첫번째 리턴값은 사용하지 않으려고 함.
	_, result4 := sumAndDiff(5, 15)
	fmt.Println(result4)
}

// 함수 정의
// C, C++와 달리 함수 정의 위치가 중요하지 않음.
// func 함수명(매개변수...) {}
func hello() {
	fmt.Println("Hello World")
}

// 정수 2자리의 합을 구하는 함수
func sum(x int, b int) int {
	return x + b
}

// 리턴값이 2개인 함수
// 다른 언어와 다르게 리턴을 여러개 보낼 수 있음.
// 합과 차를 구함
func sumAndDiff(x int, b int) (int, int) {
	sum := x + b
	diff := x - b
	return sum, diff
}
