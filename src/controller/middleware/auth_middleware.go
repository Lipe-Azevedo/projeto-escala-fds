package middleware

import (
	"fmt"
	"os"
	"strings"

	"github.com/Lipe-Azevedo/escala-fds/src/configuration/logger"
	"github.com/Lipe-Azevedo/escala-fds/src/configuration/rest_err"
	"github.com/Lipe-Azevedo/escala-fds/src/model/domain"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Removida a variável de pacote: var jwtAuthSecretKey = []byte(os.Getenv("JWT_SECRET_KEY"))

// AuthMiddleware é o middleware para autenticação JWT
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Busca a chave JWT_SECRET_KEY do ambiente dinamicamente a cada requisição
		// Isso garante que ela seja lida após o godotenv.Load() em main.go
		jwtSecretFromEnv := os.Getenv("JWT_SECRET_KEY")
		if jwtSecretFromEnv == "" {
			errMsg := "CRITICAL: JWT_SECRET_KEY environment variable not set or empty for middleware. Authentication cannot proceed."
			logger.Error(errMsg, nil, zap.String("journey", "authMiddleware"))
			// Retorna um erro interno do servidor, pois é uma falha de configuração do servidor.
			errResp := rest_err.NewInternalServerError("Server authentication mechanism is misconfigured.")
			c.JSON(errResp.Code, errResp)
			c.Abort()
			return
		}
		currentJwtSecretKey := []byte(jwtSecretFromEnv)

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			logger.Warn("Authorization header missing", zap.String("journey", "authMiddleware"))
			err := rest_err.NewUnauthorizedError("Authorization header is required")
			c.JSON(err.Code, err)
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
			logger.Warn("Invalid Authorization header format", zap.String("journey", "authMiddleware"), zap.String("header", authHeader))
			err := rest_err.NewUnauthorizedError("Authorization header format must be Bearer {token}")
			c.JSON(err.Code, err)
			c.Abort()
			return
		}

		tokenString := parts[1]
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return currentJwtSecretKey, nil // Usa a chave obtida dinamicamente
		})

		if err != nil {
			validationErr, ok := err.(*jwt.ValidationError)
			if ok {
				if validationErr.Errors&jwt.ValidationErrorMalformed != 0 {
					logger.Warn("Malformed token received", zap.String("token", tokenString), zap.String("journey", "authMiddleware"))
					errRest := rest_err.NewUnauthorizedError("Malformed token")
					c.JSON(errRest.Code, errRest)
				} else if validationErr.Errors&jwt.ValidationErrorExpired != 0 {
					logger.Info("Expired token received", zap.String("journey", "authMiddleware"))
					errRest := rest_err.NewUnauthorizedError("Token has expired")
					c.JSON(errRest.Code, errRest)
				} else if validationErr.Errors&jwt.ValidationErrorNotValidYet != 0 {
					logger.Warn("Token not active yet", zap.String("journey", "authMiddleware"))
					errRest := rest_err.NewUnauthorizedError("Token not active yet")
					c.JSON(errRest.Code, errRest)
				} else {
					logger.Warn("Invalid token (validation error)", zap.Error(err), zap.String("journey", "authMiddleware"))
					errRest := rest_err.NewUnauthorizedError("Invalid token")
					c.JSON(errRest.Code, errRest)
				}
			} else {
				logger.Error("Error parsing JWT token (non-validation error)", err, zap.String("journey", "authMiddleware"))
				errRest := rest_err.NewUnauthorizedError("Invalid token")
				c.JSON(errRest.Code, errRest)
			}
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			userID, okUserID := claims["userID"].(string)
			userTypeStr, okUserTypeStr := claims["userType"].(string)

			if !okUserID || !okUserTypeStr {
				logger.Error("Invalid token claims type or missing fields", nil,
					zap.String("journey", "authMiddleware"),
					zap.Any("claims", claims))
				errRest := rest_err.NewUnauthorizedError("Invalid token claims")
				c.JSON(errRest.Code, errRest)
				c.Abort()
				return
			}

			userType := domain.UserType(userTypeStr)
			if userType != domain.UserTypeCollaborator && userType != domain.UserTypeMaster {
				logger.Error("Invalid userType in token claims", nil,
					zap.String("journey", "authMiddleware"),
					zap.String("receivedUserType", userTypeStr))
				errRest := rest_err.NewForbiddenError("Invalid user type in token")
				c.JSON(errRest.Code, errRest)
				c.Abort()
				return
			}

			c.Set("userID", userID)
			c.Set("userType", userType)

			logger.Info("JWT token validated successfully",
				zap.String("userID", userID),
				zap.String("userType", string(userType)),
				zap.String("journey", "authMiddleware"))
			c.Next()
		} else {
			logger.Warn("Invalid JWT token or claims (post-parse check failed, token.Valid is false or claims type assertion failed)",
				zap.Bool("token.Valid", token.Valid),
				zap.Any("claims", token.Claims),
				zap.String("journey", "authMiddleware"))
			errRest := rest_err.NewUnauthorizedError("Invalid token")
			c.JSON(errRest.Code, errRest)
			c.Abort()
		}
	}
}
