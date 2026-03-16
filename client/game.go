package client

import (
	"ecs-test/client/config"
	"ecs-test/client/scene"
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

type GameClient struct {
	switcher *scene.SceneSwitcher

	width, height int
}

func NewGameClient(config config.GameClientConfig, switcher *scene.SceneSwitcher) *GameClient {
	return &GameClient{
		switcher: switcher,
		width:    config.Width,
		height:   config.Height,
	}
}

func (g *GameClient) Update() error {
	if g.switcher == nil {
		return nil
	}

	return g.switcher.Update()
}

func (g *GameClient) Draw(screen *ebiten.Image) {
	if g.switcher == nil {
		return
	}
	g.switcher.Draw(screen)
}

func (g *GameClient) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	g.width, g.height = outsideWidth, outsideHeight
	return g.width, g.height
}

func (g *GameClient) Start() {
	ebiten.SetWindowSize(g.width, g.height)
	ebiten.SetCursorMode(ebiten.CursorModeHidden)
	ebiten.SetVsyncEnabled(true)
	ebiten.SetTPS(60)

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
