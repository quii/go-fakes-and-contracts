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
		var (
			ctx             = context.Background()
			someIngredients = []ingredients.Ingredient{
				{Name: "Bananas", Quantity: 2},
				{Name: "Flour", Quantity: 1},
				{Name: "Eggs", Quantity: 2},
			}
			sut = s.NewPantry()
		)
		assert.NoError(t, sut.Store(ctx, someIngredients...))

		storedIngredients, err := sut.GetIngredients(ctx)
		assert.NoError(t, err)
		assert.Equal(t, storedIngredients, someIngredients)
	})

	t.Run("it adds to the quantity of ingredients", func(t *testing.T) {
		var (
			ctx    = context.Background()
			orange = ingredients.Ingredient{
				Name:     "Orange",
				Quantity: 1,
			}
			sut = s.NewPantry()
		)
		assert.NoError(t, sut.Store(ctx, orange))
		assert.NoError(t, sut.Store(ctx, orange))
		assert.NoError(t, sut.Store(ctx, orange))

		got, err := sut.GetIngredients(ctx)
		assert.NoError(t, err)
		assert.Equal(t, got.NumberOf(orange.Name), 3)
	})
}
