package modifier

import "ecs-test/shared/rules/ability"

type Modifier struct {
	Changes []Change
	// TODO add rules for applying changes
}

func PlusStrength(value uint) Change {
	return Change{
		Attribute: ability.Strength,
		Value:     value,
	}
}

func PlusDexterity(value uint) Change {
	return Change{
		Attribute: ability.Dexterity,
		Value:     value,
	}
}

func PlusConstitution(value uint) Change {
	return Change{
		Attribute: ability.Constitution,
		Value:     value,
	}
}

func PlusCharisma(value uint) Change {
	return Change{
		Attribute: ability.Charisma,
		Value:     value,
	}
}

func PlusIntelligence(value uint) Change {
	return Change{
		Attribute: ability.Intelligence,
		Value:     value,
	}
}

func PlusWisdom(value uint) Change {
	return Change{
		Attribute: ability.Wisdom,
		Value:     value,
	}
}

type Change struct {
	Attribute ability.Score
	Value     uint
}
