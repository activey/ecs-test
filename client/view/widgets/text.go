package widgets

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/tinne26/etxt"
	"github.com/yohamta/furex/v2"
	"image"
	"image/color"
)

type TextAlign int

const (
	LeftAlign TextAlign = iota
	CenterAlign
	RightAlign
)

type Text struct {
	text     string
	renderer *etxt.Renderer

	scaleX, scaleY float64
	verticalShift  float64
	alignment      TextAlign // Alignment using the TextAlign type
	shadow         bool
	color          color.Color
}

func NewText(text string, renderer *etxt.Renderer, scale float64) *Text {
	return &Text{
		text:      text,
		renderer:  renderer,
		scaleX:    scale,
		scaleY:    scale,
		alignment: LeftAlign, // Default to left alignment
	}
}
func (t *Text) WithVerticalShift(shift float64) *Text {
	t.verticalShift = shift
	return t
}

func (t *Text) WithAlignment(alignment TextAlign) *Text {
	t.alignment = alignment
	return t
}

func (t *Text) WithShadow() *Text {
	t.shadow = true
	return t
}

func (t *Text) SetText(s string) {
	t.text = s
}

func (t *Text) WithColor(color color.Color) *Text {
	t.color = color
	return t
}

func (t *Text) Update(v *furex.View) {
	textBounds := t.renderer.Measure(t.text)
	textWidth := textBounds.Width().ToInt()
	textHeight := textBounds.Height().ToInt()

	v.SetWidth(textWidth)
	v.SetHeight(textHeight)
}

func (t *Text) Draw(screen *ebiten.Image, frame image.Rectangle, view *furex.View) {
	textBounds := t.renderer.Measure(t.text)
	textWidth := textBounds.Width().ToInt()
	textHeight := textBounds.Height().ToInt()

	containerHeight := float64(frame.Dy())
	containerWidth := float64(frame.Dx())

	var x float64
	switch t.alignment {
	case CenterAlign:
		x = float64(frame.Min.X) + (containerWidth-float64(textWidth))/2
	case RightAlign:
		x = float64(frame.Min.X) + containerWidth - float64(textWidth)
	default:
		// LeftAlign or default
		x = float64(frame.Min.X)
	}

	y := float64(frame.Min.Y) + (containerHeight+float64(textHeight))/2 - float64(textHeight/4) + t.verticalShift
	if t.shadow {
		shadowOffsetX := 2
		shadowOffsetY := 2

		previousColor := t.renderer.GetColor()
		t.renderer.SetColor(color.RGBA{
			A: 255,
		})
		t.renderer.Draw(screen, t.text, int(x)+shadowOffsetX, int(y)+shadowOffsetY)
		t.renderer.SetColor(previousColor)
	}

	if t.color != nil {
		previousColor := t.renderer.GetColor()
		t.renderer.SetColor(t.color)
		t.renderer.Draw(screen, t.text, int(x), int(y))
		t.renderer.SetColor(previousColor)
	} else {
		t.renderer.Draw(screen, t.text, int(x), int(y))
	}
}

func (t *Text) SetColor(color color.Color) {
	t.color = color
}
