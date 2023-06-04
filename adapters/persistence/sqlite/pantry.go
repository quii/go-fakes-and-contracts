package sqlite

import (
	"context"
	_ "github.com/mattn/go-sqlite3"
	"github.com/quii/go-fakes-and-contracts/adapters/persistence/sqlite/ent"
	"github.com/quii/go-fakes-and-contracts/adapters/persistence/sqlite/ent/ingredient"
	"github.com/quii/go-fakes-and-contracts/adapters/persistence/sqlite/ent/pantry"
	"github.com/quii/go-fakes-and-contracts/domain/ingredients"
)

type Pantry struct {
	client *ent.Client
}

func NewPantry(client *ent.Client) *Pantry {
	return &Pantry{
		client: client,
	}
}

func (i Pantry) GetIngredients(ctx context.Context) (ingredients.Ingredients, error) {
	persistedPantry, err := i.client.Pantry.Query().WithIngredient().All(ctx)
	if err != nil {
		return nil, err
	}

	var allIngredients ingredients.Ingredients
	for _, pantryItem := range persistedPantry {
		if pantryItem.Quantity == 0 {
			continue
		}
		allIngredients = append(allIngredients, ingredients.Ingredient{
			Name:     pantryItem.Edges.Ingredient.Name,
			Quantity: uint(pantryItem.Quantity),
		})
	}
	return allIngredients, nil
}

func (i Pantry) Store(ctx context.Context, ingredients ...ingredients.Ingredient) error {
	for _, newIngredient := range ingredients {
		if err := i.addOrIncrementIngredient(ctx, newIngredient); err != nil {
			return err
		}
	}
	return nil
}

func (i Pantry) Remove(ctx context.Context, toRemove ...ingredients.Ingredient) error {
	// decrement the count for each ingredient
	for _, ii := range toRemove {
		err := i.client.Pantry.Update().
			Where(pantry.HasIngredientWith(ingredient.Name(ii.Name))).
			AddQuantity(-int(ii.Quantity)).
			Exec(ctx)

		if err != nil {
			return err
		}
	}
	return nil
}

func (i Pantry) addOrIncrementIngredient(ctx context.Context, newIngredient ingredients.Ingredient) error {
	savedIngredient, err := CreateIngredientIfNotExists(ctx, i.client, newIngredient)
	if err != nil {
		return err
	}

	err = i.client.Pantry.Create().
		SetIngredientID(savedIngredient.ID).
		SetQuantity(int(newIngredient.Quantity)).
		Exec(ctx)

	if ent.IsConstraintError(err) {
		err = i.client.Pantry.Update().
			Where(pantry.HasIngredientWith(ingredient.ID(savedIngredient.ID))).
			AddQuantity(int(newIngredient.Quantity)).
			Exec(ctx)

		if err != nil {
			return err
		}
	}
	return nil
}
