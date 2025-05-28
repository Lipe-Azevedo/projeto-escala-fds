package comment

import (
	"context"
	"fmt"
	"os"

	"github.com/Lipe-Azevedo/escala-fds/src/configuration/logger"
	"github.com/Lipe-Azevedo/escala-fds/src/configuration/rest_err"

	// "github.com/Lipe-Azevedo/escala-fds/src/model/domain" // Não usado no stub
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.uber.org/zap"
)

func (cr *commentRepository) DeleteComment(commentID string) *rest_err.RestErr {
	logger.Info("Init DeleteComment repository", zap.String("commentID", commentID))

	collectionName := os.Getenv(MONGODB_COMMENTS_COLLECTION_ENV_KEY)
	collection := cr.databaseConnection.Collection(collectionName)

	objID, err := primitive.ObjectIDFromHex(commentID)
	if err != nil {
		return rest_err.NewBadRequestError("Invalid comment ID format")
	}

	result, err := collection.DeleteOne(context.Background(), bson.M{"_id": objID})
	if err != nil {
		logger.Error("Error deleting comment from repository", err, zap.String("commentID", commentID))
		return rest_err.NewInternalServerError("Error deleting comment")
	}

	if result.DeletedCount == 0 {
		return rest_err.NewNotFoundError(fmt.Sprintf("Comment with ID %s not found for deletion", commentID))
	}
	return nil
}
