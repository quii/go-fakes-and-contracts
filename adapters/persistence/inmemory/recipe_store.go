package inmemory

import (
	"context"
	"github.com/quii/go-fakes-and-contracts/domain/recipe"
)

type RecipeStore struct {
	Recipes []recipe.Recipe
}

func NewRecipeStore() *RecipeStore {
	return &RecipeStore{}
}

func (s *RecipeStore) GetRecipes(_ context.Context) ([]recipe.Recipe, error) {
	return s.Recipes, nil
}

func (s *RecipeStore) AddRecipes(_ context.Context, r ...recipe.Recipe) error {
	s.Recipes = append(s.Recipes, r...)
	return nil
}
