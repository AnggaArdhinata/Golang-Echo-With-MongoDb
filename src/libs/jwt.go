package libs

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type claims struct {
	UserId primitive.ObjectID
	IsAdmin bool
	jwt.RegisteredClaims
}

var myKey = []byte(os.Getenv("JWT_KEY"))

func NewToken(id primitive.ObjectID, role bool) *claims {
	return &claims{
		UserId: id,
		IsAdmin: role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 10)),
		},
	}
}

func (c *claims)  Create() (string, error){
	tokens := jwt.NewWithClaims(jwt.SigningMethodHS384, c)

	return tokens.SignedString(myKey)
}

func ChekToken(token string) (*claims, error)  {
	tokens, err := jwt.ParseWithClaims(token, &claims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(myKey), nil
	})
	if err != nil {
		return nil, err
	}
	claims := tokens.Claims.(*claims)
	return claims, nil
}