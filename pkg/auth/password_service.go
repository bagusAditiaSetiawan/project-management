package auth

type PasswordService interface {
	Hashing(password string) (string, error)
	Compare(password string, hashed string) bool
}
