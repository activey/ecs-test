package animations

import (
	"bytes"
	_ "embed"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/yohamta/ganim8/v2"
	"log"
)

var (
	//go:embed flag_sheet.png
	flagSheetData []byte

	FlagSprite *ganim8.Sprite
)

func MustLoadFlagAnimation() {
	spriteSheet, _, err := ebitenutil.NewImageFromReader(bytes.NewReader(flagSheetData))
	if err != nil {
		log.Panic(err)
	}

	spriteGrid := ganim8.NewGrid(64, 64, spriteSheet.Bounds().Dx(), spriteSheet.Bounds().Dy())
	FlagSprite = ganim8.NewSprite(spriteSheet, spriteGrid.Frames("1-2", "1-2"))
}

func NewFlagAnimation() *ganim8.Animation {
	return ganim8.NewAnimation(FlagSprite, 160)
}
