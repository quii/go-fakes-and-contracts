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
	// Run the auto migration tool.
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
		allIngredients = append(allIngredients, ingredients.Ingredient{
			Name:     ingredient.Name,
			Quantity: ingredient.Quantity,
		})
	}
	return allIngredients, nil
}

func (i IngredientStore) Store(ctx context.Context, ingredients ...ingredients.Ingredient) error {
	for _, ingredient := range ingredients {
		_, err := i.client.Ingredient.Create().SetName(ingredient.Name).SetQuantity(ingredient.Quantity).Save(ctx)
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
