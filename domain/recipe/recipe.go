package recipe

import "github.com/quii/go-fakes-and-contracts/domain/ingredients"

type Recipe struct {
	Name        string
	Ingredients []ingredients.Ingredient
}
