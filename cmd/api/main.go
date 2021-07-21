package main

import (
	"casorder/cmd"
	"casorder/utils/wsgi"
)

func main() {
	cmd.Initialize()

	wsgi.Initialize()
}
