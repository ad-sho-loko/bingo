package main

import (
	"fmt"
	"log"
	"testing"
)


func TestLeftBracket(t *testing.T) {
	l := NewLexer()
	actual := l.tokenize([]byte{'<', '>'})

	expected := []*Token{
		&Token{kind:LeftBracket, value:""},
		&Token{kind:RightBracket, value:""},
	}

	if actual[0] != expected[0]{
		fmt.Println(actual[0])
		fmt.Println(expected[0])
		log.Fatal("Ops!!!")
	}
}
