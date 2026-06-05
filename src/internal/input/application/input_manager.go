package application

import "github.com/InakiGT/white-dragon-engine/src/internal/input/domain"

type InputManager struct {
	driver   domain.DeviceDriver
	bindings map[domain.Action][]domain.Key
}

func NewInputManager(driver domain.DeviceDriver) *InputManager {
	return &InputManager{
		driver:   driver,
		bindings: make(map[domain.Action][]domain.Key),
	}
}

func (m *InputManager) BindAction(action domain.Action, keys ...domain.Key) {
	m.bindings[action] = keys
}

func (m *InputManager) IsActionActive(action domain.Action) bool {
	keys, exists := m.bindings[action]
	if !exists {
		return false
	}

	for _, key := range keys {
		if m.driver.IsKeyDown(key) {
			return true
		}
	}

	return false
}
