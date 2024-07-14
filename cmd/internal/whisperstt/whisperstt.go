package whisperstt

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httputil"
	"os"

	viper "Flo-AI/cmd/internal/viper"
)

// CONSTRUCTOR
type WhisperSTT struct {
	apiKey string
}

func New() (*WhisperSTT, error) {
	token, err := viper.GetString("OPENAI_TOKEN")
	if err != nil {
		return nil, errors.New("WHISPER - No response from Viper for token")
	}

	return &WhisperSTT{apiKey: token}, nil
}

// END CONSTRUCTOR

func (whisperObj *WhisperSTT) TranscribeStreamToText(audioStream io.Reader) (string, error) {
	// Function takes an io.Reader as a parameter that can either be an audio file or directly from a microphone

	// Create Multipart Writer that will hold and write the stream to the HTTP request
	body := &bytes.Buffer{}
	wr := multipart.NewWriter(body)

	// Create writer form the streamed audio file
	partWriter, err := wr.CreateFormFile("file", "audio_files/Recording.m4a")
	if err != nil {
		wr.Close()
		return "", fmt.Errorf("WHISPER_01 - Failed to create multipart writer from file. Error: %v", err)
	}

	// Copy file contents to multipart writer
	_, err = io.Copy(partWriter, audioStream)
	if err != nil {
		wr.Close()
		return "", fmt.Errorf("WHISPER_02 - Failed to copy stream contents to multipart writer. Error: %v", err)
	}

	// Add model configuration and close writer
	err = wr.WriteField("model", "whisper-1")
	if err != nil {
		wr.Close()
		return "", fmt.Errorf("WHISPER_03 - Failed to write model configuration. Error: %v", err)
	}
	err = wr.WriteField("response_format", "text")
	if err != nil {
		wr.Close()
		return "", fmt.Errorf("WHISPER_03 - Failed to write model configuration. Error: %v", err)
	}

	wr.Close()

	// Create HTTP request with config
	request, err := http.NewRequest("POST", "https://api.openai.com/v1/audio/transcriptions", body)
	if err != nil {
		return "", fmt.Errorf("WHISPER_04 - Failed to create HTTP request. Error: %v", err)
	}

	request.Header.Set("Authorization", "Bearer "+whisperObj.apiKey)
	request.Header.Set("Content-Type", wr.FormDataContentType())

	// Send HTTP request
	response, err := http.DefaultClient.Do(request)
	if err != nil && response.StatusCode != http.StatusOK {
		response.Body.Close()

		// DUMP FOR DEBUG
		dumpRequest, _ := httputil.DumpRequest(request, true)
		writeDumpToFile(dumpRequest)
		dumpResponse, _ := httputil.DumpResponse(response, true)
		writeDumpToFile(dumpResponse)
		return "", fmt.Errorf("WHISPER_05 - Bad HTTP response. Status: %d", response.StatusCode)
	}

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return "", fmt.Errorf("WHISPER_06 - Error reading HTTP response. Status: %d. Error: %v", response.StatusCode, err)
	}
	defer response.Body.Close()

	return string(responseBody), nil
}

func writeDumpToFile(dump []byte) {
	// Open a file for writing
	file, err := os.OpenFile("http_dumps/debug_dump.txt", os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Printf("Error opening file for writing: %v\n", err)
		return
	}
	defer file.Close()

	_, err = file.Write(dump)
	if err != nil {
		fmt.Printf("Error writing dump to file: %v\n", err)
		return
	}
}
