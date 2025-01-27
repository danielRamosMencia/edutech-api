package utils

import (
	"time"

	"github.com/danielRamosMencia/edutech-api/internal/constans"
	"github.com/danielRamosMencia/edutech-api/internal/models"
	"github.com/danielRamosMencia/edutech-api/internal/zap_logger"
	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
)

func GenerateJWT(data models.SessionData) (string, int64, error) {
	now := time.Now()
	exp := now.Add(time.Hour * time.Duration(constans.Envs.JwtTime)).Unix()
	iat := now.Unix()
	nbf := now.Unix()
	maxAge := exp - iat

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = exp
	claims["iat"] = iat
	claims["nbf"] = nbf
	claims["maxAge"] = maxAge
	claims["id"] = data.Id
	claims["user_name"] = data.Username
	claims["email"] = data.Email
	claims["active"] = data.Active
	claims["role_id"] = data.RoleId
	claims["role"] = data.Role

	signedToken, err := token.SignedString([]byte(constans.Envs.JwtSecret))
	if err != nil {
		zap_logger.Logger.Info("Error creating token", zap.Error(err))
		return "", 0, err
	}

	return signedToken, maxAge, nil
}
