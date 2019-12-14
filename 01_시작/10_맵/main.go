package main

import "fmt"

// fmt 패키지 : 기본적인 입출력을 담당하는 패키지
// Golang은 기본적으로 세미콜론(;) 사용 안함

// 맵은 해시 테이블 ,딕셔너리라고도 함
// 키-값 형태 자료 저장
// 레퍼런스 타입
func main() {
	a := make(map[string]int) // key - string, value - int
	a["asd"] = 1
	fmt.Println(a)
	fmt.Println(a["asd"])

	// 맵에 데이터 저장하고 조회하기
	// 미국달러 기준 해외송금환율
	frRmtExR := make(map[string]float64)
	frRmtExR["USD"] = 1
	frRmtExR["KRW"] = 0.031
	frRmtExR["VND"] = 0.000049

	fmt.Println("미국달러 환율:", frRmtExR["USD"])

	// 맵 순회하기
	for key, value := range frRmtExR {
		fmt.Println(key, value)
	}

	// 맵 데이터 삭제
	delete(frRmtExR, "VND")
	fmt.Println(frRmtExR)
}
