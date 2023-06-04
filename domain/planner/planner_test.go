package planner_test

import (
	"context"
	"github.com/alecthomas/assert/v2"
	"github.com/google/uuid"
	"github.com/quii/go-fakes-and-contracts/adapters/persistence/inmemory"
	"github.com/quii/go-fakes-and-contracts/adapters/persistence/sqlite"
	"github.com/quii/go-fakes-and-contracts/domain/ingredients"
	"github.com/quii/go-fakes-and-contracts/domain/planner"
	"github.com/quii/go-fakes-and-contracts/domain/recipe"
	"math/rand"
	"testing"
)

func TestRecipeMatcher(t *testing.T) {
	// for local, snappy integration test with a fake (which we can be confident is correct due to it conforming to the store contract)
	t.Run("with in memory store", func(t *testing.T) {
		RecipeMatcherTest{
			NewRecipeBook: func() planner.RecipeBook {
				return inmemory.NewRecipeStore()
			},
			NewPantry: func() planner.Pantry {
				return inmemory.NewPantry()
			},
		}.Test(t)
	})

	// we can run a broader integration test with a "real" db if we wish, using this contract approach
	t.Run("with sqlite", func(t *testing.T) {
		client := sqlite.NewSQLiteClient()
		t.Cleanup(func() {
			assert.NoError(t, client.Close())
		})

		if !testing.Short() {
			RecipeMatcherTest{
				NewRecipeBook: func() planner.RecipeBook {
					return sqlite.NewRecipeStore(client)
				},
				NewPantry: func() planner.Pantry {
					return sqlite.NewPantry(client)
				},
			}.Test(t)
		}
	})
}

type RecipeMatcherTest struct {
	NewPantry     func() planner.Pantry
	NewRecipeBook func() planner.RecipeBook
}

func (r RecipeMatcherTest) Test(t *testing.T) {
	t.Run("planning meals", func(t *testing.T) {

		t.Run("happy path, have ingredients for a recipe", func(t *testing.T) {
			ctx := context.Background()
			lasagna := randomRecipe()

			store := r.NewPantry()
			recipeBook := r.NewRecipeBook()

			assert.NoError(t, recipeBook.AddRecipes(ctx, lasagna))
			assert.NoError(t, store.Store(ctx, lasagna.Ingredients...))

			planner := planner.New(recipeBook, store)
			recipes, err := planner.SuggestRecipes(ctx)
			assert.NoError(t, err)
			assert.Equal(t, []recipe.Recipe{lasagna}, recipes)
		})

	})

	t.Run("suggesting recipes", func(t *testing.T) {

		t.Run("if don't have the ingredients for a meal, we cant make it", func(t *testing.T) {
			ctx := context.Background()
			store := r.NewPantry()
			recipeBook := r.NewRecipeBook()
			planner := planner.New(recipeBook, store)

			pie := randomRecipe()
			assert.NoError(t, recipeBook.AddRecipes(ctx, pie))

			recipes, err := planner.SuggestRecipes(ctx)
			assert.NoError(t, err)
			assertDoesntHaveRecipe(t, recipes, pie)
		})

		t.Run("if we have the ingredients for a recipe we can make it", func(t *testing.T) {
			ctx := context.Background()
			store := r.NewPantry()
			recipeBook := r.NewRecipeBook()
			planner := planner.New(recipeBook, store)

			bananaBread := randomRecipe()
			aRecipeWeWontHaveIngredientsFor := randomRecipe()
			assert.NoError(t, recipeBook.AddRecipes(ctx, bananaBread, aRecipeWeWontHaveIngredientsFor))

			assert.NoError(t, store.Store(
				ctx,
				bananaBread.Ingredients...,
			))

			recipes, err := planner.SuggestRecipes(ctx)
			assert.NoError(t, err)
			assertHasRecipe(t, recipes, bananaBread)
			assertDoesntHaveRecipe(t, recipes, aRecipeWeWontHaveIngredientsFor)
		})

		t.Run("if we have ingredients for 2 recipes, we can make both", func(t *testing.T) {
			ctx := context.Background()
			store := r.NewPantry()
			recipeBook := r.NewRecipeBook()
			planner := planner.New(recipeBook, store)

			bananaBread := randomRecipe()
			bananaMilkshake := randomRecipe()

			assert.NoError(t, recipeBook.AddRecipes(ctx, bananaBread, bananaMilkshake))

			assert.NoError(t, store.Store(
				ctx,
				append(bananaBread.Ingredients, bananaMilkshake.Ingredients...)...,
			))

			recipes, err := planner.SuggestRecipes(ctx)
			assert.NoError(t, err)
			assertHasRecipe(t, recipes, bananaBread)
			assertHasRecipe(t, recipes, bananaMilkshake)
		})
	})
}

func randomRecipe() recipe.Recipe {
	return recipe.Recipe{
		Name:        uuid.New().String(),
		Ingredients: random3ingredients(),
	}
}

func randomIngredient() ingredients.Ingredient {
	return ingredients.Ingredient{
		Name:     uuid.New().String(),
		Quantity: uint(rand.Intn(10)),
	}
}

func random3ingredients() []ingredients.Ingredient {
	return []ingredients.Ingredient{
		randomIngredient(),
		randomIngredient(),
		randomIngredient(),
	}
}

func assertHasRecipe(t *testing.T, recipes recipe.Recipes, expected recipe.Recipe) {
	t.Helper()
	_, found := recipes.FindByName(expected.Name)
	assert.True(t, found)
}

func assertDoesntHaveRecipe(t *testing.T, recipes recipe.Recipes, expected recipe.Recipe) {
	t.Helper()
	_, found := recipes.FindByName(expected.Name)
	assert.False(t, found)
}
