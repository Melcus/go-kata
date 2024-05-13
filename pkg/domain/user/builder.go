package user

type builder struct {
	user *user
}

func newUserBuilder() *builder {
	return &builder{user: &user{}}
}

func (b *builder) SetName(name string) *builder {
	b.user.Name = name
	return b
}

func (b *builder) SetEmail(email string) *builder {
	b.user.Email = email
	return b
}

func (b *builder) SetPassword(password string) *builder {
	b.user.Password = password
	return b
}

func (b *builder) Build() *user {
	return b.user
}
