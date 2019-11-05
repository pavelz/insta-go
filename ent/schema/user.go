package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int("age").
			Positive(),
		field.String("name").
			Default("unknown"),
		field.String("nick").
			Default("unknown"),
		field.String("nick2").
			Default("unknown"),
		field.Bytes("image").
			Optional().
			Default([]byte("code")),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
