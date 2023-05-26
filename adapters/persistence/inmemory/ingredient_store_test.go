package inmemory_test

import (
	"github.com/quii/go-fakes-and-contracts/adapters/persistence/inmemory"
	"github.com/quii/go-fakes-and-contracts/domain/recipe"
	"testing"
)

func TestInMemoryIngredientStore(t *testing.T) {
	recipe.StoreContract{
		NewStore: func() recipe.CloseableStore {
			return inmemory.NewIngredientStore()
		},
	}.Test(t)
}
