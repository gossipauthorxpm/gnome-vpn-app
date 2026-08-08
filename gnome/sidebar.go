package gnome

import (
	"github.com/diamondburned/gotk4-adwaita/pkg/adw"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

type Sidebar struct {
	view *adw.NavigationPage
}

func CreateSideBar() *Sidebar {
	listBoxWidget := buildSidebarList()
	sidebar := adw.NewNavigationPage(listBoxWidget, "Меню")
	return &Sidebar{view: sidebar}
}

func (self *Sidebar) GetView() *adw.NavigationPage {
	return self.view
}

func (self *Sidebar) Collapse() {

}

func (self *Sidebar) Expose() {

}

func buildSidebarList() *adw.PreferencesGroup {

	group := adw.NewPreferencesGroup() // todo: Посмотерть, возможно есть уже готовый компонент для списка сайдбара. Пример настройки Gnome

	connection := gtk.NewButtonWithLabel("Подключение")
	routing := gtk.NewButtonWithLabel("Роутинг")
	provider := gtk.NewButtonWithLabel("Провайдер")
	settings := gtk.NewButtonWithLabel("Настройки")

	group.Add(connection)
	group.Add(routing)
	group.Add(provider)
	group.Add(settings)

	return group
}
