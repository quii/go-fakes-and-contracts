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

func (i *Ingredients) Remove(ingredient Ingredient) {
	oldIngredients := *i
	for idx, pantryIngredient := range oldIngredients {
		if pantryIngredient.Name == ingredient.Name {
			oldIngredients[idx].Quantity -= ingredient.Quantity
			if oldIngredients[idx].Quantity == 0 {
				*i = append(oldIngredients[:idx], oldIngredients[idx+1:]...)
			}
		}
	}
}
