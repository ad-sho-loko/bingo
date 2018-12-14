package main

import (
	"fmt"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"io/ioutil"
	"log"
)

type RenderEngine struct{
	renderTree *RenderTree
}

func (e *RenderEngine) paintAll() CustomWidget{
	// e.renderTree.PaintAll()
	return CustomWidget{}
}

func (e *RenderEngine) run(){
	var inTE *walk.TextEdit
	const (
		maxHeight = 400
		maxWidth  = 600
	)

	if _, err := (MainWindow{
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
							var _, err = ioutil.ReadFile(inTE.Text())
							if err != nil {
								fmt.Printf("Cannot find %s", inTE.Text())
							}
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

			// Rendering Area of node tree. better Canvas??
			CustomWidget{

			},
		},
	}.Run()); err != nil{
		log.Fatal(err)
	}
}