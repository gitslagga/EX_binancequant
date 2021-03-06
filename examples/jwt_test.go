package examples

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

// For HMAC signing method, the key can be any []byte. It is recommended to generate
// a key using crypto/rand or something equivalent. You need the same key for signing
// and validating.
var hmacSampleSecret []byte

func init() {
	// Load sample key data
	hmacSampleSecret = []byte("bitway-todo_block")
}

// Example creating, signing, and encoding a JWT token using the HMAC signing method
func ExampleNew_hmac() {
	timeByte, _ := json.Marshal(time.Now().Unix())
	jti := fmt.Sprintf("%x", md5.Sum(timeByte)) //将[]byte转成16进制

	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS384, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour).Unix(),
		Id:        jti,
		IssuedAt:  time.Now().Unix(),
		Issuer:    "Bitway",
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(hmacSampleSecret)

	fmt.Println(tokenString, err)
	// Output: eyJhbGciOiJIUzM4NCIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1OTg0NDE1NTksImp0aSI6IjY3N2I2OTJjZGM2MGYwYjQ3MGNhY2QzMTNkZjQ0NjQ4IiwiaWF0IjoxNTk4NDM3OTU5LCJpc3MiOiJCaXR3YXkifQ.fEodlAS1Ov_6PGdk7Id2ZgOTZritQuPgV87_iiW06Z9Fd3wxXHVWe6YXT2STDeSW <nil>
}

// Example parsing and validating a token using the HMAC signing method
func ExampleParse_hmac() {
	// sample token string taken from the New example
	tokenString := "eyJhbGciOiJIUzM4NCIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1OTg0NDE1NTksImp0aSI6IjY3N2I2OTJjZGM2MGYwYjQ3MGNhY2QzMTNkZjQ0NjQ4IiwiaWF0IjoxNTk4NDM3OTU5LCJpc3MiOiJCaXR3YXkifQ.fEodlAS1Ov_6PGdk7Id2ZgOTZritQuPgV87_iiW06Z9Fd3wxXHVWe6YXT2STDeSW"

	// Parse takes the token string and a function for looking up the key. The latter is especially
	// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
	// head of the token to identify which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return hmacSampleSecret, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims["exp"], claims["jti"], claims["iat"], claims["iss"])
	} else {
		fmt.Println(err)
	}

	// Output: 1.598441559e+09 677b692cdc60f0b470cacd313df44648 1.598437959e+09 Bitway
}
