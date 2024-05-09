module OpieFlo-AI

go 1.22.3
require (
    github.com/Azure/cognitive-services-go-sdk/speech v1.22.0-preview.1 // Azure Speech Services (STT & TTS)
    github.com/sashabaranov/go-openai // OpenAI Go Wrapper
    github.com/Microsoft/cognitive-services-speech-sdk-go/audio
    github.com/Microsoft/cognitive-services-speech-sdk-go/common
    github.com/Microsoft/cognitive-services-speech-sdk-go/speech // Dialogflow Client Library
)
