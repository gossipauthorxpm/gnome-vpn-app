package subscription

import (
	"github.com/diamondburned/gotk4-adwaita/pkg/adw"
	"gossipauthorxpm.ru/vpn/gnome/pages"
)

var viewPageBuilder pages.PageViewBuilder = &PageBuilder{}

type SubcribtionPage struct {
	page *adw.PreferencesPage
}

func Create() *SubcribtionPage {
	page := viewPageBuilder.BuildPageView()
	return &SubcribtionPage{page: page}
}

func (self *SubcribtionPage) Get() *adw.PreferencesPage {
	return self.page
}

func (self *SubcribtionPage) GetNavigationPage() *adw.NavigationPage {
	return adw.NewNavigationPage(self.page, "Управление подпиской")
}

func (self *SubcribtionPage) InjectInSidebar(stack *adw.ViewStack) {
	routingStackPage := stack.AddTitled(self.GetNavigationPage(), "subscribtion", "Управление подпиской")
	routingStackPage.SetIconName("user-home-symbolic")
}
