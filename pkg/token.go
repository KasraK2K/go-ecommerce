package pkg

import (
	"errors"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"app/config"
)

type token struct {
	Platform Platform `json:"platform"  bson:"platform"`
	UserType UserType `json:"user_type" bson:"user_type"`
}

type Platform struct {
	Web int `json:"web" bson:"web"`
	App int `json:"app" bson:"app"`
}

type UserType struct {
	Portal int `json:"portal" bson:"portal"`
	User   int `json:"user"   bson:"user"`
}

type PayloadClaims struct {
	ID         uint   `json:"id"         bson:"id"`
	Permission string `json:"permission" bson:"permission"`
	Platform   int    `json:"platform"   bson:"platform"`
	UserType   int    `json:"user_type"  bson:"user_type"`
	jwt.RegisteredClaims
}

type Payload struct {
	ID         uint   `json:"id"         bson:"id"`
	Permission string `json:"permission" bson:"permission"`
	Platform   int    `json:"platform"   bson:"platform"`
	UserType   int    `json:"user_type"  bson:"user_type"`
	ExpiresAt  int64  `json:"exp"        bson:"exp"`
}

var platform = Platform{
	Web: 1,
	App: 2,
}

var userType = UserType{
	Portal: 1,
	User:   2,
}

var Token = token{
	Platform: platform,
	UserType: userType,
}

// CreateToken /* ----------------------------------------------------------- */
/*                                Create Token                                */
/* -------------------------------------------------------------------------- */
// payloadClaims := pkg.PayloadClaims{
// 	ID:   1,
// 	Permission: "1111111",
// }
// token, err := pkg.CreateToken(payloadClaims)
// if err != nil {
// 	fmt.Println("t error:", err)
// }
// fmt.Println("token:", token)
/* -------------------------------------------------------------------------- */
func (t *token) CreateToken(payloadClaims PayloadClaims) (string, error) {
	// Set the expiration time for the token
	expiresIn := time.Hour * 24
	expirationTime := time.Now().Add(expiresIn).Unix()

	// Convert the Unix timestamp to a time.Time value
	expTime := time.Unix(expirationTime, 0)

	// Set the claims for the token
	claims := &PayloadClaims{
		ID:         payloadClaims.ID,
		Permission: payloadClaims.Permission,
		Platform:   payloadClaims.Platform,
		UserType:   payloadClaims.UserType,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   strconv.Itoa(int(payloadClaims.ID)),
		},
	}

	// Create the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the specified key
	signingKey := []byte(config.AppConfig.JwtSigningKey)
	tokenString, err := token.SignedString(signingKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ParseToken /* ------------------------------------------------------------ */
/*                                 Parse Token                                */
/* -------------------------------------------------------------------------- */
// payload, err := pkg.ParseToken(token)
// if err != nil {
// 	fmt.Println("p error:", err)
// }
// fmt.Println("payload:", payload)
// fmt.Println("time:", time.Unix(payload.ExpiresAt, 0))
/* -------------------------------------------------------------------------- */
func (t *token) ParseToken(tokenString string) (*Payload, error) {
	token, err := jwt.ParseWithClaims(tokenString, &PayloadClaims{}, func(token *jwt.Token) (interface{}, error) {
		signingKey := []byte(config.AppConfig.JwtSigningKey)
		return signingKey, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*PayloadClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	payload := &Payload{
		ExpiresAt: claims.ExpiresAt.Time.Unix(),
	}

	// id, err := strconv.ParseUint(claims.Subject, 10, 32)
	// if err != nil {
	// 	return nil, err
	// }
	payload.ID = claims.ID
	payload.Permission = claims.Permission
	payload.Platform = claims.Platform
	payload.UserType = claims.UserType

	return payload, nil
}
