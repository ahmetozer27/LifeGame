package main

import (
	"fmt"
	"life_game/configs"
	"life_game/internal/person"
	"life_game/internal/person/models"
	"life_game/internal/simulation"
	"log"
	"time"
)

func printHumanStatus(lifecycle *simulation.LifecycleManager) {
	for _, h := range lifecycle.Humans {
		fmt.Printf(
		"                   -----------------------------------\n"+
		"%s, yaş: %d, hayatta mı? %v, doğum yılı: %v, ölüm yılı: %v\n Sağlık verileri: %v, Akıl becerileri: %v \n"+
		"             ------------------------------------------", h.Name, h.Age, h.Status.Alive,h.BirthDate,h.DeathDate,h.Health,h.Traits)
	}
}

func main() {
	err := configs.LoadSettings("../configs/settings.yaml")
	if err != nil {
		log.Fatalf("Ayarlar yüklenemedi: %v", err)
	}

	fmt.Println("Başlatılıyor...")

	// Lifecycle manager'dan insanları yönet
	lifecycle := simulation.NewLifecycleManager()

	// Olay yöneticisi oluştur 
	eventmanager := simulation.NewEventManager()

	// Yeni engine oluştur
	engine := simulation.NewSimEngine(lifecycle,eventmanager)


	// Örnek insan oluştur ve ekle
	human := person.NewHuman("Ahmet", models.Male, time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC), "", "", "Istanbul", "Marmara")
	lifecycle.AddHuman(human)


	// Simülasyonu çalıştır
	engine.Start(100)

	engine.Stop() // Simülasyonu durdur
	// Simülasyon bittikten sonra insanların durumunu yazdır
	printHumanStatus(lifecycle)
	
}
