package main

import (
	"os"

	"gossipauthorxpm.ru/vpn/bootstrap"
)

const applicationSid = "gossipauthorxpm.ru.gnome-vpn"

func main() {
	bootstrap.Setup(applicationSid, os.Args)
}
