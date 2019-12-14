package main

import (
	test1 "fmt"
	test2 "runtime"
)

// fmt 패키지 : 기본적인 입출력을 담당하는 패키지
// 패키지를 가져올 때 패키지 이름 앞에 별칭을 붙여 사용할 수 있음.

func main() {
	test1.Println("CPU Count: ", test2.NumCPU())
}
