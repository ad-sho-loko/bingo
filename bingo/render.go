// Copyright 2018 Shogo Arakawa. Released under the MIT license.

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

type RenderTree struct{
}

type RenderObect interface {
	Node() Node
	Render(paintInfo PaintInfo)
}

type h1 struct{
}

type p struct{
}

type PaintInfo struct{

}

