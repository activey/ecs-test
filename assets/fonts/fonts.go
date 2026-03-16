package fonts

import (
	_ "embed"
	"github.com/tinne26/etxt/font"
	"golang.org/x/image/font/sfnt"
)

var (
	//go:embed gamer.ttf
	mainFontData []byte

	//go:embed joystix.otf
	secondaryFontData []byte

	MainFont      *sfnt.Font
	SecondaryFont *sfnt.Font
)

func MustLoadFonts() {
	f, _, err := font.ParseFromBytes(mainFontData)
	if err != nil {
		panic(err)
	}
	MainFont = f

	f, _, err = font.ParseFromBytes(secondaryFontData)
	if err != nil {
		panic(err)
	}
	SecondaryFont = f
}
