package connection

import (
	"github.com/diamondburned/gotk4-adwaita/pkg/adw"
	"gossipauthorxpm.ru/vpn/gnome/pages"
)

var viewPageBuilder pages.PageViewBuilder = &PageBuilder{}

type ConnectionPage struct {
	page *adw.PreferencesPage
}

func Create() *ConnectionPage {
	page := viewPageBuilder.BuildPageView()
	return &ConnectionPage{page: page}
}

func (self *ConnectionPage) Get() *adw.PreferencesPage {
	return self.page
}

func (self *ConnectionPage) GetNavigationPage() *adw.NavigationPage {
	return adw.NewNavigationPage(self.page, "Подключение")
}

func (self *ConnectionPage) InjectInSidebar(stack *adw.ViewStack) {
	connectionStackPage := stack.AddTitled(self.GetNavigationPage(), "connection", "Подключение")
	connectionStackPage.SetIconName("network-wired-symbolic")
}
