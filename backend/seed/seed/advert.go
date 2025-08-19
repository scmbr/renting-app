package seed

import (
	"log"
	"math/rand"
	"time"

	"github.com/scmbr/renting-app/internal/domain"
	"gorm.io/gorm"
)

const (
	advertSeedCount = 100
)

var (
	rentalTypes = []string{"short-term", "long-term", "daily"}
	adTitles    = []string{
		"Уютная квартира в центре",
		"Светлая квартира с видом",
		"Квартира с ремонтом и мебелью",
		"Недорогая квартира рядом с метро",
		"Комфортное жилье для семьи",
	}
)

func SeedAdverts(db *gorm.DB, cfg interface{}) {
	rand := rand.New(rand.NewSource(time.Now().UnixNano()))

	var users []domain.User
	if err := db.Find(&users).Error; err != nil {
		log.Fatalf("failed to load users: %s", err)
	}

	var apartments []domain.Apartment
	if err := db.Find(&apartments).Error; err != nil {
		log.Fatalf("failed to load apartments: %s", err)
	}

	if len(users) == 0 || len(apartments) == 0 {
		log.Fatal("no users or apartments found to create adverts")
	}

	var adverts []domain.Advert

	for i := 0; i < advertSeedCount; i++ {
		user := users[rand.Intn(len(users))]
		apartment := apartments[rand.Intn(len(apartments))]

		advert := domain.Advert{
			UserID:         user.ID,
			ApartmentID:    apartment.ID,
			Status:         "active",
			Title:          adTitles[rand.Intn(len(adTitles))],
			Pets:           rand.Intn(2) == 1,
			Babies:         rand.Intn(2) == 1,
			Smoking:        rand.Intn(2) == 1,
			Internet:       rand.Intn(2) == 1,
			WashingMachine: rand.Intn(2) == 1,
			TV:             rand.Intn(2) == 1,
			Conditioner:    rand.Intn(2) == 1,
			Dishwasher:     rand.Intn(2) == 1,
			Concierge:      apartment.Concierge,
			Rent:           float64(rand.Intn(50000) + 20000),
			Deposit:        float64(rand.Intn(50000) + 20000),
			RentalType:     rentalTypes[rand.Intn(len(rentalTypes))],
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
		}

		adverts = append(adverts, advert)
	}

	if err := db.Create(&adverts).Error; err != nil {
		log.Fatalf("failed to insert adverts: %s", err)
	}

	log.Printf("Seeded %d adverts", advertSeedCount)
}
