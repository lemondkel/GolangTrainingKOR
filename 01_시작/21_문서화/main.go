package _1_기본

type Data struct {
	a, b int
}

func main() {

}

func Sum(a int, b int) int {
	return a + b
}

// CLI에서 godoc -http=:6060 실행하면 6060포트로 문서화된 소스 웹으로 확인 가능.
