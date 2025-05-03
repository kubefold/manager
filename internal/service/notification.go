package service

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sns"
)

type NotificationService interface {
	SendNotification(phones []string, message string) error
}

type notificationService struct {
	snsClient *sns.Client
}

func newNotificationService() NotificationService {
	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		return nil
	}
	snsClient := sns.NewFromConfig(cfg)
	return &notificationService{
		snsClient: snsClient,
	}
}

func (n *notificationService) SendNotification(phones []string, message string) error {
	ctx := context.Background()
	for _, phone := range phones {
		_, err := n.snsClient.Publish(ctx, &sns.PublishInput{
			Message:     aws.String(message),
			PhoneNumber: aws.String(phone),
		})
		if err != nil {
			return err
		}
	}
	return nil
}
