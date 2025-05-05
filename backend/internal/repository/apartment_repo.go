package repository

import (
	"context"
	"errors"

	"github.com/scmbr/renting-app/internal/dto"
	"github.com/scmbr/renting-app/internal/models"
	"gorm.io/gorm"
)

type ApartmentRepo struct {
	db *gorm.DB
}

func NewApartmentRepo(db *gorm.DB) *ApartmentRepo {
	return &ApartmentRepo{db: db}
}
func (r *ApartmentRepo) GetAllApartments(ctx context.Context, userId int) ([]*dto.GetApartmentResponse, error) {
	var apartments []models.Apartment
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
			District:         apartment.District,
			House:            apartment.House,
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
		getApartmentDTOs = append(getApartmentDTOs, &getApartmentDTO)
	}

	tx.Commit()

	return getApartmentDTOs, nil

}
func (r *ApartmentRepo) GetApartmentById(ctx context.Context, userId int, id int) (*dto.GetApartmentResponse, error) {
	var apartment models.Apartment

	err := r.db.WithContext(ctx).
		Where("id = ? AND user_id = ?", id, userId).
		First(&apartment).Error
	if err != nil {
		return nil, err
	}

	getApartmentDTO := dto.GetApartmentResponse{
		ID:               apartment.ID,
		UserID:           apartment.UserID,
		City:             apartment.City,
		Street:           apartment.Street,
		District:         apartment.District,
		House:            apartment.House,
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

	return &getApartmentDTO, nil
}
func (r *ApartmentRepo) CreateApartment(ctx context.Context, userId int, input dto.CreateApartmentInput) error {
	tx := r.db.WithContext(ctx).Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	apartmentGorm := models.Apartment{
		UserID:           uint(userId),
		City:             input.City,
		Street:           input.Street,
		District:         input.District,
		House:            input.House,
		Building:         input.Building,
		Floor:            input.Floor,
		ApartmentNumber:  input.ApartmentNumber,
		Longitude:        input.Longitude,
		Latitude:         input.Latitude,
		Rooms:            input.Rooms,
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
	result := tx.Create(&apartmentGorm)
	if result.Error != nil {

		tx.Rollback()
		return result.Error
	}
	tx.Commit()

	return nil
}
func (r *ApartmentRepo) DeleteApartment(ctx context.Context, userId int, id int) error {
	tx := r.db.WithContext(ctx).Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	var apartment models.Apartment
	result := tx.First(&apartment, "id = ? AND user_id = ?", id, userId)
	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}
	if result.RowsAffected == 0 {
		tx.Rollback()
		return errors.New("user not found")
	}
	if err := tx.Delete(&apartment).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
func (r *ApartmentRepo) UpdateApartment(ctx context.Context, userId int, id int, input *dto.UpdateApartmentInput) error {
	tx := r.db.WithContext(ctx).Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var apartment models.Apartment
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
	if input.District != nil {
		apartment.District = *input.District
	}
	if input.House != nil {
		apartment.House = *input.House
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
