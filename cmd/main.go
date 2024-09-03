package main

import (
	"fmt"
	"log"
	"os"

	"Flo-AI/cmd/internal/openai"
	"Flo-AI/cmd/internal/viper"
	"Flo-AI/cmd/internal/whisperstt"
)

func main() {
	// LOAD Viper
	if err := viper.LoadConfig(); err != nil {
		panic(err.Error())
	}

	// LOAD OpenAI
	apiKey, err := viper.GetString("OPENAI_TOKEN")
	if err != nil {
		log.Fatalf("MAIN - viper could not get OPENAI_TOKEN: %s", err)
	}
	config := openai.Config{
		APIKey: apiKey,
	}

	client := openai.NewClient(config)
	if err != nil {
		log.Fatalf("MAIN - could not create OpenAI Client: %s", err)
	}

	// LOAD Whisper
	// Initialize Whisper instance
	wh, err := whisperstt.New()
	if err != nil {
		panic(err.Error())
	}

	// TRANSCRIBE AUDIO FILE AND SEND TO GPT
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

	fmt.Println(client.MakeRequest(transcript))
}
