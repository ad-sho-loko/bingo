package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	// [HTML]
	// Read html file.
	f, err := ioutil.ReadFile("./test/test1.html")
	if err != nil {
		log.Fatal(err)
	}

	// Tokenize html
	p := NewParser()
	tokens := p.tokenize(f)
	/*
	for _, t := range tokens {
		fmt.Print(*t)
	}*/

	// Parse Tokens
	nodeTree := p.parse(tokens)
	/*
	for _, n := range nodes {fmt.Print(n)
	}
	*/

	nodeTree.Walk(func(n Node){
		fmt.Println(n)
	})

	// [CSS]
	// Read css file
	// Tokenize css
	// Parse Tokens

	// [JavaScript]
	// Read JavaScript
	// Tokenize
	// Make AST
	// Generate VM Code
	// Execute?

	// Make Rendering tree by using dom tree.


	// [Rendering]
	// Finally, Rendering by walking rendering tree.
}
