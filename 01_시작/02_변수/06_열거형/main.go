package main

import "fmt"

// fmt 패키지 : 기본적인 입출력을 담당하는 패키지
// Golang은 기본적으로 세미콜론(;) 사용 안함

func main() {
	const (
		Sunday = iota // 0
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
		numberOfDays
	)

	fmt.Println(Sunday)       // 0
	fmt.Println(Monday)       // 1
	fmt.Println(Tuesday)      // 2
	fmt.Println(Wednesday)    // 3
	fmt.Println(Thursday)     // 4
	fmt.Println(Friday)       // 5
	fmt.Println(Saturday)     // 6
	fmt.Println(numberOfDays) // 7

}
