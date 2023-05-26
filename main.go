package main

import (
	"fmt"

	// Routes
	"elf-talk/config"
	"elf-talk/routes"
)

func main() {
	fmt.Println("Server runnning on", config.Conf.Port, ".")

	r := routes.SetupRouter()
	r.Run(fmt.Sprintf(":%d", config.Conf.Port))
}
