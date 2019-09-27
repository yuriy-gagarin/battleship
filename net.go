package main

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

func ReadMessage(conn *websocket.Conn) Message {
	var msg Message
	var err error

	for tries := TRIES; tries > 0; tries-- {
		err = conn.ReadJSON(&msg)
		if err != nil {
			log.Println(err)
			continue
		}

		err = validate(msg)
		if err != nil {
			log.Println(err)
			continue
		}

		break
	}

	if err != nil {
		log.Panic(err)
	}

	return msg
}

func WriteMessage(conn *websocket.Conn, msg Message) {
	var err error

	for tries := TRIES; tries > 0; tries-- {
		err = conn.WriteJSON(&msg)
		if err != nil {
			log.Println(err)
			continue
		}

		err = validate(msg)
		if err != nil {
			log.Println(err)
			continue
		}

		break
	}

	if err != nil {
		log.Panic(err)
	}
}

func validate(msg Message) error {
	switch msg.MsgType {
	case MsgAck:
		return nil
	case MsgAnswer:
		return isAnswer(msg)
	case MsgAttack:
		return isAttack(msg)
	}

	return fmt.Errorf("not a message")
}

func isAnswer(m Message) error {
	if m.MsgType != MsgAnswer {
		return fmt.Errorf("message is not an answer")
	}

	if !inRange(m.CoordX) || !inRange(m.CoordY) {
		return fmt.Errorf("coordinates out of range")
	}

	if m.Result != AtkMiss && m.Result != AtkHit {
		return fmt.Errorf("message has incorrect result")
	}

	return nil
}

func isAttack(m Message) error {
	if m.MsgType != MsgAttack {
		return fmt.Errorf("message is not an attack")
	}

	if !inRange(m.CoordX) || !inRange(m.CoordY) {
		return fmt.Errorf("coordinates out of range")
	}

	return nil
}

func isAck(m Message) error {
	if m.MsgType != MsgAck {
		return fmt.Errorf("message is not ack")
	}

	return nil
}
