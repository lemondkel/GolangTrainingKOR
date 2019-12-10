package main

import "fmt"

// fmt 패키지 : 기본적인 입출력을 담당하는 패키지
// Golang은 기본적으로 세미콜론(;) 사용 안함

func main() {
	exam1() // 예제1
	exam2() // 예제2
}

/**
예제1 - 본인의 팀 확인하기
*/
func exam1() {
	var team = "Aston Villa"
	var text = "회원님은 " + team + " 팬입니다."

	fmt.Println(text)
}

/**
예제2 - 핸드폰 기종 확인 (상수 변수 활용)
*/
func exam2() {
	const phoneModel = "SAMSUNG A50"
	fmt.Println("PHONE MODEL: ", phoneModel)

	// phoneModel = "SAMSUNG A500" // 문법 에러 - 상수(const) 값은 변경 X
}
