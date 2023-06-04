package inmemory

import (
	"context"
	"github.com/quii/go-fakes-and-contracts/domain/ingredients"
)

type Pantry struct {
	ingredients ingredients.Ingredients
}

func (p *Pantry) Remove(ctx context.Context, i ...ingredients.Ingredient) error {
	for _, ingredient := range i {
		p.ingredients.Remove(ingredient)
	}
	return nil
}

func NewPantry() *Pantry {
	return &Pantry{}
}

func (p *Pantry) GetIngredients(_ context.Context) (ingredients.Ingredients, error) {
	return p.ingredients, nil
}

func (p *Pantry) Store(_ context.Context, ingredients ...ingredients.Ingredient) error {
	for idx, ingredient := range ingredients {
		if p.ingredients.Has(ingredient) {
			p.ingredients[idx].Quantity += ingredient.Quantity
		} else {
			p.ingredients = append(p.ingredients, ingredient)
		}
	}
	return nil
}
