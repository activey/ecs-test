package item

type Rarity string

const (
	Common    Rarity = "Common"
	Uncommon  Rarity = "Uncommon"
	Rare      Rarity = "Rare"
	VeryRare  Rarity = "Very Rare"
	Legendary Rarity = "Legendary"
	StoryItem Rarity = "Story Item"
)

type Type string

func (t Type) String() string {
	return string(t)
}

const (
	Weapon = Type("Weapon")
	Potion = Type("Potion")
)

type Index string

type Item interface {
	Index() Index
	Type() Type
	Name() string
	Description() string
	Rarity() Rarity
	Weight() float64
	Price() uint
}

type BaseItem struct {
	BaseIndex       Index
	BaseType        Type
	BaseName        string
	BaseDescription string
	BaseRarity      Rarity
	BaseWeight      float64
	BasePrice       uint
}

func (b BaseItem) Index() Index {
	return b.BaseIndex
}

func (b BaseItem) Type() Type {
	return b.BaseType
}

func (b BaseItem) Name() string {
	return b.BaseName
}

func (b BaseItem) Description() string {
	return b.BaseDescription
}

func (b BaseItem) Rarity() Rarity {
	return b.BaseRarity
}

func (b BaseItem) Weight() float64 {
	return b.BaseWeight
}

func (b BaseItem) Price() uint {
	return b.BasePrice
}
