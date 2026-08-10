package gnome

import (
	"gossipauthorxpm.ru/vpn/gnome/pages"
	"gossipauthorxpm.ru/vpn/gnome/pages/connection"
	"gossipauthorxpm.ru/vpn/gnome/pages/routing"
	"gossipauthorxpm.ru/vpn/gnome/pages/subscription"
)

type Content struct {
	connection   pages.Page
	routing      pages.Page
	subscription pages.Page
}

func CreateContent() *Content {
	connectionPage := connection.Create()
	routingPage := routing.Create()
	subscription := subscription.Create()
	return &Content{connection: connectionPage, routing: routingPage, subscription: subscription}
}

func (self *Content) GetPages() []pages.Page {
	return []pages.Page{self.connection, self.routing, self.subscription}
}

func (self *Content) GetConnectionPage() *pages.Page {
	return &self.connection
}

func (self *Content) GetRoutingPage() *pages.Page {
	return &self.routing
}
