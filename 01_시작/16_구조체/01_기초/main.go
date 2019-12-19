package main

import "fmt"

// fmt 패키지 : 기본적인 입출력을 담당하는 패키지
// Golang은 기본적으로 세미콜론(;) 사용 안함

func main() {
	cd1 := Card{cdNo: "1029384719273847", cdImg: "https://blahblahblah.com/blahblah.jpg", higlnthCdYn: "Y"}
	cd2 := Card{cdNo: "3847291734826384", cdImg: "https://blahblahblah.com/blahblah2.jpg", higlnthCdYn: "N"}
	fmt.Println(cd1)
	fmt.Println(cd2)

	cd3 := NewCard("2903123437372626", "https://bah.com/bah.jpg", "N")
	cd3.higlnthCdYn = "Y"
	fmt.Println(cd3)
	cd3.destroy()
}

type Card struct {
	cdNo        string // 카드번호
	cdImg       string // 카드이미지
	higlnthCdYn string // 세로카드여부
}

func NewCard(cdNo, cdImg, higlnthCdYn string) *Card {
	return &Card{cdNo, cdImg, higlnthCdYn}
}

// 구조체에 메소드 연결하기
// func (리시버명 *구조체타입) 함수명() 리턴값_자료형 {}
// 리시버명 뒤 구조체 타입이 포인터형이면 원래의 값에 영향을 미침
// 포인터 형이 아니면, 원래의 값에 영향도 X
// 리시버 변수를 사용하지 않는다면 _로 생략 가능
func (card Card) destroy() {
	fmt.Println(card.cdNo, "번 카드를 터뜨렸습니다.")
}
