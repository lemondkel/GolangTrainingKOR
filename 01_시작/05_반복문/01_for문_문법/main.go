package main

import "fmt"

// fmt 패키지 : 기본적인 입출력을 담당하는 패키지

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	// 중괄호 위치가 이상하거나, 생략하면 if문과 동일하게 컴파일 에러 발생

	/*for {
		// 무한루프
		fmt.Println("Hello, World!")
	}*/

	i := 0
	for {
		if i > 4 {
			break
		}

		fmt.Println(i)
		i++
	}

Loop: // 레이블 지정
	// fmt.Println("Hello") // 들어가면 안되는 코드 (레이블과 for문 사이 엄격한 문법)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if j == 2 {
				break Loop
				// 레이블 탈출
			}

			fmt.Println(i, j)
		}
	}

	fmt.Println("Hello, World!")

	// 반복문에서 여러개의 변수 사용
	for i, j := 0, 0; i < 10; i, j = i+1, j+2 {
		fmt.Println(i, j)
	}
}
