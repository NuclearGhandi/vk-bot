package main

import (
	"context"
	"log"

	"github.com/SevereCloud/vksdk/v3/api/params"
	"github.com/SevereCloud/vksdk/v3/events"
)

func NewMessageHandler(_ context.Context, obj events.MessageNewObject) {
	log.Printf("%d: %s", obj.Message.PeerID, obj.Message.Text)
	b := params.NewMessagesSendBuilder()

	b.Message("123")
	b.RandomID(0)
	b.PeerID(obj.Message.PeerID)

	_, err := rnt.vk.MessagesSend(b.Params)
	if err != nil {
		log.Fatal(err)
	}
}
