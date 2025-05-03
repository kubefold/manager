package main

import (
	"os"
	"strings"

	"github.com/kubefold/manager/internal/dto"
	"github.com/kubefold/manager/internal/service"
	"github.com/sirupsen/logrus"
)

func main() {
	inputPath := os.Getenv("INPUT_PATH")
	outputPath := os.Getenv("OUTPUT_PATH")
	encodedInput := os.Getenv("ENCODED_INPUT")
	bucket := os.Getenv("BUCKET")
	notificationPhones := strings.Split(os.Getenv("NOTIFICATION_PHONES"), ",")
	notificationMessage := os.Getenv("NOTIFICATION_MESSAGE")

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

	if bucket != "" {
		logrus.Infof("Uploading artifacts to bucket: %s", bucket)
		err := services.Upload().UploadArtifacts(bucket)
		if err != nil {
			logrus.Fatalf("unable to upload artifacts: %v", err)
		}
		return
	}

	if len(notificationPhones) > 0 && notificationPhones[0] != "" && notificationMessage != "" {
		logrus.Infof("Sending notification to %d phone numbers", len(notificationPhones))
		err := services.Notification().SendNotification(notificationPhones, notificationMessage)
		if err != nil {
			logrus.Fatalf("unable to send notification: %v", err)
		}
		return
	}

	logrus.Fatalf("no action specified")
}
