package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Ingredient holds the schema definition for the Ingredient entity.
type Ingredient struct {
	ent.Schema
}

// Fields of the Ingredient.
func (Ingredient) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Unique(),
		field.Bool("vegan").Default(false),
	}
}

// Edges of the Ingredient.
func (Ingredient) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("pantry", Pantry.Type).Ref("ingredient").Unique(),
		edge.From("recipeingredient", RecipeIngredient.Type).
			Ref("ingredient").Unique(),
	}
}
