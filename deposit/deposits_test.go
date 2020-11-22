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
	singleTimePlanPortfolios := []model.PortfolioPlan{
		model.PortfolioPlan{
			Portfolio:                         highRisk,
			MaxAmountToDepositPerTransacction: 10000,
		},
		model.PortfolioPlan{
			Portfolio:                         retirement,
			MaxAmountToDepositPerTransacction: 500,
		},
	}

	montlyPlanPortfolio := []model.PortfolioPlan{
		model.PortfolioPlan{
			Portfolio:                         highRisk,
			MaxAmountToDepositPerTransacction: 0,
		},
		model.PortfolioPlan{
			Portfolio:                         retirement,
			MaxAmountToDepositPerTransacction: 100,
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
	expectation["High Risk"] = 10000 // 10,000 from 10,500 goes to this account
	expectation["Retirement"] = 600  // 500 from first deposit goes here + 100 from the second deposit
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
	expectations["High Risk"] = 10000 // 10,000 goes here from first deposit
	expectations["Retirement"] = 600  // 500 from first deposit + 100 from second deposit (touch limit per month)
	expectations["Default"] = 100     // remaining money, is allocated into default account (no risk)
	validatePortfolioExpectations(t, portfolios, expectations)
}

// In this case, client send less money than he should
func TestExampleDespositWhenUserSendLessMoneyThanHeShould(t *testing.T) {
	setup()
	janDeposit := model.Deposit{
		Amount: 10000,
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
	expectations := make(map[string]int16)
	expectations["High Risk"] = 10000 // all first deposit goes straight to high risk account. in this case, there is no money to assign to second portfolio (retirement)
	expectations["Retirement"] = 100  // second deposit goes here
	expectations["Default"] = 0       // no extra money
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
