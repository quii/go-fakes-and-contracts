package plannertest

import (
	"github.com/quii/go-fakes-and-contracts/domain/planner/internal/expect"
	"github.com/quii/go-fakes-and-contracts/domain/recipe"
	"testing"
)

func AssertHasRecipe(t *testing.T, recipes recipe.Recipes, expected recipe.Recipe) {
	t.Helper()
	_, found := recipes.FindByName(expected.Name)
	expect.True(t, found)
}

func AssertDoesntHaveRecipe(t *testing.T, recipes recipe.Recipes, expected recipe.Recipe) {
	t.Helper()
	_, found := recipes.FindByName(expected.Name)
	expect.False(t, found)
}
