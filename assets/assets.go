package assets

import (
	"ecs-test/assets/fonts"
	"ecs-test/assets/shaders"
	"ecs-test/assets/sprites"
	"ecs-test/assets/sprites/animations"
	"ecs-test/assets/sprites/ui"
	"embed"
	_ "embed"
)

var (
	//go:embed scenes/*
	ScenesAssets embed.FS
)

func MustLoadAssets() {
	fonts.MustLoadFonts()
	sprites.MustLoadSprites()
	ui.MustLoadUiAssets()

	animations.MustLoadPlayerAnimations()
	animations.MustLoadFlagAnimation()

	shaders.MustLoadShaders()
}
