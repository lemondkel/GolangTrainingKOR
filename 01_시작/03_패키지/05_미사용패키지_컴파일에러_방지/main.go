package main

import (
	"fmt"
	_ "runtime"
)

// fmt 패키지 : 기본적인 입출력을 담당하는 패키지
// runtime 패키지 앞에 _를 붙여 컴파일 에러 방지 (미사용 패키지)

func main() {
	fmt.Println("Hello World")
}
