package domain

type EntityID string

type Entity struct {
	Id         EntityID
	Parent     *Entity
	Nodes      []*Entity
	Transform  Transform
	Components map[ComponentType]Component
}

func NewEntity(id string, parent *Entity, transform Transform) *Entity {
	return &Entity{
		Id:         EntityID(id),
		Parent:     parent,
		Nodes:      make([]*Entity, 0),
		Transform:  transform,
		Components: make(map[ComponentType]Component),
	}
}

func (e *Entity) AddNode(node *Entity) {
	e.Nodes = append(e.Nodes, node)
}

func (e *Entity) ID() EntityID {
	return e.Id
}

func (e *Entity) Destroy() []EntityID {
	destroyed := []EntityID{}

	for _, child := range e.Nodes {
		destroyed = append(destroyed, child.Destroy()...)
	}

	for _, component := range e.Components {
		component.OnDestroy()
	}

	destroyed = append(destroyed, e.Id)
	return destroyed
}

func (e *Entity) AddComponent(c Component) {
	e.Components[c.Type()] = c
}
