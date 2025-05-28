package comment

import (
	"context"
	"fmt"
	"os"

	"github.com/Lipe-Azevedo/escala-fds/src/configuration/logger"
	"github.com/Lipe-Azevedo/escala-fds/src/configuration/rest_err"
	"github.com/Lipe-Azevedo/escala-fds/src/model/domain"
	commentconv "github.com/Lipe-Azevedo/escala-fds/src/model/repository/entity/converter/comment"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (cr *commentRepository) CreateComment(
	commentDomain domain.CommentDomainInterface,
) (domain.CommentDomainInterface, *rest_err.RestErr) {
	logger.Info("Init CreateComment repository",
		zap.String("journey", "createComment"),
		zap.String("collaboratorId", commentDomain.GetCollaboratorID()))

	collectionName := os.Getenv(MONGODB_COMMENTS_COLLECTION_ENV_KEY)
	if collectionName == "" {
		errMsg := fmt.Sprintf("Environment variable %s not set for comments collection", MONGODB_COMMENTS_COLLECTION_ENV_KEY)
		logger.Error(errMsg, nil, zap.String("journey", "createComment"))
		return nil, rest_err.NewInternalServerError("Database configuration error")
	}
	collection := cr.databaseConnection.Collection(collectionName)

	commentEntity := commentconv.ConvertCommentDomainToEntity(commentDomain)

	result, err := collection.InsertOne(context.Background(), commentEntity)
	if err != nil {
		logger.Error("Error trying to create comment in repository", err,
			zap.String("journey", "createComment"))
		return nil, rest_err.NewInternalServerError("Error creating comment in database")
	}

	commentID := result.InsertedID.(primitive.ObjectID)
	commentDomain.SetID(commentID.Hex()) // Atualiza o domínio com o ID gerado

	logger.Info("CreateComment repository executed successfully",
		zap.String("commentId", commentDomain.GetID()),
		zap.String("journey", "createComment"))

	return commentDomain, nil
}
