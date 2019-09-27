package main

import "math/rand"

type Player struct {
	Tracking Board
	Primary  Board
}

type Board [BOARD_SIZE][BOARD_SIZE]State

const (
	StateEmpty State = iota
	StateShip
	StateTagged
	StateWreckage
)

type State int

const (
	DirVertical Direction = iota
	DirHorizontal
)

type Direction int

func NewRandomPlayer() *Player {
	p := Player{}

	ships := []int{5, 4, 3, 3, 2, 2, 2}

	for _, shipSize := range ships {
		for tries := 0; tries < 40; tries++ {
			x, y := rand.Intn(BOARD_SIZE), rand.Intn(BOARD_SIZE)
			dir := Direction(rand.Intn(2))
			if p.placeShip(shipSize, x, y, dir) {
				break
			}
		}
	}

	return &p
}

func (p *Player) TrackResult(x, y int, result AtkResult) {
	if !inRange(x) {
		x = 0
	}

	if !inRange(y) {
		y = 0
	}

	if result == AtkHit {
		p.Tracking[x][y] = StateWreckage
	} else if result == AtkMiss {
		p.Tracking[x][y] = StateTagged
	}
}

func (p *Player) ReceiveAttack(x, y int) AtkResult {
	if x < 0 || x >= BOARD_SIZE {
		x = 0
	}

	if y < 0 || y >= BOARD_SIZE {
		y = 0
	}

	switch p.Primary[x][y] {
	case StateEmpty:
		p.Primary[x][y] = StateTagged
		return AtkMiss
	case StateShip:
		p.Primary[x][y] = StateWreckage
		return AtkHit
	default:
		return AtkMiss
	}
}

func (p *Player) placeShip(shipSize int, startx, starty int, direction Direction) bool {
	if !p.Primary.isValidPlacement(startx, starty, shipSize, direction) {
		return false
	}

	for x, y, s := startx, starty, 0; s < shipSize; s++ {
		p.Primary[x][y] = StateShip
		if direction == DirHorizontal {
			y++
		} else {
			x++
		}
	}

	return true
}

// isFree: is the square free for the purposes of placing a ship near it
// Square is free if it's Empty or out of bounds
func (b Board) isFree(x, y int) bool {
	if x < 0 || x > BOARD_SIZE-1 || y < 0 || y > BOARD_SIZE-1 {
		return true
	}

	if b[x][y] == StateEmpty {
		return true
	}

	return false
}

func (g Board) isValidPlacement(x, y, size int, dir Direction) bool {
	var n, m, a, b = x - 1, y - 1, x + size, y + 1
	var along = x

	if dir == DirHorizontal {
		a, b = x+1, y+size
		along = y
	}

	for i := n; i <= a; i++ {
		for j := m; j <= b; j++ {
			if !g.isFree(i, j) {
				return false
			}
		}
	}

	if along < 0 || BOARD_SIZE < along+size {
		return false
	}

	return true
}

func (g Board) transpose() Board {
	ylen, xlen := len(g), len(g[0])
	result := Board{}

	for x := 0; x < xlen; x++ {
		for y := 0; y < ylen; y++ {
			result[x][y] = g[y][x]
		}
	}

	return result
}

// returns true if grids are equal
func compareGrids(grid1, grid2 Board) bool {
	if len(grid1) != len(grid2) {
		return false
	}

	for i := range grid1 {
		if len(grid1[i]) != len(grid2[i]) {
			return false
		}

		for j := range grid2 {
			if grid1[i][j] != grid2[i][j] {
				return false
			}
		}
	}

	return true
}
