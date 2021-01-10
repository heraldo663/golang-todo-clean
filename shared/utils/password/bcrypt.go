package password

import "golang.org/x/crypto/bcrypt"

type IBcrypt interface {
	Generate(raw string) string
	Verify(hash string, raw string) error
}

type bycriptImpl struct{}

func NewBcrypt() IBcrypt {
	return &bycriptImpl{}
}

// Generate return a hashed password
func (b *bycriptImpl) Generate(raw string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(raw), 10)

	if err != nil {
		panic(err)
	}

	return string(hash)
}

// Verify compares a hashed password with plaintext password
func (b *bycriptImpl) Verify(hash string, raw string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(raw))
}
