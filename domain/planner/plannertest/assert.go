package plannertest

import (
	"github.com/quii/go-fakes-and-contracts/domain/planner/expect"
	"github.com/quii/go-fakes-and-contracts/domain/recipe"
	"testing"
)

func ExpectHasRecipe(t *testing.T, recipes recipe.Recipes, expected recipe.Recipe) {
	t.Helper()
	_, found := recipes.FindByName(expected.Name)
	expect.True(t, found)
}

func ExpectDoesntHaveRecipe(t *testing.T, recipes recipe.Recipes, expected recipe.Recipe) {
	t.Helper()
	_, found := recipes.FindByName(expected.Name)
	expect.False(t, found)
}
