package gnome

import (
	"gossipauthorxpm.ru/vpn/gnome/pages"
	"gossipauthorxpm.ru/vpn/gnome/pages/connection"
	"gossipauthorxpm.ru/vpn/gnome/pages/routing"
)

type Content struct {
	connection pages.Page
	routing    pages.Page
}

func CreateContent() *Content {
	connectionPage := connection.Create()
	routingPage := routing.Create()
	return &Content{connection: connectionPage, routing: routingPage}
}

func (self *Content) GetPages() []pages.Page {
	return []pages.Page{self.connection, self.routing}
}

func (self *Content) GetConnectionPage() *pages.Page {
	return &self.connection
}

func (self *Content) GetRoutingPage() *pages.Page {
	return &self.routing
}
