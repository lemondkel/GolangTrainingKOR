package main

import "fmt"

// fmt 패키지 : 기본적인 입출력을 담당하는 패키지

func main() {
	i := 10
	if i >= 5 {
		// 조건의 결과가 참(true)이라 해당 블록 실행
		fmt.Println("5 이상")
	}

	/*if i >= 5
	{
		fmt.Println("5 이상")
	}*/
	// 컴파일 에러 발생. 여는 중괄호 위치가 무조건 조건문을 시작하는 줄에 있어야 함.

	/*if >= 5
	fmt.Println("5 이상")*/
	// 컴파일 에러. 중괄호 생략 불가능
}
