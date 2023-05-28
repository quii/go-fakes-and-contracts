package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// RecipeIngredient holds the schema definition for the RecipeIngredient entity.
type RecipeIngredient struct {
	ent.Schema
}

// Fields of the RecipeIngredient.
func (RecipeIngredient) Fields() []ent.Field {
	return []ent.Field{
		field.Int("quantity"),
	}
}

// Edges of the RecipeIngredient.
func (RecipeIngredient) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("ingredient", Ingredient.Type),
	}
}
