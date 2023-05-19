package main

import "github.com/VictorOliveiraPy/configs"

func main() {
	///...
	config, _ := configs.LoadConfig(".")
	println(config.DBDriver)
}