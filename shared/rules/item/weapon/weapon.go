package weapon

import (
	"ecs-test/shared/rules/item"
)

type Category uint64

const (
	Simple Category = iota
	Martial
)

type Type struct {
	Category Category
	Name     string
}

func NewSimpleType(name string) Type {
	return Type{
		Category: Simple,
		Name:     name,
	}
}

func NewMartialType(name string) Type {
	return Type{
		Category: Martial,
		Name:     name,
	}
}

var (
	Clubs          = NewSimpleType("Clubs")
	Daggers        = NewSimpleType("Daggers")
	Handaxes       = NewSimpleType("Handaxes")
	Javelins       = NewSimpleType("Javelins")
	LightHammers   = NewSimpleType("Light Hammers")
	Maces          = NewSimpleType("Maces")
	Sickles        = NewSimpleType("Sickles")
	Quarterstaves  = NewSimpleType("Quarterstaves")
	Spears         = NewSimpleType("Spears")
	Greatclubs     = NewSimpleType("Greatclubs")
	LightCrossbows = NewSimpleType("Light Crossbows")
	Shortbows      = NewSimpleType("Shortbows")

	Greataxes      = NewMartialType("Greataxes")
	Flails         = NewMartialType("Flails")
	Morningstars   = NewMartialType("Morningstars")
	Scimitars      = NewMartialType("Scimitars")
	WarPicks       = NewMartialType("War Picks")
	Tridents       = NewMartialType("Tridents")
	Glaives        = NewMartialType("Glaives")
	Shortswords    = NewMartialType("Shortswords")
	Longswords     = NewMartialType("Longswords")
	Greatswords    = NewMartialType("Greatswords")
	Halberds       = NewMartialType("Halberds")
	Mauls          = NewMartialType("Mauls")
	Pikes          = NewMartialType("Pikes")
	Rapiers        = NewMartialType("Rapiers")
	Battleaxes     = NewMartialType("Battleaxes")
	Warhammers     = NewMartialType("Warhammers")
	HandCrossbows  = NewMartialType("Hand Crossbows")
	HeavyCrossbows = NewMartialType("Heavy Crossbows")
	Longbows       = NewMartialType("Longbows")
)

type Weapon struct {
	item.BaseItem
	Damage     []item.Damage
	WeaponType Type
}
