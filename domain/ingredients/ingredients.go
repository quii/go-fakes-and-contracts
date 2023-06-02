package ingredients

type Ingredient struct {
	Name     string
	Quantity uint
}

type Ingredients []Ingredient

func (i Ingredients) Has(ingredient Ingredient) bool {
	for _, pantryIngredient := range i {
		if pantryIngredient.Name == ingredient.Name {
			return true
		}
	}
	return false
}

func (i Ingredients) NumberOf(ingredientName string) uint {
	for _, pantryIngredient := range i {
		if pantryIngredient.Name == ingredientName {
			return pantryIngredient.Quantity
		}
	}
	return 0
}
