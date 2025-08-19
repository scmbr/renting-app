package repository

import (
	"context"

	"github.com/scmbr/renting-app/internal/domain"
	"github.com/scmbr/renting-app/internal/dto"
	"gorm.io/gorm"
)

type ApartmentPhotoRepo struct {
	db *gorm.DB
}

func NewApartmentPhotoRepo(db *gorm.DB) *ApartmentPhotoRepo {
	return &ApartmentPhotoRepo{db: db}
}

// Получить все фото по ID квартиры
func (r *ApartmentPhotoRepo) GetAllPhotos(ctx context.Context, apartmentId int) ([]*dto.GetApartmentPhoto, error) {
	var photos []*domain.ApartmentPhoto
	err := r.db.WithContext(ctx).Where("apartment_id = ?", apartmentId).Find(&photos).Error
	if err != nil {
		return nil, err
	}

	var result []*dto.GetApartmentPhoto
	for _, p := range photos {
		result = append(result, &dto.GetApartmentPhoto{
			ID:          p.ID,
			ApartmentID: p.ApartmentID,
			URL:         p.URL,
			IsCover:     p.IsCover,
		})
	}
	return result, nil
}

// Получить конкретное фото
func (r *ApartmentPhotoRepo) GetPhotoById(ctx context.Context, apartmentId, photoId int) (*dto.GetApartmentPhoto, error) {
	var photo domain.ApartmentPhoto
	err := r.db.WithContext(ctx).First(&photo, "id = ? AND apartment_id = ?", photoId, apartmentId).Error
	if err != nil {
		return nil, err
	}

	return &dto.GetApartmentPhoto{
		ID:          photo.ID,
		ApartmentID: photo.ApartmentID,
		URL:         photo.URL,
		IsCover:     photo.IsCover,
	}, nil
}

// Добавить пачку фотографий
func (r *ApartmentPhotoRepo) AddPhotoBatch(ctx context.Context, userId, apartmentId int, inputs []dto.CreatePhotoInput) error {
	tx := r.db.WithContext(ctx).Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	for _, input := range inputs {
		photo := domain.ApartmentPhoto{
			ApartmentID: uint(apartmentId),
			URL:         input.URL,
			IsCover:     input.IsCover,
		}
		if err := tx.Create(&photo).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}

// Удалить фотографию
func (r *ApartmentPhotoRepo) DeletePhoto(ctx context.Context, userId, apartmentId, photoId int) error {
	tx := r.db.WithContext(ctx).Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var photo domain.ApartmentPhoto
	err := tx.First(&photo, "id = ? AND apartment_id = ?", photoId, apartmentId).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Delete(&photo).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

// Назначить фото обложкой
func (r *ApartmentPhotoRepo) SetCover(ctx context.Context, userId, apartmentId, photoId int) error {
	tx := r.db.WithContext(ctx).Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Сброс старой обложки
	if err := tx.Model(&domain.ApartmentPhoto{}).
		Where("apartment_id = ? AND is_cover = true", apartmentId).
		Update("is_cover", false).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Установка новой
	if err := tx.Model(&domain.ApartmentPhoto{}).
		Where("id = ? AND apartment_id = ?", photoId, apartmentId).
		Update("is_cover", true).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
func (r *ApartmentPhotoRepo) HasCoverPhoto(apartmentId int) (bool, error) {
	var count int64
	err := r.db.Model(&domain.ApartmentPhoto{}).
		Where("apartment_id = ? AND is_cover = true", apartmentId).
		Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *ApartmentPhotoRepo) ReplaceAllPhotos(ctx context.Context, userId, apartmentId int, inputs []dto.CreatePhotoInput) error {
	tx := r.db.WithContext(ctx).Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Where("apartment_id = ?", apartmentId).Delete(&domain.ApartmentPhoto{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	for _, input := range inputs {
		photo := domain.ApartmentPhoto{
			ApartmentID: uint(apartmentId),
			URL:         input.URL,
			IsCover:     input.IsCover,
		}
		if err := tx.Create(&photo).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}
