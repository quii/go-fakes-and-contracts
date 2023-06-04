package inmemory

import (
	"github.com/quii/go-fakes-and-contracts/domain/planner"
	"testing"
)

func TestInMemoryCalendar(t *testing.T) {
	planner.CalendarContract{
		NewCalendar: func() planner.Calendar {
			return NewCalendar()
		},
	}.Test(t)
}
