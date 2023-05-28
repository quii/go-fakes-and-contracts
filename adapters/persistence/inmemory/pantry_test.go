package inmemory_test

import (
	"github.com/quii/go-fakes-and-contracts/adapters/persistence/inmemory"
	"github.com/quii/go-fakes-and-contracts/domain/planner"
	"testing"
)

func TestInMemoryIngredientStore(t *testing.T) {
	planner.PantryContract{
		NewPantry: func() planner.CloseablePantry {
			return inmemory.NewPantry()
		},
	}.Test(t)
}
