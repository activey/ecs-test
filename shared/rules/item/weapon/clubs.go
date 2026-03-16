package weapon

import (
	"ecs-test/shared/rules/ability"
	"ecs-test/shared/rules/item"
)

var (
	Club = Weapon{
		BaseItem: item.BaseItem{
			BaseIndex: "club",
			BaseName:  "Club",
			BaseDescription: "Club is a mundane, nonmagical variant of the *Clubs* family of weapons. " +
				"It is a *simple melee weapon* wielded in one hand. It's a *light* weapon that anyone can dual-wield without special training. " +
				"It is more or less a temporary weapon considering other simple weapons like *Mace* are superior to it in damage.",
			BaseType:   item.Weapon,
			BaseRarity: item.Common,
			BaseWeight: 0.9,
			BasePrice:  10,
		},

		Damage: []item.Damage{
			{
				Name:         "Damage",
				Type:         item.Bludgeoning,
				Dice:         "1d4",
				PlusModifier: &ability.Strength,
			},
			{
				Name: "Extra damage",
				Type: item.Cold,
				Dice: "1d4",
			},
		},
		WeaponType: Clubs,
	}

	AllClubs = []item.Item{
		&Club,
	}
)
