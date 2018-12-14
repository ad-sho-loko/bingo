// Copyright 2018 Shogo Arakawa. Released under the MIT license.

package main

type RenderTree struct{
	RenderObjects []RenderObject
}

func NewRenderTree(nodeTree *NodeTree) *RenderTree{
	return &RenderTree{
	}
}

func (t *RenderTree) PaintAll(){
}

// RenderObject.
// eg) h1 tag implements 2 interfaces which are RenderObject and Block.
// RenderObject <-- Block <-- h1
type RenderObject interface {
	Node() Node

	// paintInfoにはタグの属性(h1,pとか)、中身（helloとか）、cssスタイル(displayとか）が詰められている.
	// それを参考にGUIライブラリ側にマッピングをしていく作業が必要になる.
	Render(paintInfo PaintInfo)
}

// Style of RenderObject Collection.

type InlineBlock struct{
	node Node
}

func (b *InlineBlock) Node() Node{
	return b.node
}

func (b *InlineBlock) Render(paintInfo PaintInfo){
}

func NewInlineBlock(node Node) RenderObject{
	return &InlineBlock{
		node:node,
	}
}

type Block struct{
}

// Element Tag Collections.

type H1 struct{
}

type Paragraph struct{
}

type PaintInfo struct{
}

