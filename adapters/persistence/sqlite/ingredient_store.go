package sqlite

import (
	"context"
	"github.com/quii/go-fakes-and-contracts/adapters/persistence/sqlite/ent"
	"github.com/quii/go-fakes-and-contracts/domain/ingredients"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type IngredientStore struct {
	client *ent.Client
}

func NewIngredientStore() *IngredientStore {
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	return &IngredientStore{client: client}
}

func (i IngredientStore) GetIngredients(ctx context.Context) ([]ingredients.Ingredient, error) {
	all, err := i.client.Ingredient.Query().All(ctx)
	if err != nil {
		return nil, err
	}
	var allIngredients []ingredients.Ingredient

	for _, ingredient := range all {
		allIngredients = append(allIngredients, mapDBIngredientToDomain(ingredient))
	}
	return allIngredients, nil
}

func (i IngredientStore) Store(ctx context.Context, ingredients ...ingredients.Ingredient) error {
	for _, newIngredient := range ingredients {
		err := i.client.Ingredient.
			Create().
			SetName(newIngredient.Name).
			SetQuantity(newIngredient.Quantity).
			OnConflict().
			AddQuantity(newIngredient.Quantity).
			Exec(ctx)

		if err != nil {
			return err
		}
	}
	return nil
}

func (i IngredientStore) Close() {
	if err := i.client.Close(); err != nil {
		log.Println("couldn't close sqlite3 client", err)
	}
}

func mapDBIngredientToDomain(ingredient *ent.Ingredient) ingredients.Ingredient {
	return ingredients.Ingredient{
		Name:     ingredient.Name,
		Quantity: ingredient.Quantity,
	}
}