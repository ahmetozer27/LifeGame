package simulation

import (
	"fmt"
	"life_game/configs" // Ayarları okumak için
	"log"
	"time"
)
var tickLogger = log.New(log.Writer(), "TICK: ", log.LstdFlags)

// SimEngine simülasyonun ana motorudur.
// Zaman adımlarını (tick) yönetir ve simülasyonun ilerlemesini sağlar.
type SimEngine struct {
	TickInterval time.Duration   // Her tick arasında geçen süre
	running      bool            // Simülasyonun çalışıp çalışmadığı durumu
	Lifecycle    *LifecycleManager // Simülasyon içindeki insanları yöneten yapı
	EventManager *EventManager // Olayları yöneten yapı (isteğe bağlı, eğer olaylar eklenirse kullanılabilir)
}

// NewSimEngine ayarlardaki tick süresi (milisaniye cinsinden) alınarak yeni bir SimEngine oluşturur.
// lifecycle parametresi simülasyondaki insanları yönetir.
func NewSimEngine(lifecycle *LifecycleManager, EventManager *EventManager) *SimEngine {
	tickMs := configs.LoadedSettings.Simulation.TickIntervalMs // Ayarlardan tick süresini al
	return &SimEngine{
		TickInterval: time.Duration(tickMs) * time.Millisecond, // time.Duration türüne dönüştür (milisaniye cinsinden)
		Lifecycle:    lifecycle,                                 // LifecycleManager'ı ata
		EventManager: EventManager, 						   // Olay yöneticisini ata 
	}
}

// Start engine çalıştırır ve tick döngüsünü başlatır.
// Her tick süresinde Lifecycle içindeki insanları günceller.
func (e *SimEngine) Start(finish int) {
	e.running = true
	ticker := time.NewTicker(e.TickInterval) // Ayarlardaki tick süresi ile zamanlayıcı oluşturulur
	defer ticker.Stop()                      // Fonksiyon sonunda zamanlayıcı durdurulur

	for e.running && finish > 0 {
		<-ticker.C       // Tick süresi kadar beklenir
		e.Tick()         // Her tick'te simülasyon güncelleme işlemi yapılır
		finish--         // Her tick'te finish sayacı azaltılır
	}
}

// Tick fonksiyonu simülasyondaki her zaman adımında yapılacak işleri yürütür.
// Şu anda sadece insan yaşam döngüsü güncelleniyor.
func (e *SimEngine) Tick() {
	tickLogger.Println("Simülasyon tick atıldı")
	configs.LoadedSettings.Simulation.CurrentYear++ // +1 Yıl güncellemesi
	e.EventManager.Emit(EventNewYear, fmt.Sprintf("Yeni yıl başladı: %d", configs.LoadedSettings.Simulation.CurrentYear), nil)
	e.Lifecycle.UpdateAllHumans() // Tüm insanların yaşını ve sağlık durumlarını güncelle
	// İleride başka sistemlerin de tick fonksiyonları buraya eklenebilir (örneğin, dünya durumu, olaylar vb.)
}

// Stop simülasyonu durdurur.
func (e *SimEngine) Stop() {
	e.running = false
}
