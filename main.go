package main

import (
	"skrshop-api/core"
	"skrshop-api/model"
	"skrshop-api/route"
)

func main() {
	r := route.InitRouter()

	model.InitDB()
	_ = core.InitRedis()
	core.InitValidate()
	r.Run(core.HTTPPort)
}
