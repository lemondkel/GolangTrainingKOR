package main

import "fmt"

// fmt 패키지 : 기본적인 입출력을 담당하는 패키지
// Golang은 기본적으로 세미콜론(;) 사용 안함

func main() {
	var a [5]int   // int 형 길이 5 배열 선언
	a[2] = 7       // 배열 세번째 요소에 7 대입
	fmt.Println(a) // 0 0 7 0 0

	var b = [5]int{1, 2, 3, 4, 5}
	fmt.Println(b) // 1 2 3 4 5

	var c = [5]string{"A", "B", "C", "", "E"}
	fmt.Println(c)

	// 배열에 ... 사용시 자동으로 크기 결정됨
	var d = [...]int{2, 3, 4}
	d[2] = 1
	fmt.Println(d)

	// 배열의 길이를 구하는 함수 len(target)
	e := [5]int{1, 2, 3, 4, 5}
	for i := 0; i < len(e); i++ {
		fmt.Println(e[i])
	}

	// i 에는 인덱스, value 에는 배열의 요소 값이 들어감
	// for 문 range 사용시, 매개변수 1 에는 인덱스, 매개변수 2 에는 배열 값
	for i, value := range e {
		fmt.Println(i, value)
	}

	// 인덱스 사용하지 않을 때 _로 선언.
	for _, value := range e {
		fmt.Println(value)
	}

	// 배열 복사하기
	f := e
	fmt.Println(e)
	fmt.Println(f)
}
