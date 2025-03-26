package service

import (
	"context"
	"mime/multipart"

	"github.com/vasya/renting-app/internal/dto"
	"github.com/vasya/renting-app/internal/repository"
	"github.com/vasya/renting-app/pkg/storage"
)

type UsersService struct {
	repo    repository.Users
	storage storage.Provider
}

func NewUsersService(repo repository.Users, storage storage.Provider) *UsersService {
	return &UsersService{
		repo:    repo,
		storage: storage,
	}
}

func (s *UsersService) GetAllUsers() ([]dto.GetUser, error) {

	return s.repo.GetAllUsers()
}
func (s *UsersService) GetUserById(id int) (*dto.GetUser, error) {

	return s.repo.GetUserById(id)
}
func (s *UsersService) DeleteUserById(id int) (*dto.GetUser, error) {

	return s.repo.DeleteUserById(id)
}
func (s *UsersService) UpdateUserById(input *dto.GetUser) (*dto.GetUser, error) {
	return s.repo.UpdateUserById(input)
}
func (s *UsersService) UploadAvatarToS3(ctx context.Context, fileHeader *multipart.FileHeader) (string, error) {
	file, err := fileHeader.Open()
	if err != nil {
		return "", err
	}
	defer file.Close()
	return s.storage.Upload(ctx, storage.UploadInput{
		File:        file,
		Name:        fileHeader.Filename,
		Size:        fileHeader.Size,
		ContentType: "image/png",
	})

}
func (s *UsersService) UpdateAvatar(userId int, avatarURL string) error {

	return s.repo.UpdateAvatar(userId, avatarURL)
}
