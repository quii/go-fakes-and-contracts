package inmemory

import (
	"github.com/quii/go-fakes-and-contracts/domain/planner"
	"testing"
)

func TestInmemoryAPI1(t *testing.T) {
	planner.API1Contract{NewAPI1: func() planner.API1 {
		return NewAPI1()
	}}.Test(t)
}
