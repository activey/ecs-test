package rules

import (
	"ecs-test/shared/rules/item"
	"ecs-test/shared/rules/item/weapon"
)

type AllItemsType []item.Item

func (a AllItemsType) CollectItems(items []item.Item) AllItemsType {
	added := append(a, items...)
	return added
}

func (a AllItemsType) ByIndex(index item.Index) item.Item {
	for _, c := range a {
		if c.Index() == index {
			return c
		}
	}
	return nil
}

func (a AllItemsType) PrintAll() {
	for _, w := range a {
		println(w.Name)
	}
}

var AllItems = AllItemsType{}

func init() {
	AllItems = AllItems.CollectItems(weapon.AllClubs)
	AllItems = AllItems.CollectItems(weapon.AllGreatAxes)
}
