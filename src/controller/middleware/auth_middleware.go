package middleware

import (
	"fmt"
	"os"
	"strings"

	"github.com/Lipe-Azevedo/escala-fds/src/configuration/logger"
	"github.com/Lipe-Azevedo/escala-fds/src/configuration/rest_err"
	"github.com/Lipe-Azevedo/escala-fds/src/model/domain" // Para domain.UserType
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// jwtSecretKey é lida da variável de ambiente.
var jwtSecretKey = []byte(os.Getenv("JWT_SECRET_KEY"))

// AuthMiddleware é o middleware para autenticação JWT
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if jwtSecretKey == nil || len(jwtSecretKey) == 0 {
			logger.Error("JWT_SECRET_KEY environment variable not set or empty in middleware.", nil, zap.String("journey", "authMiddleware"))
			// Este é um erro de configuração do servidor, não um erro do cliente.
			errResp := rest_err.NewInternalServerError("JWT secret key not configured on server")
			c.JSON(errResp.Code, errResp)
			c.Abort()
			return
		}

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			logger.Warn("Authorization header missing", zap.String("journey", "authMiddleware"))
			err := rest_err.NewUnauthorizedError("Authorization header is required")
			c.JSON(err.Code, err)
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") { // Usar EqualFold para "bearer" ou "Bearer"
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
			return jwtSecretKey, nil
		})

		if err != nil {
			logger.Error("Error parsing JWT token", err, zap.String("journey", "authMiddleware"))
			validationErr, ok := err.(*jwt.ValidationError)
			if ok {
				if validationErr.Errors&jwt.ValidationErrorMalformed != 0 {
					errRest := rest_err.NewUnauthorizedError("Malformed token")
					c.JSON(errRest.Code, errRest)
				} else if validationErr.Errors&jwt.ValidationErrorExpired != 0 {
					errRest := rest_err.NewUnauthorizedError("Token has expired")
					c.JSON(errRest.Code, errRest)
				} else if validationErr.Errors&jwt.ValidationErrorNotValidYet != 0 {
					errRest := rest_err.NewUnauthorizedError("Token not active yet")
					c.JSON(errRest.Code, errRest)
				} else {
					errRest := rest_err.NewUnauthorizedError("Invalid token")
					c.JSON(errRest.Code, errRest)
				}
			} else {
				// Erro genérico de parse
				errRest := rest_err.NewUnauthorizedError("Invalid token")
				c.JSON(errRest.Code, errRest)
			}
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			userID, okUserID := claims["userID"].(string)
			userTypeStr, okUserTypeStr := claims["userType"].(string)
			// userEmail, okUserEmail := claims["email"].(string) // Se você adicionou email ao token e precisa dele

			if !okUserID || !okUserTypeStr {
				logger.Error("Invalid token claims type or missing fields", nil,
					zap.String("journey", "authMiddleware"),
					zap.Any("claims", claims))
				errRest := rest_err.NewUnauthorizedError("Invalid token claims")
				c.JSON(errRest.Code, errRest)
				c.Abort()
				return
			}

			// Validar se userTypeStr é um UserType válido
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

			// Injetar informações do usuário no contexto do Gin
			c.Set("userID", userID)
			c.Set("userType", userType) // Agora é do tipo domain.UserType
			// if okUserEmail { c.Set("userEmail", userEmail) }

			logger.Info("JWT token validated successfully",
				zap.String("userID", userID),
				zap.String("userType", string(userType)),
				zap.String("journey", "authMiddleware"))
			c.Next()
		} else {
			logger.Warn("Invalid JWT token or claims (post-parse check failed)", zap.String("journey", "authMiddleware"))
			errRest := rest_err.NewUnauthorizedError("Invalid token")
			c.JSON(errRest.Code, errRest)
			c.Abort()
		}
	}
}
