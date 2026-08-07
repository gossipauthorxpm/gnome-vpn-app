package main

import (
	"os"

	"gossipauthorxpm.ru/vpn/adapter"
)

const applicationSid = "gossipauthorxpm.ru.gnome-vpn"

func main() {
	adapter.Setup(applicationSid, os.Args)
}
