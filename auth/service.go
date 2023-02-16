package auth

import (
	"errors"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Service interface {
	GenerateToken(userID int) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

type MyClaims struct {
	UserID     int `json:"user_id"`
	Created_at int `json:"created_at"`
	jwt.StandardClaims
}

type jwtService struct {
}

type Claims struct {
	UserID string
	jwt.StandardClaims
}

func AuthSecretKey() []byte {

	var AUTH_SECRETKEY = []byte("" + os.Getenv("AUTHSECRETKEY"))

	return AUTH_SECRETKEY
}

func EmailSecretKey() []byte {
	var EMAIL_SECRETKEY = []byte(os.Getenv("" + os.Getenv("EMAILSECRETKEY")))
	return EMAIL_SECRETKEY
}

func NewService() *jwtService {
	return &jwtService{}
}

func (s *jwtService) GenerateToken(userID int) (string, error) {
	AUTH_SECRETKEY := AuthSecretKey()
	// expirationTime := time.Now().Add(24 * time.Hour)
	claim := &MyClaims{
		UserID:     userID,
		Created_at: int(time.Now().Unix()),
		// StandardClaims: jwt.StandardClaims{
		// 	ExpiresAt: expirationTime.Unix(),
		// },
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	signedToken, err := token.SignedString(AUTH_SECRETKEY)
	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}
func (s *jwtService) ValidateToken(encodedToken string) (*jwt.Token, error) {
	AUTH_SECRETKEY := AuthSecretKey()
	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("invalid token")
		}

		return []byte(AUTH_SECRETKEY), nil
	})

	if err != nil {
		return token, err
	}

	return token, nil
}

// GenerateNonAuthToken handles generation of a jwt code
// @returns string -> token and error -> err
func GenerateNonAuthToken(userID string) (string, error) {
	EMAIL_SECRETKEY := EmailSecretKey()
	// Define token expiration time
	expirationTime := time.Now().Add(1 * time.Hour)
	// Define the payload and exp time
	claims := &Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Generate token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign token with secret key encoding
	tokenString, err := token.SignedString(EMAIL_SECRETKEY)

	return tokenString, err
}

// DecodeNonAuthToken handles decoding a jwt token
func DecodeNonAuthToken(tkStr string) (string, error) {
	EMAIL_SECRETKEY := EmailSecretKey()
	claims := &Claims{}

	// Decode token based on parameters provided, if it fails throw err
	tkn, err := jwt.ParseWithClaims(tkStr, claims, func(token *jwt.Token) (interface{}, error) {
		return EMAIL_SECRETKEY, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return "", err
		}
		return "", err
	}

	if !tkn.Valid {
		return "", err
	}

	// Return encoded email
	return claims.UserID, nil
}

//
