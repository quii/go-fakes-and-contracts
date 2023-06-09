package main

import (
	"fmt"
	"github.com/quii/go-fakes-and-contracts/adapters/driving/http"
	adapters "github.com/quii/go-fakes-and-contracts/cmd"
	"github.com/quii/go-fakes-and-contracts/domain/planner"
	"github.com/quii/go-fakes-and-contracts/specifications"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWebServer(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	var (
		port            = "8081"
		driver, cleanup = http.NewMealPlannerDriver(fmt.Sprintf("http://localhost:%s", port))
	)

	t.Cleanup(func() {
		assert.NoError(t, cleanup())
	})

	adapters.StartDockerServer(t, port, "webserver")
	specifications.MealPlanner{CreateDependencies: func() (planner.RecipeBook, planner.Pantry, planner.MealPlanner) {
		return driver, driver, driver
	}}.Test(t)
}
