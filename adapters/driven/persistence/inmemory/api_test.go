package inmemory_test

import (
	"testing"

	"github.com/quii/go-fakes-and-contracts/adapters/driven/persistence/inmemory"
	"github.com/quii/go-fakes-and-contracts/domain/planner"
)

func TestInmemoryAPI1(t *testing.T) {
	planner.API1Contract{NewAPI1: func() planner.API1 {
		return inmemory.NewAPI1()
	}}.Test(t)
}
