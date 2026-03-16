package view

import (
	"ecs-test/assets/fonts"
	"ecs-test/client/view/widgets"
	"github.com/tinne26/etxt"
	"github.com/yohamta/furex/v2"
	"image/color"
)

var (
	redButtonTextColor = color.RGBA{
		R: 218,
		G: 134,
		B: 62,
		A: 255,
	}

	greenButtonTextColor = color.RGBA{
		R: 209,
		G: 251,
		B: 147,
		A: 255,
	}
)

type ButtonView struct {
	variant   widgets.ButtonVariant
	labelText string
	scale     float64
	disabled  bool

	renderer    *etxt.Renderer
	buttonView  *furex.View
	button      *widgets.Button
	buttonLabel *widgets.Text
}

func NewButtonView(
	variant widgets.ButtonVariant,
	labelText string,
	scale float64,
) *ButtonView {
	r := etxt.NewRenderer()
	r.SetFont(fonts.MainFont)
	r.SetSize(32)

	bv := &ButtonView{
		variant:   variant,
		labelText: labelText,
		scale:     scale,
		renderer:  r,
	}

	return bv.initView()
}

func (v *ButtonView) initView() *ButtonView {
	textColor := greenButtonTextColor
	if v.variant == widgets.ButtonVariantRed {
		textColor = redButtonTextColor
	}

	v.buttonLabel = widgets.
		NewText(v.labelText, v.renderer, v.scale).
		WithColor(textColor).
		WithVerticalShift(-2)
	v.button = widgets.
		NewButton(v.variant, v.scale).
		WithLastFrameFix(8)

	v.buttonView = &furex.View{
		Direction:    furex.Row,
		AlignContent: furex.AlignContentCenter,
		Justify:      furex.JustifyCenter,
		AlignItems:   furex.AlignItemCenter,
		Width:        150,
		Handler:      v.button,
	}
	v.buttonView.AddChild(&furex.View{
		Handler: v.buttonLabel,
	})

	return v
}

func (v *ButtonView) View() *furex.View {
	return v.buttonView
}

func (v *ButtonView) Disable() {
	v.disabled = true
	v.buttonLabel.SetColor(color.RGBA{
		R: 21,
		G: 27,
		B: 51,
		A: 255})
	v.button.Disable()
}

func (v *ButtonView) Enable() {
	v.disabled = false

	textColor := greenButtonTextColor
	if v.variant == widgets.ButtonVariantRed {
		textColor = redButtonTextColor
	}
	v.buttonLabel.SetColor(textColor)
	v.button.Enable()
}

func (v *ButtonView) OnClick(onClick func()) *ButtonView {
	v.button.OnClick(func() {
		if v.disabled {
			return
		}
		onClick()
	})
	return v
}

func (v *ButtonView) Width(width int) *ButtonView {
	v.buttonView.Width = width
	return v
}

func (v *ButtonView) MarginRight(marginRight int) *ButtonView {
	v.buttonView.MarginRight = marginRight
	return v
}
