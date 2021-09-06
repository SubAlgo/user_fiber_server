package password

import "golang.org/x/crypto/bcrypt"

func Hash(password string) (hashPassword string, err error) {
	var bytePassword []byte
	bytePassword, err = bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return hashPassword, err
	}
	hashPassword = string(bytePassword)
	return hashPassword, err
}

func Compare(hashPassword string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
}
