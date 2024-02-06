package inmemory_test

import (
	"testing"

	"github.com/quii/go-fakes-and-contracts/adapters/driven/persistence/inmemory"
	"github.com/quii/go-fakes-and-contracts/domain/planner"
)

func TestInMemoryCalendar(t *testing.T) {
	planner.CalendarContract{
		NewCalendar: func() planner.Calendar {
			return inmemory.NewCalendar()
		},
	}.Test(t)
}
