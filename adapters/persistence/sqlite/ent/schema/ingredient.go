package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Ingredient holds the schema definition for the Ingredient entity.
type Ingredient struct {
	ent.Schema
}

// Fields of the Ingredient.
func (Ingredient) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.Uint("quantity"),
	}
}

// Edges of the Ingredient.
func (Ingredient) Edges() []ent.Edge {
	return nil
}
