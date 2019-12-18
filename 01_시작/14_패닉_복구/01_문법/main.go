package main

import "fmt"

// fmt 패키지 : 기본적인 입출력을 담당하는 패키지
// Golang은 기본적으로 세미콜론(;) 사용 안함

func f() {
	// 지연호출 사용 (defer)
	defer func() {
		// 변수 := recover() - 패닉이 발생했을 때 프로그램 종료되지 않고 예외처리를 할 수 있음.
		// 타 언어의 try - catch 구문과 비슷함.
		s := recover() // panic 함수에서 설정한 에러메시지를 받아옴.

		fmt.Println(s, "Zz")
	}()

	// panic 부터 설정됨.
	// panic(에러메시지) => 문법적 에러는 아니지만 로직에 따라 에러로 처리하고 싶을 때 사용
	panic("Error !!")
}

func main() {
	f()

	fmt.Println("Hello, World!")
}
