package inmemory_test

import (
	"github.com/quii/go-fakes-and-contracts/adapters/persistence/inmemory"
	"github.com/quii/go-fakes-and-contracts/domain/planner"
	"testing"
)

func TestInMemoryIngredientStore(t *testing.T) {
	planner.StoreContract{
		NewStore: func() planner.CloseableStore {
			return inmemory.NewIngredientStore()
		},
	}.Test(t)
}
