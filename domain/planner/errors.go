package planner

import (
	"fmt"
	"github.com/quii/go-fakes-and-contracts/domain/ingredients"
)

type ErrorMissingIngredients struct {
	MissingIngredients ingredients.Ingredients
}

func (e ErrorMissingIngredients) Error() string {
	return fmt.Sprintf("missing ingredients: %v", e.MissingIngredients)
}
