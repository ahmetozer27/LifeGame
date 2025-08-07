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

## 📅 Simülasyon Akışı

* Belirlenen yıl ve zaman aralığına göre “saatlik” şekilde ilerler.
* Her adımda insanlar yaşlanır, ölebilir, çocuk sahibi olabilir.
* Olaylar (afet, salgın, savaş) belli olasılıklarla tetiklenir.

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
