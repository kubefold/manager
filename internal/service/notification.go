package service

type NotificationService interface{}

type notificationService struct{}

func newNotificationService() NotificationService {
	return &notificationService{}
}
