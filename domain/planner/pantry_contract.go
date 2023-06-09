package planner

import (
	"context"
	"github.com/quii/go-fakes-and-contracts/domain/ingredients"
	"github.com/quii/go-fakes-and-contracts/domain/planner/expect"
	"github.com/quii/go-fakes-and-contracts/domain/planner/plannertest"
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
			someIngredients = plannertest.RandomIngredients()
			sut             = s.NewPantry()
		)
		expect.NoErr(t, sut.Store(ctx, someIngredients...))

		storedIngredients, err := sut.GetIngredients(ctx)
		expect.NoErr(t, err)
		expect.SliceEqual(t, storedIngredients, someIngredients)
	})

	t.Run("it adds to the quantity of ingredients", func(t *testing.T) {
		var (
			ctx    = context.Background()
			orange = plannertest.RandomIngredient()
			sut    = s.NewPantry()
		)
		orange.Quantity = 1

		expect.NoErr(t, sut.Store(ctx, orange))
		expect.NoErr(t, sut.Store(ctx, orange))
		expect.NoErr(t, sut.Store(ctx, orange))

		got, err := sut.GetIngredients(ctx)
		expect.NoErr(t, err)
		expect.Equal(t, got.NumberOf(orange.Name), 3)
	})

	t.Run("it removes quantities of ingredients", func(t *testing.T) {
		var (
			ctx   = context.Background()
			apple = plannertest.RandomIngredient()
			sut   = s.NewPantry()
		)
		apple.Quantity = 1

		expect.NoErr(t, sut.Store(ctx, apple))
		expect.NoErr(t, sut.Store(ctx, apple))
		expect.NoErr(t, sut.Store(ctx, apple))

		expect.NoErr(t, sut.Remove(ctx, apple))

		got, err := sut.GetIngredients(ctx)
		expect.NoErr(t, err)
		expect.Equal(t, got.NumberOf(apple.Name), 2)
	})

	t.Run("if you run out of an ingredient entirely, it is removed from the pantry", func(t *testing.T) {
		var (
			ctx    = context.Background()
			banana = plannertest.RandomIngredient()
			sut    = s.NewPantry()
		)

		expect.NoErr(t, sut.Store(ctx, banana))
		expect.NoErr(t, sut.Remove(ctx, banana))

		got, err := sut.GetIngredients(ctx)
		expect.NoErr(t, err)
		expect.Equal(t, got.NumberOf(banana.Name), 0)
	})
}
