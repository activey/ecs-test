package animations

import (
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/yohamta/ganim8/v2"
	"log"
	"time"
)

const (
	playerAnimationTime = time.Millisecond * 80
)

var (
	IdleDownSprite     *ganim8.Sprite
	IdleUpSprite       *ganim8.Sprite
	IdleRightSprite    *ganim8.Sprite
	IdleLeftSprite     *ganim8.Sprite
	WalkingRightSprite *ganim8.Sprite
	WalkingLeftSprite  *ganim8.Sprite
	WalkingUpSprite    *ganim8.Sprite
	WalkingDownSprite  *ganim8.Sprite

	DustSprite *ganim8.Sprite
)

func MustLoadPlayerAnimations() {
	spriteSheet, _, err := ebitenutil.NewImageFromFile("assets/scenes/tilesets/woods/sprites/characters/player.png")
	if err != nil {
		log.Panic(err)
	}

	playerAnimationGrid := ganim8.NewGrid(48, 48, spriteSheet.Bounds().Dx(), spriteSheet.Bounds().Dy())
	IdleDownSprite = ganim8.NewSprite(spriteSheet, playerAnimationGrid.Frames("1-6", 1))
	IdleRightSprite = ganim8.NewSprite(spriteSheet, playerAnimationGrid.Frames("1-6", 2))
	IdleUpSprite = ganim8.NewSprite(spriteSheet, playerAnimationGrid.Frames("1-6", 3))
	IdleLeftSprite = ganim8.NewSprite(spriteSheet, playerAnimationGrid.Frames("1-6", 2))
	IdleLeftSprite.FlipH()

	WalkingRightSprite = ganim8.NewSprite(spriteSheet, playerAnimationGrid.Frames("1-6", 5))
	WalkingLeftSprite = ganim8.NewSprite(spriteSheet, playerAnimationGrid.Frames("1-6", 5))
	WalkingLeftSprite.FlipH()
	WalkingDownSprite = ganim8.NewSprite(spriteSheet, playerAnimationGrid.Frames("1-6", 4))
	WalkingUpSprite = ganim8.NewSprite(spriteSheet, playerAnimationGrid.Frames("1-6", 6))

	spriteSheet, _, err = ebitenutil.NewImageFromFile("assets/scenes/tilesets/woods/sprites/particles/dust_particles_01.png")
	dustGrid := ganim8.NewGrid(12, 12, spriteSheet.Bounds().Dx(), spriteSheet.Bounds().Dy())
	DustSprite = ganim8.NewSprite(spriteSheet, dustGrid.Frames("2-3", 1))
}

func NewWalkingRightAnimation() *ganim8.Animation {
	return ganim8.NewAnimation(WalkingRightSprite, playerAnimationTime)
}

func NewWalkingLeftAnimation() *ganim8.Animation {
	return ganim8.NewAnimation(WalkingLeftSprite, playerAnimationTime)
}

func NewWalkingDownAnimation() *ganim8.Animation {
	return ganim8.NewAnimation(WalkingDownSprite, playerAnimationTime)
}

func NewWalkingUpAnimation() *ganim8.Animation {
	return ganim8.NewAnimation(WalkingUpSprite, playerAnimationTime)
}

func NewIdleUpAnimation() *ganim8.Animation {
	return ganim8.NewAnimation(IdleUpSprite, playerAnimationTime)
}

func NewIdleDownAnimation() *ganim8.Animation {
	return ganim8.NewAnimation(IdleDownSprite, playerAnimationTime)
}

func NewIdleLeftAnimation() *ganim8.Animation {
	return ganim8.NewAnimation(IdleLeftSprite, playerAnimationTime)
}

func NewIdleRightAnimation() *ganim8.Animation {
	return ganim8.NewAnimation(IdleRightSprite, playerAnimationTime)
}

func NewDustAnimation() *ganim8.Animation {
	return ganim8.NewAnimation(DustSprite, playerAnimationTime)
}
