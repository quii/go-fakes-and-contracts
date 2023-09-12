package planner

import (
	"context"
	"github.com/quii/go-fakes-and-contracts/domain/ingredients"
)

type PantryDelegate struct {
	GetIngredientsFunc func(ctx context.Context) (ingredients.Ingredients, error)
	StoreFunc          func(context.Context, ...ingredients.Ingredient) error
	RemoveFunc         func(ctx context.Context, i ...ingredients.Ingredient) error
	delegate           Pantry
}

func NewPantryDelegate(delegate Pantry) *PantryDelegate {
	return &PantryDelegate{delegate: delegate}
}

func (p *PantryDelegate) GetIngredients(ctx context.Context) (ingredients.Ingredients, error) {
	if p.GetIngredientsFunc != nil {
		return p.GetIngredientsFunc(ctx)
	}
	return p.delegate.GetIngredients(ctx)
}

func (p *PantryDelegate) Store(ctx context.Context, ingredients ...ingredients.Ingredient) error {
	if p.StoreFunc != nil {
		return p.StoreFunc(ctx, ingredients...)
	}
	return p.delegate.Store(ctx, ingredients...)
}

func (p *PantryDelegate) Remove(ctx context.Context, i ...ingredients.Ingredient) error {
	if p.RemoveFunc != nil {
		return p.RemoveFunc(ctx, i...)
	}
	return p.delegate.Remove(ctx, i...)
}
