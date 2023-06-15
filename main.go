package main

import (
	"mygin/model"
	"mygin/routes"
)

func main() {
	//初始化数据库
	model.InitDb()
	routes.InitRouter()
}
