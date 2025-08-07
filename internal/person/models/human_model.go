package models

import (
	"fmt"
	"time"
)

// Gender tipini enum gibi kullanmak için
type Gender string

const (
    Male   Gender = "male"
    Female Gender = "female"
)

// Disease hastalık bilgisi
type Disease struct {
	Name      string
	Severity  int       // 0-100
	StartDate time.Time
}

// Health sağlık durumu
type Health struct {
	CurrentHealth int
	MaxHealth     int
	Diseases      []Disease
	Immunity      map[string]float64 // hastalıklara karşı bağışıklık oranı 0-1 arası
}

// Needs temel ihtiyaçlar
type Needs struct {
	Hunger  int // 0-100, 0 çok tok, 100 çok aç
	Sleep   int
	Social  int
	Hygiene int
}

// Skills yetenekler
type Skills struct {
	Communication int
	Crafting      int
	Farming       int
}

// Traits karakter özellikleri
type Traits struct {
	Intelligence int
	Strength     int
	Empathy      int
	Creativity   int
	Resilience   int
}

// Relationships ilişkiler (diğer kişi ID'leri ile)
type Relationships struct {
	Parents  []string
	Children []string
	Partner  string
	Friends  []string
}

// Genetics genetik bilgiler
type Genetics struct {
	DNACode      string
	MutationImpact float64
}

// Location konum bilgisi
type Location struct {
	City        string
	Region      string
	Coordinates [2]float64
}

// Status canlılık, hamilelik vb. durumlar
type Status struct {
	Alive        bool
	Pregnant     bool
	Incarcerated bool
}

// Human temel insan yapısı
type Human struct {
	ID            string
	Name          string
	Gender        Gender
	BirthDate     time.Time
	DeathDate     *time.Time
	Age           int // yıl cinsinden
	Generation    int
	Traits        Traits
	Health        Health
	Needs         Needs
	Skills        Skills
	Relationships Relationships
	Location      Location
	JobTitle      string
	Education     string
	Income        int
	Status        Status
	Genetics      Genetics
}

// AgeOneYear yaşı 1 yıl arttırır ve sağlık durumunu etkiler
func (p *Human) AgeOneYear() {
	if !p.Status.Alive {
		return
	}
	p.Age++
	p.NeedTick()
	p.HealthTick()
	
}

// NeedTick temel ihtiyaçları günceller
func (p *Human) NeedTick() {
	// Açlık, uyku, sosyal, hijyen ihtiyacı artar
	p.Needs.Hunger += 10
	p.Needs.Sleep += 8
	p.Needs.Social += 5
	p.Needs.Hygiene += 7

	// Limit aşımı kontrolü
	if p.Needs.Hunger > 100 {
		p.Needs.Hunger = 100
		p.Health.CurrentHealth -= 10
	}
	if p.Needs.Sleep > 100 {
		p.Needs.Sleep = 100
		p.Health.CurrentHealth -= 8
	}
	if p.Needs.Social > 100 {
		p.Needs.Social = 100
		p.Health.CurrentHealth -= 5
	}
	if p.Needs.Hygiene > 100 {
		p.Needs.Hygiene = 100
		p.Health.CurrentHealth -= 6
	}
}

// HealthTick sağlık kontrolü, hastalıklar ve bağışıklıkla ilgili işlemler
func (p *Human) HealthTick() {
	// Sağlık Kontrolü
	if p.Health.CurrentHealth <= 0 {
		p.Die("health is under 0")
		return
	}

	// Örnek: 80 yaş üstü sağlık azalır
	if p.Age > 80 {
		p.Health.CurrentHealth -= 5
		if p.Health.CurrentHealth <= 0 {
			p.Die("natural causes")
		}
	}


	// Hastalıklar varsa sağlık azalır
	for _, d := range p.Health.Diseases {
		p.Health.CurrentHealth -= d.Severity / 10
		if p.Health.CurrentHealth <= 0 {
			p.Die(fmt.Sprintf("disease: %s", d.Name))
			return
		}
	}
	// Bağışıklık etkisi vb. burada hesaplanabilir
}

// Die kişinin ölümü gerçekleşir
func (p *Human) Die(reason string) {
	p.Status.Alive = false
	now := time.Now()
	p.DeathDate = &now
	fmt.Printf("%s has died due to %s at age %d\n", p.Name, reason, p.Age)
}