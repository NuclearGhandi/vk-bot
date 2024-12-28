package main

import (
	"log"

	"github.com/SevereCloud/vksdk/v3/api"

	"github.com/SevereCloud/vksdk/v3/longpoll-bot"
)

func longpollInit() {
	rnt.vk = api.NewVK(cfg.VKTOKEN)

	// get information about the group
	group, err := rnt.vk.GroupsGetByID(nil)
	if err != nil {
		log.Fatal(err)
	}

	// Initializing Long Poll
	lp, err := longpoll.NewLongPoll(rnt.vk, group.Groups[0].ID)
	if err != nil {
		log.Fatal(err)
	}

	// New message event
	lp.MessageNew(NewMessageHandler)

	// Run Bots Long Poll
	log.Println("Start Long Poll")
	if err := lp.Run(); err != nil {
		log.Fatal(err)
	}
}
