package adapter

import (
	"fmt"

	"gossipauthorxpm.ru/vpn/gnome"
)

func Setup(applicationSid string, osArgs []string) {
	fmt.Println("Bootstrap") // todo: replace with app logger
	gnome.InitShell(applicationSid, osArgs)
}
