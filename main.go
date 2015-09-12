package main

import (
	"github.com/Unknwon/log"
	"github.com/Unknwon/macaron"

	"github.com/Unknwon/fws/modules/middleware"
	"github.com/Unknwon/fws/modules/setting"
	"github.com/Unknwon/fws/routers/v1"
)

const APP_VER = "0.0.1.0912"

func main() {
	log.Info("App Version: %s", APP_VER)
	m := macaron.Classic()
	m.Use(macaron.Renderer())
	m.Use(middleware.Contexter())

	m.Get("/fibonacci", v1.Fibonacci)

	m.Run(setting.HTTPPort)
}
