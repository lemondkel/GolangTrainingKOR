package main

import "fmt"

// fmt 패키지 : 기본적인 입출력을 담당하는 패키지
// Golang은 기본적으로 세미콜론(;) 사용 안함

func main() {
	var numPtr *int // 포인터형 변수를 선언하면 nil로 초기화됨
	fmt.Println(numPtr)

	// 빈 포인터형 변수는 바로 사용할 수 없으므로  new 함수로 메모리 할당 필요
	numPtr2 := new(int)
	fmt.Println(numPtr2)

	// 포인터형 변수에 대입하거나 값을 가져오려면 역참조 사용
	var numPtr3 *int = new(int) // new 함수로 공간 할당
	*numPtr3 = 1                // 역참조로 포인터형 변수에 값 대입
	fmt.Println(*numPtr3)       // 포인터형 변수에서 값 가져오기

	// 일반 변수에 참조(레퍼런스)를 사용하면 포인터형 변수에 대입 가능
	num := 1
	numPtr4 := &num
	fmt.Println(numPtr4)
	fmt.Println(&num)
}
