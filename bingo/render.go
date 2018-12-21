// Copyright 2018 Shogo Arakawa. Released under the MIT license.

package main

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

type RenderTree struct{
	RenderObjects []RenderObject
}

func MakeRenderTree(nodeTree *NodeTree) *RenderTree{
	// skipByBody
	var objs []RenderObject
	nodeTree.Walk(func(me Node) {
		if me.Kind() == Element{
			e := me.(*ElementNode)
			switch e.name {
			case "h1":
				objs = append(objs, NewH1(*e))
			}
		}
	})

	return &RenderTree{
		RenderObjects:objs,
	}
}

func (r *RenderTree) walk(c walk.Container){
	r.walkPaint(c, r.RenderObjects[0])
}

func (r *RenderTree) walkPaint(c walk.Container, obj RenderObject){
	obj.Paint(c, PaintInfo{})
	for _, child := range obj.Children(){
		r.walkPaint(c, child)
	}
}

// RenderObject.
// eg) h1 element implements 2 interfaces which are RenderObject and Block.
// RenderObject <-- Block <-- h1
type RenderObject interface {
	Node() Node
	Children() []RenderObject
	Paint(c walk.Container, info PaintInfo)
}

// Style of RenderObject Collection.

type InlineBlock struct{
	node Node
}

func (b *InlineBlock) Node() Node{
	return b.node
}

func (b *InlineBlock) Children() []RenderObject{
	return nil
}

func (b *InlineBlock) Splitter() *HSplitter{
	return &HSplitter{}
}

func NewInlineBlock(node Node) *InlineBlock{
	return &InlineBlock{
		node:node,
	}
}

type Block struct{
}

// Element Tag Collections.

type H1 struct{
	*InlineBlock
	text string
}

func NewH1(n ElementNode) *H1{
	text := ""
	for _, child := range n.Children(){
		if child.Kind() == Text{
			text += child.(*TextNode).value
		}
	}
	return &H1{
		text:text,
	}
}

func (h *H1) Paint(c walk.Container, info PaintInfo) {
	l, _ := walk.NewLabel(c)
	l.SetText(h.text)
}

type Paragraph struct{
	*InlineBlock
}

func (p *Paragraph) Paint(c walk.Container, info PaintInfo){
}

type PaintInfo struct{
}

