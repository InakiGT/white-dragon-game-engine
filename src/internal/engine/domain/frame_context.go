package domain

import "time"

type FrameContext struct {
	DeltaTime float32
	TotalTime time.Duration
}

func NewFramwContext(deltaTime float32, totalTime time.Duration) *FrameContext {
	return &FrameContext{deltaTime, totalTime}
}
