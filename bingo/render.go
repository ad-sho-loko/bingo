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
	objs := []RenderObject{
			&H1{},
	}
	return &RenderTree{
		RenderObjects:objs,
	}
}

func (r *RenderTree) walk(c walk.Container){
	// FIXME
	r.RenderObjects[0].Paint(c, PaintInfo{})
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
}

func (h *H1) Paint(c walk.Container, info PaintInfo) {
	walk.NewTextEdit(c)
}

type Paragraph struct{
	*InlineBlock
}

func (p *Paragraph) Paint(c walk.Container, info PaintInfo){
}

type PaintInfo struct{
}

