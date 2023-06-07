package sqlite

import (
	"context"
	"fmt"
	ent2 "github.com/quii/go-fakes-and-contracts/adapters/driven/persistence/sqlite/ent"
	"github.com/quii/go-fakes-and-contracts/domain/ingredients"
)

func CreateIngredientIfNotExists(ctx context.Context, client *ent2.Client, newIngredient ingredients.Ingredient) (*ent2.Ingredient, error) {
	id, err := client.Ingredient.Create().SetName(newIngredient.Name).OnConflict().Ignore().ID(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create ingredient %v: %w", newIngredient, err)
	}

	return client.Ingredient.GetX(ctx, id), nil
}
