package retro

import "github.com/hadamrd/retrodata/retrotyp"

type Spell struct {
	Id          int
	Name        string
	Description string
	Levels      []retrotyp.SpellLevel
}
