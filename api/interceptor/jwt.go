package interceptor

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"time"

	"github.com/labstack/echo/v4"
)

type JWTManager struct {
	jwtSecret string
	jwtIssuer string
}

type JWTConfig struct {
	UserID int64 `json:"user_id"`
	jwt.RegisteredClaims
}

func NewJWTManager(secret, issuer string) (*JWTManager, error) {
	manager := &JWTManager{
		jwtSecret: secret,
		jwtIssuer: issuer,
	}
	return manager, nil
}

func (m *JWTManager) Issue(userID int64) (string, int64, error) {
	//lgr := utils.LoggerForMethod(m.Logger, "issueToken")
	//Issue new token
	now := time.Now()
	expiredAt := now.Add(time.Hour * 72)

	claims := JWTConfig{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{Time: expiredAt},
			IssuedAt:  &jwt.NumericDate{Time: now},
			Issuer:    m.jwtIssuer,
		},
	}
	jwtClaimToken := jwt.NewWithClaims(cfg.JWTClaimMethod, claims)
	token, err := jwtClaimToken.SignedString([]byte(m.jwtSecret))
	if err != nil {
		//lgr.Error("cannot issue token for user ", zap.Error(err))
		return "", 0, err
	}

	return token, expiredAt.Unix(), nil
}

func (m *JWTManager) GetClaims(c echo.Context) (*cfg.CustomClaims, error) {

	token, ok := c.Get("user").(*jwt.Token) // by default token is stored under `user` key
	if !ok {
		return nil, errors.New("JWT token missing or invalid")
	}

	claims := token.Claims.(*JWTConfig)
	return claims, nil
}

func (m *JWTManager) Skipper(c echo.Context) bool {

	return false
}
