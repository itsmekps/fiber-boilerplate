package dtos

import "go.mongodb.org/mongo-driver/bson/primitive"

// DTO for validating an integer ID
type GetUserByIDRequest struct {
    ID int `param:"id" validate:"required"`
}

// DTO for validating a MongoDB ObjectID
type GetUserByMongoID struct {
	ID primitive.ObjectID `param:"id" validate:"required"`
}

// type GetUserByMongoIDRequest struct {
// 	ID primitive.ObjectID `param:"id" validate:"required"` // Parameter is a MongoDB ObjectID
// }

// DTO for validating a string ID
type GetUserByStringIDRequest struct {
    ID string `param:"id" validate:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`   // Email must be provided and valid
	Password string `json:"password" validate:"required,min=6"` // Password must be at least 6 characters long
}
