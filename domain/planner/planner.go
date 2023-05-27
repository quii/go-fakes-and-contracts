package planner

import (
	"context"
	"github.com/quii/go-fakes-and-contracts/domain/ingredients"
	"github.com/quii/go-fakes-and-contracts/domain/recipe"
)

type Book interface {
	GetRecipes() []recipe.Recipe
}

type Planner struct {
	recipeBook      Book
	ingredientStore Store
}

func New(recipes Book, ingredientStore Store) *Planner {
	return &Planner{recipeBook: recipes, ingredientStore: ingredientStore}
}

func (p Planner) SuggestRecipes(ctx context.Context) ([]recipe.Recipe, error) {
	availableIngredients, err := p.ingredientStore.GetIngredients(ctx)
	if err != nil {
		return nil, err
	}

	var suggestions []recipe.Recipe
	for _, r := range p.recipeBook.GetRecipes() {
		if haveIngredients(availableIngredients, r) {
			suggestions = append(suggestions, r)
		}
	}
	return suggestions, nil
}

func haveIngredients(availableIngredients ingredients.Ingredients, recipe recipe.Recipe) bool {
	for _, ingredient := range recipe.Ingredients {
		if !availableIngredients.Has(ingredient) {
			return false
		}
	}
	return true
}
