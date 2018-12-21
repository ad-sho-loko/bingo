// Copyright 2018 Shogo Arakawa. Released under the MIT license.

package main

import (
	"io/ioutil"
	"log"
)

// when clicked,
func clicked(){
	// request("")
}

func main() {
	// GUI RUN

	// if clicked..
	// [HTTP]
	// send request
	// receive response
	// extract HTML, CSS, JavaScript

	// [HTML]
	// Read html file.
	f, err := ioutil.ReadFile("./example/test1.html")
	if err != nil {
		log.Fatal(err)
	}

	// Tokenize html
	l := NewLexer(f)
	tokens := l.tokenize()

	// Parse Tokens
	p := NewParser()
	nodeTree := p.parse(tokens)
	nodeTree.Print()

	// [CSS]
	// Read css file
	// Tokenize css
	// Parse Tokens

	// [Attach html, css for node]
	// attach html & css

	// [JavaScript]
	// Read JavaScript
	// Tokenize
	// Make AST
	// Generate VM Code
	// Execute?

	// Make Rendering tree by using dom tree.

	// [Rendering]
	// Finally, Rendering by walking rendering tree.
	// r := NewRenderTree(nodeTree)
	// r.PaintAll()

	// Bridge of GUI and RenderTree
	// need to implement running render engine async.
	renderTree := MakeRenderTree(nodeTree)
	engine := NewRenderEngine(renderTree)
	engine.run()
}
