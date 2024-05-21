package main

import (
	"fmt"

	"Flo-AI/cmd/internal/viper"
	"Flo-AI/cmd/internal/vosk"
)

func main() {

	viper.LoadConfig()
	vosk.Init()

	fmt.Println("bruh")

}
