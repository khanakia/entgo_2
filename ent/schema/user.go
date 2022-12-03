package schema

import (
	"os/user"
)

type User struct {
	user.User
}

// User holds the schema definition for the User entity.
// type User struct {
// 	ent.Schema
// }

// // Fields of the User.
// func (User) Fields() []ent.Field {
// 	return []ent.Field{
// 		field.Int("age").
// 			Positive(),
// 		field.String("name").
// 			Default("unknown"),
// 	}
// }
