package view

import (
	"ecs-test/assets/fonts"
	"ecs-test/client/view/widgets"
	"github.com/tinne26/etxt"
	"github.com/yohamta/furex/v2"
	"image/color"
)

type ErrorDialog struct {
	scale    float64
	renderer *etxt.Renderer

	rootView *furex.View
}

func NewErrorDialog(scale float64) *ErrorDialog {
	r := etxt.NewRenderer()
	r.SetFont(fonts.MainFont)
	r.SetSize(32)
	r.SetColor(color.RGBA{
		R: 217,
		G: 179,
		B: 132,
		A: 255,
	})

	dialog := &ErrorDialog{
		scale:    scale,
		renderer: r,
	}

	return dialog.init()
}

func (e *ErrorDialog) init() *ErrorDialog {
	e.rootView = &furex.View{
		Hidden:    true,
		Direction: furex.Column,
		Height:    100,

		AlignContent: furex.AlignContentCenter,
		Justify:      furex.JustifyStart,
		AlignItems:   furex.AlignItemCenter,
		Handler:      widgets.NewPanel(e.scale, 400),
	}

	e.initHeader()

	errorRow := &furex.View{
		MarginTop:    10,
		Direction:    furex.Row,
		AlignContent: furex.AlignContentCenter,
		Justify:      furex.JustifyStart,
		AlignItems:   furex.AlignItemCenter,
	}
	errorRow.AddChild(&furex.View{
		MarginRight: 10,
		Handler: widgets.NewText("An error occurred", e.renderer, e.scale).
			WithColor(color.RGBA{
				R: 218,
				G: 134,
				B: 62,
				A: 255,
			}),
	})

	buttonView := NewButtonView(widgets.ButtonVariantRed, "Dang it", 2.0).
		Width(100).
		OnClick(e.Hide)

	errorRow.AddChild(buttonView.View())
	e.rootView.AddChild(errorRow)
	return e
}

func (e *ErrorDialog) initHeader() {
	topRow := &furex.View{
		Direction:    furex.Row,
		AlignContent: furex.AlignContentStart,
		Justify:      furex.JustifyStart,
		AlignItems:   furex.AlignItemStart,
		WidthInPct:   100,
	}

	header := &furex.View{
		Direction:    furex.Row,
		Handler:      widgets.NewHeader(e.scale),
		MarginLeft:   8,
		MarginRight:  8,
		MarginTop:    8,
		WidthInPct:   100,
		AlignContent: furex.AlignContentCenter,
		Justify:      furex.JustifyCenter,
		AlignItems:   furex.AlignItemCenter,
	}
	header.AddChild(&furex.View{
		Handler: widgets.
			NewText("Error", e.renderer, e.scale).
			WithColor(color.RGBA{
				R: 165,
				G: 48,
				B: 48,
				A: 255,
			}).
			WithVerticalShift(-3),
	})

	topRow.AddChild(header)
	e.rootView.AddChild(topRow)
}

func (e *ErrorDialog) Show() {
	e.rootView.Hidden = false
}

func (e *ErrorDialog) Hide() {
	e.rootView.Hidden = true
}

func (e *ErrorDialog) View() *furex.View {
	return e.rootView
}
