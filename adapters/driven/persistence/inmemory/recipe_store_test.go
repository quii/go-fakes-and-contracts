package inmemory

import (
	"github.com/quii/go-fakes-and-contracts/domain/planner"
	"testing"
)

func TestInMemoryRecipeStore(t *testing.T) {
	planner.RecipeBookContract{NewBook: func() planner.RecipeBook {
		return NewRecipeStore()
	}}.Test(t)
}
