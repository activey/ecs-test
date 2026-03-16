package ability

type AllScoresType []Score

func (a AllScoresType) Size() int {
	return len(a)
}

type Score string

func (a Score) String() string {
	return string(a)
}

var (
	Strength     Score = "Strength"
	Dexterity    Score = "Dexterity"
	Constitution Score = "Constitution"
	Intelligence Score = "Intelligence"
	Wisdom       Score = "Wisdom"
	Charisma     Score = "Charisma"

	AllScores = AllScoresType{
		Strength, Dexterity, Constitution, Intelligence, Wisdom, Charisma,
	}
)
