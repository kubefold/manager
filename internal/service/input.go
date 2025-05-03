package service

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"path"

	"github.com/kubefold/manager/internal/dto"
	"github.com/sirupsen/logrus"
)

type InputService interface {
	PlaceInput(encodedInput string) error
}

type inputService struct {
	config dto.Config
}

func newInputService(config dto.Config) InputService {
	return &inputService{
		config: config,
	}
}

func (i inputService) PlaceInput(encodedInput string) error {
	err := os.Mkdir(i.config.InputPath, 0750)
	if err != nil && !os.IsExist(err) {
		return err
	}
	err = os.Mkdir(i.config.OutputPath, 0750)
	if err != nil && !os.IsExist(err) {
		return err
	}

	filename := path.Join(i.config.InputPath, "fold_input.json")

	data, err := base64.StdEncoding.DecodeString(encodedInput)
	if err != nil {
		return err
	}

	logrus.Infof("decoded input: %s", string(data))

	var payload interface{}
	err = json.Unmarshal(data, &payload)
	if err != nil {
		return fmt.Errorf("fold input is invalid as cannot be decoded into json: %w", err)
	}

	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return err
	}
	logrus.Infof("placed fold input to %s", filename)
	return nil
}
