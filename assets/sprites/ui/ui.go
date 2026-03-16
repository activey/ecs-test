package ui

import (
	"bytes"
	_ "embed"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/yohamta/ganim8/v2"
	"log"
)

var (
	//go:embed ui_sheet.png
	uiSheetData []byte

	PanelSprite     *ganim8.Sprite
	HeaderSprite    *ganim8.Sprite
	TextInputSprite *ganim8.Sprite
	CursorSprite    *ganim8.Sprite

	GreenButtonSpriteNormal *ganim8.Sprite
	GreenButtonSpriteHover  *ganim8.Sprite
	GreenButtonSpriteDown   *ganim8.Sprite

	RedButtonSpriteNormal *ganim8.Sprite
	RedButtonSpriteHover  *ganim8.Sprite
	RedButtonSpriteDown   *ganim8.Sprite

	ProgressBarSprite       *ganim8.Sprite
	ProgressBarFillerSprite *ganim8.Sprite
)

func MustLoadUiAssets() {
	spriteSheet, _, err := ebitenutil.NewImageFromReader(bytes.NewReader(uiSheetData))
	if err != nil {
		log.Panic(err)
	}

	mustLoadPanelSprite(spriteSheet)
	mustLoadHeaderSprite(spriteSheet)
	mustLoadTextInputSprite(spriteSheet)
	mustLoadGreenButtonSprite(spriteSheet)
	mustLoadRedButtonSprite(spriteSheet)
	mustLoadCursorSprite(spriteSheet)
	mustLoadProgressBarSprite(spriteSheet)
}

func mustLoadPanelSprite(spriteSheet *ebiten.Image) {
	panelSpriteGrid := ganim8.NewGrid(16, 16, spriteSheet.Bounds().Dx(), spriteSheet.Bounds().Dy())
	PanelSprite = ganim8.NewSprite(spriteSheet, panelSpriteGrid.Frames("2-4", "2-4"))
}

func mustLoadHeaderSprite(spriteSheet *ebiten.Image) {
	headerSpriteGrid := ganim8.NewGrid(16, 16, spriteSheet.Bounds().Dx(), spriteSheet.Bounds().Dy())
	HeaderSprite = ganim8.NewSprite(spriteSheet, headerSpriteGrid.Frames("2-4", 6))
}

func mustLoadTextInputSprite(spriteSheet *ebiten.Image) {
	inputSpriteGrid := ganim8.NewGrid(16, 16, spriteSheet.Bounds().Dx(), spriteSheet.Bounds().Dy())
	TextInputSprite = ganim8.NewSprite(spriteSheet, inputSpriteGrid.Frames("12-16", 21))
}

func mustLoadGreenButtonSprite(spriteSheet *ebiten.Image) {
	spriteGrid := ganim8.NewGrid(16, 16, spriteSheet.Bounds().Dx(), spriteSheet.Bounds().Dy())
	GreenButtonSpriteNormal = ganim8.NewSprite(spriteSheet, spriteGrid.Frames("6-9", "19-20"))
	GreenButtonSpriteHover = ganim8.NewSprite(spriteSheet, spriteGrid.Frames("6-9", "21-22"))
	GreenButtonSpriteDown = ganim8.NewSprite(spriteSheet, spriteGrid.Frames("6-9", "23-24"))
}

func mustLoadRedButtonSprite(spriteSheet *ebiten.Image) {
	spriteGrid := ganim8.NewGrid(16, 16, spriteSheet.Bounds().Dx(), spriteSheet.Bounds().Dy())
	RedButtonSpriteNormal = ganim8.NewSprite(spriteSheet, spriteGrid.Frames("2-5", "19-20"))
	RedButtonSpriteHover = ganim8.NewSprite(spriteSheet, spriteGrid.Frames("2-5", "21-22"))
	RedButtonSpriteDown = ganim8.NewSprite(spriteSheet, spriteGrid.Frames("2-5", "23-24"))
}

func mustLoadCursorSprite(spriteSheet *ebiten.Image) {
	cursorGrid := ganim8.NewGrid(16, 16, spriteSheet.Bounds().Dx(), spriteSheet.Bounds().Dy())
	CursorSprite = ganim8.NewSprite(spriteSheet, cursorGrid.Frames("23-24", "20-23"))
}

func mustLoadProgressBarSprite(spriteSheet *ebiten.Image) {
	progressBarGrid := ganim8.NewGrid(16, 16, spriteSheet.Bounds().Dx(), spriteSheet.Bounds().Dy())
	ProgressBarSprite = ganim8.NewSprite(spriteSheet, progressBarGrid.Frames("19-29", "9-10"))
	ProgressBarFillerSprite = ganim8.NewSprite(spriteSheet, progressBarGrid.Frames("19-29", 12))
}
