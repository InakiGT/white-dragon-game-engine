package domain

type EntityRegistry struct {
	entities map[EntityID]*Entity
}

func NewEntityRegistry(entities map[EntityID]*Entity) *EntityRegistry {
	registry := &EntityRegistry{
		entities: make(map[EntityID]*Entity),
	}

	if entities != nil {
		registry.entities = entities
	}

	return registry
}

func (r *EntityRegistry) Remove(id EntityID) {
	delete(r.entities, id)
}

func (r *EntityRegistry) Get(id EntityID) (*Entity, bool) {
	e, ok := r.entities[id]
	return e, ok
}

func (r *EntityRegistry) GetAll() []*Entity {
	entities := make([]*Entity, 0)

	for _, entity := range r.entities {
		entities = append(entities, entity)
	}

	return entities
}
