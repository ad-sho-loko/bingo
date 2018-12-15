package main

import (
	"testing"
)

func assertToken(expected *Token, actual *Token, tb testing.TB){
	if *expected != *actual{
		tb.Errorf("\n[expected] \n %s, \n[actual] \n %s", actual, expected)
	}
}

func assertTokenSlice(expected []*Token, actual []*Token, tb testing.TB){
	if len(expected) != len(actual){
		tb.Errorf("len is different.....")
	}
	for i:=0; i<len(expected); i++{
		assertToken(expected[i], actual[i], tb)
	}
}

func TestBracketToken(t *testing.T) {
	expected := []*Token{
		{kind: LeftBracket, value: ""},
		{kind: RightBracket, value: ""},
	}
	actual := NewLexer([]byte{'<', '>'}).tokenize()
	assertTokenSlice(expected, actual, t)
}

func TestLeftBracketToken(t *testing.T){
	expected := []*Token{
		{kind: LeftBracketWithSlash, value: ""},
		{kind: RightBracket, value: ""},
	}
	actual := NewLexer([]byte{'<', '/', '>'}).tokenize()
	assertTokenSlice(expected, actual, t)
}

func TestBracketWithSpaceToken(t *testing.T){
	expected := []*Token{
		{kind: LeftBracket, value: ""},
		{kind: Space, value: " "},
		{kind: RightBracket, value: ""},
	}
	actual := NewLexer([]byte{'<', ' ', '>'}).tokenize()
	assertTokenSlice(expected, actual, t)
}

func TestTextStringToken(t *testing.T){
	expected := []*Token{
		{kind: LeftBracket, value: ""},
		{kind: TextString, value: "h1"},
		{kind: RightBracket, value: ""},
	}
	actual := NewLexer([]byte{'<', 'h', '1', '>'}).tokenize()
	assertTokenSlice(expected, actual, t)
}

func TestStartEndToken(t *testing.T){
	expected := []*Token{
		{kind: LeftBracket, value: ""},
		{kind: TextString, value: "h1"},
		{kind: RightBracket, value: ""},
		{kind: LeftBracketWithSlash, value: ""},
		{kind: TextString, value: "h1"},
		{kind: RightBracket, value: ""},
	}
	actual := NewLexer([]byte{'<', 'h', '1', '>', '<', '/', 'h', '1', '>'}).tokenize()
	assertTokenSlice(expected, actual, t)
}
