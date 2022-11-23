package lexer

import (
	"fmt"
	"strconv"

	"github.com/lukekeum/garithmetic/store"
)

func strToInt(s string) int {
	v, _ := strconv.Atoi(s)
	return v
}

func Execute(content string) []store.Token {
	line := 1
	parsingLine := 1
	bracketNum := 0

	tstore := []store.Token{}

	for i := 0; i < len(content); i++ {

		if content[i] == '\n' {

			if bracketNum != 0 {
				panic(fmt.Sprintf("[Error] Bracket expected token ), line number %d", line))
			}

			line += 1
			bracketNum = 0
		} else if content[i] >= '0' && content[i] <= '9' {
			value := fmt.Sprintf("%c", content[i])

			for true {
				if content[i+1] >= '0' && content[i+1] <= '9' {
					value += fmt.Sprintf("%c", content[i])
				} else {
					break
				}
				i += 1
			}

			fmt.Printf("[CONST] %s \n", value)

			tstore = append(tstore, *store.New(store.CONST, line, parsingLine, strToInt(value)))
		} else if content[i] == '/' {
			if content[i+1] == '/' {
				for true {
					if i+1 >= len(content) {
						break
					}
					if content[i+1] == '\n' {
						break
					}
					i += 1
				}
			}
		} else if content[i] == '+' || content[i] == '-' || content[i] == '*' || content[i] == '/' {
			fmt.Printf("[OPER] %c \n", content[i])

			tstore = append(tstore, *store.New(store.OPER, line, parsingLine, int(rune(content[i]))))
		} else if content[i] == '(' {
			fmt.Printf("[OPER] %c \n", content[i])

			tstore = append(tstore, *store.New(store.OPER, line, parsingLine, int('(')))
			bracketNum += 1
		} else if content[i] == ')' {
			fmt.Printf("[OPER] %c \n", content[i])
			tstore = append(tstore, *store.New(store.OPER, line, parsingLine, int(')')))
			bracketNum -= 1
		} else if content[i] == ';' {
			fmt.Printf("[SEPER] %c \n", content[i])
			tstore = append(tstore, *store.New(store.SEPER, line, parsingLine, int(';')))
			parsingLine += 1
		} else if content[i] == ' ' {
			continue
		} else {
			panic(fmt.Sprintf("[Error] Unexpected token scanned, line number %d, found: %c", line, content[i]))
		}
	}

	return tstore
}
