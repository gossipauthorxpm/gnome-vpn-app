package routing

import (
	"github.com/diamondburned/gotk4-adwaita/pkg/adw"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

type RoutingPage struct {
	page *adw.PreferencesPage
}

func Create() *RoutingPage {
	page := buildPage()
	return &RoutingPage{page: page}
}

func (self *RoutingPage) Get() *adw.PreferencesPage {
	return self.page
}

func (self *RoutingPage) GetNavigationPage() *adw.NavigationPage {
	return adw.NewNavigationPage(self.page, "Маршрутизация")
}

func (self *RoutingPage) InjectInSidebar(stack *adw.ViewStack) {
	routingStackPage := stack.AddTitled(self.GetNavigationPage(), "routing", "Роутинг")
	routingStackPage.SetIconName("network-workgroup-symbolic")
}

func buildPage() *adw.PreferencesPage {
	page := adw.NewPreferencesPage()

	group := adw.NewPreferencesGroup()
	group.SetTitle("Общие")
	group.SetDescription("Настройки внешнего вида и поведения")

	// Switch Row (Переключатель)
	switchRow := adw.NewSwitchRow()
	switchRow.SetTitle("Использовать темные цвета")
	group.Add(switchRow)

	// Combo Row (Выпадающий список)
	comboRow := adw.NewComboRow()
	comboRow.SetTitle("Выберите язык приложения")
	strList := gtk.NewStringList([]string{"Русский", "English", "Deutsch"})
	comboRow.SetModel(strList)
	group.Add(comboRow)

	page.Add(group)
	return page
}
