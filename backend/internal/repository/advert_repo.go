package repository

import (
	"context"
	"errors"

	"github.com/scmbr/renting-app/internal/dto"
	"github.com/scmbr/renting-app/internal/models"
	"gorm.io/gorm"
)

type AdvertRepo struct {
	db *gorm.DB
}

func NewAdvertRepo(db *gorm.DB) *AdvertRepo {
	return &AdvertRepo{db: db}
}
func (r *AdvertRepo) GetAllAdverts(ctx context.Context, userId int) ([]*dto.GetAdvertResponse, error) {
	var adverts []models.Advert
	tx := r.db.WithContext(ctx).Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	result := tx.
		Where("user_id = ?", userId).
		Find(&adverts)
	if result.Error != nil {
		tx.Rollback()
		return nil, result.Error
	}
	var getAdvertDTOs []*dto.GetAdvertResponse

	for _, advert := range adverts {
		getAdvertDTO := dto.GetAdvertResponse{
			ID:             advert.ID,
			UserID:         advert.UserID,
			ApartmentID:    advert.ApartmentID,
			CreatedAt:      advert.CreatedAt,
			UpdatedAt:      advert.UpdatedAt,
			Status:         advert.Status,
			Title:          advert.Title,
			Pets:           advert.Pets,
			Babies:         advert.Babies,
			Smoking:        advert.Smoking,
			Internet:       advert.Internet,
			WashingMachine: advert.WashingMachine,
			TV:             advert.TV,
			Conditioner:    advert.Conditioner,
			Concierge:      advert.Concierge,
			Rent:           advert.Rent,
			Deposit:        advert.Deposit,
			RentalType:     advert.RentalType,
		}
		getAdvertDTOs = append(getAdvertDTOs, &getAdvertDTO)
	}
	tx.Commit()

	return getAdvertDTOs, nil
}
func (r *AdvertRepo) GetAdvertById(ctx context.Context, userId int, id int) (*dto.GetAdvertResponse, error) {
	var advert models.Advert

	err := r.db.WithContext(ctx).
		Where("id = ? AND user_id = ?", id, userId).
		First(&advert).Error
	if err != nil {
		return nil, err
	}
	getAdvertDTO := dto.GetAdvertResponse{
		ID:             advert.ID,
		UserID:         advert.UserID,
		ApartmentID:    advert.ApartmentID,
		CreatedAt:      advert.CreatedAt,
		UpdatedAt:      advert.UpdatedAt,
		Status:         advert.Status,
		Title:          advert.Title,
		Pets:           advert.Pets,
		Babies:         advert.Babies,
		Smoking:        advert.Smoking,
		Internet:       advert.Internet,
		WashingMachine: advert.WashingMachine,
		TV:             advert.TV,
		Conditioner:    advert.Conditioner,
		Concierge:      advert.Concierge,
		Rent:           advert.Rent,
		Deposit:        advert.Deposit,
		RentalType:     advert.RentalType,
	}
	return &getAdvertDTO, nil
}
func (r *AdvertRepo) CreateAdvert(ctx context.Context, userId int, input dto.CreateAdvertInput) error {
	tx := r.db.WithContext(ctx).Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	advertGorm := models.Advert{
		UserID:         uint(userId),
		ApartmentID:    input.ApartmentID,
		Status:         "active",
		Title:          input.Title,
		Pets:           input.Pets,
		Babies:         input.Babies,
		Smoking:        input.Smoking,
		Internet:       input.Internet,
		WashingMachine: input.WashingMachine,
		TV:             input.TV,
		Conditioner:    input.Conditioner,
		Concierge:      input.Concierge,
		Rent:           input.Rent,
		Deposit:        input.Deposit,
		RentalType:     input.RentalType,
	}
	result := tx.Create(&advertGorm)
	if result.Error != nil {

		tx.Rollback()
		return result.Error
	}
	tx.Commit()

	return nil
}
func (r *AdvertRepo) DeleteAdvert(ctx context.Context, userId int, id int) error {
	tx := r.db.WithContext(ctx).Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	var advert models.Advert
	result := tx.First(&advert, "id = ? AND user_id = ?", id, userId)
	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}
	if result.RowsAffected == 0 {
		tx.Rollback()
		return errors.New("user not found")
	}
	if err := tx.Delete(&advert).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
func (r *AdvertRepo) UpdateAdvert(ctx context.Context, userId int, id int, input *dto.UpdateAdvertInput) error {
	tx := r.db.WithContext(ctx).Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var advert models.Advert
	// Проверка существования и принадлежности
	err := tx.First(&advert, "id = ? AND user_id = ?", id, userId).Error
	if err != nil {
		tx.Rollback()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("apartment not found or does not belong to user")
		}
		return err
	}
	if input.Title != nil {
		advert.Title = *input.Title
	}
	if input.Pets != nil {
		advert.Pets = *input.Pets
	}
	if input.Babies != nil {
		advert.Babies = *input.Babies
	}
	if input.Smoking != nil {
		advert.Smoking = *input.Smoking
	}
	if input.Internet != nil {
		advert.Internet = *input.Internet
	}
	if input.WashingMachine != nil {
		advert.WashingMachine = *input.WashingMachine
	}
	if input.TV != nil {
		advert.TV = *input.TV
	}
	if input.Conditioner != nil {
		advert.Conditioner = *input.Conditioner
	}
	if input.Dishwasher != nil {
		advert.Dishwasher = *input.Dishwasher
	}
	if input.Concierge != nil {
		advert.Concierge = *input.Concierge
	}
	if input.Rent != nil {
		advert.Rent = *input.Rent
	}
	if input.Deposit != nil {
		advert.Deposit = *input.Deposit
	}
	if input.RentalType != nil {
		advert.RentalType = *input.RentalType
	}
	if input.Concierge != nil {
		advert.Concierge = *input.Concierge
	}
	if input.Rent != nil {
		advert.Rent = *input.Rent
	}
	if input.Status != nil {
		advert.Status = *input.Status
	}
	if err := tx.Save(&advert).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}
func (r *AdvertRepo) GetAllAdvertsAdmin(ctx context.Context) ([]*dto.GetAdvertResponse, error) {
	var adverts []models.Advert
	tx := r.db.WithContext(ctx).Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	result := tx.Find(&adverts)
	if result.Error != nil {
		tx.Rollback()
		return nil, result.Error
	}

	var getAdvertDTOs []*dto.GetAdvertResponse
	for _, advert := range adverts {
		getAdvertDTO := dto.GetAdvertResponse{
			ID:             advert.ID,
			UserID:         advert.UserID,
			ApartmentID:    advert.ApartmentID,
			CreatedAt:      advert.CreatedAt,
			UpdatedAt:      advert.UpdatedAt,
			Status:         advert.Status,
			Title:          advert.Title,
			Pets:           advert.Pets,
			Babies:         advert.Babies,
			Smoking:        advert.Smoking,
			Internet:       advert.Internet,
			WashingMachine: advert.WashingMachine,
			TV:             advert.TV,
			Conditioner:    advert.Conditioner,
			Concierge:      advert.Concierge,
			Rent:           advert.Rent,
			Deposit:        advert.Deposit,
			RentalType:     advert.RentalType,
		}
		getAdvertDTOs = append(getAdvertDTOs, &getAdvertDTO)
	}
	tx.Commit()
	return getAdvertDTOs, nil
}

func (r *AdvertRepo) GetAdvertByIdAdmin(ctx context.Context, id int) (*dto.GetAdvertResponse, error) {
	var advert models.Advert
	err := r.db.WithContext(ctx).First(&advert, "id = ?", id).Error
	if err != nil {
		return nil, err
	}

	getAdvertDTO := dto.GetAdvertResponse{
		ID:             advert.ID,
		UserID:         advert.UserID,
		ApartmentID:    advert.ApartmentID,
		CreatedAt:      advert.CreatedAt,
		UpdatedAt:      advert.UpdatedAt,
		Status:         advert.Status,
		Title:          advert.Title,
		Pets:           advert.Pets,
		Babies:         advert.Babies,
		Smoking:        advert.Smoking,
		Internet:       advert.Internet,
		WashingMachine: advert.WashingMachine,
		TV:             advert.TV,
		Conditioner:    advert.Conditioner,
		Concierge:      advert.Concierge,
		Rent:           advert.Rent,
		Deposit:        advert.Deposit,
		RentalType:     advert.RentalType,
	}
	return &getAdvertDTO, nil
}

func (r *AdvertRepo) UpdateAdvertAdmin(ctx context.Context, id int, input *dto.UpdateAdvertInput) error {
	tx := r.db.WithContext(ctx).Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var advert models.Advert
	err := tx.First(&advert, "id = ?", id).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	// Обновляем только те поля, что не nil
	if input.Title != nil {
		advert.Title = *input.Title
	}
	if input.Pets != nil {
		advert.Pets = *input.Pets
	}
	if input.Babies != nil {
		advert.Babies = *input.Babies
	}
	if input.Smoking != nil {
		advert.Smoking = *input.Smoking
	}
	if input.Internet != nil {
		advert.Internet = *input.Internet
	}
	if input.WashingMachine != nil {
		advert.WashingMachine = *input.WashingMachine
	}
	if input.TV != nil {
		advert.TV = *input.TV
	}
	if input.Conditioner != nil {
		advert.Conditioner = *input.Conditioner
	}
	if input.Dishwasher != nil {
		advert.Dishwasher = *input.Dishwasher
	}
	if input.Concierge != nil {
		advert.Concierge = *input.Concierge
	}
	if input.Rent != nil {
		advert.Rent = *input.Rent
	}
	if input.Deposit != nil {
		advert.Deposit = *input.Deposit
	}
	if input.RentalType != nil {
		advert.RentalType = *input.RentalType
	}
	if input.Status != nil {
		advert.Status = *input.Status
	}

	if err := tx.Save(&advert).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (r *AdvertRepo) DeleteAdvertAdmin(ctx context.Context, id int) error {
	tx := r.db.WithContext(ctx).Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var advert models.Advert
	result := tx.First(&advert, "id = ?", id)
	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}
	if result.RowsAffected == 0 {
		tx.Rollback()
		return errors.New("advert not found")
	}
	if err := tx.Delete(&advert).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
