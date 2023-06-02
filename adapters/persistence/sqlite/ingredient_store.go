package sqlite

import (
	"context"
	"fmt"
	"github.com/quii/go-fakes-and-contracts/adapters/persistence/sqlite/ent"
	"github.com/quii/go-fakes-and-contracts/domain/ingredients"
)

func CreateIngredientIfNotExists(ctx context.Context, client *ent.Client, newIngredient ingredients.Ingredient) (*ent.Ingredient, error) {
	id, err := client.Ingredient.Create().SetName(newIngredient.Name).OnConflict().Ignore().ID(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create ingredient %v: %w", newIngredient, err)
	}

	return client.Ingredient.GetX(ctx, id), nil
}
