package inmemory_test

import (
	"github.com/quii/go-fakes-and-contracts/adapters/persistence/inmemory"
	"github.com/quii/go-fakes-and-contracts/domain/planner"
	"testing"
)

func TestInMemoryIngredientStore(t *testing.T) {
	planner.IngredientStoreContract{
		NewStore: func() planner.CloseableIngredientStore {
			return inmemory.NewIngredientStore()
		},
	}.Test(t)
}
