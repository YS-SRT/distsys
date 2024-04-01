package utils

import (
	"distsys/base"
	"distsys/global"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
)

func BuildClaims(uid int, uname string) jwt.MapClaims {

	claims := jwt.MapClaims{
		"sub": uid,
		"iss": "distsys",
	}

	return claims
}

func GenerateAccessToken(claims jwt.MapClaims) (string, *base.CustomError) {
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(global.GCONFIG.JWT.AccessExpiredMinutes)).UnixMilli()
	accessKey := ([]byte(global.GCONFIG.JWT.AccessKey))
	return generateToken(claims, accessKey)
}

func GenerateRefrashToken(claims jwt.MapClaims) (string, *base.CustomError) {
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(global.GCONFIG.JWT.RefreshExpiredMinutes)).UnixMilli()
	refreshKey := ([]byte(global.GCONFIG.JWT.RefreshKey))
	return generateToken(claims, refreshKey)
}

func ParseAccessToken(tokenString string) (jwt.MapClaims, *base.CustomError) {
	accessKey := ([]byte(global.GCONFIG.JWT.AccessKey))
	return parseToken(tokenString, accessKey)
}

func ParseRefrashToken(tokenString string) (jwt.MapClaims, *base.CustomError) {

	refreshKey := ([]byte(global.GCONFIG.JWT.RefreshKey))
	return parseToken(tokenString, refreshKey)
}

func parseToken(tokenString string, key []byte) (jwt.MapClaims, *base.CustomError) {

	if token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {

		return key, nil

	}, jwt.WithIssuer("distsys"), jwt.WithJSONNumber()); err != nil {

		zap.L().Error("parse token error", zap.Error(err))
		return nil, base.InternalCustomError(err)

	} else {

		if token.Valid {

			return token.Claims.(jwt.MapClaims), nil

		} else {

			zap.L().Sugar().Info("token already expired", token.Claims)
			return nil, base.BuildCustomError(base.CodeTokenExpired, "token expired")
		}

	}

}

func generateToken(claims jwt.Claims, key []byte) (string, *base.CustomError) {

	if token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(key); err != nil {
		return "", base.InternalCustomError(err)
	} else {
		return token, nil
	}
}
