package cmd

import (
	"ecs-test/assets"
	"ecs-test/client"
	"ecs-test/client/archetype"
	"ecs-test/client/config"
	"ecs-test/client/middleware/broadcast"
	"ecs-test/client/middleware/player"
	"ecs-test/client/middleware/session"
	"ecs-test/client/middleware/socket"
	"ecs-test/client/scene"
	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/math"
	"go.uber.org/dig"
)

var (
	rootCmd = &cobra.Command{
		Use: "game",
		Run: RunGame,
	}

	serverHost = "localhost"
)

func NewGameClientConfig() config.GameClientConfig {
	return config.NewGameClientConfig(
		800,
		600,
		serverHost,
	)
}

func RunGame(cmd *cobra.Command, args []string) {
	assets.MustLoadAssets()

	container := dig.New()
	provideMiddlewareComponents(container)
	provideGameComponents(container)
	provideGameScenes(container)

	err := container.Invoke(startGame)
	if err != nil {
		log.Fatal(err)
	}
}

func provideGameComponents(container *dig.Container) {
	Provide(container, NewGameClientConfig)
	Provide(container, donburi.NewWorld)
	Provide(container, client.NewGameClient)
	Provide(container, scene.NewSceneSwitcher)

	err := container.Invoke(addArchetypes)
	if err != nil {
		log.Fatal(err)
	}
}

func addArchetypes(world donburi.World) {
	archetype.NewDebug(world, false)
	archetype.NewCursor(world, math.Vec2{})
	archetype.NewSession(world)
	archetype.NewPlayerLock(world)
	archetype.NewCamera(world, math.Vec2{}, 0.05, 0.2)
	archetype.NewWorldMap(world)
}

func provideGameScenes(container *dig.Container) {
	Provide(container, scene.NewMenuSceneFactory)
	Provide(container, scene.NewWorldDisplaySceneFactory)

	err := container.Invoke(func(s *scene.SceneSwitcher, scene *scene.MenuSceneFactory) {
		s.AddScene("menu", scene)
	})
	if err != nil {
		log.Fatal(err)
	}

	err = container.Invoke(func(s *scene.SceneSwitcher, scene *scene.WorldDisplaySceneFactory) {
		s.AddScene("world_display", scene)
	})
	if err != nil {
		log.Fatal(err)
	}
}

func provideMiddlewareComponents(container *dig.Container) {
	err := session.ProvideModuleComponents(container)
	if err != nil {
		log.Fatal(err)
	}

	err = player.ProvideModuleComponents(container)
	if err != nil {
		log.Fatal(err)
	}

	err = socket.ProvideModuleComponents(container)
	if err != nil {
		log.Fatal(err)
	}

	err = broadcast.ProvideModuleComponents(container)
	if err != nil {
		log.Fatal(err)
	}
}

func startGame(g *client.GameClient) {
	g.Start()
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func Provide(c *dig.Container, constructor interface{}) {
	err := c.Provide(constructor)
	if err != nil {
		log.Fatal(err)
		return
	}
}

func init() {
	rootCmd.
		PersistentFlags().
		StringVarP(&serverHost, "server-host", "s", "localhost", "server host to connect to")

}
