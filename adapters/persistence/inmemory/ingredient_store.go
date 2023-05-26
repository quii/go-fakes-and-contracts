package inmemory

import (
	"context"
	"github.com/quii/go-fakes-and-contracts/domain/ingredients"
)

type IngredientStore struct {
	ingredients ingredients.Ingredients
}

func NewIngredientStore() *IngredientStore {
	return &IngredientStore{}
}

func (i *IngredientStore) GetIngredients(ctx context.Context) ([]ingredients.Ingredient, error) {
	return i.ingredients, nil
}

func (i *IngredientStore) Store(ctx context.Context, ingredients ...ingredients.Ingredient) error {
	for idx, ingredient := range ingredients {
		if i.ingredients.Has(ingredient) {
			i.ingredients[idx].Quantity += ingredient.Quantity
		} else {
			i.ingredients = append(i.ingredients, ingredient)
		}
	}
	return nil
}

func (i *IngredientStore) Close() {
}
