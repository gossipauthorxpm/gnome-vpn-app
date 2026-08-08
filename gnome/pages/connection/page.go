package connection

import (
	"github.com/diamondburned/gotk4-adwaita/pkg/adw"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

type ConnectionPage struct {
	page *adw.PreferencesPage
}

func Create() *ConnectionPage {
	page := buildPage()
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

	// Action Row (Кликабельная строка)
	actionRow := adw.NewActionRow()
	actionRow.SetTitle("О программе")
	actionRow.SetSubtitle("Лицензия и авторы")
	actionRow.SetActivatable(true)
	actionRow.AddSuffix(gtk.NewImageFromIconName("go-next-symbolic"))
	group.Add(actionRow)

	page.Add(group)
	return page
}
