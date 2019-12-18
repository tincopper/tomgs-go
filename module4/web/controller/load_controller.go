package controller

import "tomgs-go/module4/web"

func LoadControllers() {
    web.AddRoute("/basic", basicHandler)
    web.AddRoute("/demo", demoHandler)
}