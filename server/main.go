package main

import (
	"os"
	"server/global"
	"server/http"
)

func main() {
	global.Init()
	if err := http.Run(); err != nil {
		os.Exit(1)
	}
}
