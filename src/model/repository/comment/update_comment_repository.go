package comment

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/Lipe-Azevedo/escala-fds/src/configuration/logger"
	"github.com/Lipe-Azevedo/escala-fds/src/configuration/rest_err"
	"github.com/Lipe-Azevedo/escala-fds/src/model/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (cr *commentRepository) UpdateComment(commentID string, commentDomain domain.CommentDomainInterface) (domain.CommentDomainInterface, *rest_err.RestErr) {
	logger.Info("Init UpdateComment repository", zap.String("commentID", commentID))

	collectionName := os.Getenv(MONGODB_COMMENTS_COLLECTION_ENV_KEY)
	collection := cr.databaseConnection.Collection(collectionName)

	objID, err := primitive.ObjectIDFromHex(commentID)
	if err != nil {
		return nil, rest_err.NewBadRequestError("Invalid comment ID format")
	}

	// Atualiza apenas os campos permitidos (ex: Text, UpdatedAt)
	commentDomain.SetUpdatedAt(time.Now()) // Define o tempo de atualização
	updateData := bson.M{
		"$set": bson.M{
			"text":       commentDomain.GetText(),
			"updated_at": commentDomain.GetUpdatedAt(),
		},
	}

	result, err := collection.UpdateByID(context.Background(), objID, updateData)
	if err != nil {
		logger.Error("Error updating comment in repository", err, zap.String("commentID", commentID))
		return nil, rest_err.NewInternalServerError("Error updating comment")
	}

	if result.MatchedCount == 0 {
		return nil, rest_err.NewNotFoundError(fmt.Sprintf("Comment with ID %s not found for update", commentID))
	}
	// Para retornar o domínio atualizado, precisaríamos buscá-lo novamente ou assumir que commentDomain reflete o estado atualizado
	// Por simplicidade, retornaremos o commentDomain passado, que agora tem UpdatedAt preenchido.
	// O ID não muda.
	return commentDomain, nil
}
