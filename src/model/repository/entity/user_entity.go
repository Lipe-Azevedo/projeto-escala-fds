package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserEntity struct {
    ID       primitive.ObjectID `bson:"_id,omitempty"`
    Email    string             `bson:"email"`
    Password string             `bson:"password"`
    Name     string             `bson:"name"`
    UserType string             `bson:"user_type"`
    WorkInfo *WorkInfoEntity    `bson:"work_info,omitempty"`
}

type WorkInfoEntity struct {
    Team          string `bson:"team"`
    Position      string `bson:"position"`
    DefaultShift  string `bson:"default_shift"`
    WeekdayOff    string `bson:"weekday_off"`
    WeekendDayOff string `bson:"weekend_day_off"`
    SuperiorID    string `bson:"superior_id"`
}