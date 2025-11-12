package retro

import (
	"github.com/hadamrd/retrodata/retrotyp"
)

type Market struct {
	Id            string
	GameServerId  int
	Quantity1     int
	Quantity2     int
	Quantity3     int
	Types         []retrotyp.ItemType
	Fee           float32
	MaxLevel      int
	MaxPerAccount int
	MaxHours      int
}
