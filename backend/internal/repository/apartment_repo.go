package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/scmbr/renting-app/internal/domain"
	"github.com/scmbr/renting-app/internal/dto"
	"gorm.io/gorm"
)

type ApartmentRepo struct {
	db *gorm.DB
}

func NewApartmentRepo(db *gorm.DB) *ApartmentRepo {
	return &ApartmentRepo{db: db}
}
func (r *ApartmentRepo) GetAllApartments(ctx context.Context, userId int) ([]*dto.GetApartmentResponse, error) {
	var apartments []domain.Apartment
	tx := r.db.WithContext(ctx).Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	result := tx.
		Where("user_id = ?", userId).
		Find(&apartments)
	if result.Error != nil {
		tx.Rollback()
		return nil, result.Error
	}

	var getApartmentDTOs []*dto.GetApartmentResponse

	for _, apartment := range apartments {
		getApartmentDTO := dto.GetApartmentResponse{
			ID:               apartment.ID,
			UserID:           apartment.UserID,
			City:             apartment.City,
			Street:           apartment.Street,
			Building:         apartment.Building,
			Floor:            apartment.Floor,
			ApartmentNumber:  apartment.ApartmentNumber,
			Longitude:        apartment.Longitude,
			Latitude:         apartment.Latitude,
			Rooms:            apartment.Rooms,
			Area:             apartment.Area,
			Elevator:         apartment.Elevator,
			GarbageChute:     apartment.GarbageChute,
			BathroomType:     apartment.BathroomType,
			Concierge:        apartment.Concierge,
			ConstructionYear: apartment.ConstructionYear,
			ConstructionType: apartment.ConstructionType,
			Remont:           apartment.Remont,
			CreatedAt:        apartment.CreatedAt,
			UpdatedAt:        apartment.UpdatedAt,
			Rating:           apartment.Rating,
			Status:           apartment.Status,
		}
		getApartmentDTOs = append(getApartmentDTOs, &getApartmentDTO)
	}

	tx.Commit()

	return getApartmentDTOs, nil

}
func (r *ApartmentRepo) GetApartmentById(ctx context.Context, userId int, id int) (*dto.GetApartmentResponse, error) {
	var apartment domain.Apartment

	err := r.db.WithContext(ctx).
		Where("id = ? AND user_id = ?", id, userId).
		First(&apartment).Error
	if err != nil {
		return nil, err
	}

	getApartmentDTO := dto.GetApartmentResponse{
		ID:     apartment.ID,
		UserID: apartment.UserID,
		City:   apartment.City,
		Street: apartment.Street,

		Building:         apartment.Building,
		Floor:            apartment.Floor,
		ApartmentNumber:  apartment.ApartmentNumber,
		Longitude:        apartment.Longitude,
		Latitude:         apartment.Latitude,
		Rooms:            apartment.Rooms,
		Area:             apartment.Area,
		Elevator:         apartment.Elevator,
		GarbageChute:     apartment.GarbageChute,
		BathroomType:     apartment.BathroomType,
		Concierge:        apartment.Concierge,
		ConstructionYear: apartment.ConstructionYear,
		ConstructionType: apartment.ConstructionType,
		Remont:           apartment.Remont,
		CreatedAt:        apartment.CreatedAt,
		UpdatedAt:        apartment.UpdatedAt,
		Rating:           apartment.Rating,
		Status:           apartment.Status,
	}

	return &getApartmentDTO, nil
}
func (r *ApartmentRepo) CreateApartment(ctx context.Context, userId int, input dto.CreateApartmentInput) (*domain.Apartment, error) {
	apartmentGorm := domain.Apartment{
		UserID:           uint(userId),
		City:             input.City,
		Street:           input.Street,
		Building:         input.Building,
		Floor:            input.Floor,
		Longitude:        input.Longitude,
		Latitude:         input.Latitude,
		Rooms:            input.Rooms,
		Area:             input.Area,
		Elevator:         input.Elevator,
		GarbageChute:     input.GarbageChute,
		BathroomType:     input.BathroomType,
		Concierge:        input.Concierge,
		ConstructionYear: input.ConstructionYear,
		ConstructionType: input.ConstructionType,
		Remont:           input.Remont,
		Rating:           0,
		Status:           "active",
	}

	if err := r.db.WithContext(ctx).Create(&apartmentGorm).Error; err != nil {
		return nil, err
	}

	return &apartmentGorm, nil
}

func (r *ApartmentRepo) DeleteApartment(ctx context.Context, userId int, id int) error {
	tx := r.db.WithContext(ctx).Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var apartment domain.Apartment
	result := tx.First(&apartment, "id = ? AND user_id = ?", id, userId)
	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}
	if result.RowsAffected == 0 {
		tx.Rollback()
		return errors.New("apartment not found or not owned by user")
	}

	if err := tx.Where("apartment_id = ?", apartment.ID).Delete(&domain.ApartmentPhoto{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	var advertIDs []int
	if err := tx.Model(&domain.Advert{}).
		Where("apartment_id = ?", apartment.ID).
		Pluck("id", &advertIDs).Error; err != nil {
		tx.Rollback()
		return err
	}

	if len(advertIDs) > 0 {
		if err := tx.Where("advert_id IN ?", advertIDs).Delete(&domain.Favorites{}).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	if err := tx.Where("apartment_id = ?", apartment.ID).Delete(&domain.Advert{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Delete(&apartment).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (r *ApartmentRepo) UpdateApartment(ctx context.Context, userId int, id int, input *dto.UpdateApartmentInput) error {
	tx := r.db.WithContext(ctx).Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var apartment domain.Apartment
	// Проверка существования и принадлежности
	err := tx.First(&apartment, "id = ? AND user_id = ?", id, userId).Error
	if err != nil {
		tx.Rollback()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("apartment not found or does not belong to user")
		}
		return err
	}

	if input.City != nil {
		apartment.City = *input.City
	}
	if input.Street != nil {
		apartment.Street = *input.Street
	}

	if input.Building != nil {
		apartment.Building = *input.Building
	}
	if input.Floor != nil {
		apartment.Floor = *input.Floor
	}
	if input.ApartmentNumber != nil {
		apartment.ApartmentNumber = *input.ApartmentNumber
	}
	if input.Longitude != nil {
		apartment.Longitude = *input.Longitude
	}
	if input.Latitude != nil {
		apartment.Latitude = *input.Latitude
	}
	if input.Rooms != nil {
		apartment.Rooms = *input.Rooms
	}
	if input.Elevator != nil {
		apartment.Elevator = *input.Elevator
	}
	if input.GarbageChute != nil {
		apartment.GarbageChute = *input.GarbageChute
	}
	if input.BathroomType != nil {
		apartment.BathroomType = *input.BathroomType
	}
	if input.Concierge != nil {
		apartment.Concierge = *input.Concierge
	}
	if input.ConstructionYear != nil {
		apartment.ConstructionYear = *input.ConstructionYear
	}
	if input.ConstructionType != nil {
		apartment.ConstructionType = *input.ConstructionType
	}
	if input.Remont != nil {
		apartment.Remont = *input.Remont
	}
	if input.Status != nil {
		apartment.Status = *input.Status
	}
	if input.Area != nil {
		apartment.Area = *input.Area
	}

	if err := tx.Save(&apartment).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}
func (r *ApartmentRepo) GetAllApartmentsAdmin(ctx context.Context) ([]*dto.GetApartmentResponse, error) {
	var apartments []domain.Apartment

	tx := r.db.WithContext(ctx).Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Find(&apartments).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	var getApartmentDTOs []*dto.GetApartmentResponse
	for _, apartment := range apartments {
		getApartmentDTO := dto.GetApartmentResponse{
			ID:     apartment.ID,
			UserID: apartment.UserID,
			City:   apartment.City,
			Street: apartment.Street,

			Building:         apartment.Building,
			Floor:            apartment.Floor,
			ApartmentNumber:  apartment.ApartmentNumber,
			Longitude:        apartment.Longitude,
			Latitude:         apartment.Latitude,
			Rooms:            apartment.Rooms,
			Elevator:         apartment.Elevator,
			GarbageChute:     apartment.GarbageChute,
			BathroomType:     apartment.BathroomType,
			Concierge:        apartment.Concierge,
			ConstructionYear: apartment.ConstructionYear,
			ConstructionType: apartment.ConstructionType,
			Remont:           apartment.Remont,
			CreatedAt:        apartment.CreatedAt,
			UpdatedAt:        apartment.UpdatedAt,
			Rating:           apartment.Rating,
			Status:           apartment.Status,
		}
		fmt.Println(apartment.ID)
		getApartmentDTOs = append(getApartmentDTOs, &getApartmentDTO)
	}

	tx.Commit()
	return getApartmentDTOs, nil
}

func (r *ApartmentRepo) GetApartmentByIdAdmin(ctx context.Context, id int) (*dto.GetApartmentResponse, error) {
	var apartment domain.Apartment

	tx := r.db.WithContext(ctx).Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	err := tx.First(&apartment, id).Error
	if err != nil {
		tx.Rollback()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("apartment not found")
		}
		return nil, err
	}

	getApartmentDTO := dto.GetApartmentResponse{
		ID:     apartment.ID,
		UserID: apartment.UserID,
		City:   apartment.City,
		Street: apartment.Street,

		Building:         apartment.Building,
		Floor:            apartment.Floor,
		ApartmentNumber:  apartment.ApartmentNumber,
		Longitude:        apartment.Longitude,
		Latitude:         apartment.Latitude,
		Rooms:            apartment.Rooms,
		Elevator:         apartment.Elevator,
		GarbageChute:     apartment.GarbageChute,
		BathroomType:     apartment.BathroomType,
		Concierge:        apartment.Concierge,
		ConstructionYear: apartment.ConstructionYear,
		ConstructionType: apartment.ConstructionType,
		Remont:           apartment.Remont,
		CreatedAt:        apartment.CreatedAt,
		UpdatedAt:        apartment.UpdatedAt,
		Rating:           apartment.Rating,
		Status:           apartment.Status,
	}

	tx.Commit()
	return &getApartmentDTO, nil
}

func (r *ApartmentRepo) UpdateApartmentAdmin(ctx context.Context, id int, input *dto.UpdateApartmentInput) error {
	tx := r.db.WithContext(ctx).Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var apartment domain.Apartment
	err := tx.First(&apartment, id).Error
	if err != nil {
		tx.Rollback()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("apartment not found")
		}
		return err
	}

	if input.City != nil {
		apartment.City = *input.City
	}
	if input.Street != nil {
		apartment.Street = *input.Street
	}

	if input.Building != nil {
		apartment.Building = *input.Building
	}
	if input.Floor != nil {
		apartment.Floor = *input.Floor
	}
	if input.ApartmentNumber != nil {
		apartment.ApartmentNumber = *input.ApartmentNumber
	}
	if input.Longitude != nil {
		apartment.Longitude = *input.Longitude
	}
	if input.Latitude != nil {
		apartment.Latitude = *input.Latitude
	}
	if input.Rooms != nil {
		apartment.Rooms = *input.Rooms
	}
	if input.Elevator != nil {
		apartment.Elevator = *input.Elevator
	}
	if input.GarbageChute != nil {
		apartment.GarbageChute = *input.GarbageChute
	}
	if input.BathroomType != nil {
		apartment.BathroomType = *input.BathroomType
	}
	if input.Concierge != nil {
		apartment.Concierge = *input.Concierge
	}
	if input.ConstructionYear != nil {
		apartment.ConstructionYear = *input.ConstructionYear
	}
	if input.ConstructionType != nil {
		apartment.ConstructionType = *input.ConstructionType
	}
	if input.Remont != nil {
		apartment.Remont = *input.Remont
	}
	if input.Status != nil {
		apartment.Status = *input.Status
	}

	if err := tx.Save(&apartment).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func (r *ApartmentRepo) DeleteApartmentAdmin(ctx context.Context, id int) error {
	tx := r.db.WithContext(ctx).Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var apartment domain.Apartment
	result := tx.First(&apartment, id)
	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}
	if result.RowsAffected == 0 {
		tx.Rollback()
		return errors.New("apartment not found")
	}

	if err := tx.Delete(&apartment).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}
