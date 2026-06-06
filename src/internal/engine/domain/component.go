package domain

type ComponentType string

const (
	ComponentRenderer ComponentType = "renderer"
	ComponentBehavoir ComponentType = "behavior"
	ComponentCollider ComponentType = "collider"
)

type Component interface {
	Type() ComponentType
	OnDestroy()
}
