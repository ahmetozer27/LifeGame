package simulation

import (
	"life_game/internal/person/models"
)



// LifecycleManager simüle edilen insanları yönetir
type LifecycleManager struct {
	Humans []*models.Human

}

// NewLifecycleManager yeni LifecycleManager oluşturur
func NewLifecycleManager() *LifecycleManager {
	return &LifecycleManager{
		Humans: []*models.Human{},
	}
}

// AddHuman listeye yeni insan ekler
func (lm *LifecycleManager) AddHuman(h *models.Human) {
	lm.Humans = append(lm.Humans, h)
}

// UpdateAllHumans her tick'te tüm insanların yaşını arttırır ve sağlık durumlarını günceller
func (lm *LifecycleManager) UpdateAllHumans() {
	for _, human := range lm.Humans {
		if human.Status.Alive {
			human.AgeOneYear() // Her çağrıldığında 1 yıl eklemek gerçek senaryoda ayarlanabilir
		}
	}
}


