package specifications

import (
	"context"
	"github.com/quii/go-fakes-and-contracts/domain/planner"
	"github.com/quii/go-fakes-and-contracts/domain/planner/expect"
	"github.com/quii/go-fakes-and-contracts/domain/planner/plannertest"
	"testing"
	"time"
)

type (
	MealPlanner struct {
		CreateDependencies func() (planner.RecipeBook, planner.Pantry, planner.MealPlanner)
	}
)

func (m MealPlanner) Test(t *testing.T) {
	t.Run("planning meals", func(t *testing.T) {

		t.Run("when we have ingredients for a meal, we can schedule it", func(t *testing.T) {
			var (
				ctx                             = context.Background()
				lasagna                         = plannertest.RandomRecipe()
				recipeBook, pantry, mealPlanner = m.CreateDependencies()
			)

			expect.NoErr(t, recipeBook.AddRecipes(ctx, lasagna))
			expect.NoErr(t, pantry.Store(ctx, lasagna.Ingredients...))

			recipes, err := mealPlanner.SuggestRecipes(ctx)
			expect.NoErr(t, err)
			lasagnaFromBook, found := recipes.FindByName(lasagna.Name)
			expect.True(t, found)

			expect.NoErr(t, mealPlanner.ScheduleMeal(ctx, lasagnaFromBook, time.Now()))
			remainingIngredients, err := pantry.GetIngredients(ctx)
			expect.NoErr(t, err)
			expect.Equal(t, 0, len(remainingIngredients))
		})
	})

}
