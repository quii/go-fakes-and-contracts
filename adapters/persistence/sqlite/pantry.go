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
		allIngredients = append(allIngredients, ingredients.Ingredient{
			Name:     pantryItem.Edges.Ingredient.Name,
			Quantity: uint(pantryItem.Quantity),
		})
	}
	return allIngredients, nil
}

func (i Pantry) Store(ctx context.Context, ingredients ...ingredients.Ingredient) error {
	for _, newIngredient := range ingredients {
		// create ingredient if it doesn't exist
		savedIngredient, err := CreateIngredientIfNotExists(ctx, i.client, newIngredient)
		if err != nil {
			return err
		}

		// this sucks ass, im sure i've got the schema wrong for this to be needed
		pantryItem, err := i.client.Pantry.Query().Where(pantry.HasIngredientWith(ingredient.ID(savedIngredient.ID))).All(ctx)
		if err != nil {
			return err
		}
		if len(pantryItem) == 0 {
			err = i.client.Pantry.Create().SetIngredientID(savedIngredient.ID).SetQuantity(int(newIngredient.Quantity)).Exec(ctx)
			if err != nil {
				return err
			}
			continue
		}
		err = pantryItem[0].Update().AddQuantity(int(newIngredient.Quantity)).Exec(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}
