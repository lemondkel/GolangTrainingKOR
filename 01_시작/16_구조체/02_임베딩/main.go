package main

import "fmt"

// fmt 패키지 : 기본적인 입출력을 담당하는 패키지
// Golang은 기본적으로 세미콜론(;) 사용 안함

type Person struct {
	name string
	age  int
}

func (p *Person) greeting() {
	fmt.Println("Hello~")
}

type Student struct {
	Person // 임베딩 효과. (Is-a) 포함하는 관계. 학생은 사람이다.
	school string
	grade  int
}

func (p *Student) greeting() {
	// 메소드 오버라이딩
	fmt.Println("Hello Students~")
}

func main() {
	var s Student

	// 구조체 타입을 통해 호출하거나 바로 호출할 수 있음
	s.Person.greeting()
	s.greeting()
}
