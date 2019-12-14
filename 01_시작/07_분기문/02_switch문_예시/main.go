package main

import (
	"fmt"
	"math/rand"
	"time"
)

// fmt 패키지 : 기본적인 입출력을 담당하는 패키지
// math, rand: 무작위(랜덤) 패키지
// time : 시간 패키지
// time.Now() - 현재 시간을 구함
// UnixNano() -> 유닉스 시간을 나노 초 단위로 리턴
// Seed() => 시드값을 설정하는 함수.
// Intn(n) -> 랜덤 값을 생성 (0부터 n까지)

func main() {
	rand.Seed(time.Now().UnixNano()) // 현재 시간으로 Seed 값 설정
	switch i := rand.Intn(10); {
	case i >= 3 && i < 6:
		fmt.Println("3 이상, 6 미만")
	case i == 9:
		fmt.Println(9)
	default:
		fmt.Println(i)
	}
}
