package main

import (
	"fmt"
)

type Node interface {
	Children() []Node
	AddChild(child Node)
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

type ElementNode struct {
	name     string
	attribute map[string]string
	children []Node
}

func (n *ElementNode) String() string{
	return fmt.Sprintf("ElementNode[%p]%+v\n", n, *n)
}

func (n *ElementNode) Children() []Node {
	return n.children
}

func (n *ElementNode) AddChild(child Node) {
	n.children = append(n.children, child)
}

type TextNode struct {
	value    string
	children []Node
}

func (n *TextNode) String() string{
	return fmt.Sprintf("TextNode[%p]%+v\n", n, *n)
}

func (n *TextNode) Children() []Node {
	return n.children
}

func (n *TextNode) AddChild(child Node) {
	n.children = append(n.children, child)
}

type CommentNode struct{
	value string
	/* children []*Node */
}

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

		case Text, Space:
			node := &TextNode{value: "", children: []Node{}}
			for ; !(tokens[p.pos].kind == LeftBracket || tokens[p.pos].kind == LeftBracketWithSlash); p.pos++ {
				node.value += tokens[p.pos].value
			}
			p.pos--
			p.currentParentNode().AddChild(node)
			nodes = append(nodes, node)

		default:
			// FIXME: Skip tab, LF, CR...(if you need to parse, add case of Text, Space....
		}
	}
	return &NodeTree{nodeList:nodes}
}

func (p *Parser) parseTag(tokens []*Token) Node {
	p.pos++
	p.skipSpace(tokens)

	if tokens[p.pos].kind != Text {
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
	Text
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
	case Text:
		return "Text"
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
func (p *Parser) tokenize(bytes []byte) []*Token {
	var tokens []*Token
	var buf []byte
	for i, b := range bytes {
		switch {
		// case b == '/' && bytes[i+1] == '*':
		//	outIfBufExist()
		//  tokens = Comment!
		case b == ' ':
			outIfBufExist(&tokens, &buf)
			tokens = append(tokens, newToken(Space, " "))
		case b == '<':
			outIfBufExist(&tokens, &buf)
			if bytes[i+1] == '/' {
				tokens = append(tokens, newToken(LeftBracketWithSlash, ""))
			} else {
				tokens = append(tokens, newToken(LeftBracket, ""))
			}
		case b == '>':
			outIfBufExist(&tokens, &buf)
			tokens = append(tokens, newToken(RightBracket, ""))
		case (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z') || (b >= '0' && b <= '9'):
			buf = append(buf, b)
		default:
			// do nothing
		}
	}
	return tokens
}

func outIfBufExist(tokens *[]*Token, buf *[]byte) bool {
	if len(*buf) != 0 {
		*tokens = append(*tokens, newToken(Text, string(*buf)))
		*buf = nil
		return true
	}
	return false
}

type NodeTree struct {
	nodeList []Node
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



