module OpieFlo-AI

go 1.22.3
require (
    github.com/sashabaranov/go-openai v1.23.1 // OpenAI Go Wrapper
    github.com/Microsoft/cognitive-services-speech-sdk-go/audio v1.32.1 // Azure Speech Services (STT & TTS)
    github.com/Microsoft/cognitive-services-speech-sdk-go/common v1.32.1 // Azure Speech Services (STT & TTS)
    github.com/Microsoft/cognitive-services-speech-sdk-go/speech v1.32.1 // Azure Speech Services (STT & TTS)
    cloud.google.com/go/dialogflow/apiv2 v1.53  // Dialogflow Client Library
)
