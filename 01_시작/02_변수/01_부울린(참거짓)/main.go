package main

import "fmt"

// fmt 패키지 : 기본적인 입출력을 담당하는 패키지
// Golang은 기본적으로 세미콜론(;) 사용 안함

func main() {
	exam1() // 예제1
	exam2() // 예제2
}

/**
예제1 - isChecked 부울린 값에 따른 텍스트 변경
*/
func exam1() {
	var isChecked bool
	isChecked = true
	var isCheckedText string

	if isChecked {
		isCheckedText = "체크"
	} else {
		isCheckedText = "미체크"
	}

	fmt.Println("체크박스가 [", isCheckedText, "] 상태 입니다.")
}

/**
예제2 - 부울린 변수 2개를 이용한 조건문
*/
func exam2() {
	var onlyCheckCard = true
	var isLogin = true

	if onlyCheckCard && isLogin {
		// 체크카드 사용자에 로그인한 경우
		fmt.Println("회원님은 체크카드 사용자입니다.")
	} else {
		if !isLogin {
			// 미로그인 상태인 경우
			fmt.Println("로그인이 필요합니다.")
		} else {
			// 체크카드만 보유한 고객이 아닌 경우
			fmt.Println("회원님은 대상이 아닙니다.")
		}
	}
}
