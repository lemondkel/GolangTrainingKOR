package main

import (
	"fmt"
	"unicode/utf8"
)

// fmt 패키지 : 기본적인 입출력을 담당하는 패키지
// Golang은 기본적으로 세미콜론(;) 사용 안함

func main() {
	exam1() // 예제1
	exam2() // 예제2
	exam3() // 예제3
	exam4() // 예제4
	exam5() // 예제5
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

/**
예제3 - 문자열 길이 구하기
*/
func exam3() {
	str := "한글"
	fmt.Println(len(str))                    // UTF-8인 경우 정상적인 길이를 가져올 수 없어 아래 메소드 사용
	fmt.Println(utf8.RuneCountInString(str)) // 문자열의 실제 길이를 구함
}

/**
예제4 - 문자열 비교, 추가
*/
func exam4() {
	var s1 string = "l2jong"
	var s2 string = "l2jong"
	var s3 = s2

	fmt.Println(s1 == s3)
	fmt.Println(s1 + s2 + s3)
}

/**
예제5 - 문자열 중 문자 하나 출력
*/
func exam5() {
	str := "ABCDEFG"
	fmt.Printf("%c\n", str[1]) // B
}
