package main

import "fmt"

// fmt 패키지 : 기본적인 입출력을 담당하는 패키지
// Golang은 기본적으로 세미콜론(;) 사용 안함

// 값이나 인스턴스의 실제 타입은 상관하지 않고 구현된 메서드로만 타입을 판단하는 방식
// Duck typing
// 덕 테스트에서 유래됨

// 오리 구조체 정의
type Duck struct {
}

// 오리의 quack 메서드 정의
func (d Duck) quack() {
	fmt.Println("꽥!")
}

func (d Duck) feathers() {
	fmt.Println("오리는 흰색과 회색 털을 가지고 있습니다.")
}

// 사람 구조체 정의
type Person struct {
}

// 사람의 quack 메서드 정의
func (p Person) quack() {
	fmt.Println("사람이 오리를 흉내냅니다. 꽥!")
}

// 사람의 feathers 메서드 정의
func (p Person) feathers() {
	fmt.Println("사람은 땅에서 깃털을 주워서 보여줍니다.")
}

// quack, feathers 메서드를 가지는 Quacker 인터페이스 정의
type Quacker interface {
	quack()
	feathers()
}

func inTheForest(q Quacker) {
	q.quack()    // Quacker 인터페이스로 quack 메서드 호출
	q.feathers() // Quacker 인터페이스로 feathers 메서드 호출
}

func main() {
	var donald Duck // 오리 인스턴스 생성
	var john Person // 사람 인스턴스 생성

	inTheForest(donald)
	// 타입이 특정 인터페이스를 구현하는지 검사하는 구문
	// interface{}(인스턴스).(인터페이스)
	// 구현하고 있다면 true
	// 구현하고 있지 않다면 false
	if _, ok := interface{}(john).(Quacker); ok {
		inTheForest(john)
	}
}
