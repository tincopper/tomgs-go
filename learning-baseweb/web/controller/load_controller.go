package controller

import "tomgs-go/learning-baseweb/web"

func LoadControllers() {
    web.AddRoute("/basic", basicHandler)
    web.AddRoute("/demo-base", demoHandler)
}