package user

import (
	"net/http"
	"time"

	// "os" // Não precisamos mais de os.Getenv diretamente aqui para a chave

	"github.com/Lipe-Azevedo/escala-fds/src/configuration/logger"
	"github.com/Lipe-Azevedo/escala-fds/src/configuration/rest_err"
	"github.com/Lipe-Azevedo/escala-fds/src/model/domain"
	"github.com/dgrijalva/jwt-go"
	"go.uber.org/zap"
)

// Removida a variável de pacote: var jwtSecretKey = []byte(os.Getenv("JWT_SECRET_KEY"))

func (uds *userDomainService) LoginUserServices(email, password string) (string, domain.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init LoginUserServices",
		zap.String("journey", "loginUser"),
		zap.String("email", email))

	// A verificação da chave agora é feita implicitamente pela presença de uds.jwtSecret.
	// Se uds.jwtSecret estivesse vazio, o construtor NewUserDomainService já teria causado um panic.
	// Não precisamos mais do panic "CRITICAL: JWT_SECRET_KEY environment variable not set or empty." aqui.

	if email == "" || password == "" {
		return "", nil, rest_err.NewBadRequestError("Email and password are required")
	}

	userDomain, serviceErr := uds.userRepository.FindUserByEmail(email)
	if serviceErr != nil {
		if serviceErr.Code == http.StatusNotFound {
			logger.Warn("User not found by email for login",
				zap.String("email", email),
				zap.String("journey", "loginUser"))
			return "", nil, rest_err.NewUnauthorizedError("Invalid email or password")
		}
		logger.Error("Error finding user by email for login", serviceErr,
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
	// Usa a chave secreta armazenada na struct do serviço
	jwtSecretBytes := []byte(uds.jwtSecret)
	tokenString, signErr := token.SignedString(jwtSecretBytes)
	if signErr != nil {
		logger.Error("Error signing JWT token", signErr, zap.String("journey", "loginUser"))
		return "", nil, rest_err.NewInternalServerError("Error generating authentication token")
	}

	logger.Info("LoginUserServices executed successfully, token generated",
		zap.String("userId", userDomain.GetID()),
		zap.String("journey", "loginUser"))

	return tokenString, userDomain, nil
}
