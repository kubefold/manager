package service

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/kubefold/manager/internal/dto"
	"github.com/sirupsen/logrus"
	"os"
	"path"
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
	filename := path.Join(i.config.InputPath, "fold_input.json")
	logrus.Infof("placing fold input to %s", filename)

	data, err := base64.StdEncoding.DecodeString(encodedInput)
	if err != nil {
		return err
	}

	logrus.Debugf("decoded input: %s", string(data))

	var payload interface{}
	err = json.Unmarshal(data, &payload)
	if err != nil {
		return fmt.Errorf("fold input is invalid as cannot be decoded into json: %w", err)
	}

	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return err
	}
	return nil
}
