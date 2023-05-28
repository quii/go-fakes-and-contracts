// Code generated by ent, DO NOT EDIT.

package recipeingredient

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the recipeingredient type in the database.
	Label = "recipe_ingredient"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldQuantity holds the string denoting the quantity field in the database.
	FieldQuantity = "quantity"
	// EdgeIngredient holds the string denoting the ingredient edge name in mutations.
	EdgeIngredient = "ingredient"
	// Table holds the table name of the recipeingredient in the database.
	Table = "recipe_ingredients"
	// IngredientTable is the table that holds the ingredient relation/edge.
	IngredientTable = "ingredients"
	// IngredientInverseTable is the table name for the Ingredient entity.
	// It exists in this package in order to avoid circular dependency with the "ingredient" package.
	IngredientInverseTable = "ingredients"
	// IngredientColumn is the table column denoting the ingredient relation/edge.
	IngredientColumn = "recipe_ingredient_ingredient"
)

// Columns holds all SQL columns for recipeingredient fields.
var Columns = []string{
	FieldID,
	FieldQuantity,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "recipe_ingredients"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"recipe_recipeingredient",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

// OrderOption defines the ordering options for the RecipeIngredient queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByQuantity orders the results by the quantity field.
func ByQuantity(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldQuantity, opts...).ToFunc()
}

// ByIngredientCount orders the results by ingredient count.
func ByIngredientCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newIngredientStep(), opts...)
	}
}

// ByIngredient orders the results by ingredient terms.
func ByIngredient(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newIngredientStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newIngredientStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(IngredientInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, IngredientTable, IngredientColumn),
	)
}
