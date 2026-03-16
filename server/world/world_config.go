package world

type Configuration struct {
	MapFile string
}

func NewWorldConfiguration(mapFile string) *Configuration {
	return &Configuration{
		MapFile: mapFile,
	}
}
