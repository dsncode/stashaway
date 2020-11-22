package deposit

import (
	"log"
	"testing"

	"github.com/dsncode/stash/model"
)

var (
	depositPlans []*model.DespositPlan
)

func setup() {
	// create portfolios
	highRisk := model.CreatePortfolio("High Risk")
	retirement := model.CreatePortfolio("Retirement")

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

	depositPlans = []*model.DespositPlan{singleTime, montly}
}

// Example test from Document (happy case)
func TestExampleDesposit(t *testing.T) {
	setup()
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
	expectation := make(map[string]int16)
	expectation["High Risk"] = 10000
	expectation["Retirement"] = 600
	expectation["Default"] = 0

	validatePortfolioExpectations(t, portfolios, expectation)
}

// Example where customer send 100 more than he should. in this case, extra money goes to default account (no risk)
func TestExampleDespositWhenUserSendMoreMoneyThanHeShould(t *testing.T) {
	setup()
	janDeposit := model.Deposit{
		Amount: 10500,
		Month:  1,
		Year:   2020,
	}

	febDeposit := model.Deposit{
		Amount: 200, // user in this case, send 100 more than he should
		Month:  2,
		Year:   2020,
	}

	deposits := []model.Deposit{janDeposit, febDeposit}

	portfolios := ComputeSavingsDistribution(depositPlans, deposits)
	expectations := make(map[string]int16)
	expectations["High Risk"] = 10000
	expectations["Retirement"] = 600
	expectations["Default"] = 100
	validatePortfolioExpectations(t, portfolios, expectations)
}

func validatePortfolioExpectations(t *testing.T, portfolios []*model.Portfolio, expectations map[string]int16) {
	log.Println("validating...")
	for _, portfolio := range portfolios {

		if expectedValue, exists := expectations[portfolio.Name]; exists {
			if portfolio.Total != expectedValue {
				t.Fatalf("%s got wrong amount. should be %d, but got %d", portfolio.Name, expectedValue, portfolio.Total)
			}
		}
	}

}
