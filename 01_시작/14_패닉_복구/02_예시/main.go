package main

import "fmt"

// fmt 패키지 : 기본적인 입출력을 담당하는 패키지
// Golang은 기본적으로 세미콜론(;) 사용 안함

func f() {
	defer func() {
		s := recover()
		fmt.Println(s)
	}()

	a := [...]int{1, 2} // 정수가 2개 저장된 배열
	for i := 0; i < 5; i++ {
		fmt.Println(a[i])
	}
}

func main() {
	f()

	fmt.Println("Hello, World!")
}
