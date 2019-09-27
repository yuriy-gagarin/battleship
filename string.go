package main

import "fmt"

var repr = map[State]string{
	StateEmpty:    "  ",
	StateShip:     "██",
	StateTagged:   "»«",
	StateWreckage: "▒▒",
}

func (d Direction) String() string {
	if d == DirVertical {
		return "Vertical"
	} else {
		return "Horizontal"
	}
}

func (g Board) String() string {
	var str = "  1 2 3 4 5 6 7 8 9 10\n"

	for i := 0; i < BOARD_SIZE; i++ {
		str += fmt.Sprintf("%c ", i+65)

		for j := 0; j < BOARD_SIZE; j++ {
			str += repr[g[i][j]]
		}

		if i < BOARD_SIZE-1 {
			str += "\n"
		}
	}

	return str
}

func (p Player) String() string {
	var str = "  1 2 3 4 5 6 7 8 9 10\t  1 2 3 4 5 6 7 8 9 10\n"

	for i := 0; i < BOARD_SIZE; i++ {
		str += fmt.Sprintf("%c ", i+65)

		for j := 0; j < BOARD_SIZE; j++ {
			str += repr[p.Primary[i][j]]
		}

		str += "\t"
		str += fmt.Sprintf("%c ", i+65)

		for j := 0; j < BOARD_SIZE; j++ {
			str += repr[p.Tracking[i][j]]
		}

		if i < BOARD_SIZE-1 {
			str += "\n"
		}
	}

	return str
}
