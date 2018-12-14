// Copyright 2018 Shogo Arakawa. Released under the MIT license.

package main

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"log"
)

type RenderEngine struct{
	window *walk.MainWindow
	renderTree *RenderTree
}

func (e *RenderEngine) paintAll(){
	// item := H1{}.Paint(PaintInfo{})
}

func NewRenderEngine() *RenderEngine{
	return nil // todo
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
						Text:          "test/index.html", AssignTo: &inTE,
					},
					PushButton{
						Text: "Go",
						OnClicked: func() {
							/*
							var _, err = ioutil.ReadFile(inTE.Text())
							if err != nil {
								fmt.Printf("Cannot find %s", inTE.Text())
							}*/

							// mock
							e.paintAll()

							// Parse 3 files
							// go html parser
							// go css parser
							// go js parser(lexer)

							// construct rendering tree

							// eg1. h1 tag, content = hello

							// TODO
							// h1 tag -> TextWidget(h1用のTextWidgetが必要か)にマッピング
							// TextWidgetのTextに`Hello`を挿入

							// Render by walking rendering tree

						},
					},
				},
			},
		},
	}).Run()

	if err != nil{
		log.Fatal(err)
	}
}