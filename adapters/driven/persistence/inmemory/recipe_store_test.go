package inmemory_test

import (
	"testing"

	"github.com/quii/go-fakes-and-contracts/adapters/driven/persistence/inmemory"
	"github.com/quii/go-fakes-and-contracts/domain/planner"
)

func TestInMemoryRecipeStore(t *testing.T) {
	planner.RecipeBookContract{NewBook: func() planner.RecipeBook {
		return inmemory.NewRecipeStore()
	}}.Test(t)
}
