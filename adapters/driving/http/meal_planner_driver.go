package http

import (
	"context"
	"github.com/go-rod/rod"
	"github.com/quii/go-fakes-and-contracts/domain/ingredients"
	"github.com/quii/go-fakes-and-contracts/domain/recipe"
	"time"
)

type MealPlannerDriver struct {
	baseURL string
	browser *rod.Browser
}

func (m MealPlannerDriver) ScheduleMeal(ctx context.Context, r recipe.Recipe, time time.Time) error {
	//TODO implement me
	panic("implement me")
}

func (m MealPlannerDriver) SuggestRecipes(ctx context.Context) (recipe.Recipes, error) {
	//TODO implement me
	panic("implement me")
}

func (m MealPlannerDriver) GetIngredients(ctx context.Context) (ingredients.Ingredients, error) {
	//TODO implement me
	panic("implement me")
}

func (m MealPlannerDriver) Store(ctx context.Context, ingredient ...ingredients.Ingredient) error {
	//TODO implement me
	panic("implement me")
}

func (m MealPlannerDriver) Remove(ctx context.Context, i ...ingredients.Ingredient) error {
	//TODO implement me
	panic("implement me")
}

func (m MealPlannerDriver) GetRecipes(ctx context.Context) ([]recipe.Recipe, error) {
	//TODO implement me
	panic("implement me")
}

func (m MealPlannerDriver) AddRecipes(ctx context.Context, recipe ...recipe.Recipe) error {
	//TODO implement me
	panic("implement me")
}

func NewMealPlannerDriver(baseURL string) (*MealPlannerDriver, func() error) {
	browser := rod.New().MustConnect()
	return &MealPlannerDriver{baseURL: baseURL, browser: browser}, browser.Close
}
