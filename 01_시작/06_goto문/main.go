package main

import "fmt"

// fmt 패키지 : 기본적인 입출력을 담당하는 패키지

// goto 키워드는 정해진 레이블로 곧장 이동합니다.
// goto 레이블명
// 레이블명:
func main() {
	goto LABEL                // 이동할 레이블을 지정합니다.
	fmt.Println("My Name Is") // 실행 X

LABEL:
	fmt.Println("Slim Shady")

	var a int = 1

	if a == 1 {
		fmt.Println("Error 1")
		return
	}

	if a == 2 {
		fmt.Println("Error 1")
		return
	}

	if a == 3 {
		fmt.Println("Error 1") // 중복 코드
		return
	}
	fmt.Println(a)
	fmt.Println("Success")

	if a == 1 {
		goto ERROR1
	}

	if a == 2 {
		goto ERROR2
	}

	if a == 3 {
		goto ERROR1
	}
	fmt.Println(a)
	fmt.Println("Success")

ERROR1:
	fmt.Println("ERROR 1")
	return
ERROR2:
	fmt.Println("ERROR 2")
	return
}
