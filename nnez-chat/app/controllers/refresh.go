package controllers

import (
	"github.com/robfig/revel"
	"nnez-contacts/nnez-chat/app/chatroom"
)

type Refresh struct {
	*revel.Controller
}

func (c Refresh) Index(user string) revel.Result {
	chatroom.Join(user)
	return c.Room(user)
}

func (c Refresh) Room(user string) revel.Result {
	subscription := chatroom.Subscribe()
	defer subscription.Cancel()
	events := subscription.Archive
	for i, _ := range events {
		if events[i].User == user {
			events[i].User = "you"
		}
	}
	return c.Render(user, events)
}

func (c Refresh) Say(user, message string) revel.Result {
	chatroom.Say(user, message)
	return c.Room(user)
}

func (c Refresh) Leave(user string) revel.Result {
	chatroom.Leave(user)
	return c.Redirect(Application.Index)
}
