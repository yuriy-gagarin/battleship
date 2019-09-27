package main

import "testing"

func TestTranspose(t *testing.T) {
	var boardBefore = Board{
		{E, E, E, E, E, E, E, E, E, E},
		{E, E, E, E, E, E, E, E, E, E},
		{E, E, E, E, E, E, E, E, E, E},
		{E, E, E, E, E, E, E, E, E, E},
		{E, E, E, E, E, E, E, E, E, E},
		{E, E, E, E, E, E, E, E, E, E},
		{E, E, E, E, E, E, E, E, E, E},
		{E, E, E, E, E, E, E, E, E, E},
		{S, E, E, E, E, E, E, E, E, E},
		{E, E, E, E, E, E, E, E, E, E},
	}

	var boardAfter = Board{
		{E, E, E, E, E, E, E, E, S, E},
		{E, E, E, E, E, E, E, E, E, E},
		{E, E, E, E, E, E, E, E, E, E},
		{E, E, E, E, E, E, E, E, E, E},
		{E, E, E, E, E, E, E, E, E, E},
		{E, E, E, E, E, E, E, E, E, E},
		{E, E, E, E, E, E, E, E, E, E},
		{E, E, E, E, E, E, E, E, E, E},
		{E, E, E, E, E, E, E, E, E, E},
		{E, E, E, E, E, E, E, E, E, E},
	}

	r := boardBefore.transpose()

	if !compareGrids(r, boardAfter) {
		t.Fatalf("expected %v to equal %v", boardAfter, boardBefore)
	}
}

func TestPlaceShip(t *testing.T) {
	var boardBefore = Board{
		{S, S, E, E, S, S, S, E, E, E},
		{E, E, E, E, E, E, E, E, S, E},
		{S, E, E, E, E, S, S, E, S, E},
		{S, E, S, E, E, E, E, E, S, E},
		{E, E, S, E, E, E, E, E, S, E},
		{E, E, S, E, E, E, E, E, E, E},
		{E, E, S, E, E, E, E, S, E, E},
		{S, E, S, E, E, E, E, S, E, E},
		{S, E, E, E, E, E, E, E, E, E},
		{S, E, S, S, S, E, E, E, E, E},
	}

	var boardAfter = Board{
		{S, S, E, E, S, S, S, E, E, E},
		{E, E, E, E, E, E, E, E, S, E},
		{S, E, E, E, E, S, S, E, S, E},
		{S, E, S, E, E, E, E, E, S, E},
		{E, E, S, E, S, E, E, E, S, E},
		{E, E, S, E, S, E, E, E, E, E},
		{E, E, S, E, S, E, E, S, E, E},
		{S, E, S, E, S, E, E, S, E, E},
		{S, E, E, E, E, E, E, E, E, E},
		{S, E, S, S, S, E, E, E, E, E},
	}

	plr := Player{}
	plr.Primary = boardBefore

	if !plr.placeShip(4, 4, 4, DirVertical) {
		t.Fatalf("can't place ship")
	}

	if !compareGrids(plr.Primary, boardAfter) {
		t.Fatalf("expected:\n%vgot:\n%v", boardAfter, plr.Primary)
	}

	if plr.placeShip(4, 4, 4, DirHorizontal) && compareGrids(plr.Primary, boardAfter) {
		t.Fatalf("incorrect ship placement should fail")
	}
}

func TestCheckBounds(t *testing.T) {
	grid := Board{
		{E, E, E, E, E, E, E, E, E, E},
		{E, S, E, E, E, S, E, E, E, E},
		{E, S, E, E, E, S, E, E, E, E},
		{E, S, E, E, E, S, E, E, E, E},
		{E, S, E, E, E, S, E, E, E, E},
		{E, S, E, E, E, S, E, E, E, E},
		{E, E, E, E, E, E, E, E, E, E},
		{E, E, E, E, E, E, E, E, E, E},
		{E, E, E, E, E, E, E, E, E, E},
		{E, E, E, E, E, E, E, E, E, E},
	}

	if grid.isValidPlacement(1, 1, 1, DirVertical) {
		t.Fatal("false true")
	}

	if grid.isValidPlacement(1, 2, 4, DirVertical) {
		t.Fatal("false true")
	}

	if !grid.isValidPlacement(1, 3, 4, DirVertical) {
		t.Fatal("true false")
	}

	if grid.isValidPlacement(1, 3, 4, DirHorizontal) {
		t.Fatal("false true")
	}

	if !grid.isValidPlacement(7, 1, 3, DirVertical) {
		t.Fatal("true false")
	}
}

func TestCompareGrids(t *testing.T) {
	grid1 := Board{
		{S, S, E, E, S, S, S, E, E, E},
		{E, E, E, E, E, E, E, E, S, E},
		{S, E, E, E, E, S, S, E, S, E},
		{S, E, S, E, E, E, E, E, S, E},
		{E, E, S, E, E, E, E, E, S, E},
		{E, E, S, E, E, E, E, E, E, E},
		{E, E, S, E, E, E, E, S, E, E},
		{S, E, S, E, E, E, E, S, E, E},
		{S, E, E, E, E, E, E, E, E, E},
		{S, E, S, S, S, E, E, E, E, E},
	}

	grid2 := Board{
		{S, S, E, E, S, S, S, E, E, E},
		{E, E, E, E, E, E, E, E, S, E},
		{S, E, E, E, E, S, S, E, S, E},
		{S, E, S, E, E, E, E, E, S, E},
		{E, E, S, E, E, E, E, E, S, E},
		{E, E, S, E, E, E, E, E, E, E},
		{E, E, S, E, E, E, E, S, E, E},
		{S, E, S, E, E, E, E, S, E, E},
		{S, E, E, E, E, E, E, E, E, E},
		{S, E, S, S, S, E, E, E, E, E},
	}

	if !compareGrids(grid1, grid2) {
		t.Fatal("expected true, got false")
	}

	grid2[2][2] = S

	if compareGrids(grid1, grid2) {
		t.Fatal("expected false, got true")
	}
}
