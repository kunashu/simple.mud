package entity

import (
	"fmt"
	"io"
)

type Attribute int

const (
	STRENGTH Attribute = iota
	HEALTH
	AGILITY
	MAXHITPOINTS
	ACCURACY
	DODGING
	STRIKEDAMAGE
	DAMAGEABSORB
	HPREGEN
)
const NUMATTRIBUTES uint = HPREGEN + 1

func (a Attribute) String() string {
	switch a {
	case STRENGTH:
		return "STRENGTH"
	case HEALTH:
		return "HEALTH"
	case AGILITY:
		return "AGILITY"
	case MAXHITPOINTS:
		return "MAXHITPOINTS"
	case ACCURACY:
		return "ACCURACY"
	case DODGING:
		return "DODGING"
	case STRIKEDAMAGE:
		return "STRIKEDAMAGE"
	case DAMAGEABSORB:
		return "DAMAGEABSORB"
	case HPREGEN:
		return "HPREGEN"
	}

	return fmt.Sprintf("UNKNOWN(%d)", uint(a))
}

type AttributeSet [NUMATTRIBUTES]int
