package user

import (
	"net/http" // Import para http.StatusNotFound
	"os"
	"time"

	"github.com/Lipe-Azevedo/escala-fds/src/configuration/logger"
	"github.com/Lipe-Azevedo/escala-fds/src/configuration/rest_err"
	"github.com/Lipe-Azevedo/escala-fds/src/model/domain"
	"github.com/dgrijalva/jwt-go"
	"go.uber.org/zap"
)

var jwtSecretKey = []byte(os.Getenv("JWT_SECRET_KEY"))

func (uds *userDomainService) LoginUserServices(email, password string) (string, domain.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init LoginUserServices",
		zap.String("journey", "loginUser"),
		zap.String("email", email))

	if len(jwtSecretKey) == 0 {
		// CORREÇÃO APLICADA AQUI:
		errMsg := "CRITICAL: JWT_SECRET_KEY environment variable not set or empty. Service cannot operate."
		logger.Error(errMsg, nil, zap.String("journey", "loginUser")) // nil para o campo de erro Go padrão
		panic(errMsg)                                                 // Interrompe a execução, pois é uma falha crítica de configuração
	}

	if email == "" || password == "" {
		return "", nil, rest_err.NewBadRequestError("Email and password are required")
	}

	userDomain, serviceErr := uds.userRepository.FindUserByEmail(email)
	if serviceErr != nil {
		if serviceErr.Code == http.StatusNotFound { // Comparar pelo código HTTP
			logger.Warn("User not found by email for login",
				zap.String("email", email),
				zap.String("journey", "loginUser"))
			return "", nil, rest_err.NewUnauthorizedError("Invalid email or password")
		}
		// O serviceErr já é do tipo *rest_err.RestErr, que implementa a interface error.
		// Podemos passá-lo diretamente para logger.Error se quisermos usar o campo "error" do Zap.
		logger.Error("Error finding user by email for login", serviceErr, // Passando serviceErr aqui
			zap.String("email", email),
			zap.String("journey", "loginUser"))
		return "", nil, serviceErr
	}

	if !userDomain.CheckPassword(password) {
		logger.Warn("Invalid password attempt for login",
			zap.String("email", email),
			zap.String("journey", "loginUser"))
		return "", nil, rest_err.NewUnauthorizedError("Invalid email or password")
	}

	claims := jwt.MapClaims{
		"userID":   userDomain.GetID(),
		"email":    userDomain.GetEmail(),
		"userType": string(userDomain.GetUserType()),
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
		"iat":      time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, signErr := token.SignedString(jwtSecretKey)
	if signErr != nil {
		logger.Error("Error signing JWT token", signErr, zap.String("journey", "loginUser"))
		return "", nil, rest_err.NewInternalServerError("Error generating authentication token")
	}

	logger.Info("LoginUserServices executed successfully, token generated",
		zap.String("userId", userDomain.GetID()),
		zap.String("journey", "loginUser"))

	return tokenString, userDomain, nil
}
