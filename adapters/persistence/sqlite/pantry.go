package sqlite

import (
	"context"
	"entgo.io/ent/dialect"
	"fmt"
	"github.com/quii/go-fakes-and-contracts/adapters/persistence/sqlite/ent"
	"github.com/quii/go-fakes-and-contracts/adapters/persistence/sqlite/ent/ingredient"
	"github.com/quii/go-fakes-and-contracts/adapters/persistence/sqlite/ent/pantry"
	"github.com/quii/go-fakes-and-contracts/domain/ingredients"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Pantry struct {
	client *ent.Client
}

func NewPantry() *Pantry {
	client, err := ent.Open(dialect.SQLite, "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	return &Pantry{client: client}
}

func (i Pantry) GetIngredients(ctx context.Context) ([]ingredients.Ingredient, error) {
	persistedPantry, err := i.client.Pantry.Query().WithIngredient().All(ctx)
	if err != nil {
		return nil, err
	}

	var allIngredients []ingredients.Ingredient
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
		err := i.client.Ingredient.Create().SetName(newIngredient.Name).OnConflict().DoNothing().Exec(ctx)
		if err != nil {
			return fmt.Errorf("failed to create ingredient: %w", err)
		}

		// get it out the db again (kinda sucks, but not sure how to do this better with ent)
		savedIngredient, err := i.client.Ingredient.Query().Where(ingredient.Name(newIngredient.Name)).First(ctx)
		if err != nil {
			return fmt.Errorf("failed to find ingredient %q: %w", newIngredient.Name, err)
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
		pantryItem[0].Update().AddQuantity(int(newIngredient.Quantity)).Exec(ctx)
	}
	return nil
}

func (i Pantry) Close() {
	if err := i.client.Close(); err != nil {
		log.Println("couldn't close sqlite3 client", err)
	}
}
