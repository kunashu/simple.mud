package entity

type Money uint64

type ItemType int

const (
	WEAPON ItemType = iota
	ARMOR
	HEALING
)

type Item struct {
	baseEntity
	itemType ItemType
	// Weapon: min damage used
	// Healing: min damage healed
	min int
	// Weapon: max damage used
	// Healing: max damage healed
	max int
	// Weapon: pause between swings in secs.
	speed      int
	price      Money
	attributes AttributeSet
}
