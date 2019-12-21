package main

import (
	"fmt"
	"runtime"
)

// fmt 패키지 : 기본적인 입출력을 담당하는 패키지
// Golang은 기본적으로 세미콜론(;) 사용 안함

func main() {
	runtime.GOMAXPROCS(1)

	s := "Hello, World!"

	for i := 0; i < 100; i++ {
		go func() {
			fmt.Println(s, i) // 반복문의 변수를 클로저에서 바로 출력
		}()
	}

	fmt.Scanln()
}
