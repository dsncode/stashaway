package deposit

import (
	"testing"

	"github.com/dsncode/stash/model"
)

// Example test from Document (happy case)
func TestExampleDesposit(t *testing.T) {
	// create portfolios

	highRisk := model.CreatePortfolio("High Risk", 10000, model.SingleTime)
	retirement := model.CreatePortfolio("Retirement", 100, model.Montly)
	defaultPortfolio := model.CreatePortfolio("Default", model.UnlimitedAmount, model.Default)
}
