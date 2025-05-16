package repository

import (
    "context"
    "os"

    "github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/logger"
    "github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/rest_err"
    "github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.uber.org/zap"
)

func (ur *userRepository) UpdateWorkInfo(
    userId string,
    workInfo model.WorkInfoInterface,
) *rest_err.RestErr {
    logger.Info("Init updateWorkInfo repository",
        zap.String("journey", "updateWorkInfo"))

    collection_name := os.Getenv(MONGODB_USER_DB)
    collection := ur.dataBaseConnection.Collection(collection_name)

    userIdHex, err := primitive.ObjectIDFromHex(userId)
    if err != nil {
        return rest_err.NewBadRequestError("Invalid user ID")
    }

    update := bson.M{
        "$set": bson.M{
            "work_info": bson.M{
                "team":           workInfo.GetTeam(),
                "position":      workInfo.GetPosition(),
                "default_shift": workInfo.GetDefaultShift(),
                "weekday_off":   workInfo.GetWeekdayOff(),
                "weekend_day_off": workInfo.GetWeekendDayOff(),
                "superior_id":    workInfo.GetSuperiorID(),
            },
        },
    }

    _, err = collection.UpdateByID(
        context.Background(),
        userIdHex,
        update,
    )
    if err != nil {
        logger.Error("Error trying to update work info",
            err,
            zap.String("journey", "updateWorkInfo"))
        return rest_err.NewInternalServerError(err.Error())
    }

    logger.Info("WorkInfo updated successfully",
        zap.String("userId", userId),
        zap.String("journey", "updateWorkInfo"))
    return nil
}