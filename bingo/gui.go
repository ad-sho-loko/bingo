// Copyright 2018 Shogo Arakawa. Released under the MIT license.

package main

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"log"
)

type RenderEngine struct{
	window *walk.MainWindow
	widgets *walk.CustomWidget
	renderTree *RenderTree
}

func NewRenderEngine(tree *RenderTree) *RenderEngine{
	return &RenderEngine{
		renderTree:tree,
	}
}

func (e *RenderEngine) paintAll(){
	e.renderTree.walk(e.widgets.Parent())
}

func (e *RenderEngine) run(){
	var inTE *walk.TextEdit
	const (
		maxHeight = 400
		maxWidth  = 600
	)

	_, err := (MainWindow{
		AssignTo:&e.window,
		Title:   "bingo",
		MinSize: Size{Height: maxHeight, Width: maxWidth},
		Layout:  VBox{},
		Children: []Widget{
			HSplitter{
				MaxSize: Size{Height: maxHeight / 20, Width: maxWidth},
				MinSize: Size{Height: maxHeight / 20, Width: maxWidth},
				Children: []Widget{
					TextEdit{
						StretchFactor: 100,
						Text:          "example/index.html", AssignTo: &inTE,
					},
					PushButton{
						Text: "Go",
						OnClicked: func() {

							clicked()
							/*
							var _, err = ioutil.ReadFile(inTE.Text())
							if err != nil {
								fmt.Printf("Cannot find %s", inTE.Text())
							}*/

							// Parse 3 files
							// go html parser
							// go css parser
							// go js parser(lexer)

							// construct rendering tree
							// eg1. h1 tag, content = hello

							// Render by walking rendering tree
							e.paintAll()
						},
					},
				},
			},
			// Custom Widget
			CustomWidget{
				AssignTo:&e.widgets,
			},
		},
	}).Run()

	if err != nil{
		log.Fatal(err)
	}
}