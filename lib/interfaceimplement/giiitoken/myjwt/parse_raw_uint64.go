package myjwt

import (
	_interface "github.com/ghafbu/padegan-backend/lib/interface"
	"github.com/golang-jwt/jwt/v5"
)

func parseRawUint64[T interface{}](secret, requestToken string) (*_interface.CustomClaims[T], error) {
	// token, err := jwt.Parse(requestToken, func(token *jwt.Token) (interface{}, error) {
	//	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
	//		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	//	}
	//	return []byte(secret), nil // generate custom secret
	// })

	token, err := jwt.ParseWithClaims(
		requestToken,
		&_interface.CustomClaims[T]{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil // secret is generateCustomTokenSecret
		})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*_interface.CustomClaims[T]); ok && token.Valid {
		return claims, nil
	}

	return nil, err
}
