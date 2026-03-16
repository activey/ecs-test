package race

import (
	"ecs-test/shared/rules/item/weapon"
	"ecs-test/shared/rules/modifier"
)

var (
	Dragonborn = Race{
		Index:       "dragonborn",
		Name:        "Dragonborn",
		Description: "Proud and determined, dragonborn are powerful draconic humanoids with innate elemental breath and a heritage of legendary dragons.",
		Traits: []modifier.Modifier{
			{
				Changes: []modifier.Change{
					modifier.PlusStrength(2),
					modifier.PlusCharisma(1),
				},
			},
		},
		Speed: Normal,
	}

	DwarfHill = Race{
		Index:       "dwarf_hill",
		Name:        "Dwarf, Hill",
		Description: "Resilient and wise, Hill Dwarves are known for their tough endurance, deep wisdom, and affinity for the earth and stone.",
		Traits: []modifier.Modifier{
			{
				Changes: []modifier.Change{
					modifier.PlusConstitution(2),
					modifier.PlusWisdom(1),
				},
			},
		},
		Speed: Slow,
		Features: []Feature{
			{
				Name: "Dwarven Combat Training",
				WeaponProficiency: &weapon.WeaponProficiency{
					Proficiency: []weapon.Type{
						weapon.Battleaxes,
						weapon.Handaxes,
						weapon.LightHammers,
						weapon.Warhammers,
					},
				},
			},
		},
	}

	DwarfMountain = Race{
		Index:       "dwarf_mountain",
		Name:        "Dwarf, Mountain",
		Description: "Strong and hardy, Mountain Dwarves are skilled warriors and craftsmen, revered for their physical strength and stoic determination.",
		Traits: []modifier.Modifier{
			{
				Changes: []modifier.Change{
					modifier.PlusConstitution(2),
					modifier.PlusStrength(2),
				},
			},
		},
		Features: []Feature{
			{
				Name: "Dwarven Combat Training",
				WeaponProficiency: &weapon.WeaponProficiency{
					Proficiency: []weapon.Type{
						weapon.Battleaxes,
						weapon.Handaxes,
						weapon.LightHammers,
						weapon.Warhammers,
					},
				},
			},
		},
		Speed: Slow,
	}

	ElfHigh = Race{
		Index:       "elf_high",
		Name:        "Elf, High",
		Description: "Graceful and intelligent, High Elves are masters of arcane magic, often residing in ancient cities where knowledge and culture thrive.",
		Traits: []modifier.Modifier{
			{
				Changes: []modifier.Change{
					modifier.PlusDexterity(2),
					modifier.PlusIntelligence(1),
				},
			},
		},
		Speed: Normal,
		Features: []Feature{
			{
				Name: "Elven Weapon Training",
				WeaponProficiency: &weapon.WeaponProficiency{
					Proficiency: []weapon.Type{
						weapon.Longswords,
						weapon.Shortswords,
						weapon.Longbows,
						weapon.Shortbows,
					},
				},
			},
		},
	}

	ElfWood = Race{
		Index:       "elf_wood",
		Name:        "Elf, Wood",
		Description: "Fleet and attuned to nature, Wood Elves are expert archers and trackers, thriving in forested areas where they live in harmony with the wild.",
		Traits: []modifier.Modifier{
			{
				Changes: []modifier.Change{
					modifier.PlusDexterity(2),
					modifier.PlusWisdom(1),
				},
			},
		},
		Speed: Normal,
		Features: []Feature{
			{
				Name: "Elven Weapon Training",
				WeaponProficiency: &weapon.WeaponProficiency{
					Proficiency: []weapon.Type{
						weapon.Longswords,
						weapon.Shortswords,
						weapon.Longbows,
						weapon.Shortbows,
					},
				},
			},
		},
	}

	ElfDrow = Race{
		Index:       "elf_drow",
		Name:        "Elf, Drow",
		Description: "Mysterious and deadly, Drow are dark-skinned elves known for their affinity with shadows, dark magic, and their complex subterranean society.",
		Traits: []modifier.Modifier{
			{
				Changes: []modifier.Change{
					modifier.PlusDexterity(2),
					modifier.PlusCharisma(1),
				},
			},
		},
		Speed: Normal,
		Features: []Feature{
			{
				Name: "Drow Weapon Training",
				WeaponProficiency: &weapon.WeaponProficiency{
					Proficiency: []weapon.Type{
						weapon.Rapiers,
						weapon.Shortswords,
						weapon.HandCrossbows,
					},
				},
			},
		},
	}

	GnomeForest = Race{
		Index:       "gnome_forest",
		Name:        "Gnome, Forest",
		Description: "Quick-witted and resourceful, Forest Gnomes are known for their affinity with nature, their cleverness, and their natural aptitude for illusion magic.",
		Traits: []modifier.Modifier{
			{
				Changes: []modifier.Change{
					modifier.PlusDexterity(1),
					modifier.PlusIntelligence(2),
				},
			},
		},
		Speed: Normal,
	}

	GnomeRock = Race{
		Index:       "gnome_rock",
		Name:        "Gnome, Rock",
		Description: "Inventive and industrious, Rock Gnomes are master tinkerers with a deep curiosity for mechanics, gadgets, and the mysteries of the arcane.",
		Traits: []modifier.Modifier{
			{
				Changes: []modifier.Change{
					modifier.PlusConstitution(1),
					modifier.PlusIntelligence(2),
				},
			},
		},
		Speed: Normal,
	}

	HalfElf = Race{
		Index:       "half-elf",
		Name:        "Half-Elf",
		Description: "Blending the traits of both humans and elves, Half-Elves are versatile and charismatic, often serving as bridges between their two heritages.",
		Traits: []modifier.Modifier{
			{
				Changes: []modifier.Change{
					modifier.PlusCharisma(2),
					//modifier.PlusChoice(1),
					//modifier.PlusChoice(1),
				},
			},
		},
		Speed: Normal,
		Features: []Feature{
			{
				Name: "Civil Militia",
				WeaponProficiency: &weapon.WeaponProficiency{
					Proficiency: []weapon.Type{
						weapon.Pikes,
						weapon.Spears,
						weapon.Halberds,
						weapon.Glaives,
					},
				},
			},
		},
	}

	HalfOrc = Race{
		Index:       "half-orc",
		Name:        "Half-Orc",
		Description: "Driven by a fierce spirit, Half-Orcs combine the physical might of their orcish ancestry with a tenacious resilience, making them formidable warriors.",
		Traits: []modifier.Modifier{
			{
				Changes: []modifier.Change{
					modifier.PlusStrength(2),
					modifier.PlusConstitution(1),
				},
			},
		},
		Speed: Normal,
	}

	HalflingLightfoot = Race{
		Index:       "halfling_lightfoot",
		Name:        "Halfling, Lightfoot",
		Description: "Nimble and stealthy, Lightfoot Halflings are known for their ability to move unseen, often relying on their charm and quick thinking.",
		Traits: []modifier.Modifier{
			{
				Changes: []modifier.Change{
					modifier.PlusDexterity(2),
					modifier.PlusCharisma(1),
				},
			},
		},
		Speed: Slow,
	}

	HalflingStout = Race{
		Index:       "halfling_stout",
		Name:        "Halfling, Stout",
		Description: "Tough and determined, Stout Halflings are renowned for their bravery and resilience, often facing challenges with unwavering courage.",
		Traits: []modifier.Modifier{
			{
				Changes: []modifier.Change{
					modifier.PlusDexterity(2),
					modifier.PlusConstitution(1),
				},
			},
		},
		Speed: Slow,
	}

	Human = Race{
		Index:       "human",
		Name:        "Human",
		Description: "Adaptive and diverse, Humans are known for their ambition and versatility, capable of excelling in nearly any role or environment.",
		Traits: []modifier.Modifier{
			{
				Changes: []modifier.Change{
					modifier.PlusStrength(1),
					modifier.PlusDexterity(1),
					modifier.PlusConstitution(1),
					modifier.PlusIntelligence(1),
					modifier.PlusWisdom(1),
					modifier.PlusCharisma(1),
				},
			},
		},
		Speed: Normal,
		Features: []Feature{
			{
				Name: "Civil Militia",
				WeaponProficiency: &weapon.WeaponProficiency{
					Proficiency: []weapon.Type{
						weapon.Pikes,
						weapon.Spears,
						weapon.Halberds,
						weapon.Glaives,
					},
				},
			},
		},
	}

	Tiefling = Race{
		Index:       "tiefling",
		Name:        "Tiefling",
		Description: "Marked by their infernal heritage, Tieflings are often misunderstood, yet they possess a deep inner strength and an innate connection to dark magic.",
		Traits: []modifier.Modifier{
			{
				Changes: []modifier.Change{
					modifier.PlusCharisma(2),
					modifier.PlusIntelligence(1),
				},
			},
		},
		Speed: Normal,
	}
)

type Index string

type Race struct {
	Index             Index
	Name, Description string
	Traits            []modifier.Modifier
	Features          []Feature
	Speed             Speed // distance in feet per turn
}
