package main

import (
	"github.com/SevereCloud/vksdk/v3/events"
)

func handler(obj events.MessageNewObject) string {
	if obj.Message.Text == "ping" {
		return "pong"
	} else if obj.Message.Text == "pong" {
		return "ping"
	} else {
		return "idk"
	}
}
