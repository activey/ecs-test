package cmd

import (
	"ecs-test/server/character"
	"ecs-test/server/infra/broadcast"
	"ecs-test/server/infra/database"
	"ecs-test/server/infra/http"
	"ecs-test/server/infra/socket"
	"ecs-test/server/player"
	"ecs-test/server/session"
	"ecs-test/server/user"
	"ecs-test/server/world"
	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
	"go.uber.org/dig"
	"os"
	"os/signal"
	"syscall"
)

var (
	RunServerCommand = &cobra.Command{
		Use: "server",
		Run: RunServer,
	}
)

func RunServer(cmd *cobra.Command, args []string) {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	container := dig.New()

	var err error
	// infrastructure
	err = container.Provide(database.NewDatabaseConfiguration)
	err = container.Provide(database.NewDatabaseConnection)
	err = container.Provide(http.NewServerConfig)
	err = container.Provide(http.NewHttpServer)
	err = container.Provide(socket.NewSocketServerConfig)
	err = container.Provide(socket.NewServer)
	err = container.Provide(broadcast.NewBroadcastServerConfig)
	err = container.Provide(broadcast.NewServer)

	// modules
	err = user.ProvideModuleComponents(container)
	if err != nil {
		log.Fatal(err)
	}

	err = character.ProvideModuleComponents(container)
	if err != nil {
		log.Fatal(err)
	}

	err = session.ProvideModuleComponents(container)
	if err != nil {
		log.Fatal(err)
	}

	err = world.ProvideModuleComponents(container)
	if err != nil {
		log.Fatal(err)
	}

	err = player.ProvideModuleComponents(container)
	if err != nil {
		log.Fatal(err)
	}

	// startup
	err = container.Invoke(startHttpServer)
	if err != nil {
		log.Fatal(err)
	}

	err = container.Invoke(startSocketServer)
	if err != nil {
		log.Fatal(err)
	}

	err = container.Invoke(startBroadcastServer)
	if err != nil {
		log.Fatal(err)
	}

	<-stop

	log.Info("Shutting down servers...")
	err = container.Invoke(stopServers)
}

func startHttpServer(s *http.Server) {
	go s.Start()
}

func startSocketServer(s *socket.Server) {
	go s.Start()
}

func startBroadcastServer(s *broadcast.Server) {
	go s.Start()
}

func stopServers(s *socket.Server, b *broadcast.Server) {
	s.Stop()
	b.Stop()
}

func init() {
	rootCmd.AddCommand(RunServerCommand)
}
