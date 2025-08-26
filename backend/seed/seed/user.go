package seed

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/scmbr/renting-app/internal/config"
	"github.com/scmbr/renting-app/internal/domain"
	"github.com/scmbr/renting-app/pkg/hash"
	"gorm.io/gorm"
)

const (
	defaultPassword = "password"
	seedCount       = 100
)

func SeedUsers(db *gorm.DB, cfg *config.Config) {
	rand := rand.New(rand.NewSource(time.Now().UnixNano()))

	maleNames := loadLines("data/male_names_rus.txt")
	femaleNames := loadLines("data/female_names_rus.txt")
	maleSurnames := loadLines("data/male_surnames_rus.txt")

	hasher := hash.NewSHA1Hasher(cfg.Auth.PasswordSalt)
	passwordHash, err := hasher.Hash(defaultPassword)
	if err != nil {
		log.Fatalf("failed to hash password: %s", err)
	}

	var users []domain.User

	for i := 0; i < seedCount; i++ {
		gender := rand.Intn(2) + 1

		var name, surname string

		if gender == 1 {
			name = randomFrom(maleNames)
			surname = randomFrom(maleSurnames)
		} else {
			name = randomFrom(femaleNames)
			surname = feminizeSurname(randomFrom(maleSurnames))
		}

		email := fmt.Sprintf("%s.%s%d@example.com",
			strings.ToLower(name),
			strings.ToLower(surname),
			rand.Intn(100000),
		)

		user := domain.User{
			Name:         name,
			Surname:      surname,
			Email:        email,
			PasswordHash: passwordHash,
			Role:         "user",
			Birthdate:    randomBirthdate(),
			Gender:       gender,
			IsActive:     true,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		}

		users = append(users, user)
	}

	if err := db.Create(&users).Error; err != nil {
		log.Fatalf("failed to insert users: %s", err)
	}

	log.Printf("Seeded %d users", seedCount)
}

func loadLines(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("failed to open file %s: %s", path, err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			lines = append(lines, line)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("scanner error: %s", err)
	}

	return lines
}

func randomFrom(list []string) string {
	return list[rand.Intn(len(list))]
}

func feminizeSurname(s string) string {
	s = strings.TrimSpace(s)
	if !strings.HasSuffix(s, "а") {
		return s + "а"
	}
	return s
}

func randomBirthdate() time.Time {
	year := rand.Intn(30) + 1975
	month := time.Month(rand.Intn(12) + 1)
	day := rand.Intn(28) + 1
	return time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
}
