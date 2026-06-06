package application

import (
	"github.com/InakiGT/white-dragon-engine/src/internal/engine/domain"
)

type EntityManager struct {
	Registry         *domain.EntityRegistry
	pendingToDestroy []domain.EntityID
}

func NewEntityManager(registry *domain.EntityRegistry) *EntityManager {
	return &EntityManager{registry, make([]domain.EntityID, 0)}
}

func (m *EntityManager) Mark(id domain.EntityID) {
	m.pendingToDestroy = append(m.pendingToDestroy, id)
}

func (m *EntityManager) Flush() {
	for _, id := range m.pendingToDestroy {
		entity, ok := m.Registry.Get(id)
		if !ok {
			continue
		}

		destroyed := entity.Destroy()
		for _, deadID := range destroyed {
			m.Registry.Remove(deadID)
		}
	}

	m.pendingToDestroy = m.pendingToDestroy[:0]
}
