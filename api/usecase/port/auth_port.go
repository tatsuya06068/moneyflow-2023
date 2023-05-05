package port

import "context"

type IAuthInputport interface {
	Signup(ctx context.Context, userName string, password string)
}

type IAuthOutputPort interface {
	Render(bool)
	RenderError(error)
}
