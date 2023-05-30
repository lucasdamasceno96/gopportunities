package main

import (
	"gopportunities/config"
	"gopportunities/router"
)

var (
	logger *config.Logger
) 

func main() {
	logger = config.GetLogger("main")

	//initialize config
	err := config.Init()
	
	if err != nil{
	logger.Error(err)
	return
	}

	// initialize router 
	router.Initialize()
}
