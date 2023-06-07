package sqlite_test

import (
	sqlite2 "github.com/quii/go-fakes-and-contracts/adapters/driven/persistence/sqlite"
	"github.com/quii/go-fakes-and-contracts/domain/planner"
	"testing"
)

func TestRecipeStore(t *testing.T) {
	planner.RecipeBookContract{
		NewBook: func() planner.RecipeBook {
			return sqlite2.NewRecipeStore(sqlite2.NewSQLiteClient())
		},
	}.Test(t)
}
