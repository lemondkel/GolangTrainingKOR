package main

import (
	"fmt"
	"math"
)

// fmt 패키지 : 기본적인 입출력을 담당하는 패키지
// Golang은 기본적으로 세미콜론(;) 사용 안함

func main() {
	calExR := 1.1               // 계산전환율
	calDam := calExR * 803.1235 // 계산달러금액

	fmt.Println("계산전환율:", calExR)
	fmt.Println("계산달러금액:", calDam)
	fmt.Println("계산달러금액(내림):", math.Floor(calDam*100)/100)   // 소수점 두자리까지 내림
	fmt.Println("계산달러금액(가까운수):", math.Round(calDam*100)/100) // 소수점 두자리까지 가까운수
	fmt.Println("계산달러금액(올림):", math.Ceil(calDam*100)/100)    // 소수점 두자리까지 올림
	// * 100 으로 나누지 않고 10으로 나누면 소수점 첫째 자리까지.
	// 1000으로 나누면 소수점 세번째 자리까지
}
