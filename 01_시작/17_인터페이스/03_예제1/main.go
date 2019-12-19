package main

import (
	"fmt"
	"strconv"
)

// fmt 패키지 : 기본적인 입출력을 담당하는 패키지
// Golang은 기본적으로 세미콜론(;) 사용 안함

func formatString(arg interface{}) string {
	switch arg.(type) {
	case int: // arg 가 int 형이라면
		i := arg.(int)         // int 형으로 값을 가져옴
		return strconv.Itoa(i) // 숫자를 문자열로 변환
	case float32: // arg 가 float32 형이라면
		f := arg.(float32)                                  // float32 형으로 값을 가져옴
		return strconv.FormatFloat(float64(f), 'f', -1, 32) // 실수를 문자열로 변환
	case float64: // arg 가 float64 형이라면
		f := arg.(float64)                         // float64 형으로 값을 가져옴
		return strconv.FormatFloat(f, 'f', -1, 64) // 실수를 문자열로 변환
	case string:
		s := arg.(string)
		return s
	default:
		return "Error"
	}
}

func main() {
	fmt.Println(formatString(1))
	fmt.Println(formatString(2.55))
	fmt.Println(formatString("Hello World"))
	fmt.Println(formatString(true))
}
