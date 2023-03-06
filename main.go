package main

import (
	"ward-go/config"
	"ward-go/handler"
	"ward-go/utils"
)

func main() {
	utils.StartGinServer(config.Wg, handler.InitRouter)
	config.Wg.Wait()
}
