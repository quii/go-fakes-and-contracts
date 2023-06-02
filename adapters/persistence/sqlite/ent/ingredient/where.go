// Code generated by ent, DO NOT EDIT.

package ingredient

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/quii/go-fakes-and-contracts/adapters/persistence/sqlite/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Ingredient {
	return predicate.Ingredient(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Ingredient {
	return predicate.Ingredient(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Ingredient {
	return predicate.Ingredient(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Ingredient {
	return predicate.Ingredient(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Ingredient {
	return predicate.Ingredient(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Ingredient {
	return predicate.Ingredient(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Ingredient {
	return predicate.Ingredient(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Ingredient {
	return predicate.Ingredient(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Ingredient {
	return predicate.Ingredient(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Ingredient {
	return predicate.Ingredient(sql.FieldEQ(FieldName, v))
}

// Vegan applies equality check predicate on the "vegan" field. It's identical to VeganEQ.
func Vegan(v bool) predicate.Ingredient {
	return predicate.Ingredient(sql.FieldEQ(FieldVegan, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Ingredient {
	return predicate.Ingredient(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Ingredient {
	return predicate.Ingredient(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Ingredient {
	return predicate.Ingredient(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Ingredient {
	return predicate.Ingredient(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Ingredient {
	return predicate.Ingredient(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Ingredient {
	return predicate.Ingredient(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Ingredient {
	return predicate.Ingredient(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Ingredient {
	return predicate.Ingredient(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Ingredient {
	return predicate.Ingredient(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Ingredient {
	return predicate.Ingredient(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Ingredient {
	return predicate.Ingredient(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Ingredient {
	return predicate.Ingredient(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Ingredient {
	return predicate.Ingredient(sql.FieldContainsFold(FieldName, v))
}

// VeganEQ applies the EQ predicate on the "vegan" field.
func VeganEQ(v bool) predicate.Ingredient {
	return predicate.Ingredient(sql.FieldEQ(FieldVegan, v))
}

// VeganNEQ applies the NEQ predicate on the "vegan" field.
func VeganNEQ(v bool) predicate.Ingredient {
	return predicate.Ingredient(sql.FieldNEQ(FieldVegan, v))
}

// HasPantry applies the HasEdge predicate on the "pantry" edge.
func HasPantry() predicate.Ingredient {
	return predicate.Ingredient(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, PantryTable, PantryColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPantryWith applies the HasEdge predicate on the "pantry" edge with a given conditions (other predicates).
func HasPantryWith(preds ...predicate.Pantry) predicate.Ingredient {
	return predicate.Ingredient(func(s *sql.Selector) {
		step := newPantryStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasRecipeingredient applies the HasEdge predicate on the "recipeingredient" edge.
func HasRecipeingredient() predicate.Ingredient {
	return predicate.Ingredient(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, RecipeingredientTable, RecipeingredientPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRecipeingredientWith applies the HasEdge predicate on the "recipeingredient" edge with a given conditions (other predicates).
func HasRecipeingredientWith(preds ...predicate.RecipeIngredient) predicate.Ingredient {
	return predicate.Ingredient(func(s *sql.Selector) {
		step := newRecipeingredientStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Ingredient) predicate.Ingredient {
	return predicate.Ingredient(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Ingredient) predicate.Ingredient {
	return predicate.Ingredient(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Ingredient) predicate.Ingredient {
	return predicate.Ingredient(func(s *sql.Selector) {
		p(s.Not())
	})
}
