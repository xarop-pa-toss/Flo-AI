package vosk

import (
	"fmt"

	viper "Flo-AI/cmd/internal/viper"

	vosk "github.com/alphacep/vosk-api/go"
)

func Init() {
	fmt.Println("Initiated vosk.")

	modelPath, err := viper.GetString("VOSK_MODEL_PATH")
	if err != nil {
		model, err := vosk.NewModel(modelPath)
		fmt.Println("VOSK - Model loaded")
	}

}
