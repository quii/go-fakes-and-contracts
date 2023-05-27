package planner

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

type CloseableStore interface {
	Store
	Close()
}

type StoreContract struct {
	NewStore func() CloseableStore
}

func (s StoreContract) Test(t *testing.T) {
	t.Run("it returns what is put in", func(t *testing.T) {
		ctx := context.Background()
		store := s.NewStore()
		t.Cleanup(store.Close)

		want := []ingredients.Ingredient{
			{Name: "Bananas", Quantity: 2},
			{Name: "Flour", Quantity: 1},
			{Name: "Eggs", Quantity: 2},
		}

		err := store.Store(ctx, want...)
		assert.NoError(t, err)

		got, err := store.GetIngredients(ctx)
		assert.NoError(t, err)
		assert.Equal(t, got, want)
	})

	t.Run("it adds to the quantity of ingredients", func(t *testing.T) {
		ctx := context.Background()
		store := s.NewStore()
		t.Cleanup(store.Close)

		assert.NoError(t, store.Store(ctx, ingredients.Ingredient{
			Name:     "Orange",
			Quantity: 1,
		}))
		assert.NoError(t, store.Store(ctx, ingredients.Ingredient{
			Name:     "Orange",
			Quantity: 1,
		}))
		assert.NoError(t, store.Store(ctx, ingredients.Ingredient{
			Name:     "Orange",
			Quantity: 1,
		}))

		got, err := store.GetIngredients(ctx)
		assert.NoError(t, err)
		assert.Equal(t, len(got), 1)
		assert.Equal(t, got[0].Quantity, 3)
	})
}
