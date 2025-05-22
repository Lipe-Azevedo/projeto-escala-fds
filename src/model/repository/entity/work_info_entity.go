package entity

type WorkInfoEntity struct {
	// UserID agora será mapeado para _id no MongoDB.
	// O tipo string é apropriado se o UserID do seu sistema de usuários é uma string (como um UUID ou o Hex de um ObjectId).
	UserID        string `bson:"_id"` // Alterado de "user_id" para "_id"
	Team          string `bson:"team"`
	Position      string `bson:"position"`
	DefaultShift  string `bson:"default_shift"`
	WeekdayOff    string `bson:"weekday_off"`
	WeekendDayOff string `bson:"weekend_day_off"`
	SuperiorID    string `bson:"superior_id"`
}
