package retro

import "github.com/hadamrd/retrodata/retrotyp"

type NPCTemplate struct {
	Id      int
	Name    string
	Actions []retrotyp.NPCAction
}
