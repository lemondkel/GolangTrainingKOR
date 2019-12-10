package main

import "fmt"

// fmt 패키지 : 기본적인 입출력을 담당하는 패키지
// Golang은 기본적으로 세미콜론(;) 사용 안함

func main() {
	var a int            // 변수 a 선언과 동시에 자료형 int로 변경
	a = 10               // a에 10 대입
	fmt.Println("a:", a) // a 출력

	var b int = 15       // 변수 b 선언과 동시에 자료형 int 15 대입
	fmt.Println("b:", b) // b 출력

	var c, d = 20, 30 // 변수 c, d 선언과 동시에 자료형 int 20, 30 대입
	// * int 형은 생략 가능
	fmt.Println("c:", c) // c 출력
	fmt.Println("d:", d) // d 출력

	fmt.Println("a+b:", a+b)
}
