package user

type manager struct {
	repo userRepository
}

func newUserManager(repo userRepository) *manager {
	return &manager{repo: repo}
}

func (manager *manager) createUser(u user) (publicUser, error) {
	return manager.repo.createUser(u)
}

func (manager *manager) findUserByEmail(email string) (*publicUser, error) {
	return manager.repo.findUserByEmail(email)
}

func (manager *manager) listUsers() ([]*publicUser, error) {
	return manager.repo.listUsers()
}
