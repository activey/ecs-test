package class

import (
	"ecs-test/shared/rules/ability"
	"ecs-test/shared/rules/feature"
	"ecs-test/shared/rules/item/weapon"
)

var (
	Barbarian = Class{
		Index:       "barbarian",
		Name:        "Barbarian",
		Description: "A fierce warrior of primitive background, relying on raw physical power and rage to dominate the battlefield.",
		HitPoints:   "1d12",
		Features: []feature.Feature{
			feature.Rage, feature.UnarmoredDefense,
		},
		SaveThrows: []ability.Score{
			ability.Strength,
			ability.Constitution,
		},
		KeyAbilities: []ability.Score{
			ability.Strength,
			ability.Dexterity,
			ability.Constitution,
		},
		WeaponProficiency: weapon.WeaponProficiency{
			AllSimpleWeapons:  true,
			AllMartialWeapons: true,
		},
	}

	Bard = Class{
		Index:       "bard",
		Name:        "Bard",
		Description: "A versatile spellcaster who uses music, poetry, and magic to inspire allies and manipulate foes.",
		HitPoints:   "1d8",
		Features: []feature.Feature{
			feature.Spellcasting,
		},
		SaveThrows: []ability.Score{
			ability.Dexterity,
			ability.Charisma,
		},
		KeyAbilities: []ability.Score{
			ability.Charisma,
			ability.Constitution,
			ability.Dexterity,
		},
		WeaponProficiency: weapon.WeaponProficiency{
			AllSimpleWeapons: true,
			Proficiency: []weapon.Type{
				weapon.HandCrossbows,
				weapon.Rapiers,
				weapon.Shortswords,
				weapon.Longswords,
			},
		},
	}

	Cleric = Class{
		Index:       "cleric",
		Name:        "Cleric",
		Description: "A holy warrior with divine magic, capable of healing and supporting allies or unleashing the wrath of their deity.",
		HitPoints:   "1d8",
		Features: []feature.Feature{
			feature.Spellcasting,
		},
		SaveThrows: []ability.Score{
			ability.Wisdom,
			ability.Charisma,
		},
		KeyAbilities: []ability.Score{
			ability.Wisdom,
			ability.Constitution,
			ability.Dexterity,
		},
		WeaponProficiency: weapon.WeaponProficiency{
			AllSimpleWeapons: true,
			Proficiency: []weapon.Type{
				weapon.Flails,
				weapon.Morningstars,
			},
		},
	}

	Druid = Class{
		Index:       "druid",
		Name:        "Druid",
		Description: "A guardian of nature with the ability to transform into animals and wield elemental magic.",
		HitPoints:   "1d8",
		Features: []feature.Feature{
			feature.Spellcasting,
		},
		SaveThrows: []ability.Score{
			ability.Intelligence,
			ability.Wisdom,
		},
		KeyAbilities: []ability.Score{
			ability.Wisdom,
			ability.Constitution,
			ability.Dexterity,
		},
		WeaponProficiency: weapon.WeaponProficiency{
			Proficiency: []weapon.Type{
				weapon.Clubs,
				weapon.Daggers,
				weapon.Javelins,
				weapon.Maces,
				weapon.Quarterstaves,
				weapon.Scimitars,
				weapon.Sickles,
				weapon.Spears,
			},
		},
	}

	Fighter = Class{
		Index:       "fighter",
		Name:        "Fighter",
		Description: "A master of combat, skilled with a wide range of weapons and tactics, adaptable to any battlefield role.",
		HitPoints:   "1d10",
		Features: []feature.Feature{
			feature.FightingStyle,
		},
		SaveThrows: []ability.Score{
			ability.Strength,
			ability.Constitution,
		},
		KeyAbilities: []ability.Score{
			ability.Strength,
			ability.Dexterity,
			ability.Constitution,
		},
		WeaponProficiency: weapon.WeaponProficiency{
			AllSimpleWeapons:  true,
			AllMartialWeapons: true,
		},
	}

	Monk = Class{
		Index:       "monk",
		Name:        "Monk",
		Description: "A master of martial arts, using physical and spiritual discipline to perform extraordinary feats of speed and precision.",
		HitPoints:   "1d8",
		Features:    []feature.Feature{
			//MartialArts,
		},
		SaveThrows: []ability.Score{
			ability.Strength,
			ability.Dexterity,
		},
		KeyAbilities: []ability.Score{
			ability.Wisdom,
			ability.Constitution,
			ability.Dexterity,
		},
		WeaponProficiency: weapon.WeaponProficiency{
			AllSimpleWeapons: true,
			Proficiency: []weapon.Type{
				weapon.Shortswords,
			},
		},
	}

	Paladin = Class{
		Index:       "paladin",
		Name:        "Paladin",
		Description: "A holy knight sworn to an oath, combining martial prowess with divine magic to uphold justice and righteousness.",
		HitPoints:   "1d10",
		Features: []feature.Feature{
			feature.Spellcasting,
		},
		SaveThrows: []ability.Score{
			ability.Wisdom,
			ability.Charisma,
		},
		KeyAbilities: []ability.Score{
			ability.Strength,
			ability.Charisma,
			ability.Constitution,
		},
		WeaponProficiency: weapon.WeaponProficiency{
			AllSimpleWeapons:  true,
			AllMartialWeapons: true,
		},
	}

	Ranger = Class{
		Index:       "ranger",
		Name:        "Ranger",
		Description: "A warrior and hunter, skilled in tracking enemies and surviving in the wild, often wielding both weapons and magic.",
		HitPoints:   "1d10",
		Features: []feature.Feature{
			feature.Spellcasting,
		},
		SaveThrows: []ability.Score{
			ability.Strength,
			ability.Dexterity,
		},
		KeyAbilities: []ability.Score{
			ability.Dexterity,
			ability.Constitution,
			ability.Wisdom,
		},
		WeaponProficiency: weapon.WeaponProficiency{
			AllSimpleWeapons:  true,
			AllMartialWeapons: true,
		},
	}

	Rogue = Class{
		Index:       "rogue",
		Name:        "Rogue",
		Description: "A nimble and cunning character, excelling in stealth, lockpicking, and dealing precise strikes to vulnerable foes.",
		HitPoints:   "1d8",
		Features:    []feature.Feature{
			//SneakAttack,
		},
		SaveThrows: []ability.Score{
			ability.Dexterity,
			ability.Intelligence,
		},
		KeyAbilities: []ability.Score{
			ability.Dexterity,
			ability.Constitution,
		},
		WeaponProficiency: weapon.WeaponProficiency{
			AllSimpleWeapons: true,
			Proficiency: []weapon.Type{
				weapon.HandCrossbows,
				weapon.Longswords,
				weapon.Rapiers,
				weapon.Shortswords,
			},
		},
	}

	Sorcerer = Class{
		Index:       "sorcerer",
		Name:        "Sorcerer",
		Description: "A natural spellcaster who draws upon innate magical abilities, often wielding powerful magic with little formal training.",
		HitPoints:   "1d6",
		Features: []feature.Feature{
			feature.Spellcasting,
		},
		SaveThrows: []ability.Score{
			ability.Constitution,
			ability.Charisma,
		},
		KeyAbilities: []ability.Score{
			ability.Charisma,
			ability.Constitution,
			ability.Dexterity,
		},
		WeaponProficiency: weapon.WeaponProficiency{
			Proficiency: []weapon.Type{
				weapon.Daggers,
				weapon.Quarterstaves,
				weapon.LightCrossbows,
			},
		},
	}

	Warlock = Class{
		Index:       "warlock",
		Name:        "Warlock",
		Description: "A spellcaster who gains their magical abilities through a pact with a powerful, otherworldly entity.",
		HitPoints:   "1d8",
		Features: []feature.Feature{
			feature.Spellcasting,
		},
		SaveThrows: []ability.Score{
			ability.Wisdom,
			ability.Charisma,
		},
		KeyAbilities: []ability.Score{
			ability.Charisma,
			ability.Constitution,
			ability.Dexterity,
		},
		WeaponProficiency: weapon.WeaponProficiency{
			AllSimpleWeapons: true,
		},
	}

	Wizard = Class{
		Index:       "wizard",
		Name:        "Wizard",
		Description: "A scholarly spellcaster who learns and masters arcane magic through rigorous study and spellbooks.",
		HitPoints:   "1d6",
		Features: []feature.Feature{
			feature.Spellcasting,
		},
		SaveThrows: []ability.Score{
			ability.Intelligence,
			ability.Wisdom,
		},
		KeyAbilities: []ability.Score{
			ability.Intelligence,
			ability.Constitution,
			ability.Dexterity,
		},
		WeaponProficiency: weapon.WeaponProficiency{
			Proficiency: []weapon.Type{
				weapon.Daggers,
				weapon.Quarterstaves,
				weapon.LightCrossbows,
			},
		},
	}
)

type Index string

type Class struct {
	Index             Index
	Name, Description string

	Features          []feature.Feature
	HitPoints         string
	KeyAbilities      []ability.Score
	SaveThrows        []ability.Score
	WeaponProficiency weapon.WeaponProficiency
}

func (c Class) String() string {
	return c.Name
}
