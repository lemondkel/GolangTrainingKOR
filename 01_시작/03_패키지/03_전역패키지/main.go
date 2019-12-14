package main

import (
	. "fmt"
	. "runtime"
)

// fmt 패키지 : 기본적인 입출력을 담당하는 패키지
// 전역 패키지를 사용할 때는 import . "fmt"와 같이 패키지명 앞에 .을 붙여준다.
// ** 단, 함수, 변수, 상수 이름이 중복될 수 있어 유의해야 함.

func main() {
	Println("CPU Count: ", NumCPU())
}
