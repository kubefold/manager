package main

import (
	"github.com/kubefold/manager/internal/dto"
	"github.com/kubefold/manager/internal/service"
	"github.com/sirupsen/logrus"
	"os"
)

func main() {
	inputPath := os.Getenv("INPUT_PATH")
	outputPath := os.Getenv("OUTPUT_PATH")
	encodedInput := os.Getenv("ENCODED_INPUT")

	if inputPath == "" {
		logrus.Panicf("INPUT_PATH env var is required")
	}
	if outputPath == "" {
		logrus.Panicf("OUTPUT_PATH env var is required")
	}

	config := dto.Config{
		InputPath:  inputPath,
		OutputPath: outputPath,
	}

	services := service.NewServices(config)
	if encodedInput != "" {
		err := services.Input().PlaceInput(encodedInput)
		if err != nil {
			logrus.Fatalf("unable to place input: %v", err)
		}
		return
	}

	logrus.Fatalf("no action specified")
}
