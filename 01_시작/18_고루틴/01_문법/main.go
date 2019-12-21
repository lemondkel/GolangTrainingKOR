package main

import "fmt"

// fmt 패키지 : 기본적인 입출력을 담당하는 패키지
// Golang은 기본적으로 세미콜론(;) 사용 안함

func hello() {
	fmt.Println("Hello")
}

func main() {
	go hello() // 함수를 고루틴으로 실행

	fmt.Scanln() // main 함수가 종료되지 않도록 대기
}
