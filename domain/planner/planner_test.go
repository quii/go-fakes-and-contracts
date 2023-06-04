package planner_test

import (
	"context"
	"github.com/alecthomas/assert/v2"
	"github.com/google/uuid"
	"github.com/quii/go-fakes-and-contracts/adapters/persistence/inmemory"
	"github.com/quii/go-fakes-and-contracts/adapters/persistence/sqlite"
	"github.com/quii/go-fakes-and-contracts/domain/ingredients"
	planner "github.com/quii/go-fakes-and-contracts/domain/planner"
	"github.com/quii/go-fakes-and-contracts/domain/recipe"
	"math/rand"
	"testing"
	"time"
)

func TestRecipeMatcher(t *testing.T) {
	// for local, snappy integration test with a fake (which we can be confident is correct due to it conforming to the store contract)
	t.Run("with in memory store", func(t *testing.T) {
		RecipeMatcherTest{
			CreateDependencies: func() (planner.RecipeBook, planner.Pantry, Cleanup) {
				return inmemory.NewRecipeStore(), inmemory.NewPantry(), func() {
					// nothing to clean up
				}
			},
		}.Test(t)
	})

	// we can run a broader integration test with a "real" db if we wish, using this contract approach
	t.Run("with sqlite", func(t *testing.T) {
		if !testing.Short() {
			RecipeMatcherTest{
				CreateDependencies: func() (planner.RecipeBook, planner.Pantry, Cleanup) {
					client := sqlite.NewSQLiteClient()
					return sqlite.NewRecipeStore(client), sqlite.NewPantry(client), func() {
						if err := client.Close(); err != nil {
							t.Error(err)
						}
					}
				},
			}.Test(t)
		}
	})
}

type Cleanup func()

type RecipeMatcherTest struct {
	CreateDependencies func() (planner.RecipeBook, planner.Pantry, Cleanup)
}

func (r RecipeMatcherTest) Test(t *testing.T) {
	t.Run("planning meals", func(t *testing.T) {

		t.Run("happy path, have ingredients for a recipe, schedule it, update pantry", func(t *testing.T) {
			var (
				ctx                          = context.Background()
				lasagna                      = randomRecipe()
				recipeBook, pantry, teardown = r.CreateDependencies()
				sut                          = planner.New(recipeBook, pantry)
			)
			t.Cleanup(teardown)

			assert.NoError(t, recipeBook.AddRecipes(ctx, lasagna))
			assert.NoError(t, pantry.Store(ctx, lasagna.Ingredients...))

			assert.NoError(t, sut.ScheduleMeal(ctx, lasagna, time.Now()))
			remainingIngredients, err := pantry.GetIngredients(ctx)
			assert.NoError(t, err)
			assert.Equal(t, 0, len(remainingIngredients))
		})

		t.Run("returns a missing ingredients error if you try to schedule a meal without all the ingredients", func(t *testing.T) {
			var (
				ctx                         = context.Background()
				lasagna                     = randomRecipe()
				recipeBook, store, teardown = r.CreateDependencies()
				sut                         = planner.New(recipeBook, store)
			)
			t.Cleanup(teardown)

			assert.NoError(t, recipeBook.AddRecipes(ctx, lasagna))

			err := sut.ScheduleMeal(ctx, lasagna, time.Now())
			assert.Error(t, err)

			missingIngredientsErr, ok := err.(planner.ErrorMissingIngredients)
			assert.True(t, ok)
			assert.Equal(t, planner.ErrorMissingIngredients{
				MissingIngredients: lasagna.Ingredients,
			}, missingIngredientsErr)
		})

		t.Run("returns the specific ingredients missing if you try to schedule a meal with some missing ingredients", func(t *testing.T) {
			var (
				ctx                          = context.Background()
				recipeBook, pantry, teardown = r.CreateDependencies()
				lasagna                      = randomRecipe()
				sut                          = planner.New(recipeBook, pantry)
			)
			t.Cleanup(teardown)

			assert.NoError(t, recipeBook.AddRecipes(ctx, lasagna))

			missingIngredient, ingredientsWeHave := lasagna.Ingredients[0], lasagna.Ingredients[1:]
			assert.NoError(t, pantry.Store(ctx, ingredientsWeHave...))

			err := sut.ScheduleMeal(ctx, lasagna, time.Now())
			assert.Error(t, err)

			missingIngredientsErr, ok := err.(planner.ErrorMissingIngredients)
			assert.True(t, ok)
			assert.Equal(t, planner.ErrorMissingIngredients{
				MissingIngredients: []ingredients.Ingredient{missingIngredient},
			}, missingIngredientsErr)
		})

	})

	t.Run("suggesting recipes", func(t *testing.T) {

		t.Run("if don't have the ingredients for a meal, we cant make it", func(t *testing.T) {
			var (
				ctx                          = context.Background()
				pie                          = randomRecipe()
				recipeBook, pantry, teardown = r.CreateDependencies()
				sut                          = planner.New(recipeBook, pantry)
			)
			t.Cleanup(teardown)

			assert.NoError(t, recipeBook.AddRecipes(ctx, pie))

			recipes, err := sut.SuggestRecipes(ctx)
			assert.NoError(t, err)
			assertDoesntHaveRecipe(t, recipes, pie)
		})

		t.Run("if we have the ingredients for a recipe we can make it", func(t *testing.T) {
			var (
				ctx                             = context.Background()
				bananaBread                     = randomRecipe()
				aRecipeWeWontHaveIngredientsFor = randomRecipe()
				recipeBook, pantry, teardown    = r.CreateDependencies()
				sut                             = planner.New(recipeBook, pantry)
			)
			t.Cleanup(teardown)

			assert.NoError(t, recipeBook.AddRecipes(ctx, bananaBread, aRecipeWeWontHaveIngredientsFor))
			assert.NoError(t, pantry.Store(ctx, bananaBread.Ingredients...))

			recipes, err := sut.SuggestRecipes(ctx)
			assert.NoError(t, err)
			assertHasRecipe(t, recipes, bananaBread)
			assertDoesntHaveRecipe(t, recipes, aRecipeWeWontHaveIngredientsFor)
		})

		t.Run("if we have ingredients for 2 recipes, we can make both", func(t *testing.T) {
			var (
				ctx                         = context.Background()
				bananaBread                 = randomRecipe()
				bananaMilkshake             = randomRecipe()
				recipeBook, store, teardown = r.CreateDependencies()
				sut                         = planner.New(recipeBook, store)
			)
			t.Cleanup(teardown)

			assert.NoError(t, recipeBook.AddRecipes(ctx, bananaBread, bananaMilkshake))
			assert.NoError(t, store.Store(ctx, bananaBread.Ingredients...))
			assert.NoError(t, store.Store(ctx, bananaMilkshake.Ingredients...))

			recipes, err := sut.SuggestRecipes(ctx)
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
