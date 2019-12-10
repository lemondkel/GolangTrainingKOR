package main

import (
	"fmt"
	"strconv"
)

// fmt 패키지 : 기본적인 입출력을 담당하는 패키지
// strconv 패키지 : 문자열<->각종 데이터유형 으로 변환하는 기본 패키지
// Golang은 기본적으로 세미콜론(;) 사용 안함

func main() {
	calExR := 1.1               // 계산전환율
	calDam := calExR * 803.1235 // 계산달러금액
	calDamStr := strconv.FormatFloat(calDam, 'f', 2, 64)
	// 2번째 자리 (prec) 까지 문자열로 변환.
	// -1 입력시 원본 실수를 문자열로 변환

	fmt.Println("계산달러금액:", calDam)
	fmt.Println("계산달러금액(정수):", int(calDam)) // 정수로 형변환
	fmt.Println("계산달러금액(문자열):", calDamStr)
}
