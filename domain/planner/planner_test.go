package planner_test

import (
	"context"
	"errors"
	"github.com/quii/go-fakes-and-contracts/adapters/driven/persistence/inmemory"
	"github.com/quii/go-fakes-and-contracts/adapters/driven/persistence/sqlite"
	"github.com/quii/go-fakes-and-contracts/domain/ingredients"
	"github.com/quii/go-fakes-and-contracts/domain/planner"
	"github.com/quii/go-fakes-and-contracts/domain/planner/internal/expect"
	"github.com/quii/go-fakes-and-contracts/domain/planner/internal/plannertest"
	"testing"
	"time"
)

func TestRecipePlanner(t *testing.T) {
	// for local, snappy integration test with a fake (which we can be confident is correct due to it conforming to the store contract)
	t.Run("with in memory store", func(t *testing.T) {
		RecipePlannerTest{
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
			RecipePlannerTest{
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

type RecipePlannerTest struct {
	CreateDependencies func() (planner.RecipeBook, planner.Pantry, Cleanup)
}

func (r RecipePlannerTest) Test(t *testing.T) {
	t.Run("planning meals", func(t *testing.T) {

		t.Run("happy path, have ingredients for a recipe, schedule it, update pantry", func(t *testing.T) {
			var (
				ctx                          = context.Background()
				lasagna                      = plannertest.RandomRecipe()
				recipeBook, pantry, teardown = r.CreateDependencies()
				sut                          = planner.New(recipeBook, pantry)
			)
			t.Cleanup(teardown)

			expect.NoErr(t, recipeBook.AddRecipes(ctx, lasagna))
			expect.NoErr(t, pantry.Store(ctx, lasagna.Ingredients...))

			expect.NoErr(t, sut.ScheduleMeal(ctx, lasagna, time.Now()))
			remainingIngredients, err := pantry.GetIngredients(ctx)
			expect.NoErr(t, err)
			expect.Equal(t, 0, len(remainingIngredients))
		})

		t.Run("returns a missing ingredients error if you try to schedule a meal without all the ingredients", func(t *testing.T) {
			var (
				ctx                         = context.Background()
				lasagna                     = plannertest.RandomRecipe()
				recipeBook, store, teardown = r.CreateDependencies()
				sut                         = planner.New(recipeBook, store)
			)
			t.Cleanup(teardown)

			expect.NoErr(t, recipeBook.AddRecipes(ctx, lasagna))

			err := sut.ScheduleMeal(ctx, lasagna, time.Now())
			expect.Err(t, err)

			missingIngredientsErr, ok := err.(planner.ErrorMissingIngredients)
			expect.True(t, ok)
			expect.DeepEqual(t, planner.ErrorMissingIngredients{
				MissingIngredients: lasagna.Ingredients,
			}, missingIngredientsErr)
		})

		t.Run("when recipeBook fails to get ingredients, we get an error", func(t *testing.T) {
			var (
				ctx                          = context.Background()
				lasagna                      = plannertest.RandomRecipe()
				recipeBook, pantry, teardown = r.CreateDependencies()
				failingPantry                = planner.NewPantryDelegate(pantry)
			)
			t.Cleanup(teardown)

			failingPantry.GetIngredientsFunc = func(ctx context.Context) (ingredients.Ingredients, error) {
				return nil, errors.New("oh no")
			}

			sut := planner.New(recipeBook, failingPantry)
			expect.NoErr(t, recipeBook.AddRecipes(ctx, lasagna))
			expect.NoErr(t, pantry.Store(ctx, lasagna.Ingredients...))

			err := sut.ScheduleMeal(ctx, lasagna, time.Now())
			expect.Err(t, err)
		})

		t.Run("returns the specific ingredients missing if you try to schedule a meal with some missing ingredients", func(t *testing.T) {
			var (
				ctx                          = context.Background()
				recipeBook, pantry, teardown = r.CreateDependencies()
				lasagna                      = plannertest.RandomRecipe()
				sut                          = planner.New(recipeBook, pantry)
			)
			t.Cleanup(teardown)

			expect.NoErr(t, recipeBook.AddRecipes(ctx, lasagna))

			missingIngredient, ingredientsWeHave := lasagna.Ingredients[0], lasagna.Ingredients[1:]
			expect.NoErr(t, pantry.Store(ctx, ingredientsWeHave...))

			err := sut.ScheduleMeal(ctx, lasagna, time.Now())
			expect.Err(t, err)

			missingIngredientsErr, ok := err.(planner.ErrorMissingIngredients)
			expect.True(t, ok)
			expect.DeepEqual(t, planner.ErrorMissingIngredients{
				MissingIngredients: []ingredients.Ingredient{missingIngredient},
			}, missingIngredientsErr)
		})

	})

	t.Run("suggesting recipes", func(t *testing.T) {

		t.Run("if don't have the ingredients for a meal, we cant make it", func(t *testing.T) {
			var (
				ctx                          = context.Background()
				pie                          = plannertest.RandomRecipe()
				recipeBook, pantry, teardown = r.CreateDependencies()
				sut                          = planner.New(recipeBook, pantry)
			)
			t.Cleanup(teardown)

			expect.NoErr(t, recipeBook.AddRecipes(ctx, pie))

			recipes, err := sut.SuggestRecipes(ctx)
			expect.NoErr(t, err)
			plannertest.AssertDoesntHaveRecipe(t, recipes, pie)
		})

		t.Run("if we have the ingredients for a recipe we can make it", func(t *testing.T) {
			var (
				ctx                             = context.Background()
				bananaBread                     = plannertest.RandomRecipe()
				aRecipeWeWontHaveIngredientsFor = plannertest.RandomRecipe()
				recipeBook, pantry, teardown    = r.CreateDependencies()
				sut                             = planner.New(recipeBook, pantry)
			)
			t.Cleanup(teardown)

			expect.NoErr(t, recipeBook.AddRecipes(ctx, bananaBread, aRecipeWeWontHaveIngredientsFor))
			expect.NoErr(t, pantry.Store(ctx, bananaBread.Ingredients...))

			recipes, err := sut.SuggestRecipes(ctx)
			expect.NoErr(t, err)
			plannertest.AssertHasRecipe(t, recipes, bananaBread)
			plannertest.AssertDoesntHaveRecipe(t, recipes, aRecipeWeWontHaveIngredientsFor)
		})

		t.Run("if we have ingredients for 2 recipes, we can make both", func(t *testing.T) {
			var (
				ctx                         = context.Background()
				bananaBread                 = plannertest.RandomRecipe()
				bananaMilkshake             = plannertest.RandomRecipe()
				recipeBook, store, teardown = r.CreateDependencies()
				sut                         = planner.New(recipeBook, store)
			)
			t.Cleanup(teardown)

			expect.NoErr(t, recipeBook.AddRecipes(ctx, bananaBread, bananaMilkshake))
			expect.NoErr(t, store.Store(ctx, bananaBread.Ingredients...))
			expect.NoErr(t, store.Store(ctx, bananaMilkshake.Ingredients...))

			recipes, err := sut.SuggestRecipes(ctx)
			expect.NoErr(t, err)
			plannertest.AssertHasRecipe(t, recipes, bananaBread)
			plannertest.AssertHasRecipe(t, recipes, bananaMilkshake)
		})
	})
}
