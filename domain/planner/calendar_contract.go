package planner

import (
	"context"
	"github.com/alecthomas/assert/v2"
	"github.com/quii/go-fakes-and-contracts/domain/planner/internal/plannertest"
	"github.com/quii/go-fakes-and-contracts/domain/recipe"
	"testing"
	"time"
)

type Calendar interface {
	ScheduleMeal(ctx context.Context, recipe recipe.Recipe, date time.Time) error
	GetSchedule(ctx context.Context) (map[time.Time]recipe.Recipes, error)
}

type CalendarContract struct {
	NewCalendar func() Calendar
}

func (c CalendarContract) Test(t *testing.T) {
	t.Run("it returns what is put in", func(t *testing.T) {
		var (
			ctx         = context.Background()
			someRecipes = plannertest.RandomRecipes()
			tomorrow    = time.Now()
			sut         = c.NewCalendar()
		)
		for _, recipe := range someRecipes {
			assert.NoError(t, sut.ScheduleMeal(ctx, recipe, tomorrow))
		}
		got, err := sut.GetSchedule(ctx)
		assert.NoError(t, err)
		assert.Equal(t, someRecipes, got[tomorrow])
	})
}
