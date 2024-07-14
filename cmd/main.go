package main

import (
	"fmt"
	"log"
	"os"

	"Flo-AI/cmd/internal/viper"
	"Flo-AI/cmd/internal/whisperstt"
)

func main() {

	if err := viper.LoadConfig(); err != nil {
		panic(err.Error())
	}
	// if err := openai.LoadConfig(); err != nil {
	// 	panic(err.Error())
	// }

	// Initialize Whisper instance
	wh, err := whisperstt.New()
	if err != nil {
		panic(err.Error())
	}

	// Open audio file into memory
	audioFile, err := os.Open("audio_files/Recording.m4a")
	if err != nil {
		log.Fatalf("MAIN - Could not open audio file: %s", err)
	}
	defer audioFile.Close()

	// Transcribe audio to string
	transcript, err := wh.TranscribeStreamToText(audioFile)
	if err != nil {
		log.Fatalf("MAIN - Could not transcribe audio to text: %s\n -> ", err)
	} else {
		fmt.Println(transcript)
	}
}
