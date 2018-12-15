package main

import "testing"

func assertNode(expected Node, actual Node, tb testing.TB){
	if expected != actual{
		tb.Errorf("\n[expected] \n %s, \n[actual] \n %s", actual, expected)
	}
}

func assertNodeSlice(expected []Node, actual []Node, tb testing.TB){
	if len(expected) != len(actual){
		tb.Errorf("len is different.....")
	}
	for i:=0; i<len(expected); i++{
		assertNode(expected[i], actual[i], tb)
	}
}

func TestBracket(t *testing.T) {
	// expected := []Node{
	// }

	// actual := NewLexer().tokenize([]byte{'<', '>'})
	// assertNode(expected, actual, t)
}