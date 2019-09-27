package main

import (
	"bufio"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func inRange(c int) bool {
	if c < 0 || BOARD_SIZE-1 < c {
		return false
	}

	return true
}

func letterToCoord(c byte) int {
	if 96 < c && c < 107 {
		return int(c - 97)
	}

	if 64 < c && c < 75 {
		return int(c - 65)
	}

	return 0
}

func letterIsCoord(c byte) bool {
	if (96 < c && c < 107) || (64 < c && c < 75) {
		return true
	}

	return false
}

func readCommand(r *bufio.Reader) (x int, y byte, err error) {
	str, err := r.ReadString('\n')
	if err != nil {
		return x, y, err
	}

	strs := strings.Split(strings.Trim(str, " \t\r\n"), " ")
	if len(strs) < 2 {
		return x, y, errors.New("not enough args")
	}

	letter, number := strs[0], strs[1]

	if !letterIsCoord(letter[0]) {
		letter, number = number, letter
	}

	if len(letter) > 1 || !letterIsCoord(letter[0]) {
		return x, y, fmt.Errorf("invalid args: %s is not a valid letter", letter)
	}

	num, err := strconv.Atoi(number)
	if err != nil {
		return x, y, fmt.Errorf("invalid args: parsing %s as integer: %v", number, err)
	}

	return num, letter[0], nil
}
