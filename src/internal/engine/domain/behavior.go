package domain

type Behavior interface {
	Start()
	Update(ctx FrameContext)
}
