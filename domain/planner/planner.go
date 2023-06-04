package planner

import (
	"context"
	"errors"
	"fmt"
	"github.com/quii/go-fakes-and-contracts/domain/ingredients"
	"github.com/quii/go-fakes-and-contracts/domain/recipe"
	"log"
	"time"
)

var (
	ErrAlreadyScheduled = errors.New("recipe already scheduled")
)

type ErrorMissingIngredients struct {
	MissingIngredients ingredients.Ingredients
}

func (e ErrorMissingIngredients) Error() string {
	return fmt.Sprintf("missing ingredients: %v", e.MissingIngredients)
}

type Planner struct {
	recipeBook RecipeBook
	pantry     Pantry
}

func New(recipes RecipeBook, ingredientStore Pantry) *Planner {
	return &Planner{recipeBook: recipes, pantry: ingredientStore}
}

func (p Planner) ScheduleMeal(ctx context.Context, r recipe.Recipe, date time.Time) error {
	// record recipe in calendar

	// check ingredients are available in pantry
	availableIngredients, err := p.pantry.GetIngredients(ctx)
	if err != nil {
		return err
	}

	if !haveIngredients(availableIngredients, r) {
		missingIngredients := ingredients.Ingredients{}
		for _, ingredient := range r.Ingredients {
			if !availableIngredients.Has(ingredient) {
				missingIngredients = append(missingIngredients, ingredient)
			}
		}

		return ErrorMissingIngredients{
			MissingIngredients: missingIngredients,
		}
	}
	// remove ingredients used from pantry
	log.Println("removing ingredients from pantry", r.Ingredients)
	return p.pantry.Remove(ctx, r.Ingredients...)
}

func (p Planner) SuggestRecipes(ctx context.Context) (recipe.Recipes, error) {
	availableIngredients, err := p.pantry.GetIngredients(ctx)
	if err != nil {
		return nil, err
	}

	recipes, err := p.recipeBook.GetRecipes(ctx)
	if err != nil {
		return nil, err
	}

	var suggestions []recipe.Recipe
	for _, r := range recipes {
		if haveIngredients(availableIngredients, r) {
			suggestions = append(suggestions, r)
		}
	}
	return suggestions, nil
}

func haveIngredients(availableIngredients ingredients.Ingredients, recipe recipe.Recipe) bool {
	for _, ingredient := range recipe.Ingredients {
		if !availableIngredients.Has(ingredient) {
			return false
		}
	}
	return true
}
