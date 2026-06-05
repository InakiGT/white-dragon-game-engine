package domain

type Behavior interface {
	Update(ctx FrameContext)
}
