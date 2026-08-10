package connection

import (
	"log"

	"github.com/diamondburned/gotk4-adwaita/pkg/adw"
	"gossipauthorxpm.ru/vpn/gnome/factory"
)

type PageBuilder struct{}

func (self *PageBuilder) BuildPageView() *adw.PreferencesPage {
	return buildPage()
}

func buildPage() *adw.PreferencesPage {
	page := adw.NewPreferencesPage()

	buildHeaderPage(page)
	buildCenterPage(page)
	buildFooterPage(page)

	return page
}

func buildHeaderPage(page *adw.PreferencesPage) {

	group := adw.NewPreferencesGroup()
	group.SetTitle("Подключение")
	group.SetDescription("Выберите страну и подключитесь к VPN")

	selectControl, stack := factory.UpSelectControlFactory().CreateSelect(
		factory.SelectContolCreateParams{Title: "Выберите сервер подключения", Subtitle: ""},
	)

	factory.UpSelectControlFactory().InjectButtonInSelectStack(
		factory.SelectContolButtonParams{Title: "Россия", Name: "russia-id"}, stack,
	)

	factory.UpSelectControlFactory().InjectButtonInSelectStack(
		factory.SelectContolButtonParams{Title: "Сша", Name: "usa-id"}, stack,
	)

	factory.UpSelectControlFactory().InjectButtonInSelectStack(
		factory.SelectContolButtonParams{Title: "Китай", Name: "china-id"}, stack,
	)

	factory.UpSelectControlFactory().InjectCallbackOnChangeInSelectControl(func(vsp *adw.ViewStackPage) {
		log.Printf(
			"selected: name=%s title=%s",
			vsp.Name(),
			vsp.Title(),
		)
	}, selectControl)

	group.Add(selectControl)
	page.Add(group)
}

func buildCenterPage(page *adw.PreferencesPage) {

	group := adw.NewPreferencesGroup()
	group.SetTitle("Запуск VPN")

	tunModeSwitcher := adw.NewSwitchRow()
	tunModeSwitcher.SetTitle("TUN mode")
	group.Add(tunModeSwitcher)

	proxyModeSwitcher := adw.NewSwitchRow()
	proxyModeSwitcher.SetTitle("PROXY mode")
	group.Add(proxyModeSwitcher)

	page.Add(group)
}

func buildFooterPage(page *adw.PreferencesPage) {

}
