package utils

import (
    "fmt"
    "time"
    "github.com/golang-jwt/jwt"
    // "encoding/json"
    // "net/http"
)

//func for generate a jwt token
func GenerateJWT(id int, username string) (string, error) {
    //create a map to store claims
    claims := jwt.MapClaims{}
    claims["authorized"] = true
    claims["user_id"] = id
    claims["username"] = username
    claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

    //create a new token with HS256 algorithm
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

    //sign the token with secret key
    return token.SignedString([]byte("secret"))
}

func VerifyJWT(token string) (jwt.MapClaims, bool) {
    //parse the token
    t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
        //check the signing method
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
        }

        //return the secret key
        return []byte("secret"), nil
    })

    //if there is an error, the token must have expired
    if err != nil {
        return nil, false
    }

    //check if the token is valid
    if claims, ok := t.Claims.(jwt.MapClaims); ok && t.Valid {
        return claims, true
    }

    return nil, false
}   