package repository

type AuthRepository struct{}

func NewAuthRepository() *AuthRepository {

	return &AuthRepository{}
}

func (a *AuthRepository) SaveUserToken() {

}
