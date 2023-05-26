package inmemory

import "github.com/quii/go-fakes-and-contracts/domain/recipe"

type RecipeStore struct {
	Recipes []recipe.Recipe
}

func (s RecipeStore) GetRecipes() []recipe.Recipe {
	return s.Recipes
}
