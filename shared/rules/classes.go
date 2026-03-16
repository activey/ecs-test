package rules

import (
	"ecs-test/shared/rules/class"
)

type AllClassesType []*class.Class

func (a AllClassesType) Names() []string {
	var names []string
	for _, c := range a {
		names = append(names, c.Name)
	}
	return names
}

func (a AllClassesType) Classes() []*class.Class {
	var names []*class.Class
	for _, c := range a {
		names = append(names, c)
	}
	return names
}

func (a AllClassesType) ByIndex(index class.Index) *class.Class {
	for _, c := range a {
		if c.Index == index {
			return c
		}
	}
	return nil
}

var AllClasses = AllClassesType{
	&class.Barbarian,
	&class.Bard,
	&class.Cleric,
	&class.Druid,
	&class.Fighter,
	&class.Monk,
	&class.Paladin,
	&class.Ranger,
	&class.Rogue,
	&class.Sorcerer,
	&class.Warlock,
	&class.Wizard,
}
