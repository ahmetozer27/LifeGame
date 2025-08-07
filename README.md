# 🌍 SimWorld: İnsan Evrimi ve Dünya Simülasyonu

SimWorld, Go dilinde yazılmış bir dünya simülasyonu projesidir. İnsanların doğum, yaşlanma, ölüm gibi süreçlerini; afetler, hastalıklar ve savaş gibi olayları içeren gelişmiş bir yaşam döngüsü modeli sunar.

## 📁 Proje Yapısı

```
LifeGame/
├── cmd/
│   └── main.go
│
├── internal/
│   ├── person/
│   │   ├── models/
│   │   │   └── human_model.go
│   │   ├── person.go
│   │   ├── health.go
│   │   ├── needs.go
│   │   ├── skills.go
│   │   ├── emotions.go
│   │   └── genetics.go
│   │
│   ├── world/
│   │   ├── map.go
│   │   ├── climate.go
│   │   ├── disasters.go
│   │   ├── conflicts.go
│   │   └── events.go
│   │
│   ├── interaction/
│   │   ├── relationship.go
│   │   └── communication.go
│   │
│   ├── economy/
│   │   ├── job.go
│   │   ├── income.go
│   │   └── economy.go
│   │
│   ├── simulation/
│   │   ├── engine.go
│   │   ├── lifecycle.go
│   │   └── event_manager.go
│   │
│   └── utils/
│       ├── random.go
│       ├── id.go
│       └── logger.go
│
├── configs/
│   ├── load_settings.go
│   └── settings.yaml
│
├── assets/
│   ├── male_names.txt
│   └── female_names.txt
│
├── go.mod
├── go.sum
└── README.md

```

## ⚙️ Ayarlar (settings.yaml)

Tüm simülasyon ayarları `settings.yaml` dosyasından yüklenir. Örnek içerik:

```yaml
simulation:
  tick_interval_ms: 1000
  max_population: 10000
  start_year: 2025

world:
  disasters:
    earthquake_chance: 0.01
    epidemic_chance: 0.02
    war_chance: 0.03
  temperature:
    min: -30
    max: 50

human:
  max_age: 120
  reproduction_age_range: [18, 40]
  death_chance_per_year: 0.01
```

## 📅 Simülasyon Akışı ve Modüller

Simülasyon aşağıdaki modüller ve öncelik sırasına göre geliştirilmiştir:

### 1. Simülasyon Çekirdeği (`internal/simulation/engine.go`)

- **Amaç:**  
  - `tick_interval_ms` ile belirlenen aralıkta zaman akışı (tick) mekanizması sağlar.  
  - Her tick’te yaşayan tüm bireylerin yaş, ihtiyaç, sağlık gibi durumlarını günceller.  
  - Doğum, ölüm ve popülasyon kontrol mekanizmalarını başlatır.  

- **Neden Öncelikli?**  
  Simülasyonun temel zaman akışı olmadan diğer modüller çalışamaz.

---

### 2. Yaşam Döngüsü Yönetimi (`internal/simulation/lifecycle.go`)

- **Amaç:**  
  - İnsanların doğum, yaşlanma, ölüm süreçlerini detaylandırır.  
  - Üreme, hamilelik ve çocuk doğumu mekanizmalarını simüle eder.  
  - Yaş sınırı ve sağlık durumuna bağlı ölüm kontrolleri sağlar.  

- **Neden Öncelikli?**  
  İnsan yaşam döngüsünün doğru yönetimi simülasyonun temelidir.

---

### 3. Olay Yönetimi (`internal/simulation/event_manager.go`)

- **Amaç:**  
  - Deprem, salgın, savaş gibi rastgele ve süreli olayları simüle eder.  
  - Bu olayların bireylere ve çevreye etkilerini uygular.  

- **Neden Öncelikli?**  
  Dinamik olaylar simülasyonun gerçekçiliğini artırır.

---

### 4. Dünya Modeli (`internal/world/`)

- **Amaç:**  
  - Harita, şehirler, bölgeler ve iklim modellerini oluşturur.  
  - İnsanların konum ve hareketlerini yönetir.  

- **Neden Öncelikli?**  
  Coğrafi bağlam simülasyonun fiziksel temelini oluşturur.

---

### 5. Sosyal Etkileşimler (`internal/interaction/`)

- **Amaç:**  
  - İnsanlar arasındaki ilişkileri ve iletişimi yönetir.  
  - Aile, arkadaşlık, partnerlik gibi bağları simüle eder.  
  - Sosyal ihtiyaçları karşılamak için algoritmalar geliştirir.  

- **Neden Önemli?**  
  Sosyal bağlar birey davranışlarını ve simülasyon dinamiklerini etkiler.

---

### 6. Ekonomi (`internal/economy/`)

- **Amaç:**  
  - Meslek, gelir, harcama ve ekonomik döngüyü simüle eder.  
  - İnsanların iş ve gelir durumlarını yönetir.  

- **Neden Önemli?**  
  Ekonomik faktörler simülasyonun derinliğini artırır.

---

### 7. İsim ve Genetik Modülleri (`internal/person/genetics.go` vb.)

- **Amaç:**  
  - Genetik kodların nesilden nesile geçişini ve mutasyonları modellemek.  
  - Genetik yapıların birey özelliklerine etkisini simüle etmek.  

- **Neden Önemli?**  
  Biyolojik çeşitlilik ve kalıtım simülasyonun gerçekçiliğini sağlar.


## 👥 İsim Verisi

* İnsanların isimleri `data/names/male_names.txt` ve `data/names/female_names.txt` dosyalarından rastgele çekilir.
* Her satır bir isim olmalıdır.

## 🚀 Çalıştırmak

```bash
go run main.go
```

## 🚜 Gelecek Planlar

* Grafik arayüz eklenmesi
* Gerçek zamanlı istatistik takibi
* Çoklu dünya desteği

---

MIT Lisanslıdır © 2025
