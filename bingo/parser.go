// Copyright 2018 Shogo Arakawa. Released under the MIT license.

// Implements the html tokenizer/parser that creates NodeTree for rendering in the bingo browser.

package main

import "fmt"

type NodeKind int

const(
	Document NodeKind = iota
	Element
	Text
)

type Node interface {
	Children() []Node
	AddChild(child Node)
	Kind() NodeKind
}

// Root node in dom tree.
type DocumentNode struct {
	children []Node
}

func (n *DocumentNode) Children() []Node {
	return n.children
}

func (n *DocumentNode) AddChild(child Node) {
	n.children = append(n.children, child)
}

func (n *DocumentNode) Kind() NodeKind{
	return Document
}

type ElementNode struct {
	name     string
	attribute map[string]string
	children []Node
}

func (n *ElementNode) Children() []Node {
	return n.children
}

func (n *ElementNode) AddChild(child Node) {
	n.children = append(n.children, child)
}

func (n *ElementNode) Kind() NodeKind{
	return Element
}

func (n *ElementNode) String() string{
	return fmt.Sprintf("ElementNode[%p]%+v\n", n, *n)
}

type TextNode struct {
	value    string
	children []Node
}

func (n *TextNode) Children() []Node {
	return n.children
}

func (n *TextNode) AddChild(child Node) {
	n.children = append(n.children, child)
}

func (n *TextNode) Kind() NodeKind{
	return Text
}

func (n *TextNode) String() string{
	return fmt.Sprintf("TextNode[%p]%+v\n", n, *n)
}

/*
type CommentNode struct{
	value string
}
*/

type Parser struct {
	pos int
	parentNodeStack []Node
}

func NewParser() *Parser {
	var top Node = &DocumentNode{}
	return &Parser{
		pos:0,
		parentNodeStack: []Node{top},
	}
}

func (p *Parser) currentParentNode() Node {
	return p.parentNodeStack[0]
}

func (p *Parser) pushNode(node Node) {
	p.parentNodeStack = append([]Node{node}, p.parentNodeStack...)
}

func (p *Parser) popNode() Node {
	n := p.parentNodeStack[0]
	if len(p.parentNodeStack) > 1 {
		p.parentNodeStack = p.parentNodeStack[1:]
	}
	return n
}

func (p *Parser) parse(tokens []*Token) *NodeTree {
	var nodes []Node
	for ; p.pos < len(tokens); p.pos++{
	switch tokens[p.pos].kind {
		case LeftBracket:
			node := p.parseTag(tokens)
			p.currentParentNode().AddChild(node)
			p.pushNode(node)
			nodes = append(nodes, node)

		case LeftBracketWithSlash:
			p.parseTag(tokens)
			p.popNode()

		case TextString, Space:
			node := &TextNode{value: "", children: []Node{}}
			for ; !(tokens[p.pos].kind == LeftBracket || tokens[p.pos].kind == LeftBracketWithSlash); p.pos++ {
				node.value += tokens[p.pos].value
			}
			p.pos--
			p.currentParentNode().AddChild(node)
			nodes = append(nodes, node)

		default:
			// FIXME: Skip tab, LF, CR...(if you need to parse, add case of TextString, Space....
		}
	}
	return &NodeTree{nodeList:nodes}
}

func (p *Parser) parseTag(tokens []*Token) Node {
	p.pos++
	p.skipSpace(tokens)

	if tokens[p.pos].kind != TextString {
		panic("The next token of Left bracket should be string.")
	}

	// set tag name
	var node Node = &ElementNode{
		name:     tokens[p.pos].value,
		children: []Node{},
	}

	for tokens[p.pos].kind != RightBracket {
		p.pos++
		// TODO : implement attribute.
	}
	return node
}

func (p *Parser) skipSpace(tokens []*Token) {
	for tokens[p.pos].kind == Space {
		p.pos++
	}
}

type Lexer struct{
}

func NewLexer() *Lexer{
	return &Lexer{}
}

type Token struct {
	kind  tokenType
	value string
}

type tokenType int

const (
	Space tokenType = iota
	LeftBracket
	RightBracket
	LeftBracketWithSlash
	TextString
)

func (t tokenType) String() string {
	switch t {
	case Space:
		return "Space"
	case LeftBracket:
		return "LeftBracket"
	case RightBracket:
		return "RightBracket"
	case LeftBracketWithSlash:
		return "LeftBracketWithSlash"
	case TextString:
		return "TextString"
	default:
		return "Others"
	}
}

func newToken(k tokenType, v string) *Token {
	return &Token{
		kind:  k,
		value: v,
	}
}

// Impl to tokenize just ASCII code. not able to read 日本語(Unicode).
func (t *Lexer) tokenize(bytes []byte) []*Token {
	var tokens []*Token
	var buf []byte
	for i, b := range bytes {
		switch {
		// case b == '/' && bytes[i+1] == '*':
		//	outIfBufExist()
		//  tokens = Comment!
		case b == ' ':
			t.outIfBufExist(&tokens, &buf)
			tokens = append(tokens, newToken(Space, " "))
		case b == '<':
			t.outIfBufExist(&tokens, &buf)
			if bytes[i+1] == '/' {
				tokens = append(tokens, newToken(LeftBracketWithSlash, ""))
			} else {
				tokens = append(tokens, newToken(LeftBracket, ""))
			}
		case b == '>':
			t.outIfBufExist(&tokens, &buf)
			tokens = append(tokens, newToken(RightBracket, ""))
		case (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z') || (b >= '0' && b <= '9'):
			buf = append(buf, b)
		default:
			// do nothing
		}
	}
	return tokens
}

func (t *Lexer) outIfBufExist(tokens *[]*Token, buf *[]byte) bool {
	if len(*buf) != 0 {
		*tokens = append(*tokens, newToken(TextString, string(*buf)))
		*buf = nil
		return true
	}
	return false
}

type NodeTree struct {
	nodeList []Node
}

func (t *NodeTree) bodyNodeList() []Node{
	// find body element.
	for i, n := range t.nodeList{
		if n.Kind() == Element{
			return t.nodeList[i:]
		}
	}
	// error
	return nil
}

func (t *NodeTree) Walk(f func(me Node)){
	t.walkMap(t.nodeList[0], f)
}

func (t *NodeTree) walkRecursive(node Node) {
	for _, n := range node.Children(){
		t.walkRecursive(n)
	}
}

func (t *NodeTree) walkMap(node Node, f func(me Node)){
	f(node)
	for _, child := range node.Children(){
		t.walkMap(child, f)
	}
}

func (t *NodeTree) Print(){
	t.Walk(func(child Node){
		fmt.Print(child)
	})
}