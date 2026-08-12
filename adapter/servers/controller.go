package servers

import (
	"log"

	"gossipauthorxpm.ru/vpn/adapter/subscription"
)

type ServersController struct {
}

func (self *ServersController) GetSubscibtionServers(controller *subscription.SubscriptionController) []ServerVisual {
	return []ServerVisual{
		{Title: "Россия", Name: "russia-id"},
		{Title: "Сша", Name: "usa-id"},
		{Title: "Китай", Name: "china-id"},
	}
}

func (self *ServersController) ChangeServerCallback(server *ServerVisual) {
	log.Printf(
		"selected: name=%s title=%s",
		server.Name,
		server.Title,
	)
}
