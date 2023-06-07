package inmemory

import (
	"context"
	"github.com/quii/go-fakes-and-contracts/domain/recipe"
	"time"
)

type Calendar struct {
	schedule map[time.Time]recipe.Recipes
}

func NewCalendar() *Calendar {
	return &Calendar{}
}
func (c *Calendar) ScheduleMeal(ctx context.Context, scheduledRecipe recipe.Recipe, date time.Time) error {
	if c.schedule == nil {
		c.schedule = make(map[time.Time]recipe.Recipes)
	}

	c.schedule[date] = append(c.schedule[date], scheduledRecipe)
	return nil
}
func (c *Calendar) GetSchedule(ctx context.Context) (map[time.Time]recipe.Recipes, error) {
	return c.schedule, nil
}
