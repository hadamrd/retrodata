package retro

import "github.com/hadamrd/retrodata/retrotyp"

type CharacterItem struct {
	Item
	Position    retrotyp.CharacterItemPosition
	CharacterId int
}
