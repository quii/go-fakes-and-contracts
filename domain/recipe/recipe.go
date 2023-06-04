package recipe

import "github.com/quii/go-fakes-and-contracts/domain/ingredients"

type MealType int

// mealtype enum
const (
	Breakfast MealType = iota
	Lunch
	Dinner
)

func (m MealType) String() string {
	return [...]string{"Breakfast", "Lunch", "Dinner"}[m]
}

func NewMealTypeFromInt(i int) MealType {
	return MealType(i)
}

type Recipe struct {
	Name        string
	MealType    MealType
	Ingredients []ingredients.Ingredient
	Description string
}

type Recipes []Recipe

func (r Recipes) FindByName(name string) (Recipe, bool) {
	for _, recipe := range r {
		if recipe.Name == name {
			return recipe, true
		}
	}
	return Recipe{}, false
}
