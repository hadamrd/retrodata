package retro

import (
	"errors"
	"strings"
)

const (
	cellWidth      = 53.0
	cellHeight     = 27.0
	cellHalfWidth  = 26.5
	cellHalfHeight = 13.5
	levelHeight    = 20.0
)

var zipKey = []string{"_a", "_b", "_c", "_d", "_e", "_f", "_g", "_h", "_i", "_j", "_k", "_l", "_m", "_n", "_o", "_p",
	"_q", "_r", "_s", "_t", "_u", "_v", "_w", "_x", "_y", "_z", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K",
	"L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "0", "1", "2", "3", "4", "5", "6", "7",
	"8", "9", "-", "_"}

var zkArray = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s",
	"t", "u", "v", "w", "x", "y", "z", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P",
	"Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "-", "_"}

var hashCodes map[string]int

func init() {
	hashCodes = make(map[string]int)
	for j := len(zkArray) - 1; j >= 0; j-- {
		hashCodes[zkArray[j]] = j
	}
}

func decode64(codedValue string) int {
	return hashCodes[codedValue]
}

func encode64(value int) string {
	return zkArray[value]
}

type Cell struct {
	Id                             int
	Active                         bool
	LineOfSight                    bool
	LayerGroundRot                 int
	GroundLevel                    int
	Movement                       int
	LayerGroundNum                 int
	GroundSlope                    int
	LayerGroundFlip                bool
	LayerObject1Num                int
	LayerObject1Rot                int
	LayerObject1Flip               bool
	LayerObject2Flip               bool
	LayerObject2Interactive        bool
	LayerObject2Num                int
	PermanentLevel                 int
	LayerObjectExternal            string
	LayerObjectExternalInteractive bool
	X                              float64
	Y                              float64
	SpriteOnId                     []int
}

func defaultGameMapCell() Cell {
	return Cell{
		Id:                      0,
		Active:                  true,
		LineOfSight:             true,
		LayerGroundRot:          0,
		GroundLevel:             7,
		Movement:                4,
		LayerGroundNum:          0,
		GroundSlope:             1,
		LayerGroundFlip:         false,
		LayerObject1Num:         0,
		LayerObject1Rot:         0,
		LayerObject1Flip:        false,
		LayerObject2Flip:        false,
		LayerObject2Interactive: false,
		LayerObject2Num:         0,
		X:                       0,
		Y:                       0,
		SpriteOnId:              nil,
	}
}

func getCellHeight(groundSlope, groundLevel int) float64 {
	loc4 := 0.5
	if groundSlope == 1 {
		loc4 = 0
	}

	loc5 := groundLevel - 7

	return float64(loc5) + loc4
}

func DecompressCells(data string, forced bool) ([]Cell, error) {
	if len(data)%10 != 0 {
		return nil, errors.New("invalid length of data")
	}

	cells := make([]Cell, len(data)/10)

	n := 0
	for i := 0; n < len(data); i++ {
		cell, err := decompressCell(data[n:n+10], forced, 0)
		if err != nil {
			return nil, err
		}
		cell.Id = i

		cells[i] = cell

		n += 10
	}

	return cells, nil
}

func decompressCell(data string, forced bool, permanentLevel int) (Cell, error) {
	cell := defaultGameMapCell()

	loc7 := len(data) - 1
	loc8 := make([]int, len(data))
	for loc7 >= 0 {
		loc8[loc7] = hashCodes[string(data[loc7])]
		loc7--
	}

	if (loc8[0]&32)>>5 == 0 {
		cell.Active = false
	} else {
		cell.Active = true
	}

	if cell.Active || forced {
		cell.PermanentLevel = permanentLevel

		if loc8[0]&1 == 0 {
			cell.LineOfSight = false
		} else {
			cell.LineOfSight = true
		}

		cell.LayerGroundRot = (loc8[1] & 48) >> 4
		cell.GroundLevel = loc8[1] & 15
		cell.Movement = (loc8[2] & 56) >> 3
		cell.LayerGroundNum = ((loc8[0] & 24) << 6) + ((loc8[2] & 7) << 6) + loc8[3]
		cell.GroundSlope = (loc8[4] & 60) >> 2

		if ((loc8[4] & 2) >> 1) == 0 {
			cell.LayerGroundFlip = false
		} else {
			cell.LayerGroundFlip = true
		}

		cell.LayerObject1Num = ((loc8[0] & 4) << 11) + ((loc8[4] & 1) << 12) + (loc8[5] << 6) + loc8[6]
		cell.LayerObject1Rot = (loc8[7] & 48) >> 4

		if (loc8[7]&8)>>3 == 0 {
			cell.LayerObject1Flip = false
		} else {
			cell.LayerObject1Flip = true
		}

		if (loc8[7]&4)>>2 == 0 {
			cell.LayerObject2Flip = false
		} else {
			cell.LayerObject2Flip = true
		}

		if (loc8[7]&2)>>1 == 0 {
			cell.LayerObject2Interactive = false
		} else {
			cell.LayerObject2Interactive = true
		}

		cell.LayerObject2Num = ((loc8[0] & 2) << 12) + ((loc8[7] & 1) << 12) + (loc8[8] << 6) + loc8[9]
		cell.LayerObjectExternal = ""
		cell.LayerObjectExternalInteractive = false
	}

	return cell, nil
}

func compressCells(cells []Cell) string {
	sb := &strings.Builder{}

	for _, cell := range cells {
		sb.WriteString(compressCell(cell))
	}

	return sb.String()
}

func compressCell(cell Cell) string {
	sb := &strings.Builder{}

	loc4 := make([]int, 10)

	active := 0
	if cell.Active {
		active = 1
	}
	loc4[0] = active << 5

	lineOfSight := 0
	if cell.LineOfSight {
		lineOfSight = 1
	}
	loc4[0] = loc4[0] | lineOfSight

	loc4[0] = loc4[0] | (cell.LayerGroundNum&1536)>>6
	loc4[0] = loc4[0] | (cell.LayerObject1Num&8192)>>11
	loc4[0] = loc4[0] | (cell.LayerObject2Num&8192)>>12

	loc4[1] = (cell.LayerGroundRot & 3) << 4
	loc4[1] = loc4[1] | cell.GroundLevel&15

	loc4[2] = (cell.Movement & 7) << 3
	loc4[2] = loc4[2] | cell.LayerGroundNum>>6&7

	loc4[3] = cell.LayerGroundNum & 63

	loc4[4] = (cell.GroundSlope & 15) << 2

	layerGroundFlip := 0
	if cell.LayerGroundFlip {
		layerGroundFlip = 1
	}
	loc4[4] = loc4[4] | layerGroundFlip<<1

	loc4[4] = loc4[4] | cell.LayerObject1Num>>12&1

	loc4[5] = cell.LayerObject1Num >> 6 & 63

	loc4[6] = cell.LayerObject1Num & 63

	loc4[7] = (cell.LayerObject1Rot & 3) << 4

	layerObject1Flip := 0
	if cell.LayerObject1Flip {
		layerObject1Flip = 1
	}
	loc4[7] = loc4[7] | layerObject1Flip<<3

	layerObject2Flip := 0
	if cell.LayerObject2Flip {
		layerObject2Flip = 1
	}
	loc4[7] = loc4[7] | layerObject2Flip<<2

	layerObject2Interactive := 0
	if cell.LayerObject2Interactive {
		layerObject2Interactive = 1
	}
	loc4[7] = loc4[7] | layerObject2Interactive<<1

	loc4[7] = loc4[7] | cell.LayerObject2Num>>12&1

	loc4[8] = cell.LayerObject2Num >> 6 & 63

	loc4[9] = cell.LayerObject2Num & 63

	for _, v := range loc4 {
		sb.WriteString(encode64(v))
	}

	return sb.String()
}

func BuiltCells(cellNum *int, buildAll bool, gameMapWidth int, cells []Cell) []Cell {
	out := make([]Cell, len(cells))
	copy(out, cells)

	loc9 := -1
	loc10 := 0
	loc11 := 0.0
	loc12 := out
	loc13 := len(loc12)
	loc14 := gameMapWidth - 1

	loc16 := false
	if cellNum != nil {
		loc16 = true
	}

	for loc20 := 0; loc20 < loc13; loc20++ {
		if loc9 == loc14 {
			loc9 = 0
			loc10++
			if loc11 == 0 {
				loc11 = cellHalfWidth
				loc14--
			} else {
				loc11 = 0
				loc14++
			}
		} else {
			loc9++
		}

		if loc16 {
			if loc20 < *cellNum {
				continue
			} else if loc20 > *cellNum {
				return out
			}
		}

		loc21 := loc12[loc20]
		if loc21.Active {
			loc22 := float64(loc9)*cellWidth + loc11
			loc23 := float64(loc10)*cellHalfHeight - cellHeight*float64(loc21.GroundLevel-7)
			loc21.X = loc22
			loc21.Y = loc23

			loc12[loc20] = loc21

			continue
		}

		if buildAll {
			var loc32 = float64(loc9)*cellWidth + loc11
			var loc33 = float64(loc10) * cellHalfHeight
			loc21.X = loc32
			loc21.Y = loc33

			loc12[loc20] = loc21
		}
	}

	return out
}
