package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	// ID       int    `json:"id"`
	// MongoID  string `json:"_id"`
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"` // MongoDB ObjectID
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `jsom:"password"`
}

// Error implements error.
// func (u *User) Error() string {
// 	// panic("unimplemented")
// 	return fmt.Sprintf("User error: %v", u.Email)
// }
