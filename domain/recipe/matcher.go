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
	var suggestions []Recipe
	for _, recipe := range m.recipeBook.GetRecipes() {
		canMake, err := m.haveIngredients(ctx, recipe)
		if err != nil {
			return nil, err
		}
		if canMake {
			suggestions = append(suggestions, recipe)
		}
	}
	return suggestions, nil
}

func (m Matcher) haveIngredients(ctx context.Context, recipe Recipe) (bool, error) {
	availableIngredients, err := m.ingredientStore.GetIngredients(ctx)
	if err != nil {
		return false, err
	}

	for _, ingredient := range recipe.Ingredients {
		if !ingredients.Ingredients(availableIngredients).Has(ingredient) {
			return false, nil
		}
	}
	return true, nil
}
