package domain

type DeviceDriver interface {
	IsKeyDown(key Key) bool
}
