package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Pantry holds the schema definition for the Pantry entity.
type Pantry struct {
	ent.Schema
}

// Fields of the Pantry.
func (Pantry) Fields() []ent.Field {
	return []ent.Field{
		field.Int("quantity"),
	}
}

// Edges of the Pantry.
func (Pantry) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("ingredient", Ingredient.Type).Unique(),
	}
}
