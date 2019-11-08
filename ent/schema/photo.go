package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
)

// Photo holds the schema definition for the Photo entity.
type Photo struct {
	ent.Schema
}

// Fields of the Photo.
func (Photo) Fields() []ent.Field {
	return []ent.Field{
		field.Float("lat").
			Default(0.0),
		field.Float("lng").
			Default(0.0),
		field.Bytes("image").
			Optional(),
		field.String("fielname"),
	}
}

// Edges of the Photo.
func (Photo) Edges() []ent.Edge {
	return nil
}
