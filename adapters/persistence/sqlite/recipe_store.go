package sqlite

import (
	"context"
	"github.com/quii/go-fakes-and-contracts/domain/recipe"
)

type RecipeStore struct {
}

func (r RecipeStore) GetRecipes(ctx context.Context) ([]recipe.Recipe, error) {
	return nil, nil
}

func (r RecipeStore) AddRecipes(ctx context.Context, recipe ...recipe.Recipe) error {
	return nil
}

func (r RecipeStore) Close() {
}

func NewRecipeStore() *RecipeStore {
	return &RecipeStore{}
}
