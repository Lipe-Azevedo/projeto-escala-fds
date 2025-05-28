package comment

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/Lipe-Azevedo/escala-fds/src/configuration/logger"
	"github.com/Lipe-Azevedo/escala-fds/src/configuration/rest_err"
	"github.com/Lipe-Azevedo/escala-fds/src/model/domain"
	"github.com/Lipe-Azevedo/escala-fds/src/model/repository/entity"
	commentconv "github.com/Lipe-Azevedo/escala-fds/src/model/repository/entity/converter/comment"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

func (cr *commentRepository) FindCommentByID(commentID string) (domain.CommentDomainInterface, *rest_err.RestErr) {
	logger.Info("Init FindCommentByID repository", zap.String("commentID", commentID))

	collectionName := os.Getenv(MONGODB_COMMENTS_COLLECTION_ENV_KEY)
	collection := cr.databaseConnection.Collection(collectionName)

	objID, err := primitive.ObjectIDFromHex(commentID)
	if err != nil {
		return nil, rest_err.NewBadRequestError("Invalid comment ID format")
	}

	var commentEntity entity.CommentEntity
	filter := bson.M{"_id": objID}
	err = collection.FindOne(context.Background(), filter).Decode(&commentEntity)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, rest_err.NewNotFoundError(fmt.Sprintf("Comment with ID %s not found", commentID))
		}
		logger.Error("Error finding comment by ID", err, zap.String("commentID", commentID))
		return nil, rest_err.NewInternalServerError("Error finding comment by ID")
	}
	return commentconv.ConvertCommentEntityToDomain(commentEntity), nil
}

func (cr *commentRepository) FindCommentsByCollaboratorAndDate(collaboratorID string, date time.Time) ([]domain.CommentDomainInterface, *rest_err.RestErr) {
	logger.Info("Init FindCommentsByCollaboratorAndDate repository",
		zap.String("collaboratorID", collaboratorID), zap.Time("date", date))

	collectionName := os.Getenv(MONGODB_COMMENTS_COLLECTION_ENV_KEY)
	collection := cr.databaseConnection.Collection(collectionName)

	// Normalizar a data para o início do dia para correspondência correta
	startOfDay := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())
	endOfDay := startOfDay.Add(24 * time.Hour).Add(-time.Nanosecond) // Fim do dia

	filter := bson.M{
		"collaborator_id": collaboratorID,
		"date": bson.M{
			"$gte": startOfDay,
			"$lt":  endOfDay,
		},
	}

	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		logger.Error("Error finding comments by collaborator and date", err)
		return nil, rest_err.NewInternalServerError("Error fetching comments")
	}
	defer cursor.Close(context.Background())

	var comments []domain.CommentDomainInterface
	for cursor.Next(context.Background()) {
		var commentEntity entity.CommentEntity
		if err := cursor.Decode(&commentEntity); err != nil {
			logger.Error("Error decoding comment entity", err)
			continue // Pula este e tenta o próximo
		}
		comments = append(comments, commentconv.ConvertCommentEntityToDomain(commentEntity))
	}
	return comments, nil
}

func (cr *commentRepository) FindCommentsByCollaboratorAndDateRange(collaboratorID string, startDate time.Time, endDate time.Time) ([]domain.CommentDomainInterface, *rest_err.RestErr) {
	logger.Info("Init FindCommentsByCollaboratorAndDateRange repository")
	// Normalizar datas
	normStartDate := time.Date(startDate.Year(), startDate.Month(), startDate.Day(), 0, 0, 0, 0, startDate.Location())
	normEndDate := time.Date(endDate.Year(), endDate.Month(), endDate.Day(), 23, 59, 59, 999999999, endDate.Location())

	collectionName := os.Getenv(MONGODB_COMMENTS_COLLECTION_ENV_KEY)
	collection := cr.databaseConnection.Collection(collectionName)

	filter := bson.M{
		"collaborator_id": collaboratorID,
		"date": bson.M{
			"$gte": normStartDate,
			"$lte": normEndDate,
		},
	}
	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		logger.Error("Error finding comments by collaborator and date range", err)
		return nil, rest_err.NewInternalServerError("Error fetching comments for date range")
	}
	defer cursor.Close(context.Background())

	var comments []domain.CommentDomainInterface
	for cursor.Next(context.Background()) {
		var commentEntity entity.CommentEntity
		if err := cursor.Decode(&commentEntity); err != nil {
			logger.Error("Error decoding comment entity in range", err)
			continue
		}
		comments = append(comments, commentconv.ConvertCommentEntityToDomain(commentEntity))
	}
	return comments, nil
}
