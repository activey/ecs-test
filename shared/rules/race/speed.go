package race

type Speed struct {
	Value uint // distance in feet per turn
	Name  string
}

var (
	Slow = Speed{
		Name:  "Slow",
		Value: 25,
	}

	Normal = Speed{
		Value: 30,
		Name:  "Normal",
	}

	Fast = Speed{
		Value: 35,
		Name:  "Fast",
	}
)
