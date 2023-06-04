package plannertest

import (
	"github.com/quii/go-fakes-and-contracts/domain/recipe"
	"testing"
)

func AssertHasRecipe(t *testing.T, recipes recipe.Recipes, expected recipe.Recipe) {
	t.Helper()
	_, found := recipes.FindByName(expected.Name)
	AssertTrue(t, found)
}

func AssertDoesntHaveRecipe(t *testing.T, recipes recipe.Recipes, expected recipe.Recipe) {
	t.Helper()
	_, found := recipes.FindByName(expected.Name)
	AssertFalse(t, found)
}

func AssertEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func AssertNotEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got == want {
		t.Errorf("didn't want %v", got)
	}
}

func AssertTrue(t *testing.T, got bool) {
	t.Helper()
	if !got {
		t.Errorf("got %v, want true", got)
	}
}

func AssertFalse(t *testing.T, got bool) {
	t.Helper()
	if got {
		t.Errorf("got %v, want false", got)
	}
}

func AssertNoErr(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatal(err)
	}
}
