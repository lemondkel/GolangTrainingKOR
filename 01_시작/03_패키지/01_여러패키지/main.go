package main

import "fmt"
import "runtime"

// fmt 패키지 : 기본적인 입출력을 담당하는 패키지
// 여러 패키지를 사용할 때는 import 키워드를 여러 개 사용하여 패키지를 가져옴.

func main() {
	fmt.Println("CPU Count: ", runtime.NumCPU())
}
