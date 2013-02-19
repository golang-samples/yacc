package main

import (
    "fmt"
    "os"
)

type Lexer {}

func (l *Lexer) func Lex(lval *yySymType) int {
    
}

func (l *Lexer) func Error(e string) {
    fmt.Fprintln(os.Stderr, e)
}
