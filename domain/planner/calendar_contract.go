package planner

import (
	"context"
	"github.com/quii/go-fakes-and-contracts/domain/planner/expect"
	"github.com/quii/go-fakes-and-contracts/domain/planner/plannertest"
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
			tomorrow    = time.Now().Add(24 * time.Hour)
			sut         = c.NewCalendar()
		)
		for _, r := range someRecipes {
			expect.NoErr(t, sut.ScheduleMeal(ctx, r, tomorrow))
		}
		got, err := sut.GetSchedule(ctx)

		expect.NoErr(t, err)
		expect.DeepEqual(t, got[tomorrow], someRecipes)
	})
}
