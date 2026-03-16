package feature

import "ecs-test/shared/rules/modifier"

type Feature struct {
	Name, Description string
	Modifiers         []modifier.Modifier
}

var (
	Rage = Feature{
		Name:        "Rage",
		Description: "In battle, you fight with primal ferocity. On your turn, you can enter a rage as a bonus action.",
		Modifiers: []modifier.Modifier{
			{
				Changes: []modifier.Change{
					modifier.PlusStrength(2),
				},
			},
		},
	}
	// TODO add information on how long it lasts during the battle etc

	UnarmoredDefense = Feature{
		Name: "Unarmored Defense",
		Description: "While you are not wearing any armor, your Armor Class equals 10 + your Dexterity modifier + your Constitution modifier. " +
			"You can use a shield and still gain this benefit.",
		Modifiers: []modifier.Modifier{
			{
				Changes: []modifier.Change{
					modifier.PlusDexterity(1),
				},
			},
		},
	}

	Spellcasting = Feature{
		Name: "Spellcasting",
		Description: "An event in your past, or in the life of a parent or ancestor, " +
			"left an indelible mark on you, infusing you with arcane magic. " +
			"This font of magic, whatever its origin, fuels your spells.",
	}

	FightingStyle = Feature{
		Name: "Fighting style",
		Description: "You adopt a particular style of fighting as your specialty. " +
			"Choose one of the following options. You can’t take a Fighting Style option more than once, " +
			"even if you later get to choose again.",
	}
)
