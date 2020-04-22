package main

import (
	"runGitDag/routers"
)

func main() {
	//gin.SetMode("release")
	router := routers.InitRouter()
	endPoint := ":3333"
	router.Run(endPoint)
}
