package rules

import (
	"ecs-test/shared/rules/race"
)

type AllRacesType []*race.Race

func (a AllRacesType) Names() []string {
	var names []string
	for _, c := range a {
		names = append(names, c.Name)
	}
	return names
}

func (a AllRacesType) Values() []*race.Race {
	var races []*race.Race
	for _, r := range a {
		races = append(races, r)
	}
	return races
}

func (a AllRacesType) ByIndex(index race.Index) *race.Race {
	for _, c := range a {
		if c.Index == index {
			return c
		}
	}
	return nil
}

var AllRaces = AllRacesType{
	&race.DwarfHill,
	&race.DwarfMountain,
	&race.ElfHigh,
	&race.ElfWood,
	&race.ElfDrow,
	&race.HalflingLightfoot,
	&race.HalflingStout,
	&race.Human,
	&race.Dragonborn,
	&race.GnomeForest,
	&race.GnomeRock,
	&race.HalfElf,
	&race.HalfOrc,
	&race.Tiefling,
}
