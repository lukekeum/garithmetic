package lexer

import (
	"fmt"

	"github.com/lukekeum/garithmetic/store"
)

func Execute(content string, tstore []store.Token) []store.Token {
	line := 1
	bracketNum := 0

	for i := 0; i < len(content); i++ {

		if content[i] == '\n' {

			if tstore[len(tstore)-1].Value != ';' {
				panic(fmt.Sprintf("[Error] Expected ; but not found, line number %d", line))
			}

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

			tstore = append(tstore, *store.New(store.CONST, line, value))
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

			tstore = append(tstore, *store.New(store.OPER, line, content[i]))
		} else if content[i] == '(' {
			fmt.Printf("[OPER] %c \n", content[i])

			tstore = append(tstore, *store.New(store.OPER, line, '('))
			bracketNum += 1
		} else if content[i] == ')' {
			fmt.Printf("[OPER] %c \n", content[i])
			tstore = append(tstore, *store.New(store.OPER, line, ')'))
			bracketNum -= 1
		} else if content[i] == ';' {
			fmt.Printf("[SEPER] %c \n", content[i])
			tstore = append(tstore, *store.New(store.SEPER, line, ';'))
		} else if content[i] == ' ' {
			continue
		} else {
			panic(fmt.Sprintf("[Error] Unexpected token scanned, line number %d, found: %c", line, content[i]))
		}
	}

	return tstore
}
