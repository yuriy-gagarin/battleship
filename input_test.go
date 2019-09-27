package main

import (
	"bufio"
	"bytes"
	"testing"
)

func TestReadCommand(t *testing.T) {
	mustPass := []struct {
		in   []byte
		outx int
		outy byte
	}{
		{[]byte("A 1\n"), 1, 'A'},
		{[]byte("1 A\n"), 1, 'A'},
		{[]byte("J 10\n"), 10, 'J'},
		{[]byte("   B 2  \n"), 2, 'B'},
		{[]byte("   B 2  \n"), 2, 'B'},
	}

	mustFail := []struct {
		in []byte
	}{
		{[]byte("1 1\n")},
		{[]byte("   K 2  \n")},
		{[]byte("   B\n 2  \n")},
		{[]byte("BA 2\n")},
	}

	for _, tc := range mustPass {
		r := bufio.NewReader(bytes.NewBuffer([]byte(tc.in)))
		x, y, err := readCommand(r)
		if x != tc.outx || y != tc.outy || err != nil {
			t.Fatalf("expected %d %d, got %d %d, with input %s. possible error %v", tc.outx, tc.outy, x, y, tc.in, err)
		}
	}

	for _, tc := range mustFail {
		r := bufio.NewReader(bytes.NewBuffer([]byte(tc.in)))
		x, y, err := readCommand(r)
		if err == nil {
			t.Log(x, y)
			t.Fatalf("input %s must fail. error: %v", tc.in, err)
		}
	}
}
