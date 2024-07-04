package helper

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

var (
	SECRET_KEY []byte = []byte("20242213960490267598567908189775563674735941821979632997103732363925894510946901202218157")
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GenerateUserJWT(name, email string, exp time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"name":  name,
			"email": email,
			"exp":   time.Now().Add(exp).Unix(),
		})
	tokenString, err := token.SignedString(SECRET_KEY)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// func ValidateUserJWT(token string) bool {
// 	jwttoken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
// 		return SECRET_KEY, nil
// 	})
// 	if err != nil {
// 		return false
// 	}

// 	return jwttoken.Valid
// }

func ValidateUserJWT(tokenString string) (bool, string, error) {
    // Mendefinisikan klaim yang diharapkan
    claims := jwt.MapClaims{}

    // Mendekode dan memverifikasi token
    token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
        return SECRET_KEY, nil
    })

    if err != nil {
        return false, "", err
    }

    if !token.Valid {
        return false, "", fmt.Errorf("invalid token")
    }

    // Mengambil nilai email dari klaim
    email, ok := claims["email"].(string)
    if !ok {
        return false, "", fmt.Errorf("email not found in token")
    }

    return true, email, nil
}
