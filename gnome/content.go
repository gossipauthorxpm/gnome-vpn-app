package gnome

import (
	"github.com/diamondburned/gotk4-adwaita/pkg/adw"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

type Content struct {
	mainPage *adw.NavigationPage
}

func CreateContent() *Content {
	page := buildPage()
	mainPage := adw.NewNavigationPage(page, "Главаня страница")
	return &Content{mainPage: mainPage}
}

func (self *Content) GetMainPage() *adw.NavigationPage {
	return self.mainPage
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
