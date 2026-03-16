package widgets

import (
	"ecs-test/assets/sprites/ui"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/colorm"
	"github.com/yohamta/furex/v2"
	"github.com/yohamta/ganim8/v2"
	"image"
)

type ButtonVariant int

const (
	ButtonVariantGreen ButtonVariant = iota
	ButtonVariantRed
)

type ButtonSpriteSet struct {
	buttonSpriteNormal *ganim8.Sprite
	buttonSpriteHover  *ganim8.Sprite
	buttonSpriteDown   *ganim8.Sprite
}

type Button struct {
	currentSprite *ganim8.Sprite
	spriteSet     ButtonSpriteSet
	offscreen     *ebiten.Image

	enabled      bool
	visible      bool
	scale        float64
	lastFrameFix float64
	onClick      func()
}

func NewButton(variant ButtonVariant, scale float64) *Button {
	greenButtonSpriteSet := ButtonSpriteSet{
		buttonSpriteNormal: ui.GreenButtonSpriteNormal,
		buttonSpriteHover:  ui.GreenButtonSpriteHover,
		buttonSpriteDown:   ui.GreenButtonSpriteDown,
	}

	redButtonSpriteSet := ButtonSpriteSet{
		buttonSpriteNormal: ui.RedButtonSpriteNormal,
		buttonSpriteHover:  ui.RedButtonSpriteHover,
		buttonSpriteDown:   ui.RedButtonSpriteDown,
	}

	spriteSetMap := map[ButtonVariant]ButtonSpriteSet{
		ButtonVariantGreen: greenButtonSpriteSet,
		ButtonVariantRed:   redButtonSpriteSet,
	}

	spriteSet := spriteSetMap[variant]

	return &Button{
		currentSprite: spriteSet.buttonSpriteNormal,
		spriteSet:     spriteSet,

		scale:   scale,
		visible: true,
		enabled: true,
	}
}

func (b *Button) WithLastFrameFix(fixValue float64) *Button {
	b.lastFrameFix = fixValue
	return b
}

func (b *Button) HandleMouseEnter(x, y int) bool {
	b.currentSprite = b.spriteSet.buttonSpriteHover
	return true
}

func (b *Button) HandleMouseLeave() {
	b.currentSprite = b.spriteSet.buttonSpriteNormal
}

func (b *Button) HandleJustPressedMouseButtonLeft(x, y int) bool {
	b.currentSprite = b.spriteSet.buttonSpriteDown
	return true
}

func (b *Button) HandleJustReleasedMouseButtonLeft(x, y int) {
	b.currentSprite = b.spriteSet.buttonSpriteHover
	if b.onClick != nil {
		b.onClick()
	}
}

func (b *Button) Hide() {
	b.visible = false
}

func (b *Button) Show() {
	b.visible = true
}

func (b *Button) Disable() {
	b.enabled = false
}

func (b *Button) Enable() {
	b.enabled = true
}

func (b *Button) Draw(screen *ebiten.Image, frame image.Rectangle, v *furex.View) {
	if !b.visible {
		return
	}

	if b.offscreen == nil {
		b.offscreen = ebiten.NewImage(screen.Bounds().Dx(), screen.Bounds().Dy())
	}
	b.offscreen.Clear()

	spriteWidth := float64(b.currentSprite.Width())
	scaledSpriteWidth := spriteWidth * b.scale

	spriteHeight := float64(b.currentSprite.Height())
	scaledSpriteHeight := spriteHeight * b.scale

	startX := float64(frame.Min.X)
	startY := float64(frame.Min.Y)

	// Draw upper left corner (Index 0)
	frameOpts := ganim8.DrawOpts(startX, startY, 0, b.scale, b.scale)
	b.currentSprite.Draw(b.offscreen, 0, frameOpts)

	// Draw bottom left corner (Index 4)
	frameOpts.SetPos(startX, startY+scaledSpriteHeight)
	b.currentSprite.Draw(b.offscreen, 4, frameOpts)

	// Draw upper right corner (Index 3)
	lastFrameFix := b.lastFrameFix * b.scale
	//lastFrameFix := 10.0
	frameOpts.SetPos(startX+float64(frame.Dx())-scaledSpriteWidth+lastFrameFix, startY)
	b.currentSprite.Draw(b.offscreen, 3, frameOpts)

	// Draw bottom right corner (Index 7)
	frameOpts.SetPos(startX+float64(frame.Dx())-scaledSpriteWidth+lastFrameFix, startY+scaledSpriteHeight)
	b.currentSprite.Draw(b.offscreen, 7, frameOpts)

	// second from end top
	frameOpts.SetPos(startX+float64(frame.Dx())-2*scaledSpriteWidth+lastFrameFix, startY)
	b.currentSprite.Draw(b.offscreen, 2, frameOpts)

	// second from end bottom
	frameOpts.SetPos(startX+float64(frame.Dx())-2*scaledSpriteWidth+lastFrameFix, startY+scaledSpriteHeight)
	b.currentSprite.Draw(b.offscreen, 6, frameOpts)

	// Draw the expanded top middle (Index 1)
	middleWidth := float64(frame.Dx()) - 3*scaledSpriteWidth + lastFrameFix // Calculate the width for the expanding middle frame
	if middleWidth > 0 {
		// Set drawing options for the middle top part, with custom scaling in width
		middleFrameOpts := ganim8.DrawOpts(startX+scaledSpriteWidth, startY)
		middleFrameOpts.SetScale(middleWidth/spriteWidth, b.scale)
		b.currentSprite.Draw(b.offscreen, 1, middleFrameOpts)
	}

	// Draw the expanded bottom middle (Index 5)
	if middleWidth > 0 {
		// Set drawing options for the middle bottom part, with custom scaling in width
		middleFrameOpts := ganim8.DrawOpts(startX+scaledSpriteWidth, startY+scaledSpriteHeight)
		middleFrameOpts.SetScale(middleWidth/spriteWidth, b.scale)
		b.currentSprite.Draw(b.offscreen, 5, middleFrameOpts)
	}

	colorM := colorm.ColorM{}
	if !b.enabled {
		colorM.ChangeHSV(0, 0, 1)
	}
	colorm.DrawImage(screen, b.offscreen, colorM, nil)
}

func (b *Button) Update(v *furex.View) {
	//if !b.enabled {
	//	return
	//}

	spriteHeight := float64(b.currentSprite.Height())
	// hack because of shitty sprite :P
	spriteHeight -= 6
	scaledSpriteHeight := 2 * spriteHeight * b.scale

	v.SetHeight(int(scaledSpriteHeight))
}

func (b *Button) OnClick(onClick func()) *Button {
	b.onClick = onClick
	return b
}
