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
		Ingredients: random3ingredients(),
	}
}

func randomIngredient() ingredients.Ingredient {
	return ingredients.Ingredient{
		Name:     uuid.New().String(),
		Quantity: uint(rand.Intn(10)) + 1,
	}
}

func random3ingredients() []ingredients.Ingredient {
	return []ingredients.Ingredient{
		randomIngredient(),
		randomIngredient(),
		randomIngredient(),
	}
}
