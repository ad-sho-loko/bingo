package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

type Node interface {
	Children() []*Node
	AddChild(child *Node)
}

// Root node in dom tree.
type DocumentNode struct {
	children []*Node
}

func (n *DocumentNode) Children() []*Node {
	return n.children
}

func (n *DocumentNode) AddChild(child *Node) {
	n.children = append(n.children, child)
}

type TagNode struct {
	name     string
	elements []Element
	children []*Node
}

func (n *TagNode) Children() []*Node {
	return n.children
}

func (n *TagNode) AddChild(child *Node) {
	n.children = append(n.children, child)
}

type TextNode struct {
	value    string
	children []*Node
}

func (n *TextNode) Children() []*Node {
	return n.children
}

func (n *TextNode) AddChild(child *Node) {
	n.children = append(n.children, child)
}

type Element struct {
	attribute map[string]string
}

type CommentNode struct{
	value string
	/* children []*Node */
}

type Parser struct {
	parentNodeStack []*Node
}

func NewParser() *Parser {
	var top Node = &DocumentNode{}
	return &Parser{parentNodeStack: []*Node{
		&top,
	}}
}

func (p *Parser) currentParentNode() *Node {
	return p.parentNodeStack[0]
}

func (p *Parser) pushNode(node *Node) {
	p.parentNodeStack = append([]*Node{node}, p.parentNodeStack...)
}

func (p *Parser) popNode() *Node {
	n := p.parentNodeStack[0]
	if len(p.parentNodeStack) > 1 {
		p.parentNodeStack = p.parentNodeStack[1:]
	}
	return n
}

func (p *Parser) parse(tokens []*Token) []*Node {
	var nodes []*Node
	for i := 0; i < len(tokens); i++ {
		switch tokens[i].kind {
		case LeftBracket:
			node := p.parseTag(tokens, &i)
			(*p.currentParentNode()).AddChild(node)
			p.pushNode(node)
			nodes = append(nodes, node)

		case LeftBracketWithSlash:
			p.parseTag(tokens, &i)
			p.popNode()

		case Text, Space:
			node := &TextNode{value: "", children: []*Node{}}
			for ; !(tokens[i].kind == LeftBracket || tokens[i].kind == LeftBracketWithSlash); i++ {
				node.value += tokens[i].value
			}
			i--
			// これどうにかできないかな...
			var n Node = node
			(*p.currentParentNode()).AddChild(&n)
			nodes = append(nodes, &n)

		default:
			// FIXME: Skip tab, LF, CR...(if you need to parse, add case of Text, Space....
		}
	}
	return nodes
}

func (p *Parser) parseTag(tokens []*Token, i *int) *Node {
	*i++
	skipSpace(tokens, i)

	if tokens[*i].kind != Text {
		panic("The next token of Left bracket should be string.")
	}

	// set tag name
	var node Node = &TagNode{
		name:     tokens[*i].value,
		children: []*Node{},
	}

	for tokens[*i].kind != RightBracket {
		*i++
		// TODO : implement attribute.
	}

	return &node
}

func skipSpace(tokens []*Token, i *int) {
	for tokens[*i].kind == Space {
		*i++
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
		case (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z'):
			// TODO : how to handle the symbol of number in string?
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

func main() {
	// [HTML]
	// Read html file.
	f, err := ioutil.ReadFile("test1.html")
	if err != nil {
		log.Fatal(err)
	}

	// Tokeninze html
	p := NewParser()
	tokens := p.tokenize(f)
	for _, t := range tokens {
		fmt.Print(*t)
	}

	// Parse Tokens
	nodes := p.parse(tokens)
	for _, n := range nodes {
		fmt.Print(*n)
	}

	// [CSS]
	// Read css file
	// Tokenize css
	// Parse Tokens

	// [JavaScript]
	// Read JavaScript
	// Tokenize
	// Make AST
	// Generate VM Code
	// Execute in

	// [Rendering]
	// Finally, Rendering.
}
