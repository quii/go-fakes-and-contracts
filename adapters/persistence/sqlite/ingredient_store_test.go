package sqlite_test

import (
	"github.com/quii/go-fakes-and-contracts/adapters/persistence/sqlite"
	"github.com/quii/go-fakes-and-contracts/domain/planner"
	"testing"
)

/*
Notes

This was made _after_ we had create the store contract and verified it with the in memory fake

this gives us very clear guardrails to what to build
Make a type
Try and compile, it fails, implement interface, try running test, it'll fail, now implement it with SQL
*/

func TestSQLiteIngredientStore(t *testing.T) {
	planner.IngredientStoreContract{
		NewStore: func() planner.CloseableIngredientStore {
			return sqlite.NewIngredientStore()
		},
	}.Test(t)
}