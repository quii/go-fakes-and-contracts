// Code generated by ent, DO NOT EDIT.

package ent

import (
	"github.com/quii/go-fakes-and-contracts/adapters/persistence/sqlite/ent/ingredient"
	"github.com/quii/go-fakes-and-contracts/adapters/persistence/sqlite/ent/schema"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	ingredientFields := schema.Ingredient{}.Fields()
	_ = ingredientFields
	// ingredientDescVegan is the schema descriptor for vegan field.
	ingredientDescVegan := ingredientFields[1].Descriptor()
	// ingredient.DefaultVegan holds the default value on creation for the vegan field.
	ingredient.DefaultVegan = ingredientDescVegan.Default.(bool)
}
