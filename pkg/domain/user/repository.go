package user

type userRepository interface {
	createUser(u user) (publicUser, error)
	findUserByEmail(email string) (*publicUser, error)
	listUsers() ([]*publicUser, error)
}

type MockUserRepository struct {
	users []user
}

func NewMockUserRepository() *MockUserRepository {
	return &MockUserRepository{}
}

func (repo *MockUserRepository) createUser(u user) (publicUser, error) {
	repo.users = append(repo.users, u)

	return publicUser{
		Name:  u.Name,
		Email: u.Email,
	}, nil
}

func (repo *MockUserRepository) findUserByEmail(email string) (*publicUser, error) {
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

func (repo *MockUserRepository) listUsers() ([]*publicUser, error) {
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
