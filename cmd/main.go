package main

import (
	"fmt"
	"log"
	"os"

	"Flo-AI/cmd/internal/openai"
	"Flo-AI/cmd/internal/viper"
	"Flo-AI/cmd/internal/whisperstt"

	tea "github.com/charmbracelet/bubbletea"
)

// Model for app state
type menuModel struct {
	header  string   // Header text
	choices []string // Menu options
	cursor  int      // Cursor index-based position
	footer  string   // Footer text
}

func mainMenuModel() menuModel {
	return menuModel{
		header: "Welcome to Flo, the AI powered voice assistant!\n",
		choices: []string{
			"1 - Start Flo (with avatar)",
			"2 - Start Flo (command line only)",
			"3 - Preferences",
			"4 - Exit"},
		footer: "Ctrl+C or Q to quit...",
	}
}

// INIT
func (m menuModel) Init() tea.Cmd {
	return nil
}

// UPDATE
func (m menuModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// Key press I/O
	case tea.KeyMsg:
		// Check which key
		switch msg.String() {
		// Exit keys
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		}
	}
	// Return updated model to Bubble Tea runtime. Not ret`urning a command.
	return m, nil
}

// VIEW
func (m menuModel) View() string {
	// UI is only a string. Redrawing is managed by Bubble Tea

	// Header
	s := m.header

	// Iterate through choices
	for i, choice := range m.choices {
		// Draw cursor on selected option
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		// Render this row by adding to the final string
		s += fmt.Sprintf("%s %s\n", cursor, choice)
	}

	// Footer
	s += m.footer

	// Send UI to Bubble Tea for rendering
	return s
}

func main() {
	p := tea.NewProgram(mainMenuModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error trying to create main menu: %v", err)
		os.Exit(1)
	}

	// LOAD Viper
	fmt.Print("Viper... ")
	if err := viper.LoadConfig(); err != nil {
		panic(err.Error())
	}
	fmt.Println("OK")

	// LOAD OpenAI
	fmt.Print("Reaching Flo... ")
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
	fmt.Println("OK")

	// LOAD Whisper
	fmt.Print("Whisper... ")
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
	fmt.Print("OK")

	client.MakeRequest(transcript)
}
