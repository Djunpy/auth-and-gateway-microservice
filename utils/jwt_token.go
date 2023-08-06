package utils

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"strings"
	"time"
)

// jwt custom_errors 40-60
const (
	MISSING_TOKEN               = 40
	CREATE_TOKEN_ERR            = 41
	DECODE_PRIVATE_KEY_ERR      = 42
	PARSE_RSA_ERR               = 43
	TOKEN_UNEXPECTED_METHOD_ERR = 44
	VALIDATE_TOKEN_ERR          = 45
	INVALID_TOKEN               = 46
	TOKEN_TYPE_ERR              = 47
	TOKEN_EXPIRED               = 48
	TOKEN_NOT_STARTED           = 49
)

type JwtError struct {
	Code int
	Err  error
}

func (je JwtError) Error() string {
	return je.Err.Error()
}

func GetAccessToken(authHeader string) (string, error) {
	var accessToken string
	fields := strings.Fields(authHeader)
	if len(fields) != 0 && fields[0] == "Bearer" {
		accessToken = fields[1]
	}
	if accessToken == "" {
		return "", JwtError{Code: MISSING_TOKEN, Err: fmt.Errorf("access token is required")}
	}
	return accessToken, nil
}

func CreateToken(ttl time.Duration, payload interface{}, tokenType, privateKey string) (string, error) {
	// Создаем новый токен
	token := jwt.New(jwt.SigningMethodHS256)

	// Создаем заголовок токена
	token.Header["alg"] = "HS256"
	token.Header["typ"] = "JWT"

	now := time.Now().UTC()

	claims := token.Claims.(jwt.MapClaims)
	claims["token_type"] = tokenType
	claims["sub"] = payload
	claims["exp"] = now.Add(ttl).Unix()
	claims["iat"] = now.Unix()
	claims["nbf"] = now.Unix()

	signedToken, err := token.SignedString([]byte(privateKey))
	if err != nil {
		return "", JwtError{Code: CREATE_TOKEN_ERR, Err: err}
	}
	return signedToken, nil

}

func ValidateToken(token string, tokenType, privateKey string) (interface{}, error) {

	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, JwtError{Code: TOKEN_UNEXPECTED_METHOD_ERR, Err: fmt.Errorf("unexpected method: %s", t.Header["alg"])}
		}
		return []byte(privateKey), nil
	})

	if err != nil {
		return nil, JwtError{Code: VALIDATE_TOKEN_ERR, Err: err}
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok || !parsedToken.Valid {
		return nil, JwtError{Code: INVALID_TOKEN, Err: err}
	}

	tType := claims["token_type"]
	currentTime := time.Now().UTC().Unix()
	expirationTime := int64(claims["exp"].(float64))
	issueTime := int64(claims["iat"].(float64))

	if tType != tokenType {
		return nil, JwtError{Code: TOKEN_TYPE_ERR, Err: fmt.Errorf("does not match the token type %s", tokenType)}
	}

	if currentTime > expirationTime {
		return nil, JwtError{Code: TOKEN_EXPIRED, Err: fmt.Errorf("token expired")}
	}

	if currentTime < issueTime {
		return nil, JwtError{Code: TOKEN_NOT_STARTED, Err: fmt.Errorf("the token was sent before it started its action")}
	}
	return claims["sub"], nil
}
