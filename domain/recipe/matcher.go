package recipe

import (
	"context"
	"github.com/quii/go-fakes-and-contracts/domain/ingredients"
)

type Book interface {
	GetRecipes() []Recipe
}

type Matcher struct {
	recipeBook      Book
	ingredientStore Store
}

func NewMatcher(recipes Book, ingredientStore Store) *Matcher {
	return &Matcher{recipeBook: recipes, ingredientStore: ingredientStore}
}

func (m Matcher) SuggestRecipes(ctx context.Context) ([]Recipe, error) {
	availableIngredients, err := m.ingredientStore.GetIngredients(ctx)
	if err != nil {
		return nil, err
	}

	var suggestions []Recipe
	for _, recipe := range m.recipeBook.GetRecipes() {
		if haveIngredients(availableIngredients, recipe) {
			suggestions = append(suggestions, recipe)
		}
	}
	return suggestions, nil
}

func haveIngredients(availableIngredients ingredients.Ingredients, recipe Recipe) bool {
	for _, ingredient := range recipe.Ingredients {
		if !availableIngredients.Has(ingredient) {
			return false
		}
	}
	return true
}
