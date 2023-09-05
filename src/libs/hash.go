package libs

import "golang.org/x/crypto/bcrypt"

func HashPassword(pass string) (string, error) {
	hashPass, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashPass), nil
}

func CheckPassword(hashedPass, bodyPass string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPass), []byte(bodyPass))

	if err != nil {
		return err == nil
	}
	return true
}
