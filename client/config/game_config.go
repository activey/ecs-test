package config

type GameClientConfig struct {
	Width, Height int
	ServerAddress string
}

func NewGameClientConfig(width int, height int, serverAddress string) GameClientConfig {
	return GameClientConfig{
		Width:         width,
		Height:        height,
		ServerAddress: serverAddress,
	}
}
