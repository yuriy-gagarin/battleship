package main

import (
	"fmt"
	"testing"
)

const (
	E = StateEmpty
	S = StateShip
	T = StateTagged
	W = StateWreckage
)

func TestGridStringer(t *testing.T) {
	var expected = `  1 2 3 4 5 6 7 8 9 10
A ████  ██    ████████
B       ██            
C ██    ██  ██████  ██
D ██  »«            ██
E ██        »«»«      
F ██»«      ██  »«  ▒▒
G       »«  ██      ▒▒
H ▒▒▒▒          »«  ██
I                     
J           ██████████`

	var grid = Board{
		{S, S, E, S, E, E, S, S, S, S},
		{E, E, E, S, E, E, E, E, E, E},
		{S, E, E, S, E, S, S, S, E, S},
		{S, E, T, E, E, E, E, E, E, S},
		{S, E, E, E, E, T, T, E, E, E},
		{S, T, E, E, E, S, E, T, E, W},
		{E, E, E, T, E, S, E, E, E, W},
		{W, W, E, E, E, E, E, T, E, S},
		{E, E, E, E, E, E, E, E, E, E},
		{E, E, E, E, E, S, S, S, S, S},
	}

	s := fmt.Sprint(grid)

	if s != expected {
		t.Fatalf("expected:%s\ngot:%s", expected, s)
	}
}

func TestPlayerStringer(t *testing.T) {
	var expected = `  1 2 3 4 5 6 7 8 9 10	  1 2 3 4 5 6 7 8 9 10
A ████  ██    ████████	A                     
B       ██            	B                     
C ██    ██  ██████  ██	C                     
D ██  »«            ██	D     »«              
E ██        »«»«      	E           »«»«      
F ██»«      ██  »«  ▒▒	F   »«          »«  ▒▒
G       »«  ██      ▒▒	G       »«          ▒▒
H ▒▒▒▒          »«  ██	H ▒▒▒▒          »«    
I                     	I                     
J           ██████████	J                     `

	var primaryGrid = Board{
		{S, S, E, S, E, E, S, S, S, S},
		{E, E, E, S, E, E, E, E, E, E},
		{S, E, E, S, E, S, S, S, E, S},
		{S, E, T, E, E, E, E, E, E, S},
		{S, E, E, E, E, T, T, E, E, E},
		{S, T, E, E, E, S, E, T, E, W},
		{E, E, E, T, E, S, E, E, E, W},
		{W, W, E, E, E, E, E, T, E, S},
		{E, E, E, E, E, E, E, E, E, E},
		{E, E, E, E, E, S, S, S, S, S},
	}

	var trackingGrid = Board{
		{E, E, E, E, E, E, E, E, E, E},
		{E, E, E, E, E, E, E, E, E, E},
		{E, E, E, E, E, E, E, E, E, E},
		{E, E, T, E, E, E, E, E, E, E},
		{E, E, E, E, E, T, T, E, E, E},
		{E, T, E, E, E, E, E, T, E, W},
		{E, E, E, T, E, E, E, E, E, W},
		{W, W, E, E, E, E, E, T, E, E},
		{E, E, E, E, E, E, E, E, E, E},
		{E, E, E, E, E, E, E, E, E, E},
	}

	var player = Player{Primary: primaryGrid, Tracking: trackingGrid}

	s := fmt.Sprint(player)

	if s != expected {
		t.Fatalf("expected:%s\ngot:%s", expected, s)
	}
}
