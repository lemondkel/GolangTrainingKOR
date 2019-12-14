package main

import (
	"fmt"
	"io/ioutil"
)

// fmt 패키지 : 기본적인 입출력을 담당하는 패키지

func main() {
	var b []byte
	var err error

	b, err = ioutil.ReadFile("01_시작/04_조건문/02_if문_예시/hello.txt")
	if err == nil {
		fmt.Println("%s", b)
	}
}
