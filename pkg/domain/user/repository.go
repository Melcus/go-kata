package user

type userRepository interface {
	createUser(u user) (publicUser, error)
	findUserByEmail(email string) (*publicUser, error)
	listUsers() ([]*publicUser, error)
}

type mockUserRepository struct {
	users []user
}

func newMockUserRepository() *mockUserRepository {
	return &mockUserRepository{}
}

func (repo *mockUserRepository) createUser(u user) (publicUser, error) {
	repo.users = append(repo.users, u)

	return publicUser{
		Name:  u.Name,
		Email: u.Email,
	}, nil
}

func (repo *mockUserRepository) findUserByEmail(email string) (*publicUser, error) {
	for _, user := range repo.users {
		if user.Email == email {
			return &publicUser{
				Name:  user.Name,
				Email: user.Email,
			}, nil
		}
	}
	return nil, nil
}

func (repo *mockUserRepository) listUsers() ([]*publicUser, error) {
	publicUsers := make([]*publicUser, 0, len(repo.users))

	for _, user := range repo.users {
		publicUser := &publicUser{
			Name:  user.Name,
			Email: user.Email,
		}
		publicUsers = append(publicUsers, publicUser)
	}

	return publicUsers, nil
}
