package auth

import (
	"crypto/rsa"
	"io/ioutil"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GetRSAPublicKey() (*rsa.PublicKey, error) {
	keyData, err := ioutil.ReadFile("id_rsa.pub.pkcs8")

	if err != nil {
		return nil, err
	}

	return jwt.ParseRSAPublicKeyFromPEM(keyData)
}

func GenerateToken() (string, error) {
	keyData, err := ioutil.ReadFile("./id_rsa")

	if err != nil {
		return "", err
	}

	key, err := jwt.ParseRSAPrivateKeyFromPEM(keyData)

	if err != nil {
		return "", err
	}

	token := jwt.New(jwt.SigningMethodRS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["admin"] = true
	claims["sub"] = "uuid"
	claims["name"] = "tommy"
	claims["iat"] = time.Now()
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	t, err := token.SignedString(key)
	if err != nil {
		return "", err
	}

	return t, nil
}
