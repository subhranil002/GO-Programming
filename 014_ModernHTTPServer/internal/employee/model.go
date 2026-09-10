package course

import "go.mongodb.org/mongo-driver/v2/bson"

type Employee struct {
	ID         bson.ObjectID `json:"id" bson:"_id,omitempty"`
	Name       string        `json:"name" bson:"name"`
	Email      string        `json:"email" bson:"email"`
	Phone      string        `json:"phone" bson:"phone"`
	Department string        `json:"department" bson:"department"`
	Salary     float64       `json:"salary" bson:"salary"`
}

type CreateEmployeeRequest struct {
	Name       string  `json:"name"`
	Email      string  `json:"email"`
	Phone      string  `json:"phone"`
	Department string  `json:"department"`
	Salary     float64 `json:"salary"`
}

type UpdateEmployeeRequest struct {
	Name       *string  `json:"name"`
	Email      *string  `json:"email"`
	Phone      *string  `json:"phone"`
	Department *string  `json:"department"`
	Salary     *float64 `json:"salary"`
}
