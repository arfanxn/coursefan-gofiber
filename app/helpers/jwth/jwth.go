package jwth

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// Encode encodes/generates a JWT token by the given signature (secret key) and claims (payload)
func Encode(signature string, payload map[string]any) (string, error) {
	tokenizer := jwt.New(jwt.SigningMethodHS256)
	claims := tokenizer.Claims.(jwt.MapClaims)

	// set authorized as true
	payload["authorized"] = true

	// check if expiration exists and if not give default expiration time
	_, ok := payload["exp"]
	if !ok {
		payload["exp"] = time.Now().Add(time.Minute * 30).Unix() // expires in N minutes
	}

	// fill claims with payload
	for key, value := range payload {
		claims[key] = value
	}
	// assign claims to tokenizer
	tokenizer.Claims = claims

	// sign token with the given signature
	token, err := tokenizer.SignedString([]byte(signature))

	return token, err
}

// Decode decodes/parse a JWT token by the given signature (secret key) and token (access token)
func Decode(signature string, token string) (*jwt.Token, error) {
	tokenizer, err := jwt.Parse(token, func(tokenizer *jwt.Token) (any, error) {
		_, ok := tokenizer.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			err := jwt.ErrTokenSignatureInvalid
			return nil, err
		}

		return []byte(signature), nil
	})

	return tokenizer, err
}
