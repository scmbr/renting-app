package seed

import (
	"encoding/json"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/scmbr/renting-app/internal/config"
	"github.com/scmbr/renting-app/internal/domain"
	"gorm.io/gorm"
)

type City struct {
	Name      string  `json:"name"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

const seedApartmentCount = 100

func SeedApartments(db *gorm.DB, cfg *config.Config, cityName string) {
	rand := rand.New(rand.NewSource(time.Now().UnixNano()))

	cities := loadCities("data/cities.json")

	var city *City
	if cityName != "" {
		city = findCityByName(cities, cityName)
		if city == nil {
			log.Fatalf("unknown city for apartments seed: %s", cityName)
		}
	} else {
		city = &cities[rand.Intn(len(cities))]
	}

	streets := loadLines("data/streets.txt")
	if len(streets) == 0 {
		log.Fatal("streets.txt is empty or not found")
	}
	var users []domain.User
	if err := db.Find(&users).Error; err != nil {
		log.Fatalf("failed to load users: %s", err)
	}
	if len(users) == 0 {
		log.Fatal("no users found to seed apartments")
	}
	user := users[rand.Intn(len(users))].ID
	var apartments []domain.Apartment

	for i := 0; i < seedApartmentCount; i++ {
		lat, lon := generateCoords(city.Latitude, city.Longitude)
		apartment := domain.Apartment{
			UserID:           user,
			City:             city.Name,
			Street:           randomFrom(streets),
			Building:         randomBuilding(),
			Floor:            rand.Intn(16) + 1,
			ApartmentNumber:  randomApartmentNumber(),
			Latitude:         lat,
			Longitude:        lon,
			Rooms:            rand.Intn(4) + 1,
			Area:             rand.Intn(120) + 20,
			Elevator:         rand.Intn(2) == 1,
			GarbageChute:     rand.Intn(2) == 1,
			BathroomType:     randomBathroomType(),
			Concierge:        rand.Intn(2) == 1,
			ConstructionYear: rand.Intn(30) + 1990,
			ConstructionType: randomConstructionType(),
			Remont:           randomRemont(),
			CreatedAt:        time.Now(),
			UpdatedAt:        time.Now(),
			Rating:           rand.Float32() * 5,
			Status:           "active",
		}
		apartments = append(apartments, apartment)
	}

	if err := db.Create(&apartments).Error; err != nil {
		log.Fatalf("failed to seed apartments: %s", err)
	}
	for _, apartment := range apartments {
		photos := pickAndCopyPhotos(rand, 3+rand.Intn(3)) // 3–5 фото
		var photodomain []domain.ApartmentPhoto

		for i, path := range photos {
			photodomain = append(photodomain, domain.ApartmentPhoto{
				ApartmentID: apartment.ID,
				URL:         path,
				IsCover:     i == 0,
			})
		}

		if err := db.Create(&photodomain).Error; err != nil {
			log.Fatalf("failed to seed apartment photos: %s", err)
		}
	}
	log.Printf("Seeded %d apartments in city %s", seedApartmentCount, city.Name)
}

func loadCities(path string) []City {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("failed to read cities file: %s", err)
	}
	var cities []City
	if err := json.Unmarshal(data, &cities); err != nil {
		log.Fatalf("failed to parse cities file: %s", err)
	}
	return cities
}

func findCityByName(cities []City, name string) *City {
	for _, city := range cities {
		if strings.EqualFold(city.Name, name) {
			return &city
		}
	}
	return nil
}

func generateCoords(baseLat, baseLon float64) (float64, float64) {
	offset := func() float64 {
		return (rand.Float64()*2 - 1) * 0.05
	}
	return baseLat + offset(), baseLon + offset()
}

func randomBuilding() string {
	return string('A' + rune(rand.Intn(5))) // A-E
}

func randomApartmentNumber() string {
	return string(rune(rand.Intn(200) + 1))
}

func randomBathroomType() string {
	types := []string{"совмещенный", "раздельный", "два санузла"}
	return randomFrom(types)
}

func randomConstructionType() string {
	types := []string{"панельный", "кирпичный", "монолитный", "блочный"}
	return randomFrom(types)
}

func randomRemont() string {
	remonts := []string{"евроремонт", "косметический", "без ремонта"}
	return randomFrom(remonts)
}
func pickAndCopyPhotos(rand *rand.Rand, count int) []string {
	srcDir := "data/apartment_photos"
	dstDir := "../../media/apartments_photo"

	files, err := os.ReadDir(srcDir)
	if err != nil {
		log.Fatalf("failed to read photo source dir: %s", err)
	}

	if len(files) == 0 {
		log.Fatal("no photos found in data/apartment_photos/")
	}

	var selected []string
	seen := map[string]bool{}

	for len(selected) < count {
		f := files[rand.Intn(len(files))].Name()
		if seen[f] {
			continue
		}
		seen[f] = true

		srcPath := srcDir + "/" + f
		dstPath := dstDir + "/" + randomFileName(rand, f)

		_ = os.MkdirAll(dstDir, os.ModePerm)

		srcData, err := os.ReadFile(srcPath)
		if err != nil {
			log.Fatalf("failed to read source image: %s", err)
		}
		if err := os.WriteFile(dstPath, srcData, 0644); err != nil {
			log.Fatalf("failed to write image to media: %s", err)
		}

		selected = append(selected, "/media/apartments_photo/"+getFileName(dstPath))
	}

	return selected
}

func randomFileName(rand *rand.Rand, original string) string {
	ext := ""
	if idx := strings.LastIndex(original, "."); idx != -1 {
		ext = original[idx:]
	}
	return randomString(rand, 10) + ext
}

func randomString(rand *rand.Rand, n int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyz0123456789")
	var sb strings.Builder
	for i := 0; i < n; i++ {
		sb.WriteRune(letters[rand.Intn(len(letters))])
	}
	return sb.String()
}

func getFileName(path string) string {
	parts := strings.Split(path, "/")
	return parts[len(parts)-1]
}
