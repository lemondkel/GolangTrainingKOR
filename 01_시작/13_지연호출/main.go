package main

import "fmt"

// fmt 패키지 : 기본적인 입출력을 담당하는 패키지
// Golang은 기본적으로 세미콜론(;) 사용 안함

func hello() {
	fmt.Println("Hello World")
}

func main() {
	fmt.Println("HOHOHOHO!")
	defer hello() // 지연 호출 사용 // 제일 마지막
	// 현재 함수 끝나기 직전에 호출함 (main())
	fmt.Println("Hi~")

	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d", i) // 마지막에서 두번째
	}
}
