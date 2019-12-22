package main

import "fmt"

// fmt 패키지 : 기본적인 입출력을 담당하는 패키지

func sum(a int, b int, c chan int) {
	c <- a + b // 채널에 a와 b의 합을 보냄
}

func main() {
	c := make(chan int) // int형 채널을 생성
	// 채널을

	go sum(1, 2, c) // sum 함수를 고루틴으로 실행한 뒤 채널을 매개변수로 넘겨줌

	n := <-c // 채널에서 값을 꺼낸 뒤 n에 대입

	fmt.Println(n) // 3
}
