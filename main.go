package main

import (
"fmt"
	"github.com/google/gxui"
	"github.com/google/gxui/drivers/gl"
	"github.com/google/gxui/math"
	"github.com/google/gxui/samples/flags"
	"github.com/google/gxui/themes/dark"
)

func appMain(driver gxui.Driver) {
	theme := dark.CreateTheme(driver)

	window := theme.CreateWindow(500, 250, "goText")
	window.OnClose(driver.Terminate)
	window.SetScale(flags.DefaultScaleFactor)
	window.SetPadding(math.Spacing{L: 10, R: 10, T: 10, B: 10})



	text := theme.CreateCodeEditor()
	text.SetMultiline(true)
	text.SetDesiredWidth(400)
	text.SetText("text := theme.CreateCodeEditor()\ntext.SetMultiline(true)\ntext.SetDesiredWidth(400)\n fmt.Println(\"BOOM\")")
	/*l1 := gxui.CreateCodeSyntaxLayer()
	l1.SetColor(gxui.Yellow)

	l1.AddData(0, 456, "hello")
	var layers gxui.CodeSyntaxLayers
	layers = append(layers, l1)
	text.SetSyntaxLayers(layers)
	text.OnTextChanged(func(e []gxui.TextBoxEdit) {
		text.SetSyntaxLayers(layers)
		fmt.Println("Changed")
    })*/

	layout := theme.CreateLinearLayout()
	layout.SetSizeMode(gxui.Fill)
	layout.Direction().TopToBottom()
	layout.AddChild(createToolBar(theme))
	layout.AddChild(text)
	window.AddChild(layout)
}



func createToolBar(theme gxui.Theme) gxui.LinearLayout {
	button := theme.CreateButton()

	click := func() {
		fmt.Println("BOOM")
	}
	button.SetText("New File")
	button.OnClick(func(gxui.MouseEvent) { click() })

	button1 := theme.CreateButton()
	click1 := func() {
		fmt.Println("BOOM1")
	}
	button1.SetText("Open File")
	button1.OnClick(func(gxui.MouseEvent) { click1() })

	layout := theme.CreateLinearLayout()
	layout.SetDirection(gxui.LeftToRight)

	layout.AddChild(button)
	layout.AddChild(button1)
	return layout
}

func main() {
	gl.StartDriver(appMain)
}
