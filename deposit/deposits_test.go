package deposit

import (
	"testing"

	"github.com/dsncode/stash/model"
)

// Example test from Document (happy case)
func TestExampleDesposit(t *testing.T) {
	// create portfolios

	highRisk := model.CreatePortfolio("High Risk", model.Standard)
	retirement := model.CreatePortfolio("Retirement", model.Standard)

	// create deposit plans
	singleTimePlanPortfolios := []*model.PortfolioPlan{
		&model.PortfolioPlan{
			Portfolio:          highRisk,
			MaxAmountToDeposit: 10000,
		},
		&model.PortfolioPlan{
			Portfolio:          retirement,
			MaxAmountToDeposit: 500,
		},
	}

	montlyPlanPortfolio := []*model.PortfolioPlan{
		&model.PortfolioPlan{
			Portfolio:          highRisk,
			MaxAmountToDeposit: 0,
		},
		&model.PortfolioPlan{
			Portfolio:          retirement,
			MaxAmountToDeposit: 100,
		},
	}

	singleTime := model.CreateDepositPlan("Single Time", singleTimePlanPortfolios, model.SingleTime)
	montly := model.CreateDepositPlan("Montly", montlyPlanPortfolio, model.Montly)

	depositPlans := []*model.DespositPlan{singleTime, montly}

	janDeposit := model.Deposit{
		Amount: 10500,
		Month:  1,
		Year:   2020,
	}

	febDeposit := model.Deposit{
		Amount: 100,
		Month:  2,
		Year:   2020,
	}

	deposits := []model.Deposit{janDeposit, febDeposit}

	portfolios := ComputeSavingsDistribution(depositPlans, deposits)

	if len(portfolios) != 3 {
		t.Fatal("there should be 3 porfolios")
	}

	for _, portfolio := range portfolios {
		switch portfolio.Name {
		case "High Risk":
			if portfolio.Total != 10000 {
				t.Fatalf("%s got wrong amount. should be 1000, but got %d", portfolio.Name, portfolio.Total)
			}
			break
		case "Retirement":
			if portfolio.Total != 600 {
				t.Fatalf("%s got wrong amount. should be 100, but got %d", portfolio.Name, portfolio.Total)
			}
			break
		case "Default":
			if portfolio.Total != 0 {
				t.Fatalf("%s got wrong amount. should be 0, but got %d", portfolio.Name, portfolio.Total)
			}
			break
		}
	}

}
