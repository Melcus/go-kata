package user

type builder struct {
	user *user
}

func newUserBuilder() *builder {
	return &builder{user: &user{}}
}

func (b *builder) setName(name string) *builder {
	b.user.Name = name
	return b
}

func (b *builder) setEmail(email string) *builder {
	b.user.Email = email
	return b
}

func (b *builder) setPassword(password string) *builder {
	b.user.Password = password
	return b
}

func (b *builder) build() *user {
	return b.user
}
