package service

import "github.com/kubefold/manager/internal/dto"

type Services interface {
	Input() InputService
	Upload() UploadService
	Notification() NotificationService
}

type services struct {
	inputService        InputService
	uploadService       UploadService
	notificationService NotificationService
}

func NewServices(config dto.Config) Services {
	return &services{
		inputService:        newInputService(config),
		uploadService:       newUploadService(config),
		notificationService: newNotificationService(),
	}
}

func (s services) Input() InputService {
	return s.inputService
}

func (s services) Upload() UploadService {
	return s.uploadService
}

func (s services) Notification() NotificationService {
	return s.notificationService
}
