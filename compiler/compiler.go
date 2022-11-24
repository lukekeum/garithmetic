package compiler

import (
	"fmt"
	"os"

	"github.com/lukekeum/garithmetic/compiler/lexer"
	"github.com/lukekeum/garithmetic/compiler/parser"
	"github.com/lukekeum/garithmetic/store"
)

// 최초의 컴파일러 실행 함수
// 인자로 파일 이름을 받아옴
func Execute(fileName string) int {

	fmt.Println("Compiler: Start Compiling")

	// fileName의 파일 읽어오기
	data, err := os.ReadFile(fileName)

	// 읽어오는 데 오류가 있으면 오류 발생 후 프로그램 종료
	if err != nil {
		panic("[Error] Error occured while reading file")
	}

	content := string(data[:])
	store := []store.Token{}

	// 렉서 실행
	store = lexer.Execute(content)

	// 파서 실행
	p := parser.New(store)

	p.Execute()

	// 실행 종료
	return 0
}
