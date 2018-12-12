package main

type displayProp int

const(
	block displayProp = iota
	inline
	inlineBlock
)

var displayKind = map[string]displayProp{
	"h1":block,
	"p":block,
}



