package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/gorilla/websocket"
)

const (
	BOARD_SIZE = 10
)

type MsgType string

const (
	MsgAttack MsgType = "attack"
	MsgAnswer         = "answer"
	MsgAck            = "ack"
)

type AtkResult string

const (
	AtkHit  AtkResult = "hit"
	AtkMiss           = "miss"
)

type Message struct {
	MsgType MsgType   `json:"msg_type,omitempty"`
	CoordX  int       `json:"coord_x,omitempty"`
	CoordY  int       `json:"coord_y,omitempty"`
	Result  AtkResult `json:"result,omitempty"`
}

// Game loop
func Game(conn *websocket.Conn) {
	var player = NewRandomPlayer()
	var stdin = bufio.NewReader(os.Stdin)

	reader := make(chan Message)
	writer := make(chan Message)

	go func(c chan Message) {
		for {
			m := ReadMessage(conn)
			c <- m
		}
	}(reader)

	go func(c chan Message) {
		for {
			m := <-c
			WriteMessage(conn, m)
		}
	}(writer)

	var attack Message
	var answer Message
	var ack = Message{MsgType: MsgAck}

	fmt.Println(player)

	// Server goes second and has to wait for the first attack
	if SERVER {
		fmt.Println("waiting for the other player...")

		// Wait for an attack message
		attack = waitForMessage(reader, isAttack)

		// Respond with answer
		writer <- Answer(attack, player)

		fmt.Println(player)

		// Wait for the ack message
		ack = waitForMessage(reader, isAck)
	}

	for {
		// Read input and send an attack
		writer <- Attack(GetInput(stdin))

		// Wait for the answer and mark it on the board
		answer = waitForMessage(reader, isAnswer)
		MarkAnswer(answer, player)

		if answer.Result == AtkHit {
			fmt.Println("it's a hit!")
		}

		fmt.Println(player)

		// Send ack to confirm that the answer was received
		writer <- ack

		fmt.Println("waiting for the other player...")

		// Wait for an attack message
		attack = waitForMessage(reader, isAttack)

		// Respond with answer
		writer <- Answer(attack, player)

		fmt.Println(player)

		// Wait for the ack message
		ack = waitForMessage(reader, isAck)
	}
}

// Wait for the next message that passes validate function
func waitForMessage(reader chan Message, validate func(Message) error) Message {
	var msg Message
	for {
		msg = <-reader
		if err := validate(msg); err != nil {
			log.Printf("invalid message %v, %v\n", msg, err)
			continue
		}

		log.Println("got", msg)
		return msg
	}
}

type Input struct {
	letter byte
	number int
}

func GetInput(r *bufio.Reader) Input {
	var input Input

	for {
		fmt.Printf("please provide input: ")
		x, y, err := readCommand(r)
		if err != nil {
			log.Println(err)
			continue
		}

		input.number, input.letter = x, y
		break
	}

	return input
}

// Attack :: Create and attack message from input
func Attack(input Input) Message {
	x, y := letterToCoord(input.letter), input.number-1

	return Message{
		MsgType: MsgAttack,
		CoordX:  x, CoordY: y,
	}
}

// Answer :: Receive an attack and create answer
func Answer(attack Message, player *Player) Message {
	var msg = Message{
		MsgType: MsgAnswer,
		CoordX:  attack.CoordX,
		CoordY:  attack.CoordY,
	}

	msg.Result = player.ReceiveAttack(attack.CoordX, attack.CoordY)
	return msg
}

// MarkAnswer :: Mark the answer on the board
func MarkAnswer(answer Message, player *Player) {
	player.TrackResult(answer.CoordX, answer.CoordY, answer.Result)
}
