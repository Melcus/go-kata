package user

import "eon/kata/mike/pkg/kernel"

type provider struct {
	repo userRepository
}

func InitDomain(app *kernel.Application) {
	provider := &provider{
		repo: newMockUserRepository(),
	}

	loadRoutes(app, provider)
}
