package domain

type InputProvider interface {
	IsActionActive(action Action) bool
}
