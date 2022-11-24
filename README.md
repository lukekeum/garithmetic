<center><img width="240px" src="https://www.clipartmax.com/png/middle/353-3536792_go-golang-logo-png.png" /></center>
<center>고 언어로 구현한 사칙연산 컴파일러</center>

<br />

# Tech Stack

+ Golang
+ Stack
+ EBNF
+ top-down parsing

# EBNF

```shell
<S> ::= <A>’;’
<A> ::= <M> { (‘+’ | ‘-’)<M> }
<M> ::= <F> { (‘*’ | ‘/’)<F> }
<F> ::= <N> | ‘(‘<N>’)’
<N> ::= 0 | 1 | 2 | … | 9
```

# How To Run

1. Make Arithmetic File with EBNF grammer correctly
2. Run command in root directory
```
go run main.go {FILE_LOCATION}
```

# References

- Book ["컴파일러의 이해"](http://www.yes24.com/Product/Goods/24330311)
- Got Idea From [here](https://github.com/hsnks100/dreampiler)
