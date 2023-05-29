package sqlite

import (
	"context"
	"entgo.io/ent/dialect"
	"github.com/quii/go-fakes-and-contracts/adapters/persistence/sqlite/ent"
	"github.com/quii/go-fakes-and-contracts/adapters/persistence/sqlite/ent/ingredient"
	"github.com/quii/go-fakes-and-contracts/domain/ingredients"
	"github.com/quii/go-fakes-and-contracts/domain/recipe"
	"log"
)

type RecipeStore struct {
	client *ent.Client
}

func NewRecipeStore() *RecipeStore {
	client, err := ent.Open(dialect.SQLite, "file:ent?mode=memory&cache=shared&_fk=1")
	//client, err := ent.Open(dialect.SQLite, "file.db?_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	return &RecipeStore{client: client}
}

func (r RecipeStore) GetRecipes(ctx context.Context) ([]recipe.Recipe, error) {
	all, err := r.client.Recipe.Query().WithRecipeingredient().All(ctx)
	if err != nil {
		return nil, err
	}
	var recipes []recipe.Recipe
	for _, r := range all {

		var ings []ingredients.Ingredient
		log.Println("ingredients for", r.ID, r.Edges.Recipeingredient)
		for _, recipeIngredient := range r.Edges.Recipeingredient {
			first, err := recipeIngredient.QueryIngredient().First(ctx)
			if err != nil {
				return nil, err
			}
			ings = append(ings, ingredients.Ingredient{
				Name:     first.Name,
				Quantity: uint(recipeIngredient.Quantity),
			})
		}

		recipes = append(recipes, recipe.Recipe{
			Name:        r.Name,
			MealType:    recipe.MealType(r.MealType),
			Description: r.Description,
			Ingredients: ings,
		})
	}
	log.Println("got recipes", recipes)
	return recipes, nil
}

func (r RecipeStore) AddRecipes(ctx context.Context, recipes ...recipe.Recipe) error {
	for _, newRecipe := range recipes {
		var recipeIngredients []*ent.RecipeIngredient
		for _, newIngredient := range newRecipe.Ingredients {
			// create ingredient if it doesn't exist
			err := r.client.Ingredient.Create().SetName(newIngredient.Name).OnConflict().DoNothing().Exec(ctx)
			if err != nil {
				return err
			}

			// get it out the db again (kinda sucks, but not sure how to do this better with ent)
			savedIngredient, err := r.client.Ingredient.Query().Where(ingredient.Name(newIngredient.Name)).First(ctx)
			if err != nil {
				return err
			}

			ri, err := r.client.RecipeIngredient.Create().
				SetIngredient(savedIngredient).
				SetQuantity(int(newIngredient.Quantity)).
				Save(ctx)
			if err != nil {
				return err
			}

			recipeIngredients = append(recipeIngredients, ri)
		}

		// create recipe
		err := r.client.Recipe.Create().
			SetName(newRecipe.Name).
			AddRecipeingredient(recipeIngredients...).
			SetMealType(int(newRecipe.MealType)).
			SetDescription(newRecipe.Description).
			Exec(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r RecipeStore) Close() {
}