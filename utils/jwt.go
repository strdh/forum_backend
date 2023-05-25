package utils

import (
    "time"
    "github.com/golang-jwt/jwt"
    "xyzforum/models"
    // "encoding/json"
    // "net/http"
)

func GenerateToken(user models.User, key string) string {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "iss": "xyzforum",
        "sub": user.Id,
        "iat": time.Now().Unix(),
        "exp": time.Now().Add(time.Hour * 24).Unix(),
    })

    finalToken, _ := token.SignedString([]byte(key))

    return finalToken
}

func VerifyJWT(token string, key string) (jwt.MapClaims, bool) {
    claims := jwt.MapClaims{}
    t, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
        return []byte(key), nil
    })

    if err != nil {
        return nil, false
    }

    return claims, t.Valid
}