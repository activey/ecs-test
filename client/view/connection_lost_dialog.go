package view

import (
	"ecs-test/assets/fonts"
	"ecs-test/client/view/widgets"
	"github.com/tinne26/etxt"
	"github.com/yohamta/furex/v2"
	"image/color"
)

type ConnectionLostDialog struct {
	scale    float64
	renderer *etxt.Renderer

	rootView        *furex.View
	reconnectButton *ButtonView
	onReconnect     func()
	onBack          func()
}

func NewConnectionLostDialog(scale float64, onReconnect func(), onBack func()) *ConnectionLostDialog {
	r := etxt.NewRenderer()
	r.SetFont(fonts.MainFont)
	r.SetSize(32)
	r.SetColor(color.RGBA{
		R: 217,
		G: 179,
		B: 132,
		A: 255,
	})

	dialog := &ConnectionLostDialog{
		scale:       scale,
		renderer:    r,
		onReconnect: onReconnect,
		onBack:      onBack,
	}

	return dialog.initView()
}

func (d *ConnectionLostDialog) initView() *ConnectionLostDialog {
	d.rootView = &furex.View{
		Hidden:    true,
		Direction: furex.Column,
		Height:    160,

		AlignContent: furex.AlignContentCenter,
		Justify:      furex.JustifyStart,
		AlignItems:   furex.AlignItemCenter,
		Handler:      widgets.NewPanel(d.scale, 400),
	}

	d.initHeader()

	errorView := &furex.View{
		WidthInPct:   100,
		MarginTop:    20,
		MarginBottom: 15,
		Direction:    furex.Column,
		AlignContent: furex.AlignContentCenter,
		Justify:      furex.JustifyStart,
		AlignItems:   furex.AlignItemCenter,
	}
	errorView.AddChild(&furex.View{
		MarginRight: 10,
		Handler: widgets.NewText("Server connection lost!", d.renderer, d.scale).
			WithColor(color.RGBA{
				R: 218,
				G: 134,
				B: 62,
				A: 255,
			}),
	})

	buttonsView := &furex.View{
		Direction:  furex.Row,
		Justify:    furex.JustifyCenter,
		WidthInPct: 100,
		MarginTop:  10,
	}

	d.reconnectButton = NewButtonView(widgets.ButtonVariantGreen, "Reconnect", 2.0).
		OnClick(d.onReconnect).
		MarginRight(10)
	backToMenuButton := NewButtonView(widgets.ButtonVariantRed, "Back", 2.0).
		OnClick(d.onBack).
		Width(100)

	buttonsView.AddChild(d.reconnectButton.View())
	buttonsView.AddChild(backToMenuButton.View())

	errorView.AddChild(buttonsView)
	d.rootView.AddChild(errorView)
	return d
}

func (d *ConnectionLostDialog) initHeader() {
	topRow := &furex.View{
		Direction:    furex.Row,
		AlignContent: furex.AlignContentStart,
		Justify:      furex.JustifyStart,
		AlignItems:   furex.AlignItemStart,
		WidthInPct:   100,
	}

	header := &furex.View{
		Direction:    furex.Row,
		Handler:      widgets.NewHeader(d.scale),
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
			NewText("Oh no :(", d.renderer, d.scale).
			WithColor(color.RGBA{
				R: 165,
				G: 48,
				B: 48,
				A: 255,
			}).
			WithVerticalShift(-3),
	})

	topRow.AddChild(header)
	d.rootView.AddChild(topRow)
}

func (d *ConnectionLostDialog) Show() {
	d.rootView.Hidden = false
}

func (d *ConnectionLostDialog) Hide() {
	d.rootView.Hidden = true
}

func (d *ConnectionLostDialog) View() *furex.View {
	return d.rootView
}

func (d *ConnectionLostDialog) DisableRetryButton() {
	d.reconnectButton.Disable()
}

func (d *ConnectionLostDialog) EnableRetryButton() {
	d.reconnectButton.Enable()
}
