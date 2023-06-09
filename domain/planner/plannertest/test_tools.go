package plannertest

import (
	"github.com/google/uuid"
	"github.com/quii/go-fakes-and-contracts/domain/ingredients"
	"github.com/quii/go-fakes-and-contracts/domain/recipe"
	"math/rand"
)

func RandomRecipes() []recipe.Recipe {
	return []recipe.Recipe{
		RandomRecipe(),
		RandomRecipe(),
		RandomRecipe(),
	}
}

func RandomRecipe() recipe.Recipe {
	return recipe.Recipe{
		Name:        uuid.New().String(),
		Ingredients: RandomIngredients(),
	}
}

func RandomIngredients() []ingredients.Ingredient {
	return []ingredients.Ingredient{
		RandomIngredient(),
		RandomIngredient(),
		RandomIngredient(),
	}
}

func RandomIngredient() ingredients.Ingredient {
	return ingredients.Ingredient{
		Name:     uuid.New().String(),
		Quantity: uint(rand.Intn(10)) + 1,
	}
}
