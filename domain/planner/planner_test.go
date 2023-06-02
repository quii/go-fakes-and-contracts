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
	t.Run("if we have no ingredients we can't make anything", func(t *testing.T) {
		store := r.NewPantry()

		recipeBook := r.NewRecipeBook()
		assert.NoError(t, recipeBook.AddRecipes(context.Background(), randomRecipe(), randomRecipe()))

		assertAvailableRecipes(t, store, recipeBook, []recipe.Recipe{})
	})

	t.Run("if we have the ingredients for a recipe we can make it", func(t *testing.T) {
		store := r.NewPantry()
		recipeBook := r.NewRecipeBook()
		bananaBread := randomRecipe()

		assert.NoError(t, recipeBook.AddRecipes(context.Background(), bananaBread, randomRecipe()))

		assert.NoError(t, store.Store(
			context.Background(),
			bananaBread.Ingredients...,
		))
		assertAvailableRecipes(t, store, recipeBook, []recipe.Recipe{bananaBread})
	})

	t.Run("if we have ingredients for 2 recipes, we can make both", func(t *testing.T) {
		store := r.NewPantry()
		recipeBook := r.NewRecipeBook()
		bananaBread := randomRecipe()
		bananaMilkshake := randomRecipe()

		assert.NoError(t, recipeBook.AddRecipes(context.Background(), bananaBread, bananaMilkshake))

		assert.NoError(t, store.Store(
			context.Background(),
			append(bananaBread.Ingredients, bananaMilkshake.Ingredients...)...,
		))
		assertAvailableRecipes(t, store, recipeBook, []recipe.Recipe{bananaMilkshake, bananaBread})
	})

}

func assertAvailableRecipes(
	t *testing.T,
	ingredientStore planner.Pantry,
	recipeStore planner.RecipeBook,
	expectedRecipes []recipe.Recipe,
) {
	t.Helper()
	suggestions, _ := planner.New(recipeStore, ingredientStore).SuggestRecipes(context.Background())

	if len(expectedRecipes) == 0 {
		assert.Equal(t, 0, len(suggestions))
		return
	}

	// create a map to count occurrences of each recipe in the suggestions
	suggestionCounts := make(map[string]int)
	for _, suggestion := range suggestions {
		suggestionCounts[suggestion.Name]++
	}

	// check that the counts of the expected recipes match the actual counts in the suggestions
	for _, expectedRecipe := range expectedRecipes {
		actualCount, ok := suggestionCounts[expectedRecipe.Name]
		if !ok {
			t.Errorf("expected recipe %s not found in suggestions", expectedRecipe.Name)
			continue
		}
		if actualCount != 1 {
			t.Errorf("expected recipe %s to appear once in suggestions, but found %d occurrences", expectedRecipe.Name, actualCount)
		}
	}
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
