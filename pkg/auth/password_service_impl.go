package auth

import "golang.org/x/crypto/bcrypt"

type PasswordServiceImpl struct {
}

func NewPasswordServiceImpl() *PasswordServiceImpl {
	return &PasswordServiceImpl{}
}

func (service *PasswordServiceImpl) Hashing(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (service *PasswordServiceImpl) Compare(password string, hashed string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
	return err == nil
}
