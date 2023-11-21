package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/golang-jwt/jwt/v4"
	"math/rand"
	"strings"
	"time"
)

var jwtKey = []byte("swiwgq")

type Claims struct {
	UserId uint
	jwt.StandardClaims
}

func GetToken(uid uint) (string, error) {
	expirationTime := time.Now().Add(7 * 24 * time.Hour)
	claims := &Claims{
		UserId: uid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "toplink.club",
			Subject:   "userToken",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenStr, nil
}

func ParseToken(tokenStr string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(
		tokenStr,
		claims,
		func(token *jwt.Token) (interface{},
			error) {
			return jwtKey, nil
		})

	return token, claims, err
}

func RandStringAndNumber(length int) string {
	var letters = []byte("ABCDEF1234567890")
	result := make([]byte, length)

	source := rand.NewSource(time.Now().UnixNano())
	rand.New(source)
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}

	return strings.ToUpper(string(result))
}

func HashStr(str1 string, str2 string, str3 string) string {
	hasher := sha256.New()
	hasher.Write([]byte(str1))
	hasher.Write([]byte(str2))
	hasher.Write([]byte(str3))
	hashBytes := hasher.Sum(nil)
	hashHex := hex.EncodeToString(hashBytes)
	return hashHex
}
