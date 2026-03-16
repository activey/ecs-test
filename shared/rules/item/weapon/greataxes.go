package weapon

import (
	"ecs-test/shared/rules/ability"
	"ecs-test/shared/rules/item"
)

var (
	Greataxe = Weapon{
		BaseItem: item.BaseItem{
			BaseIndex:       "greataxe",
			BaseType:        item.Weapon,
			BaseName:        "Greataxe",
			BaseDescription: "Greataxe is a common variant of the *Greataxes* family of weapons.",
			BaseRarity:      item.Common,
			BaseWeight:      3.15,
			BasePrice:       65,
		},
		Damage: []item.Damage{
			{
				Name:         "Damage",
				Type:         item.Slashing,
				Dice:         "1d12",
				PlusModifier: &ability.Strength,
			},
		},
		WeaponType: Greataxes,
	}

	AllGreatAxes = []item.Item{
		&Greataxe,
	}
)
