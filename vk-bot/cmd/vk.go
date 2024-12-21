package main

import (
	"context"
	"log"

	"github.com/SevereCloud/vksdk/v3/api"
	"github.com/SevereCloud/vksdk/v3/api/params"
	"github.com/SevereCloud/vksdk/v3/events"
	"github.com/SevereCloud/vksdk/v3/longpoll-bot"
)

func longpollInit() {
	token := "vk1.a.8RlrGVwJCs7TSx55OaJHKvTTnGFCZ_fiba-t9rWJWIRtRDuxSWmH-Vlh18Np7jiRLur5-4suJNbThrT9k17RMZRjMMpe0qShwvfvv78xBIg9JCcQ2PG3NXjPegil18ZBWrex3EIL67xiDjb_ZNcbO4B-8B_IPDhR0d_kJVjWCAnH4f4P3MxZn2fq9x_wRobVZ8BRyoYhWsNPZQmzZFqmJw"
	vk := api.NewVK(token)

	// get information about the group
	group, err := vk.GroupsGetByID(nil)
	if err != nil {
		log.Fatal(err)
	}

	// Initializing Long Poll
	lp, err := longpoll.NewLongPoll(vk, group.Groups[0].ID)
	if err != nil {
		log.Fatal(err)
	}

	// New message event
	lp.MessageNew(func(_ context.Context, obj events.MessageNewObject) {
		log.Printf("%d: %s", obj.Message.PeerID, obj.Message.Text)
		b := params.NewMessagesSendBuilder()
		b.Message(handler(obj))
		b.RandomID(0)
		b.PeerID(obj.Message.PeerID)

		_, err := vk.MessagesSend(b.Params)
		if err != nil {
			log.Fatal(err)
		}
	})

	// Run Bots Long Poll
	log.Println("Start Long Poll")
	if err := lp.Run(); err != nil {
		log.Fatal(err)
	}
}
