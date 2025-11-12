package retro

import "github.com/hadamrd/retrodata/retrotyp"

type EffectTemplate struct {
	Id               int
	Description      string
	Dice             bool
	Operator         retrotyp.EffectOperator
	CharacteristicId retrotyp.CharacteristicId
	Element          retrotyp.EffectElement
}
