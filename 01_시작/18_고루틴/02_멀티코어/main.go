package main

import (
	"fmt"
	"runtime"
)

// fmt 패키지 : 기본적인 입출력을 담당하는 패키지
// Golang은 기본적으로 세미콜론(;) 사용 안함

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	// CPU 개수를 구한 뒤 사용할 최대 CPU 개수 설정

	fmt.Println(runtime.GOMAXPROCS(0)) // 설정 값 출력

	s := "Hello, World!"

	for i := 0; i < 100; i++ {
		go func(n int) {
			// 익명 함수를 고루틴으로 설정
			fmt.Println(s, n)
			// s 와 매개변수로 받은 n값 출력
			//반복문의 변수는 매개변수로 넘겨줌
		}(i)
	}

	fmt.Scanln()
}
