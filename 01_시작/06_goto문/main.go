package main

import "fmt"

// fmt 패키지 : 기본적인 입출력을 담당하는 패키지

func main() {
	// 타 언어와 달리 break 생략 가능
	i := 3

	switch i {
	case 0:
		fmt.Println(0)
	case 1:
		fmt.Println(1)
	case 2:
		fmt.Println(2)
	case 3:
		fmt.Println(3)
	case 4:
		fmt.Println(4)
	}

	s := "world"
	switch s {
	case "hello":
		fmt.Println("Hello")
	case "world":
		fmt.Println("World")
	default:
		fmt.Println("일치하는 문자열이 없습니다.")
	}

	//break 사용하기
	// 다음과 같은 상황에서는 if 조건문 안에 berka 키워드로 문장 실행을 중단할 수 있습니다.
	x := "Hello"
	y := 2

	switch y {
	case 1:
		fmt.Println(1)
	case 2:
		if x == "Hello" {
			fmt.Println(2)
			break
		}
		fmt.Println(2)
	}

	// fallthrough 키워드
	// 특정 case 의 문장을 실행한 뒤 다음 case의 문장을 실행하고 싶을 때
	z := 3
	switch z {
	case 4:
		fmt.Println("4 이상")
		fallthrough
	case 3:
		fmt.Println("3 이상")
		fallthrough
	case 2:
		fmt.Println("2 이상")
		fallthrough
	case 1:
		fmt.Println("1 이상")
		fallthrough
	case 0:
		fmt.Println("0 이상")
		// 마지막 case 에는 fallthrough 사용 불가능
	}

	// 여러 분기를 한번에 처리
	a := 3
	switch a {
	case 2, 4, 6:
		fmt.Println("짝수")
	case 1, 3, 5:
		fmt.Println("홀수")
	}

	// 조건으로 분기
	b := 7
	switch {
	case b >= 5 && b < 10:
		fmt.Println("5이상, 10미만")
	case b >= 0 && b < 5:
		fmt.Println("0이상, 5미만")
	}
}
