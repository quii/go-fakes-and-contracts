package inmemory

import (
	"context"
	"github.com/quii/go-fakes-and-contracts/domain/ingredients"
)

type IngredientStore struct {
	ingredients []ingredients.Ingredient
}

func (i *IngredientStore) GetIngredients(ctx context.Context) ([]ingredients.Ingredient, error) {
	return i.ingredients, nil
}

func (i *IngredientStore) Store(ctx context.Context, ingredients ...ingredients.Ingredient) error {
	i.ingredients = append(i.ingredients, ingredients...)
	return nil
}
