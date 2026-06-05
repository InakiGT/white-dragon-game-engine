package domain

type EntityID string

type Entity struct {
	Id        EntityID
	Transform Transform
	Nodes     []*Entity
}

func NewEntity(id string) *Entity {
	return &Entity{
		Id:        EntityID(id),
		Transform: *NewTransform(0, 0),
		Nodes:     make([]*Entity, 0),
	}
}

func (e *Entity) AddNode(node *Entity) {
	e.Nodes = append(e.Nodes, node)
}

func (e *Entity) ID() EntityID {
	return e.Id
}
