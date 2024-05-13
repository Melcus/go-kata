package user

type Manager struct {
	repo userRepository
}

func NewUserManager(repo userRepository) *Manager {
	return &Manager{repo: repo}
}

func (manager *Manager) createUser(u user) (publicUser, error) {
	return manager.repo.createUser(u)
}

func (manager *Manager) findUserByEmail(email string) (*publicUser, error) {
	return manager.repo.findUserByEmail(email)
}

func (manager *Manager) listUsers() ([]*publicUser, error) {
	return manager.repo.listUsers()
}
