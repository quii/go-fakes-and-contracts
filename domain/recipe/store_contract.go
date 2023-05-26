package recipe

import (
	"context"
	"github.com/alecthomas/assert/v2"
	"github.com/quii/go-fakes-and-contracts/domain/ingredients"
	"testing"
)

type Store interface {
	GetIngredients(ctx context.Context) ([]ingredients.Ingredient, error)
	Store(context.Context, ...ingredients.Ingredient) error
}
type StoreContract struct {
	NewStore func() Store
}

func (s StoreContract) Test(t *testing.T) {
	t.Run("it returns what is put in", func(t *testing.T) {
		ctx := context.Background()
		want := []ingredients.Ingredient{
			{Name: "Bananas", Quantity: 2},
			{Name: "Flour", Quantity: 1},
			{Name: "Eggs", Quantity: 2},
		}
		store := s.NewStore()
		err := store.Store(ctx, want...)
		assert.NoError(t, err)

		got, err := store.GetIngredients(ctx)
		assert.NoError(t, err)
		assert.Equal(t, got, want)
	})
}
