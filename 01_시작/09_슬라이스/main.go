package main

import "fmt"

// fmt 패키지 : 기본적인 입출력을 담당하는 패키지
// Golang은 기본적으로 세미콜론(;) 사용 안함

// 슬라이스는 가변 배열
// 동적으로 크기가 늘어날 수 있음
func main() {
	var a []int        // int형 슬라이스 선언
	a = make([]int, 5) // make 함수로 길이가 5인 int 슬라이스 선언
	a[0] = 1
	a[1] = 2
	// a[5] = 3 // ERROR - 길이 5로 make 함수 사용함
	// a[6] = 4 // ERROR - 길이 5로 make 함수 사용함
	fmt.Println(a)

	b := []int{
		32,
		29,
		78,
		16,
		81,
	}
	fmt.Println(b)

	// 길이5, 용량10인 슬라이스 생성
	c := make([]int, 5, 10)
	fmt.Println(cap(b)) // b의 용량
	fmt.Println(cap(c)) // c의 용량

	// 슬라이스 맨 뒤에 값 추가
	c = append(c, 3)
	fmt.Println(cap(c))
	fmt.Println(c)

	// 슬라이스는 레퍼런스 타입임
	// d, e => 배열 // f,g => 슬라이스
	d := [3]int{1, 2, 3}
	var e [3]int

	e = d      // 배열의 요소가 모두 복사
	e[0] = 123 // e의 첫번째 요소만 바뀜
	fmt.Println(d)
	fmt.Println(e)

	f := []int{1, 2, 3}
	var g []int

	g = f      // 슬라이스는 요소가 아예 복사되는게 아니라 참조를 함
	g[0] = 123 // f[0]도 123으로 바뀜.
	fmt.Println(f)
	fmt.Println(g)

	// 슬라이스 복사하기
	// copy 함수 사용 (복사해서 대입하면, 원본을 참조하는게 아니라 아예 다른걸 만듬)
	// copy(복사될 슬라이스, 원본 슬라이스)
	x := []int{1, 2, 3, 4, 5}
	y := make([]int, 3) // make 함수로 공간 3 할당
	copy(y, x)

	fmt.Println(x) // [1 2 3 4 5]
	fmt.Println(y) // [1 2 3] 슬라이스 y의 길이는 3이므로 x의 요소 3개만 복사됨

	x2 := []int{1, 2, 3}
	y2 := make([]int, 3)
	copy(y2, x2) // 슬라이스를 복사함.
	y2[0] = 123  // y2[0]만 바뀌고, x2[0]은 안바뀜

	fmt.Println(x2)
	fmt.Println(y2)
}
