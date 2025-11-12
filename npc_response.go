package retro

import "github.com/hadamrd/retrodata/retrotyp"

type NPCResponse struct {
	Id         int
	Text       string
	Action     retrotyp.NPCResponseAction
	Arguments  []string
	Conditions []string // TODO
}
