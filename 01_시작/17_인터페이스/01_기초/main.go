package main

import "fmt"

// fmt 패키지 : 기본적인 입출력을 담당하는 패키지
// Golang은 기본적으로 세미콜론(;) 사용 안함

// type 인터페이스명 interface {}
type hello interface {
	// 인터페이스 정의
	method1()          // 리턴값이 없을 때
	method2() int      // 리턴값이 있을 때 (정수형)
	method3(x int) int // 매개변수 하나와 리턴값이 있을 때 (정수형)
}

type MyInt int

func (i MyInt) method2() int {
	panic("implement me")
}

func (i MyInt) method3(x int) int {
	panic("implement me")
}

func (i MyInt) method1() {
	fmt.Println("method1!")
}

func f1(arg interface{}) {
	// 모든 타입 저장 가능한 메서드
}

type Any interface{}

func f2(arg Any) {
	// 모든 타입을 저장할 수 있음
}

func main() {
	var i MyInt = 5
	var h hello
	h = i
	h.method1()
}
