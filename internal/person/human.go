package person

import (
	"life_game/configs"
	"life_game/internal/person/models"
	"life_game/internal/utils"
	"math/rand"
	"time"
)

// NewHuman ayarlara göre yeni bir kişi oluşturur (basit rastgele ile)
func NewHuman(name string, gender models.Gender, birthDate time.Time, motherID string, fatherID string, cityOfBirth string, regionOfBirth string) *models.Human {
	settings := configs.LoadedSettings // Ayarlar global değişkeninden alınıyor

	// Yardımcı: int aralığında rastgele sayı üretir
	randomInRange := func(min, max int) int {
		if max <= min {
			return min
		}
		return rand.Intn(max-min+1) + min
	}

	return &models.Human{
		ID:        utils.CreateNewID(),       // Benzersiz ID oluştur
		Name:      name,                      // Kişi adı
		Gender:    gender,                    // Cinsiyet
		BirthDate: birthDate,                 // Doğum tarihi
		Age:       0,                        // Yaş (başlangıçta 0)
		Generation: 1,                       // Nesil (ilk nesil = 1)

		Traits: models.Traits{                      // Kişisel özellikler ayarlardan aralıklarla atanır
			Intelligence: randomInRange(settings.Attributes.Intelligence.Min, settings.Attributes.Intelligence.Max),
			Strength:     randomInRange(settings.Attributes.Strength.Min, settings.Attributes.Strength.Max),
			Empathy:      randomInRange(settings.Attributes.Empathy.Min, settings.Attributes.Empathy.Max),
			Creativity:   randomInRange(settings.Attributes.Creativity.Min, settings.Attributes.Creativity.Max),
			Resilience:   randomInRange(settings.Attributes.Resilience.Min, settings.Attributes.Resilience.Max),
		},

		Health: models.Health{
			CurrentHealth: 100,              // Başlangıç sağlığı tam
			MaxHealth:     100,              // Maksimum sağlık
			Diseases:      []models.Disease{},      // Başlangıçta hastalık yok
			Immunity:      make(map[string]float64),
		},

		Needs: models.Needs{
			Hunger:  randomInRange(0, settings.Needs.HungerMax),   // Açlık ihtiyacı
			Sleep:   randomInRange(0, settings.Needs.SleepMax),    // Uyku ihtiyacı
			Social:  randomInRange(0, settings.Needs.SocialMax),   // Sosyal ihtiyaç
			Hygiene: randomInRange(0, settings.Needs.HygieneMax),  // Hijyen ihtiyacı
		},

		Skills: models.Skills{
			Communication: randomInRange(0, settings.Skills.CommunicationMax),
			Crafting:      randomInRange(0, settings.Skills.CraftingMax),
			Farming:       randomInRange(0, settings.Skills.FarmingMax),
		},

		Relationships: models.Relationships{
			Parents:  []string{motherID, fatherID}, // Anne-baba ID'leri
			Children: []string{},                    // Başlangıçta çocuk yok
			Partner:  "",                           // Partner yok
			Friends:  []string{},                   // Arkadaşlar boş
		},

		Location: models.Location{
			City:        cityOfBirth,          // Doğduğu şehir
			Region:      regionOfBirth,        // Doğduğu bölge
			Coordinates: [2]float64{0, 0},     // Koordinatlar başlangıçta sıfır
		},

		JobTitle:  "Unemployed",  // İşsiz
		Education: "None",        // Eğitim yok
		Income:    0,             // Gelir sıfır

		Status: models.Status{
			Alive:        true,    // Hayatta
			Pregnant:     false,   // Hamile değil
			Incarcerated: false,   // Hapis değil
		},

		Genetics: models.Genetics{
			DNACode:      generateRandomDNA(100), // 100 karakter DNA
			MutationImpact: settings.Mutation.Impact,         // Mutasyon oranı ayarlardan
		},
	}
}



// generateRandomDNA basit DNA dizisi üretir
func generateRandomDNA(length int) string {
	bases := []rune{'A', 'T', 'C', 'G'}
	dna := make([]rune, length)
	for i := 0; i < length; i++ {
		dna[i] = bases[rand.Intn(len(bases))]
	}
	return string(dna)
}


