package service

type UploadService interface{}

type uploadService struct{}

func newUploadService() UploadService {
	return &uploadService{}
}
