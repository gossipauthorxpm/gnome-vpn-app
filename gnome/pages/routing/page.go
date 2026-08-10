package routing

import (
	"github.com/diamondburned/gotk4-adwaita/pkg/adw"
	"gossipauthorxpm.ru/vpn/gnome/pages"
)

var viewPageBuilder pages.PageViewBuilder = &PageBuilder{}

type RoutingPage struct {
	page *adw.PreferencesPage
}

func Create() *RoutingPage {
	page := viewPageBuilder.BuildPageView()
	return &RoutingPage{page: page}
}

func (self *RoutingPage) Get() *adw.PreferencesPage {
	return self.page
}

func (self *RoutingPage) GetNavigationPage() *adw.NavigationPage {
	return adw.NewNavigationPage(self.page, "Роутинг")
}

func (self *RoutingPage) InjectInSidebar(stack *adw.ViewStack) {
	routingStackPage := stack.AddTitled(self.GetNavigationPage(), "routing", "Роутинг")
	routingStackPage.SetIconName("network-workgroup-symbolic")
}
