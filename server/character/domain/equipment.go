package domain

import (
	"ecs-test/shared/rules/item"
)

//type EquipmentTarget func(item item.Item, equipment *Equipment) error

//var (
//	PutSlotItem = func(item item.Item, slot *EquipmentSlot, equipment *Equipment) error {
//		if !slot.CanKeepItem(item) {
//			return fmt.Errorf("unable to equip item %s in slot %s", item.Name, slot.Name)
//		}
//		equipment.PutItemIntoSlot(item, slot)
//		return nil
//	}
//
//	SlotMainHand = func(item item.Item, equipment *Equipment) error {
//		return PutSlotItem(item, equipment.MainHandSlot, equipment)
//	}
//
//	SlotOffHand = func(item item.Item, equipment *Equipment) error {
//		return PutSlotItem(item, equipment.OffHandSlot, equipment)
//	}
//)

type EquipmentSlot struct {
	Name         string
	EquippedItem item.Item
	AllowedKinds []item.Type
}

//func (es *EquipmentSlot) PutSlotItem(item *item.Item) {
//	es.EquippedItem = item
//}
//
//func (es *EquipmentSlot) CanKeepItem(item item.Item) bool {
//	//for _, allowedKind := range es.AllowedKinds {
//	//	for _, kind := range item.Kinds() {
//	//		if allowedKind == kind {
//	//			return true
//	//		}
//	//	}
//	//}
//	return false
//}
//
//func (es *EquipmentSlot) IsEmpty() bool {
//	return es.EquippedItem == nil
//}

// Equipment ---------------------------
type Equipment struct {
	AllItems     []item.Item
	HeadSlot     *EquipmentSlot
	MainHandSlot *EquipmentSlot
	OffHandSlot  *EquipmentSlot
	ChestSlot    *EquipmentSlot
	LegsSlot     *EquipmentSlot
	FeetSlot     *EquipmentSlot

	totalWeight float64
}

//func (e *Equipment) Equip(item item.Item, equipAtTarget EquipmentTarget) error {
//	return equipAtTarget(item, e)
//}

//func (e *Equipment) PutItemIntoSlot(item item.Item, slot *EquipmentSlot) {
//	slot.PutSlotItem(&item)
//}

type ItemPredicate func(item item.Item) bool

func (e *Equipment) FindItem(predicate ItemPredicate) item.Item {
	for _, i := range e.AllItems {
		if predicate(i) {
			return i
		}
	}
	return nil
}

func (e *Equipment) CollectItem(item item.Item) {
	e.AllItems = append(e.AllItems, item)
	e.totalWeight = e.totalWeight + item.Weight()
}

func (e *Equipment) TotalWeight() float64 {
	return e.totalWeight
}

func NewEquipment() Equipment {
	return Equipment{
		HeadSlot: &EquipmentSlot{
			Name: "Head",
		},
		MainHandSlot: &EquipmentSlot{
			Name: "Main hand weapon",
		},
		OffHandSlot: &EquipmentSlot{
			Name: "Off hands weapon",
		},
		ChestSlot: &EquipmentSlot{
			Name: "Chest",
		},
		LegsSlot: &EquipmentSlot{
			Name: "Legs",
		},
		FeetSlot: &EquipmentSlot{
			Name: "Feet",
		},
	}
}
