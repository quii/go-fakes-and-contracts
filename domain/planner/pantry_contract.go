package planner

import (
	"context"
	"github.com/alecthomas/assert/v2"
	"github.com/quii/go-fakes-and-contracts/domain/ingredients"
	"testing"
)

type Pantry interface {
	GetIngredients(ctx context.Context) (ingredients.Ingredients, error)
	Store(context.Context, ...ingredients.Ingredient) error
	Remove(ctx context.Context, i ...ingredients.Ingredient) error
}

type PantryContract struct {
	NewPantry func() Pantry
}

func (s PantryContract) Test(t *testing.T) {
	t.Run("it returns what is put in", func(t *testing.T) {
		ctx := context.Background()
		store := s.NewPantry()

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
		pantry := s.NewPantry()

		assert.NoError(t, pantry.Store(ctx, ingredients.Ingredient{
			Name:     "Orange",
			Quantity: 1,
		}))
		assert.NoError(t, pantry.Store(ctx, ingredients.Ingredient{
			Name:     "Orange",
			Quantity: 1,
		}))
		assert.NoError(t, pantry.Store(ctx, ingredients.Ingredient{
			Name:     "Orange",
			Quantity: 1,
		}))

		got, err := pantry.GetIngredients(ctx)
		assert.NoError(t, err)
		assert.Equal(t, got.NumberOf("Orange"), 3)
	})
}
